package models

type CustomerTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CustomerDetails struct {
	MobileNo     string        `json:"mobileNo"`
	CustomerTags []CustomerTag `json:"customerTags"`
}

type DeviceTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CustomerParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type BillDetails struct {
	BillerID       string          `json:"billerId"`
	CustomerParams []CustomerParam `json:"customerParams"`
}

type RequestBody struct {
	ChID            int             `json:"chId"`
	Amount          int             `json:"amount"`
	IsRealTimeFetch bool            `json:"isRealTimeFetch"`
	CustDetails     CustomerDetails `json:"custDetails"`
	AgentDetails    AgentDetails    `json:"agentDetails"`
	BillDetails     BillDetails     `json:"billDetails"`
}

type ResponseBody struct {
	RespCode string   `json:"respCode"`
	Status   string   `json:"status"`
	Response Response `json:"response"`
}
type AgentDetails struct {
	AgentID    string      `json:"agentId"`
	DeviceTags []DeviceTag `json:"deviceTags"`
}
type BillerResponse struct {
	CustomerName string        `json:"customerName"`
	Amount       string        `json:"amount"`
	DueDate      string        `json:"dueDate"`
	CustConvFee  string        `json:"custConvFee"`
	CustConvDesc string        `json:"custConvDesc"`
	BillDate     string        `json:"billDate"`
	BillNumber   string        `json:"billNumber"`
	BillPeriod   string        `json:"billPeriod"`
	BillTags     []interface{} `json:"billTags"`
}

type Response struct {
	ChID             int             `json:"chId"`
	RefID            string          `json:"refId"`
	ApprovalRefNum   string          `json:"approvalRefNum"`
	ResponseCode     string          `json:"responseCode"`
	ResponseReason   string          `json:"responseReason"`
	ComplianceReason string          `json:"complianceReason"`
	ComplianceRespCd string          `json:"complianceRespCd"`
	BillDetails      []CustomerParam `json:"billDetails"`
	BillerResponse   BillerResponse  `json:"billerResponse"`
	AdditionalInfo   []interface{}   `json:"additionalInfo"`
}
