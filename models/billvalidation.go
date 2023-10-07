package models

type AgentDetail struct {
	AgentId string `json:"agentId"`
}

type ReqBody struct {
	AgentDetails AgentDetail `json:"agentDetails"`
	BillDetails  BillDetails `json:"billDetails"`
}

type ResponsBody struct {
	RespCode string `json:"respCode"`
	Status   string `json:"status"`
	Response struct {
		ChId             int    `json:"chId"`
		ApprovalRefNum   string `json:"approvalRefNum"`
		ResponseCode     string `json:"responseCode"`
		ResponseReason   string `json:"responseReason"`
		ComplianceReason string `json:"complianceReason"`
	} `json:"response"`
}

type BillValidationResponse struct {
	RespCode string             `json:"respCode"`
	Status   string             `json:"status"`
	Response BillValidationData `json:"response"`
}
type BillValidationData struct {
	ChID             int    `json:"chId"`
	ApprovalRefNum   string `json:"approvalRefNum"`
	ResponseCode     string `json:"responseCode"`
	ResponseReason   string `json:"responseReason"`
	ComplianceReason string `json:"complianceReason"`
}
