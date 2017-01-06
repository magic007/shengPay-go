package shengPay

/** 盛付通支付
 * Created by Magic <zhongguovu@gmail.com>.
 * User: Magic
 * Date: 2016/12/26
 * Time: 11:35
 */
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/fiorix/wsdl2go/soap"
)

const (
	MERCHANT_KEY = "abcdefg"
	SENDER_ID    = "107537"
	PAY_CHANNEL  = "ap" //wp 微信 ap支付宝
	NOTIFY_URL   = "https://www.magic.com"
)

//盛付通查询订单
func Query(orderNO string) (*QueryOrderResponse, error) {

	url := "https://cardpay.shengpay.com/web-acquire-channel/services/queryOrderService"
	cli := soap.Client{
		URL:         url, //"http://schemas.xmlsoap.org/soap/envelope/",
		Namespace:   Namespace,
		ContentType: "application/xml;charset=utf-8",
	}
	conn := NewQueryOrderAPI(&cli)

	service := &Service{
		ServiceCode: "QUERY_ORDER_REQUEST",
		Version:     "V4.3.1.1.1",
	}
	header := &Header{
		Service: service,
		Charset: "UTF-8",
		//网关跟踪号, 保证唯一
		TraceNo: strconv.FormatInt(time.Now().Unix(), 10),
		Sender: &Sender{
			SenderId: SENDER_ID,
		},

		SendTime: time.Now().Format("20060102150405"),
	}
	extension := &Extension{
		Ext1: "",
		Ext2: "",
	}

	signature := &Signature{
		SignType: "MD5",
		SignMsg:  "",
	}

	orderQueryRequest := &OrderQueryRequest{
		Header: header,
		//商户订单号(商户根据实际情况自行更改), 需保证唯一
		OrderNo:    orderNO,
		Extension:  extension,
		MerchantNo: "107537",
		Signature:  signature,
		TransNo:    "",
	}

	sign := service.ServiceCode + service.Version +
		header.Charset + header.TraceNo + SENDER_ID + header.SendTime +
		orderQueryRequest.MerchantNo + orderQueryRequest.OrderNo + orderQueryRequest.TransNo +
		extension.Ext2 + signature.SignType + MERCHANT_KEY

	cacheKey := fmt.Sprintf("%s", sign)
	h := md5.New()
	h.Write([]byte(cacheKey))
	signature.SignMsg = hex.EncodeToString(h.Sum(nil))
	orderQueryRequest.Signature = signature
	result, err := conn.QueryOrder(&QueryOrder{Arg0: orderQueryRequest})

	if err != nil {
		return nil, err
	}

	return result, nil

}

//盛付通发起支付
func Pay(orderNO string, amoumt float64, productName string) (*ProcessB2CPayResponse, error) {

	//收单
	result, err := receiveOrderService(orderNO, amoumt, productName)
	if err != nil {
		return nil, err
	}

	tokenId := result.TokenId
	//网关支付标识
	sessionId := result.SessionId
	//网关订单号
	transNo := result.TransNo

	service := &Service{
		ServiceCode: "B2CPayment",
		Version:     "V4.1.1.1.1",
	}
	header := &Header{
		Service: service,
		Charset: "UTF-8",
		//网关跟踪号, 保证唯一
		TraceNo: strconv.FormatInt(time.Now().Unix(), 10),
		Sender: &Sender{
			SenderId: SENDER_ID,
		},
		SendTime: time.Now().Format("20060102150405"),
	}

	paymentItem := []*PaymentItem{{Key: "PAYER_IP", Value: "127.0.0.1"}, {Key: "CARD_INFO", Value: ""}}

	b2CPayment := &B2CPayment{
		PaymentType:  "PT312",
		InstCode:     "WXZF",
		PayChannel:   PAY_CHANNEL, //wp 微信 ap支付宝
		PaymentItems: paymentItem,
	}

	extension := &Extension{
		Ext1: "",
		Ext2: "",
	}
	signature := &Signature{
		SignType: "MD5",
		SignMsg:  "",
	}

	amoumtStr := fmt.Sprintf("%.2f", amoumt)
	order := &B2COrder{TransNo: transNo, OrderAmoumt: amoumtStr, OrderType: "OT001"}
	payer := &B2CPayer{PtId: "107537", PtIdType: "0004"}
	payee := &B2CPayee{}

	b2CPaymentRequest := &B2CPaymentRequest{
		Header:    header,
		Order:     order,
		Payer:     payer,
		Payee:     payee,
		Payment:   b2CPayment,
		TokenId:   tokenId,
		SessionId: sessionId,
		Extension: extension,
	}

	ptIdType := string(payer.PtIdType)

	sign := service.ServiceCode + service.Version + header.Charset + header.TraceNo +
		SENDER_ID + header.SendTime + order.TransNo + order.OrderAmoumt + order.OrderType +
		payer.PtId + ptIdType + payer.SdId + payer.MemberId + payer.AccountId + payer.AccountType +
		payer.AccountType + payer.PayableAmount + payer.PayableFee + payee.PtId + payee.SdId + payee.MemberId +
		payee.AccountId + payee.AccountType + payee.AccountType + payee.ReceivableAmount + payee.ReceivableFee +
		b2CPayment.PaymentType + b2CPayment.InstCode + b2CPayment.PayChannel

	for _, v := range paymentItem {
		sign = sign + v.Key + v.Value
	}
	sign = sign + b2CPaymentRequest.TokenId + b2CPaymentRequest.SessionId +
		extension.Ext1 + extension.Ext2 + signature.SignType + MERCHANT_KEY

	cacheKey := fmt.Sprintf("%s", sign)
	h := md5.New()
	h.Write([]byte(cacheKey))
	signature.SignMsg = hex.EncodeToString(h.Sum(nil))
	b2CPaymentRequest.Signature = signature

	url := "http://cardpay.shengpay.com/api-acquire-channel/services/paymentService"
	cli := soap.Client{
		URL:         url, //"http://schemas.xmlsoap.org/soap/envelope/",
		Namespace:   Namespace,
		ContentType: "application/xml;charset=utf-8",
	}
	conn := NewPaymentAPI(&cli)
	processB2CPayResponse, err := conn.ProcessB2CPay(&ProcessB2CPay{Arg0: b2CPaymentRequest})
	//收单结束
	if err != nil {
		return nil, err
	}
	return processB2CPayResponse, nil

}

