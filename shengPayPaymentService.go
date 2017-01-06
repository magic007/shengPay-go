package shengPay

import (
	"encoding/xml"
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
// var Namespace = "http://www.sdo.com/mas/api/payment/"

// NewPaymentAPI creates an initializes a PaymentAPI.
func NewPaymentAPI(cli *soap.Client) PaymentAPI {
	return &paymentAPI{cli}
}

// PaymentAPI was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type PaymentAPI interface {
	// ProcessPreApprovedPay was auto-generated from WSDL.
	ProcessPreApprovedPay(α *ProcessPreApprovedPay) (β *ProcessPreApprovedPayResponse, err error)

	// ProcessB2CPay was auto-generated from WSDL.
	ProcessB2CPay(α *ProcessB2CPay) (β *ProcessB2CPayResponse, err error)

	// ProcessB2CCombinatedPay was auto-generated from WSDL.
	ProcessB2CCombinatedPay(α *ProcessB2CCombinatedPay) (β *ProcessB2CCombinatedPayResponse, err error)

	// ProcessTransfer was auto-generated from WSDL.
	ProcessTransfer(α *ProcessTransfer) (β *ProcessTransferResponse, err error)
}

// MemberIdType was auto-generated from WSDL.
type memberIdType string

// Validate validates memberIdType.
func (v memberIdType) Validate() bool {
	for _, vv := range []string{
		"SD_ID",
		"PT_ID",
		"MEMBER_ID",
		"MERCHANT_NO",
		"SFT_LOGIN_ID",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AccountIdentify was auto-generated from WSDL.
type AccountIdentify struct {
	AccountNo    string `xml:"accountNo,omitempty" json:"accountNo,omitempty" yaml:"accountNo,omitempty"`
	AccountType  string `xml:"accountType,omitempty" json:"accountType,omitempty" yaml:"accountType,omitempty"`
	Extension    string `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	MemberId     string `xml:"memberId,omitempty" json:"memberId,omitempty" yaml:"memberId,omitempty"`
	MemberIdType string `xml:"memberIdType,omitempty" json:"memberIdType,omitempty" yaml:"memberIdType,omitempty"`
}

// B2COrder was auto-generated from WSDL.
type B2COrder struct {
	OrderAmoumt string `xml:"orderAmoumt,omitempty" json:"orderAmoumt,omitempty" yaml:"orderAmoumt,omitempty"`
	OrderAmount string `xml:"orderAmount,omitempty" json:"orderAmount,omitempty" yaml:"orderAmount,omitempty"`
	OrderType   string `xml:"orderType,omitempty" json:"orderType,omitempty" yaml:"orderType,omitempty"`
	TransNo     string `xml:"transNo,omitempty" json:"transNo,omitempty" yaml:"transNo,omitempty"`
}

// B2CPayee was auto-generated from WSDL.
type B2CPayee struct {
	AccountId        string `xml:"accountId,omitempty" json:"accountId,omitempty" yaml:"accountId,omitempty"`
	AccountType      string `xml:"accountType,omitempty" json:"accountType,omitempty" yaml:"accountType,omitempty"`
	MemberId         string `xml:"memberId,omitempty" json:"memberId,omitempty" yaml:"memberId,omitempty"`
	PtId             string `xml:"ptId,omitempty" json:"ptId,omitempty" yaml:"ptId,omitempty"`
	ReceivableAmount string `xml:"receivableAmount,omitempty" json:"receivableAmount,omitempty" yaml:"receivableAmount,omitempty"`
	ReceivableFee    string `xml:"receivableFee,omitempty" json:"receivableFee,omitempty" yaml:"receivableFee,omitempty"`
	SdId             string `xml:"sdId,omitempty" json:"sdId,omitempty" yaml:"sdId,omitempty"`
}

// B2CPayer was auto-generated from WSDL.
type B2CPayer struct {
	AccountId     string       `xml:"accountId,omitempty" json:"accountId,omitempty" yaml:"accountId,omitempty"`
	AccountType   string       `xml:"accountType,omitempty" json:"accountType,omitempty" yaml:"accountType,omitempty"`
	MemberId      string       `xml:"memberId,omitempty" json:"memberId,omitempty" yaml:"memberId,omitempty"`
	PayableAmount string       `xml:"payableAmount,omitempty" json:"payableAmount,omitempty" yaml:"payableAmount,omitempty"`
	PayableFee    string       `xml:"payableFee,omitempty" json:"payableFee,omitempty" yaml:"payableFee,omitempty"`
	PtId          string       `xml:"ptId,omitempty" json:"ptId,omitempty" yaml:"ptId,omitempty"`
	PtIdType      memberIdType `xml:"ptIdType,omitempty" json:"ptIdType,omitempty" yaml:"ptIdType,omitempty"`
	SdId          string       `xml:"sdId,omitempty" json:"sdId,omitempty" yaml:"sdId,omitempty"`
}

// B2CPayment was auto-generated from WSDL.
type B2CPayment struct {
	InstCode     string         `xml:"instCode,omitempty" json:"instCode,omitempty" yaml:"instCode,omitempty"`
	PayChannel   string         `xml:"payChannel,omitempty" json:"payChannel,omitempty" yaml:"payChannel,omitempty"`
	PaymentItems []*PaymentItem `xml:"paymentItems,omitempty" json:"paymentItems,omitempty" yaml:"paymentItems,omitempty"`
	PaymentType  string         `xml:"paymentType,omitempty" json:"paymentType,omitempty" yaml:"paymentType,omitempty"`
}

// B2CPaymentCombinatedDetailInfo was auto-generated from WSDL.
type B2CPaymentCombinatedDetailInfo struct {
	InstCode                string         `xml:"instCode,omitempty" json:"instCode,omitempty" yaml:"instCode,omitempty"`
	PayChannel              string         `xml:"payChannel,omitempty" json:"payChannel,omitempty" yaml:"payChannel,omitempty"`
	PayeeInfo               *B2CPayee      `xml:"payeeInfo,omitempty" json:"payeeInfo,omitempty" yaml:"payeeInfo,omitempty"`
	PayerInfo               *B2CPayer      `xml:"payerInfo,omitempty" json:"payerInfo,omitempty" yaml:"payerInfo,omitempty"`
	PaymentDetailParameters []*PaymentItem `xml:"paymentDetailParameters,omitempty" json:"paymentDetailParameters,omitempty" yaml:"paymentDetailParameters,omitempty"`
	PaymentType             string         `xml:"paymentType,omitempty" json:"paymentType,omitempty" yaml:"paymentType,omitempty"`
}

// B2CPaymentCombinatedRequest was auto-generated from WSDL.
type B2CPaymentCombinatedRequest struct {
	CombExtensionInfo            *Extension                        `xml:"combExtensionInfo,omitempty" json:"combExtensionInfo,omitempty" yaml:"combExtensionInfo,omitempty"`
	CombOrderInfo                *B2COrder                         `xml:"combOrderInfo,omitempty" json:"combOrderInfo,omitempty" yaml:"combOrderInfo,omitempty"`
	CombRequestHeader            *Header                           `xml:"combRequestHeader,omitempty" json:"combRequestHeader,omitempty" yaml:"combRequestHeader,omitempty"`
	CombSignature                *Signature                        `xml:"combSignature,omitempty" json:"combSignature,omitempty" yaml:"combSignature,omitempty"`
	CombTokenID                  string                            `xml:"combTokenID,omitempty" json:"combTokenID,omitempty" yaml:"combTokenID,omitempty"`
	CombinatedPaymentDetailInfos []*B2CPaymentCombinatedDetailInfo `xml:"combinatedPaymentDetailInfos,omitempty" json:"combinatedPaymentDetailInfos,omitempty" yaml:"combinatedPaymentDetailInfos,omitempty"`
	SessionID                    string                            `xml:"sessionID,omitempty" json:"sessionID,omitempty" yaml:"sessionID,omitempty"`
}

// B2CPaymentCombinatedResponse was auto-generated from WSDL.
type B2CPaymentCombinatedResponse struct {
	CombResponseExtension *Extension                 `xml:"combResponseExtension,omitempty" json:"combResponseExtension,omitempty" yaml:"combResponseExtension,omitempty"`
	CombResponseHeader    *Header                    `xml:"combResponseHeader,omitempty" json:"combResponseHeader,omitempty" yaml:"combResponseHeader,omitempty"`
	CombResponseInfos     []*CombPaymentResponseInfo `xml:"combResponseInfos,omitempty" json:"combResponseInfos,omitempty" yaml:"combResponseInfos,omitempty"`
	CombResponseSignature *Signature                 `xml:"combResponseSignature,omitempty" json:"combResponseSignature,omitempty" yaml:"combResponseSignature,omitempty"`
	CombReturnInfo        *ReturnInfo                `xml:"combReturnInfo,omitempty" json:"combReturnInfo,omitempty" yaml:"combReturnInfo,omitempty"`
	CustomerOrderNo       string                     `xml:"customerOrderNo,omitempty" json:"customerOrderNo,omitempty" yaml:"customerOrderNo,omitempty"`
	OrderAmount           string                     `xml:"orderAmount,omitempty" json:"orderAmount,omitempty" yaml:"orderAmount,omitempty"`
	SessionId             string                     `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	TransNo               string                     `xml:"transNo,omitempty" json:"transNo,omitempty" yaml:"transNo,omitempty"`
	TransStatus           string                     `xml:"transStatus,omitempty" json:"transStatus,omitempty" yaml:"transStatus,omitempty"`
}

// B2CPaymentRequest was auto-generated from WSDL.
type B2CPaymentRequest struct {
	Extension *Extension  `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header    *Header     `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	Order     *B2COrder   `xml:"order,omitempty" json:"order,omitempty" yaml:"order,omitempty"`
	Payee     *B2CPayee   `xml:"payee,omitempty" json:"payee,omitempty" yaml:"payee,omitempty"`
	Payer     *B2CPayer   `xml:"payer,omitempty" json:"payer,omitempty" yaml:"payer,omitempty"`
	Payment   *B2CPayment `xml:"payment,omitempty" json:"payment,omitempty" yaml:"payment,omitempty"`
	SessionId string      `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Signature *Signature  `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TokenId   string      `xml:"tokenId,omitempty" json:"tokenId,omitempty" yaml:"tokenId,omitempty"`
}

// B2CPaymentResponse was auto-generated from WSDL.
type B2CPaymentResponse struct {
	BankForm      string      `xml:"bankForm,omitempty" json:"bankForm,omitempty" yaml:"bankForm,omitempty"`
	Extension     *Extension  `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header        *Header     `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	OrderAmount   string      `xml:"orderAmount,omitempty" json:"orderAmount,omitempty" yaml:"orderAmount,omitempty"`
	OrderNo       string      `xml:"orderNo,omitempty" json:"orderNo,omitempty" yaml:"orderNo,omitempty"`
	PaidAmount    string      `xml:"paidAmount,omitempty" json:"paidAmount,omitempty" yaml:"paidAmount,omitempty"`
	PaidFee       string      `xml:"paidFee,omitempty" json:"paidFee,omitempty" yaml:"paidFee,omitempty"`
	PayableAmount string      `xml:"payableAmount,omitempty" json:"payableAmount,omitempty" yaml:"payableAmount,omitempty"`
	PayableFee    string      `xml:"payableFee,omitempty" json:"payableFee,omitempty" yaml:"payableFee,omitempty"`
	PaymentNo     string      `xml:"paymentNo,omitempty" json:"paymentNo,omitempty" yaml:"paymentNo,omitempty"`
	PaymentStatus string      `xml:"paymentStatus,omitempty" json:"paymentStatus,omitempty" yaml:"paymentStatus,omitempty"`
	PaymentTime   string      `xml:"paymentTime,omitempty" json:"paymentTime,omitempty" yaml:"paymentTime,omitempty"`
	ReceivableFee string      `xml:"receivableFee,omitempty" json:"receivableFee,omitempty" yaml:"receivableFee,omitempty"`
	ReceivedFee   string      `xml:"receivedFee,omitempty" json:"receivedFee,omitempty" yaml:"receivedFee,omitempty"`
	ReturnInfo    *ReturnInfo `xml:"returnInfo,omitempty" json:"returnInfo,omitempty" yaml:"returnInfo,omitempty"`
	SessionId     string      `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Signature     *Signature  `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TransNo       string      `xml:"transNo,omitempty" json:"transNo,omitempty" yaml:"transNo,omitempty"`
	TransStatus   string      `xml:"transStatus,omitempty" json:"transStatus,omitempty" yaml:"transStatus,omitempty"`
}

// CombPaymentResponseInfo was auto-generated from WSDL.
type CombPaymentResponseInfo struct {
	BankForm      string `xml:"bankForm,omitempty" json:"bankForm,omitempty" yaml:"bankForm,omitempty"`
	InstCode      string `xml:"instCode,omitempty" json:"instCode,omitempty" yaml:"instCode,omitempty"`
	PaidAmount    string `xml:"paidAmount,omitempty" json:"paidAmount,omitempty" yaml:"paidAmount,omitempty"`
	PaidFee       string `xml:"paidFee,omitempty" json:"paidFee,omitempty" yaml:"paidFee,omitempty"`
	PayChannel    string `xml:"payChannel,omitempty" json:"payChannel,omitempty" yaml:"payChannel,omitempty"`
	PayableAmount string `xml:"payableAmount,omitempty" json:"payableAmount,omitempty" yaml:"payableAmount,omitempty"`
	PayableFee    string `xml:"payableFee,omitempty" json:"payableFee,omitempty" yaml:"payableFee,omitempty"`
	PaymentNo     string `xml:"paymentNo,omitempty" json:"paymentNo,omitempty" yaml:"paymentNo,omitempty"`
	PaymentStatus string `xml:"paymentStatus,omitempty" json:"paymentStatus,omitempty" yaml:"paymentStatus,omitempty"`
	PaymentTime   string `xml:"paymentTime,omitempty" json:"paymentTime,omitempty" yaml:"paymentTime,omitempty"`
	PaymentType   string `xml:"paymentType,omitempty" json:"paymentType,omitempty" yaml:"paymentType,omitempty"`
	ReceivableFee string `xml:"receivableFee,omitempty" json:"receivableFee,omitempty" yaml:"receivableFee,omitempty"`
	ReceivedFee   string `xml:"receivedFee,omitempty" json:"receivedFee,omitempty" yaml:"receivedFee,omitempty"`
}

// PaySchedule was auto-generated from WSDL.
type PaySchedule struct {
	DaysOfMonth  string `xml:"daysOfMonth,omitempty" json:"daysOfMonth,omitempty" yaml:"daysOfMonth,omitempty"`
	DaysOfWeek   string `xml:"daysOfWeek,omitempty" json:"daysOfWeek,omitempty" yaml:"daysOfWeek,omitempty"`
	EndCount     string `xml:"endCount,omitempty" json:"endCount,omitempty" yaml:"endCount,omitempty"`
	EndTime      string `xml:"endTime,omitempty" json:"endTime,omitempty" yaml:"endTime,omitempty"`
	Extension    string `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Frequency    string `xml:"frequency,omitempty" json:"frequency,omitempty" yaml:"frequency,omitempty"`
	Hour         string `xml:"hour,omitempty" json:"hour,omitempty" yaml:"hour,omitempty"`
	Interval     string `xml:"interval,omitempty" json:"interval,omitempty" yaml:"interval,omitempty"`
	MonthsOfYear string `xml:"monthsOfYear,omitempty" json:"monthsOfYear,omitempty" yaml:"monthsOfYear,omitempty"`
	StartTime    string `xml:"startTime,omitempty" json:"startTime,omitempty" yaml:"startTime,omitempty"`
}

// PaymentItem was auto-generated from WSDL.
type PaymentItem struct {
	Key   string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Value string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// PreApprovedPaymentRequest was auto-generated from WSDL.
type PreApprovedPaymentRequest struct {
	Amount       string           `xml:"amount,omitempty" json:"amount,omitempty" yaml:"amount,omitempty"`
	BizCode      string           `xml:"bizCode,omitempty" json:"bizCode,omitempty" yaml:"bizCode,omitempty"`
	ContractNO   string           `xml:"contractNO,omitempty" json:"contractNO,omitempty" yaml:"contractNO,omitempty"`
	Currency     string           `xml:"currency,omitempty" json:"currency,omitempty" yaml:"currency,omitempty"`
	Ext          string           `xml:"ext,omitempty" json:"ext,omitempty" yaml:"ext,omitempty"`
	Extension    *Extension       `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header       *Header          `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	InstCode     string           `xml:"instCode,omitempty" json:"instCode,omitempty" yaml:"instCode,omitempty"`
	Memo         string           `xml:"memo,omitempty" json:"memo,omitempty" yaml:"memo,omitempty"`
	NotifyURL    string           `xml:"notifyURL,omitempty" json:"notifyURL,omitempty" yaml:"notifyURL,omitempty"`
	OrderNo      string           `xml:"orderNo,omitempty" json:"orderNo,omitempty" yaml:"orderNo,omitempty"`
	OrderTime    string           `xml:"orderTime,omitempty" json:"orderTime,omitempty" yaml:"orderTime,omitempty"`
	PayMode      string           `xml:"payMode,omitempty" json:"payMode,omitempty" yaml:"payMode,omitempty"`
	PaySchedule  *PaySchedule     `xml:"paySchedule,omitempty" json:"paySchedule,omitempty" yaml:"paySchedule,omitempty"`
	PayeeAccount *AccountIdentify `xml:"payeeAccount,omitempty" json:"payeeAccount,omitempty" yaml:"payeeAccount,omitempty"`
	PayerAccount *AccountIdentify `xml:"payerAccount,omitempty" json:"payerAccount,omitempty" yaml:"payerAccount,omitempty"`
	Signature    *Signature       `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
}

// PreApprovedPaymentResponse was auto-generated from WSDL.
type PreApprovedPaymentResponse struct {
	Amount      string      `xml:"amount,omitempty" json:"amount,omitempty" yaml:"amount,omitempty"`
	Extension   *Extension  `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header      *Header     `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	OrderNo     string      `xml:"orderNo,omitempty" json:"orderNo,omitempty" yaml:"orderNo,omitempty"`
	ReturnInfo  *ReturnInfo `xml:"returnInfo,omitempty" json:"returnInfo,omitempty" yaml:"returnInfo,omitempty"`
	Signature   *Signature  `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TransAmoumt string      `xml:"transAmoumt,omitempty" json:"transAmoumt,omitempty" yaml:"transAmoumt,omitempty"`
	TransNo     string      `xml:"transNo,omitempty" json:"transNo,omitempty" yaml:"transNo,omitempty"`
	TransStatus string      `xml:"transStatus,omitempty" json:"transStatus,omitempty" yaml:"transStatus,omitempty"`
	TransTime   string      `xml:"transTime,omitempty" json:"transTime,omitempty" yaml:"transTime,omitempty"`
}

// ProcessB2CCombinatedPay was auto-generated from WSDL.
type ProcessB2CCombinatedPay struct {
	XMLName xml.Name                     `xml:"http://www.sdo.com/mas/api/payment/ processB2CCombinatedPay" json:"-" yaml:"-"`
	Arg0    *B2CPaymentCombinatedRequest `xml:"arg0,omitempty" json:"arg0,omitempty" yaml:"arg0,omitempty"`
}

// ProcessB2CCombinatedPayResponse was auto-generated from WSDL.
type ProcessB2CCombinatedPayResponse struct {
	Return *B2CPaymentCombinatedResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ProcessB2CPay was auto-generated from WSDL.
type ProcessB2CPay struct {
	XMLName xml.Name           `xml:"processB2CPay" json:"-" yaml:"-"`
	Arg0    *B2CPaymentRequest `xml:"arg0,omitempty" json:"arg0,omitempty" yaml:"arg0,omitempty"`
}

// ProcessB2CPayResponse was auto-generated from WSDL.
type ProcessB2CPayResponse struct {
	Return *B2CPaymentResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ProcessPreApprovedPay was auto-generated from WSDL.
type ProcessPreApprovedPay struct {
	XMLName xml.Name                   `xml:"http://www.sdo.com/mas/api/payment/ processPreApprovedPay" json:"-" yaml:"-"`
	Arg0    *PreApprovedPaymentRequest `xml:"arg0,omitempty" json:"arg0,omitempty" yaml:"arg0,omitempty"`
}

// ProcessPreApprovedPayResponse was auto-generated from WSDL.
type ProcessPreApprovedPayResponse struct {
	Return *PreApprovedPaymentResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ProcessTransfer was auto-generated from WSDL.
type ProcessTransfer struct {
	XMLName xml.Name                `xml:"http://www.sdo.com/mas/api/payment/ processTransfer" json:"-" yaml:"-"`
	Arg0    *TransferPaymentRequest `xml:"arg0,omitempty" json:"arg0,omitempty" yaml:"arg0,omitempty"`
}

// ProcessTransferResponse was auto-generated from WSDL.
type ProcessTransferResponse struct {
	Return *TransferPaymentResponse `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// TransferOrder was auto-generated from WSDL.
type TransferOrder struct {
	ClientIP    string `xml:"clientIP,omitempty" json:"clientIP,omitempty" yaml:"clientIP,omitempty"`
	ContractNO  string `xml:"contractNO,omitempty" json:"contractNO,omitempty" yaml:"contractNO,omitempty"`
	Currency    string `xml:"currency,omitempty" json:"currency,omitempty" yaml:"currency,omitempty"`
	Extension   string `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Memo        string `xml:"memo,omitempty" json:"memo,omitempty" yaml:"memo,omitempty"`
	NotifyURL   string `xml:"notifyURL,omitempty" json:"notifyURL,omitempty" yaml:"notifyURL,omitempty"`
	OrderAmount string `xml:"orderAmount,omitempty" json:"orderAmount,omitempty" yaml:"orderAmount,omitempty"`
	OrderNo     string `xml:"orderNo,omitempty" json:"orderNo,omitempty" yaml:"orderNo,omitempty"`
	OrderTime   string `xml:"orderTime,omitempty" json:"orderTime,omitempty" yaml:"orderTime,omitempty"`
	ReturnURL   string `xml:"returnURL,omitempty" json:"returnURL,omitempty" yaml:"returnURL,omitempty"`
	TransType   string `xml:"transType,omitempty" json:"transType,omitempty" yaml:"transType,omitempty"`
}

// TransferPayee was auto-generated from WSDL.
type TransferPayee struct {
	AccountNo    string `xml:"accountNo,omitempty" json:"accountNo,omitempty" yaml:"accountNo,omitempty"`
	Extension    string `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	MemberId     string `xml:"memberId,omitempty" json:"memberId,omitempty" yaml:"memberId,omitempty"`
	MemberIdType string `xml:"memberIdType,omitempty" json:"memberIdType,omitempty" yaml:"memberIdType,omitempty"`
}

// TransferPayer was auto-generated from WSDL.
type TransferPayer struct {
	CardValue     string `xml:"cardValue,omitempty" json:"cardValue,omitempty" yaml:"cardValue,omitempty"`
	DynamicPwd    string `xml:"dynamicPwd,omitempty" json:"dynamicPwd,omitempty" yaml:"dynamicPwd,omitempty"`
	Extension     string `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	InstAccountNo string `xml:"instAccountNo,omitempty" json:"instAccountNo,omitempty" yaml:"instAccountNo,omitempty"`
	InstCode      string `xml:"instCode,omitempty" json:"instCode,omitempty" yaml:"instCode,omitempty"`
	MemberId      string `xml:"memberId,omitempty" json:"memberId,omitempty" yaml:"memberId,omitempty"`
	MemberIdType  string `xml:"memberIdType,omitempty" json:"memberIdType,omitempty" yaml:"memberIdType,omitempty"`
	MobileNo      string `xml:"mobileNo,omitempty" json:"mobileNo,omitempty" yaml:"mobileNo,omitempty"`
	Password      string `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	PayMode       string `xml:"payMode,omitempty" json:"payMode,omitempty" yaml:"payMode,omitempty"`
	PayerName     string `xml:"payerName,omitempty" json:"payerName,omitempty" yaml:"payerName,omitempty"`
}

// TransferPaymentRequest was auto-generated from WSDL.
type TransferPaymentRequest struct {
	Extension *Extension     `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header    *Header        `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	Order     *TransferOrder `xml:"order,omitempty" json:"order,omitempty" yaml:"order,omitempty"`
	Payee     *TransferPayee `xml:"payee,omitempty" json:"payee,omitempty" yaml:"payee,omitempty"`
	Payer     *TransferPayer `xml:"payer,omitempty" json:"payer,omitempty" yaml:"payer,omitempty"`
	Signature *Signature     `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
}

// TransferPaymentResponse was auto-generated from WSDL.
type TransferPaymentResponse struct {
	Extension   *Extension  `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
	Header      *Header     `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	OrderAmount string      `xml:"orderAmount,omitempty" json:"orderAmount,omitempty" yaml:"orderAmount,omitempty"`
	OrderNo     string      `xml:"orderNo,omitempty" json:"orderNo,omitempty" yaml:"orderNo,omitempty"`
	PayUrl      string      `xml:"payUrl,omitempty" json:"payUrl,omitempty" yaml:"payUrl,omitempty"`
	ReturnInfo  *ReturnInfo `xml:"returnInfo,omitempty" json:"returnInfo,omitempty" yaml:"returnInfo,omitempty"`
	Signature   *Signature  `xml:"signature,omitempty" json:"signature,omitempty" yaml:"signature,omitempty"`
	TransAmoumt string      `xml:"transAmoumt,omitempty" json:"transAmoumt,omitempty" yaml:"transAmoumt,omitempty"`
	TransNo     string      `xml:"transNo,omitempty" json:"transNo,omitempty" yaml:"transNo,omitempty"`
	TransStatus string      `xml:"transStatus,omitempty" json:"transStatus,omitempty" yaml:"transStatus,omitempty"`
	TransTime   string      `xml:"transTime,omitempty" json:"transTime,omitempty" yaml:"transTime,omitempty"`
}

// paymentAPI implements the PaymentAPI interface.
type paymentAPI struct {
	cli *soap.Client
}

// ProcessPreApprovedPay was auto-generated from WSDL.
func (p *paymentAPI) ProcessPreApprovedPay(α *ProcessPreApprovedPay) (β *ProcessPreApprovedPayResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ProcessPreApprovedPayResponse `xml:"ProcessPreApprovedPayResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// ProcessB2CPay was auto-generated from WSDL.
func (p *paymentAPI) ProcessB2CPay(α *ProcessB2CPay) (β *ProcessB2CPayResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ProcessB2CPayResponse `xml:"processB2CPayResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// ProcessB2CCombinatedPay was auto-generated from WSDL.
func (p *paymentAPI) ProcessB2CCombinatedPay(α *ProcessB2CCombinatedPay) (β *ProcessB2CCombinatedPayResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ProcessB2CCombinatedPayResponse `xml:"ProcessB2CCombinatedPayResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// ProcessTransfer was auto-generated from WSDL.
func (p *paymentAPI) ProcessTransfer(α *ProcessTransfer) (β *ProcessTransferResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ProcessTransferResponse `xml:"ProcessTransferResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}
