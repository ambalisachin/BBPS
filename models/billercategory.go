package models

import "database/sql"

type BillerCategory struct {
	ID                     string `json:"Id"`
	BillerID               string `json:"billerId"`
	BillerName             string `json:"billerName"`
	BillerAliasName        string `json:"billerAliasName"`
	BillerCategoryName     string `json:"billerCategoryName"`
	BillerSubCategoryName  string `json:"billerSubCategoryName"`
	BillerMode             string `json:"billerMode"`
	BillerAcceptsAdhoc     string `json:"billerAcceptsAdhoc"`
	ParentBiller           string `json:"parentBiller"`
	ParentBillerID         string `json:"parentBillerId"`
	BillerCoverage         string `json:"billerCoverage"`
	FetchRequirement       string `json:"fetchRequirement"`
	PaymentAmountExactness string `json:"paymentAmountExactness"`
	SupportBillValidation  string `json:"supportBillValidation"`
	BillerEffctvFrom       string `json:"billerEffctvFrom"`
	BillerPymtModes        []struct {
		PaymentMode string `json:"paymentMode"`
		MaxLimit    string `json:"maxLimit"`
		MinLimit    string `json:"minLimit"`
	} `json:"billerPymtModes"`
	BillerPymtChnls []struct {
		PaymentChannel string         `json:"paymentChannel"`
		MaxLimit       sql.NullString `json:"maxLimit"`
		MinLimit       sql.NullString `json:"minLimit"`
	} `json:"billerPymtChnls"`
	CustomerParams []struct {
		ParamName  string `json:"paramName"`
		DataType   string `json:"dataType"`
		Optional   string `json:"optional"`
		MinLength  string `json:"minLength"`
		MaxLength  string `json:"maxLength"`
		Regex      string `json:"regex"`
		Visibility string `json:"visibility"`
	} `json:"customerParams"`
	ReponseParams []struct {
		AmtBreakupList []struct {
			AmtBreakup string `json:"amtBreakup"`
		} `json:"amtBreakupList"`
	} `json:"reponseParams"`
	AdditonalInfo       []interface{} `json:"additonalInfo"`
	IntFeeConf          []interface{} `json:"intFeeConf"`
	IntChngFee          []interface{} `json:"intChngFee"`
	Status              string        `json:"status"`
	BillerResponseType  string        `json:"billerResponseType"`
	CustomerParamGroups struct {
		Group JSONSlice `json:"group"`
	} `json:"customerParamGroups"`
	BillerPlanResponseParams struct {
		Params        []interface{} `json:"params"`
		AmountOptions []struct {
			AmtBreakupList JSONSlice `json:"amtBreakupList"`
		} `json:"amountOptions"`
	} `json:"billerPlanResponseParams"`
	AdditonalInfoPayment []interface{} `json:"additonalInfoPayment"`
	PlanAdditionalInfo   []interface{} `json:"planAdditionalInfo"`
	PlanMdmRequirement   string        `json:"planMdmRequirement"`
}
