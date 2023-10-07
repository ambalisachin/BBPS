package models

import "encoding/json"

type Bill struct {
	ID       int `json:"id"`
	BillerID int `json:"billerid"`
}

type CustomarParam struct {
	ParamName  string `json:"paramName"`
	DataType   string `json:"dataType"`
	Optional   string `json:"optional"`
	MinLength  string `json:"minLength"`
	MaxLength  string `json:"maxLength"`
	Regex      string `json:"regex"`
	Visibility string `json:"visibility"`
}

type AmtBreakup struct {
	AmtBreakup string `json:"amtBreakup"`
}

type ResponseParam struct {
	AmtBreakupList []AmtBreakup `json:"amtBreakupList"`
}

type Biller struct {
	ID                    string `json:"Id"`
	BillerID              string `json:"billerId"`
	BillerName            string `json:"billerName"`
	BillerAliasName       string `json:"billerAliasName"`
	BillerCategoryName    string `json:"billerCategoryName"`
	BillerSubCategoryName string `json:"billerSubCategoryName"`
	BillerCustomerInfo    string `json:"billerCustomerInfo"`
	CoverageCity          string `json:"coverageCity"`
	BillerMode            string `json:"billerMode"`
	BillerUpdatedDate     string `json:"billerUpdatedDate"`
	BillerCategoryKey     string `json:"billerCategoryKey"`
	IsAvailable           string `json:"isAvailable"`
	BillerChannel         string `json:"billerChannel"`
	BillerCategory        string `json:"billerCategory"`
	// Type                     string                   `json:"Type"`
	BillerMinAmount          string                   `json:"billerMinAmount"`
	BillerFetchBill          string                   `json:"billerFetchBill"`
	BillerAcceptsAdhoc       string                   `json:"billerAcceptsAdhoc"`
	CoverageState            string                   `json:"coverageState"`
	CoveragePincode          string                   `json:"coveragePincode"`
	ParentBiller             string                   `json:"parentBiller"`
	ParentBillerID           string                   `json:"parentBillerId"`
	BillerCoverage           string                   `json:"billerCoverage"`
	FetchRequirement         string                   `json:"fetchRequirement"`
	PaymentAmountExactness   string                   `json:"paymentAmountExactness"`
	SupportBillValidation    string                   `json:"supportBillValidation"`
	BillerEffctvFrom         string                   `json:"billerEffctvFrom"`
	CustomerParams           []CustomerParam          `json:"customerParams"`
	ResponseParams           []ResponseParam          `json:"reponseParams"`
	AdditonalInfo            JSONSlice                `json:"additonalInfo"`
	IntFeeConf               JSONSlice                `json:"intFeeConf"`
	IntChngFee               JSONSlice                `json:"intChngFee"`
	BillerPymtModes          JSONSlice                `json:"billerPymtModes"`
	Status                   string                   `json:"status"`
	BillerResponseType       string                   `json:"billerResponseType"`
	CustomerParamGroups      map[string][]interface{} `json:"customerParamGroups"`
	BillerPlanResponseParams struct {
		Params        JSONSlice `json:"params"`
		AmountOptions []struct {
			AmtBreakupList JSONSlice `json:"amtBreakupList"`
		} `json:"amountOptions"`
	} `json:"billerPlanResponseParams"`
	AdditonalInfoPayment JSONSlice `json:"additonalInfoPayment"`
	PlanAdditionalInfo   JSONSlice `json:"planAdditionalInfo"`
	PlanMdmRequirement   string    `json:"planMdmRequirement"`
}

type JSONSlice []interface{}

func (js *JSONSlice) Scan(value interface{}) error {
	if value == nil {
		*js = nil
		return nil
	}
	return json.Unmarshal(value.([]byte), js)
}

type JSONMap map[string][]interface{}

func (jm *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*jm = nil
		return nil
	}
	return json.Unmarshal(value.([]byte), jm)
}
