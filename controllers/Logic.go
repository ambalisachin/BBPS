package controllers

import (
	"BBT/config"
	"BBT/models"
	"database/sql"
	"fmt"
	"log"
)

type IDatabase interface {
	Create(reqs *models.FetchBillRequest) error
	GetResp(reqs *models.FetchBillRequest) (*models.FetchBillResponse, error)
	Validate(reqs *models.FetchBillRequest) bool

	CreateBillValidation(req *models.ReqBody) error
	GetRespBillValidation(req *models.ReqBody) (*models.BillValidationResponse, error)
	ValidateBillValidation(req *models.ReqBody) bool

	CreateNonPrepaidPayment(reqs *models.ReqtBody) error
	GetRespNonPrepaidPayment(reqs *models.ReqtBody) (*models.NonPrepaidResponse, error)
	ValidateNonPrepaidPayment(reqs *models.ReqtBody) bool

	// CreatePrepaidPayment(reqt *models.PrepaidBillRequest) error
	// GetRespPrepaidPayment(reqt *models.PrepaidBillRequest) (*models.PrepaidBillResponse, error)
	// ValidatePrepaidPayment(reqt *models.PrepaidBillRequest) bool
}

type DataBase struct {
	db *sql.DB
}

func (d *DataBase) Create(reqs *models.FetchBillRequest) error {

	row, err := d.db.Query("SELECT billerId FROM fetchbills where billerId=?", reqs.BillDetail.BillerID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	var billerId string
	isEmpty := true
	for row.Next() {

		isEmpty = false
		if err := row.Scan(&billerId); err != nil {
			log.Fatal(err)
			return err
		}
	}

	if isEmpty {

		insertQuery := `INSERT INTO  fetchbills (chId ,isRealTimeFetch ,mobileNo ,email ,agendId ,initiatingchannel , mobile ,geocode ,postalcode ,terminalId,billerId , consumermobileno ,uid ,accountno )VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?);`
		_, err = d.db.Exec(insertQuery,
			reqs.ChID, reqs.IsRealTimeFetch, reqs.CustDetails.MobileNo,
			reqs.CustDetails.CustomerTags[0].Value,
			reqs.AgentDetails.AgentID, reqs.AgentDetails.DeviceTags[0].Value,
			reqs.AgentDetails.DeviceTags[1].Value,
			reqs.AgentDetails.DeviceTags[2].Value, reqs.AgentDetails.DeviceTags[3].Value,
			reqs.AgentDetails.DeviceTags[4].Value, reqs.BillDetail.BillerID,
			reqs.BillDetail.CustomerParams[0].Value,
			reqs.BillDetail.CustomerParams[1].Value,
			reqs.BillDetail.CustomerParams[2].Value,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (d *DataBase) GetResp(reqs *models.FetchBillRequest) (*models.FetchBillResponse, error) {

	//*********************** this part should be removed after real database is cerated *******start************//
	// Check if there's any record in the User table

	var resp models.FetchBillResponse

	if reqs.IsRealTimeFetch {
		resp = models.FetchBillResponse{
			RespCode: "1",
			Status:   "SUCCESS",
			Response: models.RespData{
				ChID:             reqs.ChID,
				RefID:            "LZ2EAV1BVVZ0JJQU05MA3RA59IN21301822",
				ApprovalRefNum:   "12345678",
				ResponseCode:     "000",
				ResponseReason:   "Successful",
				ComplianceReason: "",
				ComplianceRespCd: "",
				BillDetails:      reqs.BillDetail.CustomerParams,
				BillerResponse: models.BillerResponse{
					CustomerName: "PRABHA",
					Amount:       "1000",
					DueDate:      "2015-06-20",
					CustConvFee:  "",
					CustConvDesc: "",
					BillDate:     "2015-06-14",
					BillNumber:   "12303",
					BillPeriod:   "june",
					BillTags:     []interface{}{},
				},
				AdditionalInfo: []interface{}{},
			},
		}
	} else {
		resp = models.FetchBillResponse{
			RespCode: "0",
			Status:   "FAILURE",
			Response: models.RespData{
				ChID:             reqs.ChID,
				RefID:            "B8O1THXV28W2JCN8JOFGASGZUNT21301740",
				ApprovalRefNum:   "",
				ResponseCode:     "002",
				ResponseReason:   "Failure",
				ComplianceReason: "BRP042,CPR014,CPR012,AIN004",
				ComplianceRespCd: "",

				BillDetails: []models.CustomerParam{},
				BillerResponse: models.BillerResponse{
					CustomerName: "",
					Amount:       "NaN",
					DueDate:      "",
					CustConvFee:  "",
					CustConvDesc: "",
					BillDate:     "",
					BillNumber:   "",
					BillPeriod:   "",
					BillTags:     []interface{}{},
				},
				AdditionalInfo: []interface{}{},
			},
		}
	}

	return &resp, nil
}

func (d *DataBase) Validate(reqs *models.FetchBillRequest) bool {

	row, err := d.db.Query("SELECT mobile,billerId,AgendId,mobileNo FROM fetchbills where billerId=?", reqs.BillDetail.BillerID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	var mobile, billerId, AgendId, mobileNo string
	for row.Next() {

		if err := row.Scan(&mobile, &billerId, &AgendId, &mobileNo); err != nil {
			log.Fatal(err)
			return false
		}

		fmt.Println(mobile, " ", reqs.AgentDetails.DeviceTags[1].Value)

		if mobile == "" || len(reqs.AgentDetails.DeviceTags[1].Value) < 10 || mobile != reqs.AgentDetails.DeviceTags[1].Value {
			return false
		}
		if billerId == "" || billerId != reqs.BillDetail.BillerID {
			return false
		}
		if AgendId == "" || AgendId != reqs.AgentDetails.AgentID {
			return false
		}
		if mobileNo == "" || mobileNo != reqs.CustDetails.MobileNo {
			return false
		}

	}
	return true
}

func (d *DataBase) CreateBillValidation(req *models.ReqBody) error {

	row, err := d.db.Query("SELECT billerId FROM billvalidations where billerId=?", req.BillDetails.BillerID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	var billerId string
	isEmpty := true
	for row.Next() {

		isEmpty = false
		if err := row.Scan(&billerId); err != nil {
			log.Fatal(err)
			return err
		}
	}
	if isEmpty {
		insertQuery := `INSERT INTO billvalidations(agentId, billerId, registeredMobileNumber) VALUES (?, ?, ?);`
		_, err = d.db.Exec(insertQuery,
			req.AgentDetails.AgentId,
			req.BillDetails.BillerID,
			req.BillDetails.CustomerParams[0].Value,
		)
		if err != nil {
			log.Fatal(err)

		}
	}
	return nil
}

func (d *DataBase) GetRespBillValidation(req *models.ReqBody) (*models.BillValidationResponse, error) {

	//*********************** this part should be removed after real database is cerated *******start************//
	// Check if there's any record in the User table

	var resp models.BillValidationResponse

	if req.BillDetails.BillerID == "" {
		resp = models.BillValidationResponse{
			RespCode: "1",
			Status:   "SUCCESS",
			Response: models.BillValidationData{
				ChID:             1,
				ApprovalRefNum:   "",
				ResponseCode:     "000",
				ResponseReason:   "Successful",
				ComplianceReason: "",
			},
		}
	} else {
		resp = models.BillValidationResponse{
			RespCode: "0",
			Status:   "FAILURE",
			Response: models.BillValidationData{
				ChID:             1,
				ApprovalRefNum:   "AB123456",
				ResponseCode:     "200",
				ResponseReason:   "FAILURE",
				ComplianceReason: "Bill response not received",
			},
		}
	}
	return &resp, nil

}

func (d *DataBase) ValidateBillValidation(req *models.ReqBody) bool {

	row, err := d.db.Query("SELECT billerId,registeredMobileNumber FROM billvalidations where billerId=?", req.BillDetails.BillerID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	var billerId, registeredMobileNumber string
	for row.Next() {

		if err := row.Scan(&billerId, &registeredMobileNumber); err != nil {
			log.Fatal(err)
			return false
		}

		fmt.Println(billerId, " ", req.BillDetails.CustomerParams[0].Value)
		fmt.Println(registeredMobileNumber, "", req.BillDetails.CustomerParams[0].Value)

		if billerId == "" || billerId != req.BillDetails.BillerID {
			return false
		}
		if registeredMobileNumber == "" || registeredMobileNumber != req.BillDetails.CustomerParams[0].Value {
			return false
		}

	}
	return true
}

func (d *DataBase) CreateNonPrepaidPayment(reqs *models.ReqtBody) error {

	row, err := d.db.Query("SELECT billerId FROM nonprepaidpayments where billerId=?", reqs.BillDetails.BillerID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	var billerId string
	isEmpty := true
	for row.Next() {

		isEmpty = false
		if err := row.Scan(&billerId); err != nil {
			log.Fatal(err)
			return err
		}
	}

	if isEmpty {
		insertQuery := `INSERT INTO nonprepaidpayments (agentId,initiatingChannel,mobile,
			geocode,postalCode,terminalId,amount,currency,custConvFee,couCustConvFee,billerId,
			consumerMobileNo,uid,accountNo,chId,email,mobileNo,remarks,paymentMode,quickPay,
			splitPay,offusPay,refId,clientRequestId) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			 ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

		_, err = d.db.Exec(insertQuery,
			reqs.AgentDetails.AgentID, reqs.AgentDetails.DeviceTags[0].Value,
			reqs.AgentDetails.DeviceTags[1].Value, reqs.AgentDetails.DeviceTags[2].Value,
			reqs.AgentDetails.DeviceTags[3].Value, reqs.AgentDetails.DeviceTags[4].Value,
			reqs.AmountDetails.Amount, reqs.AmountDetails.Currency, reqs.AmountDetails.CustConvFee,
			reqs.AmountDetails.CouCustConvFee, reqs.BillDetails.BillerID,
			reqs.BillDetails.CustomerParams[0].Value,
			reqs.BillDetails.CustomerParams[1].Value, reqs.BillDetails.CustomerParams[2].Value,
			reqs.ChId, reqs.CustDetails.CustomerTags[0].Value, reqs.CustDetails.MobileNo,
			reqs.PaymentDetails.PaymentInfo[0].Value, reqs.PaymentDetails.PaymentMode, reqs.PaymentDetails.QuickPay,
			reqs.PaymentDetails.SplitPay, reqs.PaymentDetails.OffusPay, reqs.RefId, reqs.ClientRequestId)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (d *DataBase) GetRespNonPrepaidPayment(reqs *models.ReqtBody) (*models.NonPrepaidResponse, error) {

	//*********************** this part should be removed after real database is cerated *******start************//
	// Check if there's any record in the User table

	var respons models.NonPrepaidResponse

	if reqs.AmountDetails.Amount != "" {
		respons = models.NonPrepaidResponse{
			RespCode: "1",
			Status:   "SUCCESS",
			Response: models.ResponseData{
				ChID:           reqs.ChId,
				RefID:          "LZ2EAV1BVVZ0JJQU05MA3RA59IN21301822",
				ApprovalRefNum: "12345093",
				ResponseCode:   "000",
				ResponseReason: "Successful",
				TxnDateTime:    "2022-05-10 18:22:25",
				TxnReferenceID: "AM012130EW74Y0R0D39C",
			},
		}
	} else {
		respons = models.NonPrepaidResponse{
			RespCode: "0",
			Status:   "FAILURE",
			Response: models.ResponseData{
				ChID:             reqs.ChId,
				RefID:            "8HVJDD3ZQRSBA6N67P3S8M8G5SJ21301824",
				ApprovalRefNum:   "",
				ResponseCode:     "200",
				ResponseReason:   "Failure",
				ComplianceReason: "Payment received for the billing period - no bill due",
				ComplianceRespCd: "BFR004",
				TxnDateTime:      "2022-05-10 18:25:10",
				TxnReferenceID:   "AM0121308RXD8X7O91UR",
			},
		}
	}
	return &respons, nil

}

func (d *DataBase) ValidateNonPrepaidPayment(reqs *models.ReqtBody) bool {

	row, err := d.db.Query("SELECT mobile,billerId,AgendId,mobileNo FROM fetchbills where billerId=?", reqs.BillDetails.BillerID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	var mobile, billerId, AgendId, amount string
	for row.Next() {

		if err := row.Scan(&mobile, &billerId, &AgendId, &amount); err != nil {
			log.Fatal(err)
			return false
		}

		fmt.Println(mobile, " ", reqs.AgentDetails.DeviceTags[1].Value, reqs.AmountDetails.Amount)

		if mobile == "" || len(reqs.AgentDetails.DeviceTags[1].Value) != 10 || mobile != reqs.AgentDetails.DeviceTags[1].Value {
			return false
		}
		if billerId == "" || billerId != reqs.BillDetails.BillerID {
			return false
		}
		if AgendId == "" || AgendId != reqs.AgentDetails.AgentID {
			return false
		}
		if amount == "" || amount != reqs.AmountDetails.Amount {
			return false
		}

	}
	return true
}
func NewDataBase() IDatabase {
	return &DataBase{
		db: config.Database.ConnectToDB(),
	}
}

// func (d *DataBase) CreatePrepaidPayment(reqt *models.PrepaidBillRequest) error {

// 	row, err := d.db.Query("SELECT agentId FROM prepaidpayments where agentId=?", reqt.AgentDetails.AgentID)
// 	if err != nil {
// 		log.Fatal(err)
// 		return err
// 	}
// 	var agentId string
// 	isEmpty := true
// 	for row.Next() {
// 		// var biller models.FetchBillRequest
// 		isEmpty = false
// 		if err := row.Scan(&agentId); err != nil {
// 			log.Fatal(err)
// 			return err
// 		}
// 	}

// 	if isEmpty {

// 		insertQuery := `INSERT INTO  prepaidpayments (chId ,refId ,clientRequestId ,agentId,
// 			initiatingchannel,amount,currency, custConvFee,couCustConvFee,billerId,
// 			Mobile Number,Circle,Id,EMAIL ,AADHAAR ,PAN ,mobileNo,Remarks , paymentMode ,quickPay ,
// 			splitPay,offusPay,type,id )VALUES
// 			 (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?,?,?,?,?,?,?,?,?);`
// 		_, err = d.db.Exec(insertQuery,
// 			reqt.ChID, reqt.RefID, reqt.ClientRequestID, reqt.AgentDetails.AgentID,
// 			reqt.AgentDetails.DeviceTags[0].Value, reqt.AmountDetails.Amount,
// 			reqt.AmountDetails.Currency, reqt.AmountDetails.CustConvFee,
// 			reqt.AmountDetails.CouCustConvFee, reqt.BillDetails.BillerID,
// 			reqt.BillDetails.CustomerParams[0].Value, reqt.BillDetails.CustomerParams[1].Value,
// 			reqt.BillDetails.CustomerParams[2].Value, reqt.CustDetails.CustomerTags[0].Value,
// 			reqt.CustDetails.CustomerTags[1].Value, reqt.CustDetails.CustomerTags[2].Value,
// 			reqt.CustDetails.MobileNo, reqt.PaymentDetails.PaymentInfo[0].Value,
// 			reqt.PaymentDetails.PaymentMode, reqt.PaymentDetails.QuickPay, reqt.PaymentDetails.SplitPay,
// 			reqt.PaymentDetails.OffusPay, reqt.PlanDetails.PlanDetail.Type, reqt.PlanDetails.PlanDetail.ID)

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	return nil

// }

// func (d *DataBase) GetRespPrepaidPayment(reqt *models.PrepaidBillRequest) (*models.PrepaidBillResponse, error) {

// 	var ra models.NonPrepaidResponse

// 	if reqt.AmountDetails.Amount != "" {
// 		ra := models.PrepaidBillResponse{
// 			RespCode: "0",
// 			Status:   "FAILURE",
// 			Response: models.BillResponseData{
// 				ChID:           reqt.ChID,
// 				RefID:          reqt.RefID,
// 				ApprovalRefNum: "",
// 				ResponseCode:   "200",
// 				ResponseReason: "Amount is null",
// 			},
// 		}
// 	} else {
// 		respon = models.SuccesResponse{
// 			RespCode: "0",
// 			Status:   "FAILURE",
// 			RefID:          "3W0YSK5MJCQW9QWUEUSB3RKTY0221861853",
// 					ApprovalRefNum: "AB12345001",
// 					ResponseCode:   "000",
// 					ResponseReason: "Successful",
// 					TxnDateTime:    "2022-07-04 13:02:52",
// 					TxnReferenceID: "YB412186NKGZG30XV5RE",
// 					BillerPlanResponse: models.BillerPlanResponse{
// 						PlanInfo: []models.PlanInfo{
// 							{
// 								Type: "ACTIVATED",
// 								PlanInfoTags: []models.PlanInfoTag{
// 									{Name: "Id", Value: "10"},
// 									{Name: "Plan Type", Value: "Recharge"},
// 									{Name: "Talktime", Value: "0"},
// 									{Name: "Validity", Value: "84 Days"},
// 									{Name: "Data", Value: "100 GB"},
// 									{Name: "Circle", Value: "Andhra Pradesh"},
// 									{Name: "amountInRupees", Value: "108"},
// 									{Name: "planDescription", Value: "Tarrif Calls - Local/STD/LL @ 1P/sec"},
// 								},
// 							},
// 						},
// 					},

// 				},
// 			}
// 			return &ra, nil
// 		}

// } else {
// 		ra := models.SuccesResponse{
// 			RespCode: "1",
// 			Status:   "SUCCESS",
// 			Response: models.SuccessData{
// 				ChID:           reqt.ChID,
// 				RefID:          "3W0YSK5MJCQW9QWUEUSB3RKTY0221861853",
// 				ApprovalRefNum: "AB12345001",
// 				ResponseCode:   "000",
// 				ResponseReason: "Successful",
// 				BillerPlanResponse: models.BillerPlanResponse{
// 					PlanInfo: []models.PlanInfo{
// 						{
// 							Type: "ACTIVATED",
// 							PlanInfoTags: []models.PlanInfoTag{
// 								{Name: "Id", Value: "10"},
// 								{Name: "Plan Type", Value: "Recharge"},
// 								{Name: "Talktime", Value: "0"},
// 								{Name: "Validity", Value: "84 Days"},
// 								{Name: "Data", Value: "100 GB"},
// 								{Name: "Circle", Value: "Andhra Pradesh"},
// 								{Name: "amountInRupees", Value: "108"},
// 								{Name: "planDescription", Value: "Tarrif Calls - Local/STD/LL @ 1P/sec"},
// 							},
// 						},
// 					},
// 				},
// 				TxnDateTime:    "2022-07-04 13:02:52",
// 				TxnReferenceID: "YB412186NKGZG30XV5RE",
// 			},
// 		}

// 	}
// 	return &ra, nil
// }

// func (d *DataBase) ValidatePrepaidPayment(reqs *models.PrepaidBillRequest) bool {
// 	// var request models.FetchBillRequest

// 	row, err := d.db.Query("SELECT AgendId FROM prepaidpayments where AgendId=?", reqs.AgentDetails.AgentID)
// 	if err != nil {
// 		log.Fatal(err)
// 		return false
// 	}
// 	var AgendId string
// 	for row.Next() {
// 		// var biller models.FetchBillRequest

// 		if err := row.Scan(&AgendId); err != nil {
// 			log.Fatal(err)
// 			return false
// 		}

// 		fmt.Println(AgendId, " ", reqs.AgentDetails.AgentID)

// 		if AgendId == "" || AgendId != reqs.AgentDetails.AgentID {
// 			return false
// 		}

// 	}
// 	return true
// }

// 	var ra models.PrepaidBillResponse
// 	if reqt.AmountDetails.Amount == "" {
// 		ra := models.PrepaidBillResponse{
// 			RespCode: "0",
// 			Status:   "FAILURE",
// 			Response: models.BillResponseData{
// 				ChID:           reqt.ChID,
// 				RefID:          reqt.RefID,
// 				ApprovalRefNum: "",
// 				ResponseCode:   "200",
// 				ResponseReason: "Amount is null",
// 			},
// 		}

// 	} else {
// 		ra := models.SuccesResponse{
// 			RespCode: "1",
// 			Status:   "SUCCESS",
// 			Response: models.SuccessData{
// 				ChID:           reqt.ChID,
// 				RefID:          "3W0YSK5MJCQW9QWUEUSB3RKTY0221861853",
// 				ApprovalRefNum: "AB12345001",
// 				ResponseCode:   "000",
// 				ResponseReason: "Successful",
// 				BillerPlanResponse: models.BillerPlanResponse{
// 					PlanInfo: []models.PlanInfo{
// 						{
// 							Type: "ACTIVATED",
// 							PlanInfoTags: []models.PlanInfoTag{
// 								{Name: "Id", Value: "10"},
// 								{Name: "Plan Type", Value: "Recharge"},
// 								{Name: "Talktime", Value: "0"},
// 								{Name: "Validity", Value: "84 Days"},
// 								{Name: "Data", Value: "100 GB"},
// 								{Name: "Circle", Value: "Andhra Pradesh"},
// 								{Name: "amountInRupees", Value: "108"},
// 								{Name: "planDescription", Value: "Tarrif Calls - Local/STD/LL @ 1P/sec"},
// 							},
// 						},
// 					},
// 				},
// 				TxnDateTime:    "2022-07-04 13:02:52",
// 				TxnReferenceID: "YB412186NKGZG30XV5RE",
// 			},
// 		}

// 	}
// 	return &ra, nil
// }
