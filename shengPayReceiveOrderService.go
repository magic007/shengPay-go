package shengPay

/** 盛付通支付-收单
 * Created by Magic <zhongguovu@gmail.com>.
 * User: Magic
 * Date: 2016/12/26
 * Time: 11:35
 */
import (
	"encoding/xml"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://www.sdo.com/mas/api/receive/"

// NewReceiveOrderAPI creates an initializes a ReceiveOrderAPI.
func NewReceiveOrderAPI(cli *soap.Client) ReceiveOrderAPI {
	return &receiveOrderAPI{cli}
}

// ReceiveOrderAPI was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ReceiveOrderAPI interface {
	// ReceiveB2COrder was auto-generated from WSDL.
	ReceiveB2COrder(α *ReceiveB2COrder) (β *ReceiveB2COrderResponse, err error)
}

// ReceB2COrderRequest was auto-generated from WSDL.
type ReceB2COrderRequest struct {
	BuyerContact    string     `xml:"buyerContact,omitempty" json:"buyerContact,omitempty" yaml:"buyerContact,omitempty"`
	BuyerId         string     `xml:"buyerId,omitempty" json:"buyerId,omitempty" yaml:"buyerId,omitempty"`
	BuyerIp         string     `xml:"buyerIp,omitempty" json:"buyerIp,omitempty" yaml:"buyerIp,omitempty"`
	BuyerName       string     `xml:"buyerName,omitempty" json:"buyerName,omitempty" yaml:"buyerName,omitempty"`
	CardPayInfo     string     `xml:"cardPayInfo,omitempty" json:"cardPayInfo,omitempty" yaml:"cardPayInfo,omitempty"`
	CardValue       string     `xml:"cardValue,omitempty" json:"cardValue,omitempty" yaml:"cardValue,omitempty"`
	Currency        string     `xml:"currency,omitempty" json:"currency,omitempty" yaml:"currency,omitempty"`
	DepositId       string     `xml:"depositId,omitempty" json:"depositId,omitempty" yaml:"depositId,omitempty"`
	DepositIdType   string     `xml:"depositIdType,omitempty" json:"depositIdType,omitempty" yaml:"depositIdType,omitempty"`
	ExpireTime      string     `xml:"expireTime,omitempty" json:"expireTime,omitempty" yaml:"expireTime,omitempty"`
	Extension       *Extension `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header          *Header    `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	InstCode        string     `xml:"instCode,omitempty" json:"instCode,omitempty" yaml:"instCode,omitempty"`
	Language        string     `xml:"language,omitempty" json:"language,omitempty" yaml:"language,omitempty"`
	NotifyUrl       string     `xml:"notifyUrl,omitempty" json:"notifyUrl,omitempty" yaml:"notifyUrl,omitempty"`
	OrderAmount     string     `xml:"orderAmount,omitempty" json:"orderAmount,omitempty" yaml:"orderAmount,omitempty"`
	OrderNo         string     `xml:"orderNo,omitempty" json:"orderNo,omitempty" yaml:"orderNo,omitempty"`
	OrderTime       string     `xml:"orderTime,omitempty" json:"orderTime,omitempty" yaml:"orderTime,omitempty"`
	PageUrl         string     `xml:"pageUrl,omitempty" json:"pageUrl,omitempty" yaml:"pageUrl,omitempty"`
	PayChannel      string     `xml:"payChannel,omitempty" json:"payChannel,omitempty" yaml:"payChannel,omitempty"`
	PayType         string     `xml:"payType,omitempty" json:"payType,omitempty" yaml:"payType,omitempty"`
	PayeeId         string     `xml:"payeeId,omitempty" json:"payeeId,omitempty" yaml:"payeeId,omitempty"`
	PayerAuthTicket string     `xml:"payerAuthTicket,omitempty" json:"payerAuthTicket,omitempty" yaml:"payerAuthTicket,omitempty"`
	PayerId         string     `xml:"payerId,omitempty" json:"payerId,omitempty" yaml:"payerId,omitempty"`
	PayerMobileNo   string     `xml:"payerMobileNo,omitempty" json:"payerMobileNo,omitempty" yaml:"payerMobileNo,omitempty"`
	ProductDesc     string     `xml:"productDesc,omitempty" json:"productDesc,omitempty" yaml:"productDesc,omitempty"`
	ProductId       string     `xml:"productId,omitempty" json:"productId,omitempty" yaml:"productId,omitempty"`
	ProductName     string     `xml:"productName,omitempty" json:"productName,omitempty" yaml:"productName,omitempty"`
	ProductNum      string     `xml:"productNum,omitempty" json:"productNum,omitempty" yaml:"productNum,omitempty"`
	ProductUrl      string     `xml:"productUrl,omitempty" json:"productUrl,omitempty" yaml:"productUrl,omitempty"`
	SellerId        string     `xml:"sellerId,omitempty" json:"sellerId,omitempty" yaml:"sellerId,omitempty"`
	Signature       *Signature `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TerminalType    string     `xml:"terminalType,omitempty" json:"terminalType,omitempty" yaml:"terminalType,omitempty"`
	UnitPrice       string     `xml:"unitPrice,omitempty" json:"unitPrice,omitempty" yaml:"unitPrice,omitempty"`
}

// ReceB2COrderResponse was auto-generated from WSDL.
type ReceB2COrderResponse struct {
	CustomerLogoUrl string      `xml:"customerLogoUrl,omitempty" json:"customerLogoUrl,omitempty" yaml:"customerLogoUrl,omitempty"`
	CustomerName    string      `xml:"customerName,omitempty" json:"customerName,omitempty" yaml:"customerName,omitempty"`
	CustomerNo      string      `xml:"customerNo,omitempty" json:"customerNo,omitempty" yaml:"customerNo,omitempty"`
	Extension       *Extension  `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header          *Header     `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	OrderAmount     string      `xml:"orderAmount,omitempty" json:"orderAmount,omitempty" yaml:"orderAmount,omitempty"`
	OrderNo         string      `xml:"orderNo,omitempty" json:"orderNo,omitempty" yaml:"orderNo,omitempty"`
	OrderType       string      `xml:"orderType,omitempty" json:"orderType,omitempty" yaml:"orderType,omitempty"`
	ReturnInfo      *ReturnInfo `xml:"returnInfo,omitempty" json:"returnInfo,omitempty" yaml:"returnInfo,omitempty"`
	SessionId       string      `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Signature       *Signature  `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TokenId         string      `xml:"tokenId,omitempty" json:"tokenId,omitempty" yaml:"tokenId,omitempty"`
	TransNo         string      `xml:"transNo,omitempty" json:"transNo,omitempty" yaml:"transNo,omitempty"`
	TransStatus     string      `xml:"transStatus,omitempty" json:"transStatus,omitempty" yaml:"transStatus,omitempty"`
	TransTime       string      `xml:"transTime,omitempty" json:"transTime,omitempty" yaml:"transTime,omitempty"`
}

// Extension was auto-generated from WSDL.
type Extension struct {
	Ext1 string `xml:"ext1,omitempty" json:"ext1,omitempty" yaml:"ext1,omitempty"`
	Ext2 string `xml:"ext2,omitempty" json:"ext2,omitempty" yaml:"ext2,omitempty"`
	Ext3 string `xml:"ext3,omitempty" json:"ext3,omitempty" yaml:"ext3,omitempty"`
}

// Header was auto-generated from WSDL.
type Header struct {
	Charset  string   `xml:"charset,omitempty" json:"charset,omitempty" yaml:"charset,omitempty"`
	SendTime string   `xml:"sendTime,omitempty" json:"sendTime,omitempty" yaml:"sendTime,omitempty"`
	Sender   *Sender  `xml:"sender,omitempty" json:"sender,omitempty" yaml:"sender,omitempty"`
	Service  *Service `xml:"service,omitempty" json:"service,omitempty" yaml:"service,omitempty"`
	TraceNo  string   `xml:"traceNo,omitempty" json:"traceNo,omitempty" yaml:"traceNo,omitempty"`
}

// ReceiveB2COrder was auto-generated from WSDL.
type ReceiveB2COrder struct {
	XMLName xml.Name             `xml:"receiveB2COrder" json:"-" yaml:"-"`
	Arg0    *ReceB2COrderRequest `xml:"arg0,omitempty" json:"arg0,omitempty" yaml:"arg0,omitempty"`
}

// ReceiveB2COrderResponse was auto-generated from WSDL.
type ReceiveB2COrderResponse struct {
	Return *ReceB2COrderResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
	// Interests Interests
}

// ReturnInfo was auto-generated from WSDL.
type ReturnInfo struct {
	ErrorCode string `xml:"errorCode,omitempty" json:"errorCode,omitempty" yaml:"errorCode,omitempty"`
	ErrorMsg  string `xml:"errorMsg,omitempty" json:"errorMsg,omitempty" yaml:"errorMsg,omitempty"`
}

// Sender was auto-generated from WSDL.
type Sender struct {
	SenderId string `xml:"senderId,omitempty" json:"senderId,omitempty" yaml:"senderId,omitempty"`
}

// Service was auto-generated from WSDL.
type Service struct {
	ServiceCode string `xml:"serviceCode,omitempty" json:"serviceCode,omitempty" yaml:"serviceCode,omitempty"`
	Version     string `xml:"version,omitempty" json:"version,omitempty" yaml:"version,omitempty"`
}

// Signature was auto-generated from WSDL.
type Signature struct {
	SignMsg  string `xml:"signMsg,omitempty" json:"signMsg,omitempty" yaml:"signMsg,omitempty"`
	SignType string `xml:"signType,omitempty" json:"signType,omitempty" yaml:"signType,omitempty"`
}

// receiveOrderAPI implements the ReceiveOrderAPI interface.
type receiveOrderAPI struct {
	cli *soap.Client
}

// ReceiveB2COrder was auto-generated from WSDL.
func (p *receiveOrderAPI) ReceiveB2COrder(α *ReceiveB2COrder) (β *ReceiveB2COrderResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ReceiveB2COrderResponse `xml:"receiveB2COrderResponse"`
		}
	}{}

	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}

	return &γ.Body.M, nil
}
