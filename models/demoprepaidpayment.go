package models

//Define structs for request and response
type PrepaidBillRequest struct {
	ChID            int            `json:"chId"`
	RefID           string         `json:"refId"`
	ClientRequestID string         `json:"clientRequestId"`
	AgentDetails    AgentDetails   `json:"agentDetails"`
	AmountDetails   AmountDetails  `json:"amountDetails"`
	BillDetails     BillDetails    `json:"billDetails"`
	CustDetails     CustDetails    `json:"custDetails"`
	PaymentDetails  PaymentDetails `json:"paymentDetails"`
	PlanDetails     PlanDetails    `json:"planDetails"`
}

// type AgentDetails struct {
// 	AgentID    string      `json:"agentId"`
// 	DeviceTags []DeviceTag `json:"deviceTags"`
// }

// type DeviceTag struct {
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }

// type AmountDetails struct {
// 	Amount         string `json:"amount"`
// 	Currency       string `json:"currency"`
// 	CustConvFee    string `json:"custConvFee"`
// 	CouCustConvFee string `json:"couCustConvFee"`
// }

// type BillDetails struct {
// 	BillerID       string          `json:"billerId"`
// 	CustomerParams []CustomerParam `json:"customerParams"`
// }

// type CustomerParam struct {
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }

// type CustDetails struct {
// 	CustomerTags []CustomerTag `json:"customerTags"`
// 	MobileNo     string        `json:"mobileNo"`
// }

// type CustomerTag struct {
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }

// type PaymentDetails struct {
// 	PaymentInfo []PaymentInfo `json:"paymentInfo"`
// 	PaymentMode string        `json:"paymentMode"`
// 	QuickPay    string        `json:"quickPay"`
// 	SplitPay    string        `json:"splitPay"`
// 	OffusPay    string        `json:"offusPay"`
// }

// type PaymentInfo struct {
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }

// type PlanDetails struct {
// 	PlanDetail PlanDetail `json:"planDetail"`
// }

// type PlanDetail struct {
// 	Type string `json:"type"`
// 	ID   string `json:"id"`
// }

type PrepaidBillResponse struct {
	RespCode string           `json:"respCode"`
	Status   string           `json:"status"`
	Response BillResponseData `json:"response"`
}

type BillResponseData struct {
	ChID               int                `json:"chId"`
	RefID              string             `json:"refId"`
	ApprovalRefNum     string             `json:"approvalRefNum"`
	ResponseCode       string             `json:"responseCode"`
	ResponseReason     string             `json:"responseReason"`
	BillerPlanResponse BillerPlanResponse `json:"billerPlanResponse"`
	TxnDateTime        string             `json:"txnDateTime"`
	TxnReferenceID     string             `json:"txnReferenceId"`
}

// type BillerPlanResponse struct {
// 	PlanInfo []PlanInfo `json:"planInfo"`
// }

// type PlanInfo struct {
// 	Type         string        `json:"type"`
// 	PlanInfoTags []PlanInfoTag `json:"planInfoTags"`
// }

// type PlanInfoTag struct {
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }
