package models

type PaymentDetails struct {
	PaymentInfo []PaymentInfo `json:"paymentInfo"`
	PaymentMode string        `json:"paymentMode"`
	QuickPay    string        `json:"quickPay"`
	SplitPay    string        `json:"splitPay"`
	OffusPay    string        `json:"offusPay"`
}

type PlanDetail struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type PlanDetails struct {
	PlanDetail PlanDetail `json:"planDetail"`
}

type ReqstBody struct {
	ChID            int             `json:"chId"`
	RefID           string          `json:"refId"`
	ClientRequestID string          `json:"clientRequestId"`
	AgentDetails    AgentDetails    `json:"agentDetails"`
	AmountDetails   AmountDetails   `json:"amountDetails"`
	BillDetails     BillDetails     `json:"billDetails"`
	CustDetails     CustomerDetails `json:"custDetails"`
	PaymentDetails  PaymentDetails  `json:"paymentDetails"`
	PlanDetails     PlanDetails     `json:"planDetails"`
}

type SuccesResponse struct {
	RespCode string      `json:"respCode"`
	Status   string      `json:"status"`
	Response SuccessData `json:"response"`
}

type SuccessData struct {
	ChID               int                `json:"chId"`
	RefID              string             `json:"refId"`
	ApprovalRefNum     string             `json:"approvalRefNum"`
	ResponseCode       string             `json:"responseCode"`
	ResponseReason     string             `json:"responseReason"`
	BillerPlanResponse BillerPlanResponse `json:"billerPlanResponse"`
	TxnDateTime        string             `json:"txnDateTime"`
	TxnReferenceID     string             `json:"txnReferenceId"`
}

type BillerPlanResponse struct {
	PlanInfo []PlanInfo `json:"planInfo"`
}

type PlanInfo struct {
	Type         string        `json:"type"`
	PlanInfoTags []PlanInfoTag `json:"planInfoTags"`
}

type PlanInfoTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
