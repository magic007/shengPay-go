package shengPay

/** 盛付通支付-查询
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
// var Namespace = "http://www.sdo.com/mas/api/query"

// NewQueryOrderAPI creates an initializes a QueryOrderAPI.
func NewQueryOrderAPI(cli *soap.Client) QueryOrderAPI {
	return &queryOrderAPI{cli}
}

// QueryOrderAPI was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type QueryOrderAPI interface {
	// QueryBatchOrder was auto-generated from WSDL.
	QueryBatchOrder(α *QueryBatchOrder) (β *QueryBatchOrderResponse, err error)

	// QueryOrder was auto-generated from WSDL.
	QueryOrder(α *QueryOrder) (β *QueryOrderResponse, err error)

	// QueryAsyBatchOrder was auto-generated from WSDL.
	QueryAsyBatchOrder(α *QueryAsyBatchOrder) (β *QueryAsyBatchOrderResponse, err error)

	// QueryOrderInfo was auto-generated from WSDL.
	QueryOrderInfo(α *QueryOrderInfo) (β *QueryOrderInfoResponse, err error)

	// QueryOrderCardInfo was auto-generated from WSDL.
	QueryOrderCardInfo(α *QueryOrderCardInfo) (β *QueryOrderCardInfoResponse, err error)

	// QueryQrcodeOrderInfo was auto-generated from WSDL.
	QueryQrcodeOrderInfo(α *QueryQrcodeOrderInfo) (β *QueryQrcodeOrderInfoResponse, err error)
}

// QueryB2COrderRequest was auto-generated from WSDL.
type QueryB2COrderRequest struct {
	CustomerNo string     `xml:"customerNo,omitempty" json:"customerNo,omitempty" yaml:"customerNo,omitempty"`
	Extension  *Extension `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header     *Header    `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	OrderNo    string     `xml:"orderNo,omitempty" json:"orderNo,omitempty" yaml:"orderNo,omitempty"`
	SessionId  string     `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Signature  *Signature `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TokenId    string     `xml:"tokenId,omitempty" json:"tokenId,omitempty" yaml:"tokenId,omitempty"`
	TransNo    string     `xml:"transNo,omitempty" json:"transNo,omitempty" yaml:"transNo,omitempty"`
}

// QueryB2COrderResponse was auto-generated from WSDL.
type QueryB2COrderResponse struct {
	CustomerName string      `xml:"customerName,omitempty" json:"customerName,omitempty" yaml:"customerName,omitempty"`
	Extension    *Extension  `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header       *Header     `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	NotifyUrl    string      `xml:"notifyUrl,omitempty" json:"notifyUrl,omitempty" yaml:"notifyUrl,omitempty"`
	OrderAmount  string      `xml:"orderAmount,omitempty" json:"orderAmount,omitempty" yaml:"orderAmount,omitempty"`
	OrderNo      string      `xml:"orderNo,omitempty" json:"orderNo,omitempty" yaml:"orderNo,omitempty"`
	OrderType    string      `xml:"orderType,omitempty" json:"orderType,omitempty" yaml:"orderType,omitempty"`
	ReturnInfo   *ReturnInfo `xml:"returnInfo,omitempty" json:"returnInfo,omitempty" yaml:"returnInfo,omitempty"`
	SessionId    string      `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Signature    *Signature  `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TransNo      string      `xml:"transNo,omitempty" json:"transNo,omitempty" yaml:"transNo,omitempty"`
	TransStatus  string      `xml:"transStatus,omitempty" json:"transStatus,omitempty" yaml:"transStatus,omitempty"`
	TransTime    string      `xml:"transTime,omitempty" json:"transTime,omitempty" yaml:"transTime,omitempty"`
}

// QueryOrderCardRequest was auto-generated from WSDL.
type QueryOrderCardRequest struct {
	CustomerNo string     `xml:"customerNo,omitempty" json:"customerNo,omitempty" yaml:"customerNo,omitempty"`
	Extension  *Extension `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header     *Header    `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	OrderNo    string     `xml:"orderNo,omitempty" json:"orderNo,omitempty" yaml:"orderNo,omitempty"`
	SessionId  string     `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Signature  *Signature `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TokenId    string     `xml:"tokenId,omitempty" json:"tokenId,omitempty" yaml:"tokenId,omitempty"`
	TransNo    string     `xml:"transNo,omitempty" json:"transNo,omitempty" yaml:"transNo,omitempty"`
}

// QueryOrderCardResponse was auto-generated from WSDL.
type QueryOrderCardResponse struct {
	CardInfos  []*CardInfo `xml:"cardInfos,omitempty" json:"cardInfos,omitempty" yaml:"cardInfos,omitempty"`
	Extension  *Extension  `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header     *Header     `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	ReturnInfo *ReturnInfo `xml:"returnInfo,omitempty" json:"returnInfo,omitempty" yaml:"returnInfo,omitempty"`
	Signature  *Signature  `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
}

// QueryQrcodeB2COrderRequest was auto-generated from WSDL.
type QueryQrcodeB2COrderRequest struct {
	Extension *Extension `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header    *Header    `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	Signature *Signature `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TransNo   string     `xml:"transNo,omitempty" json:"transNo,omitempty" yaml:"transNo,omitempty"`
}

// QueryQrcodeB2COrderResponse was auto-generated from WSDL.
type QueryQrcodeB2COrderResponse struct {
	CustomerLogoUrl string      `xml:"customerLogoUrl,omitempty" json:"customerLogoUrl,omitempty" yaml:"customerLogoUrl,omitempty"`
	CustomerName    string      `xml:"customerName,omitempty" json:"customerName,omitempty" yaml:"customerName,omitempty"`
	CustomerNo      string      `xml:"customerNo,omitempty" json:"customerNo,omitempty" yaml:"customerNo,omitempty"`
	Extension       *Extension  `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header          *Header     `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	OrderAmount     string      `xml:"orderAmount,omitempty" json:"orderAmount,omitempty" yaml:"orderAmount,omitempty"`
	OrderNo         string      `xml:"orderNo,omitempty" json:"orderNo,omitempty" yaml:"orderNo,omitempty"`
	OrderType       string      `xml:"orderType,omitempty" json:"orderType,omitempty" yaml:"orderType,omitempty"`
	ProductName     string      `xml:"productName,omitempty" json:"productName,omitempty" yaml:"productName,omitempty"`
	ReturnInfo      *ReturnInfo `xml:"returnInfo,omitempty" json:"returnInfo,omitempty" yaml:"returnInfo,omitempty"`
	SessionId       string      `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Signature       *Signature  `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TokenId         string      `xml:"tokenId,omitempty" json:"tokenId,omitempty" yaml:"tokenId,omitempty"`
	TransNo         string      `xml:"transNo,omitempty" json:"transNo,omitempty" yaml:"transNo,omitempty"`
	TransStatus     string      `xml:"transStatus,omitempty" json:"transStatus,omitempty" yaml:"transStatus,omitempty"`
	TransTime       string      `xml:"transTime,omitempty" json:"transTime,omitempty" yaml:"transTime,omitempty"`
}

// BatchOrderQueryRequest was auto-generated from WSDL.
type BatchOrderQueryRequest struct {
	BeginTime     string     `xml:"beginTime,omitempty" json:"beginTime,omitempty" yaml:"beginTime,omitempty"`
	EndTime       string     `xml:"endTime,omitempty" json:"endTime,omitempty" yaml:"endTime,omitempty"`
	Extension     *Extension `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header        *Header    `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	MemberIdType  string     `xml:"memberIdType,omitempty" json:"memberIdType,omitempty" yaml:"memberIdType,omitempty"`
	OrderNos      []string   `xml:"orderNos,omitempty" json:"orderNos,omitempty" yaml:"orderNos,omitempty"`
	PageNo        string     `xml:"pageNo,omitempty" json:"pageNo,omitempty" yaml:"pageNo,omitempty"`
	PageSize      string     `xml:"pageSize,omitempty" json:"pageSize,omitempty" yaml:"pageSize,omitempty"`
	PayerMemberId string     `xml:"payerMemberId,omitempty" json:"payerMemberId,omitempty" yaml:"payerMemberId,omitempty"`
	Signature     *Signature `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TimeType      string     `xml:"timeType,omitempty" json:"timeType,omitempty" yaml:"timeType,omitempty"`
}

// BatchOrderQueryResponse was auto-generated from WSDL.
type BatchOrderQueryResponse struct {
	Extension  *Extension  `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header     *Header     `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	Orders     []*Order    `xml:"orders,omitempty" json:"orders,omitempty" yaml:"orders,omitempty"`
	ReturnInfo *ReturnInfo `xml:"returnInfo,omitempty" json:"returnInfo,omitempty" yaml:"returnInfo,omitempty"`
	Signature  *Signature  `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
}

// CardInfo was auto-generated from WSDL.
type CardInfo struct {
	Balance     string `xml:"balance,omitempty" json:"balance,omitempty" yaml:"balance,omitempty"`
	CardNo      string `xml:"cardNo,omitempty" json:"cardNo,omitempty" yaml:"cardNo,omitempty"`
	ExpiredDate string `xml:"expiredDate,omitempty" json:"expiredDate,omitempty" yaml:"expiredDate,omitempty"`
	ExtJson     string `xml:"extJson,omitempty" json:"extJson,omitempty" yaml:"extJson,omitempty"`
	PaidAmount  string `xml:"paidAmount,omitempty" json:"paidAmount,omitempty" yaml:"paidAmount,omitempty"`
}

// Order was auto-generated from WSDL.
type Order struct {
	OrderAmoumt  string `xml:"orderAmoumt,omitempty" json:"orderAmoumt,omitempty" yaml:"orderAmoumt,omitempty"`
	OrderNo      string `xml:"orderNo,omitempty" json:"orderNo,omitempty" yaml:"orderNo,omitempty"`
	RefundAmount string `xml:"refundAmount,omitempty" json:"refundAmount,omitempty" yaml:"refundAmount,omitempty"`
	TransAmoumt  string `xml:"transAmoumt,omitempty" json:"transAmoumt,omitempty" yaml:"transAmoumt,omitempty"`
	TransNo      string `xml:"transNo,omitempty" json:"transNo,omitempty" yaml:"transNo,omitempty"`
	TransStatus  string `xml:"transStatus,omitempty" json:"transStatus,omitempty" yaml:"transStatus,omitempty"`
	TransTime    string `xml:"transTime,omitempty" json:"transTime,omitempty" yaml:"transTime,omitempty"`
	TransType    string `xml:"transType,omitempty" json:"transType,omitempty" yaml:"transType,omitempty"`
}

// OrderQueryRequest was auto-generated from WSDL.
type OrderQueryRequest struct {
	Extension  *Extension `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header     *Header    `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	MerchantNo string     `xml:"merchantNo,omitempty" json:"merchantNo,omitempty" yaml:"merchantNo,omitempty"`
	OrderNo    string     `xml:"orderNo,omitempty" json:"orderNo,omitempty" yaml:"orderNo,omitempty"`
	Signature  *Signature `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TransNo    string     `xml:"transNo,omitempty" json:"transNo,omitempty" yaml:"transNo,omitempty"`
}

// OrderQueryResponse was auto-generated from WSDL.
type OrderQueryResponse struct {
	Extension   *Extension  `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header      *Header     `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	OrderAmount string      `xml:"orderAmount,omitempty" json:"orderAmount,omitempty" yaml:"orderAmount,omitempty"`
	OrderNo     string      `xml:"orderNo,omitempty" json:"orderNo,omitempty" yaml:"orderNo,omitempty"`
	ReturnInfo  *ReturnInfo `xml:"returnInfo,omitempty" json:"returnInfo,omitempty" yaml:"returnInfo,omitempty"`
	Signature   *Signature  `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TransAmoumt string      `xml:"transAmoumt,omitempty" json:"transAmoumt,omitempty" yaml:"transAmoumt,omitempty"`
	TransNo     string      `xml:"transNo,omitempty" json:"transNo,omitempty" yaml:"transNo,omitempty"`
	TransStatus string      `xml:"transStatus,omitempty" json:"transStatus,omitempty" yaml:"transStatus,omitempty"`
	TransTime   string      `xml:"transTime,omitempty" json:"transTime,omitempty" yaml:"transTime,omitempty"`
}

// QueryAsyBatchOrder was auto-generated from WSDL.
type QueryAsyBatchOrder struct {
	XMLName xml.Name                   `xml:"http://www.sdo.com/mas/api/query queryAsyBatchOrder" json:"-" yaml:"-"`
	Arg0    *QueryBatchOrderAsyRequest `xml:"arg0,omitempty" json:"arg0,omitempty" yaml:"arg0,omitempty"`
}

// QueryAsyBatchOrderResponse was auto-generated from WSDL.
type QueryAsyBatchOrderResponse struct {
	Return *QueryBatchOrderAsyResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// QueryBatchOrder was auto-generated from WSDL.
type QueryBatchOrder struct {
	XMLName xml.Name                `xml:"http://www.sdo.com/mas/api/query queryBatchOrder" json:"-" yaml:"-"`
	Arg0    *BatchOrderQueryRequest `xml:"arg0,omitempty" json:"arg0,omitempty" yaml:"arg0,omitempty"`
}

// QueryBatchOrderAsyRequest was auto-generated from WSDL.
type QueryBatchOrderAsyRequest struct {
	BeginTime     string     `xml:"beginTime,omitempty" json:"beginTime,omitempty" yaml:"beginTime,omitempty"`
	EndTime       string     `xml:"endTime,omitempty" json:"endTime,omitempty" yaml:"endTime,omitempty"`
	Extension     *Extension `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	FileType      string     `xml:"fileType,omitempty" json:"fileType,omitempty" yaml:"fileType,omitempty"`
	Header        *Header    `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	MerchantNo    string     `xml:"merchantNo,omitempty" json:"merchantNo,omitempty" yaml:"merchantNo,omitempty"`
	NotifyURL     string     `xml:"notifyURL,omitempty" json:"notifyURL,omitempty" yaml:"notifyURL,omitempty"`
	PayerMemberId string     `xml:"payerMemberId,omitempty" json:"payerMemberId,omitempty" yaml:"payerMemberId,omitempty"`
	QueryMode     string     `xml:"queryMode,omitempty" json:"queryMode,omitempty" yaml:"queryMode,omitempty"`
	Signature     *Signature `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TimeType      string     `xml:"timeType,omitempty" json:"timeType,omitempty" yaml:"timeType,omitempty"`
	TransType     string     `xml:"transType,omitempty" json:"transType,omitempty" yaml:"transType,omitempty"`
}

// QueryBatchOrderAsyResponse was auto-generated from WSDL.
type QueryBatchOrderAsyResponse struct {
	Extension  *Extension  `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header     *Header     `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	ReturnInfo *ReturnInfo `xml:"returnInfo,omitempty" json:"returnInfo,omitempty" yaml:"returnInfo,omitempty"`
	Signature  *Signature  `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	Status     string      `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
}

// QueryBatchOrderResponse was auto-generated from WSDL.
type QueryBatchOrderResponse struct {
	Return *BatchOrderQueryResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// QueryOrder was auto-generated from WSDL.
type QueryOrder struct {
	XMLName xml.Name           `xml:"queryOrder" json:"-" yaml:"-"`
	Arg0    *OrderQueryRequest `xml:"arg0,omitempty" json:"arg0,omitempty" yaml:"arg0,omitempty"`
}

// QueryOrderCardInfo was auto-generated from WSDL.
type QueryOrderCardInfo struct {
	XMLName xml.Name               `xml:"http://www.sdo.com/mas/api/query queryOrderCardInfo" json:"-" yaml:"-"`
	Arg0    *QueryOrderCardRequest `xml:"arg0,omitempty" json:"arg0,omitempty" yaml:"arg0,omitempty"`
}

// QueryOrderCardInfoResponse was auto-generated from WSDL.
type QueryOrderCardInfoResponse struct {
	Return *QueryOrderCardResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// QueryOrderInfo was auto-generated from WSDL.
type QueryOrderInfo struct {
	XMLName xml.Name              `xml:"http://www.sdo.com/mas/api/query queryOrderInfo" json:"-" yaml:"-"`
	Arg0    *QueryB2COrderRequest `xml:"arg0,omitempty" json:"arg0,omitempty" yaml:"arg0,omitempty"`
}

// QueryOrderInfoResponse was auto-generated from WSDL.
type QueryOrderInfoResponse struct {
	Return *QueryB2COrderResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// QueryOrderResponse was auto-generated from WSDL.
type QueryOrderResponse struct {
	Return *OrderQueryResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// QueryQrcodeOrderInfo was auto-generated from WSDL.
type QueryQrcodeOrderInfo struct {
	XMLName xml.Name                    `xml:"http://www.sdo.com/mas/api/query queryQrcodeOrderInfo" json:"-" yaml:"-"`
	Arg0    *QueryQrcodeB2COrderRequest `xml:"arg0,omitempty" json:"arg0,omitempty" yaml:"arg0,omitempty"`
}

// QueryQrcodeOrderInfoResponse was auto-generated from WSDL.
type QueryQrcodeOrderInfoResponse struct {
	Return *QueryQrcodeB2COrderResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// queryOrderAPI implements the QueryOrderAPI interface.
type queryOrderAPI struct {
	cli *soap.Client
}

// QueryBatchOrder was auto-generated from WSDL.
func (p *queryOrderAPI) QueryBatchOrder(α *QueryBatchOrder) (β *QueryBatchOrderResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M QueryBatchOrderResponse `xml:"QueryBatchOrderResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// QueryOrder was auto-generated from WSDL.
func (p *queryOrderAPI) QueryOrder(α *QueryOrder) (β *QueryOrderResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M QueryOrderResponse `xml:"queryOrderResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// QueryAsyBatchOrder was auto-generated from WSDL.
func (p *queryOrderAPI) QueryAsyBatchOrder(α *QueryAsyBatchOrder) (β *QueryAsyBatchOrderResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M QueryAsyBatchOrderResponse `xml:"QueryAsyBatchOrderResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// QueryOrderInfo was auto-generated from WSDL.
func (p *queryOrderAPI) QueryOrderInfo(α *QueryOrderInfo) (β *QueryOrderInfoResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M QueryOrderInfoResponse `xml:"QueryOrderInfoResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// QueryOrderCardInfo was auto-generated from WSDL.
func (p *queryOrderAPI) QueryOrderCardInfo(α *QueryOrderCardInfo) (β *QueryOrderCardInfoResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M QueryOrderCardInfoResponse `xml:"QueryOrderCardInfoResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// QueryQrcodeOrderInfo was auto-generated from WSDL.
func (p *queryOrderAPI) QueryQrcodeOrderInfo(α *QueryQrcodeOrderInfo) (β *QueryQrcodeOrderInfoResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M QueryQrcodeOrderInfoResponse `xml:"QueryQrcodeOrderInfoResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}
