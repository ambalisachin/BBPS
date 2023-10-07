package controllers

// import (
// 	"BBT/models"
// 	"database/sql"
// 	"fmt"
// 	"log"
// )

// type IDataBase interface {
// 	CreatePrepaid(reqt *models.PrepaidBillRequest) error
// 	GetRespPrepaid(reqt *models.PrepaidBillRequest) (*models.PrepaidBillResponse, error)
// 	ValidatePrepaid(reqt *models.PrepaidBillRequest) bool
// }
// type Database struct {
// 	db *sql.DB
// }

// func (d *Database) CreatePrepaid(reqt *models.PrepaidBillRequest) error {

// 	row, err := d.db.Query("SELECT agentId FROM nonprepaidpayments where agentId=?", reqt.AgentDetails.AgentID)
// 	if err != nil {
// 		log.Fatal(err)
// 		return err
// 	}
// 	var agentId string
// 	isEmpty := true
// 	for row.Next() {

// 		isEmpty = false
// 		if err := row.Scan(&agentId); err != nil {
// 			log.Fatal(err)
// 			return err
// 		}
// 	}

// 	if isEmpty {

// 		insertQuery := `INSERT INTO  prepaidpayments (chId ,refId ,clientRequestId ,agentId,
// 		initiatingchannel,amount,currency, custConvFee,couCustConvFee,billerId,
//  			Mobile Number,Circle,Id,EMAIL ,AADHAAR ,PAN ,mobileNo,Remarks , paymentMode ,quickPay ,
//  			splitPay,offusPay,type,id )VALUES
//  			 (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?,?,?,?,?,?,?,?,?);`
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

// 	var ra models.PrepaidBillResponse
// 	if reqt.AmountDetails.Amount == "" {
// 		ra := models.PrepaidBillResponse{
// 			RespCode: "0",
// 			Status:   "FAILURE",
// 			Response: models.BillResponseData{
// 				ChID:           int(reqt.ChID),
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
// 				ChID:           int(reqt.ChID),
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

// 		// c.JSON(http.StatusOK, response)SS

// 	}
// 	return &ra, nil
// }

// func (d *Database) ValidatePrepaidPayment(reqt *models.PrepaidBillRequest) bool {

// 	row, err := d.db.Query("SELECT AgendId FROM prepaidpayments where AgendId=?", reqt.AgentDetails.AgentID)
// 	if err != nil {
// 		log.Fatal(err)
// 		return false
// 	}
// 	var AgendId string
// 	for row.Next() {

// 		if err := row.Scan(&AgendId); err != nil {
// 			log.Fatal(err)
// 			return false
// 		}

// 		fmt.Println(AgendId, " ", reqt.AgentDetails.AgentID)

// 		if AgendId == "" || AgendId != reqt.AgentDetails.AgentID {
// 			return false
// 		}

// 	}
// 	return true
// }

// func NewDataBases() IDataBase {
// 	return &Database{
// 		db: config.Database.ConnectToDB(),
// 	}
// }
