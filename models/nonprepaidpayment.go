package models

type AmountDetails struct {
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	CustConvFee    string `json:"custConvFee"`
	CouCustConvFee string `json:"couCustConvFee"`
}

type CustDetails struct {
	CustomerTags []CustomerTag `json:"customerTags"`
	MobileNo     string        `json:"mobileNo"`
}

type PaymentInfo struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ReqtBody struct {
	AgentDetails    AgentDetails   `json:"agentDetails"`
	AmountDetails   AmountDetails  `json:"amountDetails"`
	BillDetails     BillDetails    `json:"billDetails"`
	ChId            int            `json:"chId"`
	CustDetails     CustDetails    `json:"custDetails"`
	PaymentDetails  PaymentDetails `json:"paymentDetails"`
	RefId           string         `json:"refId"`
	ClientRequestId string         `json:"clientRequestId"`
}

type ResponseData struct {
	ChID             int    `json:"chId"`
	RefID            string `json:"refId"`
	ApprovalRefNum   string `json:"approvalRefNum"`
	ResponseCode     string `json:"responseCode"`
	ResponseReason   string `json:"responseReason"`
	ComplianceReason string `json:"complianceReason"`
	ComplianceRespCd string `json:"complianceRespCd"`
	TxnDateTime      string `json:"txnDateTime"`
	TxnReferenceID   string `json:"txnReferenceId"`
}

type NonPrepaidResponse struct {
	RespCode string       `json:"respCode"`
	Status   string       `json:"status"`
	Response ResponseData `json:"response"`
}