func receiveOrderService(orderNO string, amoumt float64, productName string) (*ReceB2COrderResponse, error) {

	url := "http://cardpay.shengpay.com/api-acquire-channel/services/receiveOrderService"
	cli := soap.Client{
		URL:         url, //"http://schemas.xmlsoap.org/soap/envelope/",
		Namespace:   Namespace,
		ContentType: "application/xml;charset=utf-8",
	}
	conn := NewReceiveOrderAPI(&cli)

	service := &Service{
		ServiceCode: "B2CPayment",
		Version:     "V4.1.1.1.1",
	}
	header := &Header{
		Service: service,
		Charset: "UTF-8",
		//网关跟踪号, 保证唯一
		TraceNo: strconv.FormatInt(time.Now().Unix(), 10),
		Sender: &Sender{
			SenderId: SENDER_ID,
		},

		SendTime: time.Now().Format("20060102150405"),
	}
	extension := &Extension{
		Ext1: "",
		Ext2: "",
	}

	signature := &Signature{
		SignType: "MD5",
		SignMsg:  "",
	}

	receB2COrderRequest := &ReceB2COrderRequest{
		Header: header,
		//商户订单号(商户根据实际情况自行更改), 需保证唯一
		OrderNo: orderNO, //time.Now().Format("20060102150405"), //orderNo, // date("YmdHis") . uniqid(),
		//订单金额(商户根据实际情况自行更改), 最小单位(元)
		OrderAmount: fmt.Sprintf("%.2f", amoumt),
		//提交订单时间
		OrderTime: time.Now().Format("20060102150405"),
		Currency:  "CNY",
		Language:  "zh-CN",
		//页面通知地址(商户根据实际情况自行更改)
		PageUrl: "http://127.0.0.1/merchantNotify.htm",
		//后台通知地址(商户根据实际情况自行更改)
		//"notifyUrl" : "http://example.com/notify.php",

		NotifyUrl: NOTIFY_URL,
		Signature: signature,
		Extension: extension,
		//以下根据商户自身需求按照文档描述自行添加
		BuyerContact:  "",
		BuyerId:       "",
		BuyerIp:       "",
		BuyerName:     "",
		CardPayInfo:   "",
		CardValue:     "",
		DepositId:     "",
		DepositIdType: "",
		//订单过期时间
		ExpireTime:      "",
		InstCode:        "",
		PayChannel:      "",
		PayType:         "",
		PayeeId:         "",
		PayerAuthTicket: "",
		PayerId:         "",
		PayerMobileNo:   "",
		ProductDesc:     productName,
		ProductId:       "",
		ProductName:     productName,
		ProductNum:      "",
		ProductUrl:      "",
		SellerId:        "",
		TerminalType:    "",
		UnitPrice:       "",
	}

	sign := service.ServiceCode + service.Version +
		header.Charset + header.TraceNo + SENDER_ID + header.SendTime +
		receB2COrderRequest.OrderNo + receB2COrderRequest.OrderAmount + receB2COrderRequest.OrderTime +
		receB2COrderRequest.ExpireTime + receB2COrderRequest.Currency + receB2COrderRequest.PayType +
		receB2COrderRequest.PayChannel + receB2COrderRequest.InstCode + receB2COrderRequest.CardValue +
		receB2COrderRequest.Language + receB2COrderRequest.PageUrl + receB2COrderRequest.NotifyUrl +
		receB2COrderRequest.TerminalType + receB2COrderRequest.ProductId + receB2COrderRequest.ProductName +
		receB2COrderRequest.ProductNum + receB2COrderRequest.UnitPrice + receB2COrderRequest.ProductDesc +
		receB2COrderRequest.ProductUrl + receB2COrderRequest.SellerId + receB2COrderRequest.BuyerName +
		receB2COrderRequest.BuyerId + receB2COrderRequest.BuyerContact + receB2COrderRequest.BuyerIp +
		receB2COrderRequest.PayeeId + receB2COrderRequest.DepositId + receB2COrderRequest.DepositIdType +
		receB2COrderRequest.PayerId + receB2COrderRequest.CardPayInfo + receB2COrderRequest.PayerMobileNo +
		receB2COrderRequest.PayerAuthTicket + extension.Ext1 + extension.Ext2 + signature.SignType + MERCHANT_KEY

	cacheKey := fmt.Sprintf("%s", sign)
	h := md5.New()
	h.Write([]byte(cacheKey))
	signature.SignMsg = hex.EncodeToString(h.Sum(nil))
	receB2COrderRequest.Signature = signature
	result, err := conn.ReceiveB2COrder(&ReceiveB2COrder{Arg0: receB2COrderRequest})

	// fmt.Printf("result=======:%+v, err:%+v", result.Return, err)

	//收单结束
	if err != nil {
		return nil, err
	}

	return result.Return, nil
}
