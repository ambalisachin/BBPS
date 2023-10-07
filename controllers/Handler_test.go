package controllers

import (
	"BBT/models"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var Request = models.FetchBillRequest{
	ChID:            1,
	IsRealTimeFetch: true,
	CustDetails: models.CustDetails{
		MobileNo: "9004398093",
		CustomerTags: []models.CustomerTag{
			{
				Name:  "EMAIL",
				Value: "mk.chekuri@gmail.com",
			},
		},
	},
	AgentDetails: models.AgentDetails{AgentID: "AM01AM11BNK519046222",
		DeviceTags: []models.DeviceTag{
			{
				Name:  "INITIATING_CHANNEL",
				Value: "BSC",
			},
			{
				Name:  "MOBILE",
				Value: "7878787123",
			},
			{
				Name:  "GEOCODE",
				Value: "28.6139,78.5555",
			},
			{
				Name:  "POSTAL_CODE",
				Value: "600001",
			},
			{
				Name:  "TERMINAL_ID",
				Value: "3451234560",
			},
		}}, BillDetail: models.BillDetail{
		BillerID: "BESCOM000KAR01",
		CustomerParams: []models.CustomerParam{{
			Name:  "Consumer Mobile No",
			Value: "7021398105",
		},
			{
				Name:  "Account No",
				Value: "8818908000",
			},
		},
	},
}

var response_byte = []byte(`{
    "chId": 1,
    "isRealTimeFetch": true,
    "custDetails": {
        "mobileNo": "9004398093",
        "customerTags": [
            {
                "name": "EMAIL",
                "value": "mk.chekuri@gmail.com"
            }
        ]
    },
    "agentDetails": {
        "agentId": "AM01AM11BNK519046222",
        "deviceTags": [
            {
                "name": "INITIATING_CHANNEL",
                "value": "BSC"
            },
            {
                "name": "MOBILE",
                "value": "7878787123"
            },
            {
                "name": "GEOCODE",
                "value": "28.6139,78.5555"
            },
            {
                "name": "POSTAL_CODE",
                "value": "600001"
            },
            {
                "name": "TERMINAL_ID",
                "value": "3451234560"
            }
        ]
    },
    "billDetails": {
        "billerId": "BESCOM000KAR01",
        "customerParams": [
            {
                "name": "Consumer Mobile No",
                "value": "7021398105"
            },
            {
                "name": "Account No",
                "value": "8818908000"
            }
        ]
    }
}`)

type MockDatabase struct{}

func (m *MockDatabase) Validate(req *models.FetchBillRequest) bool {

	return true
}

func (m *MockDatabase) Create(req *models.FetchBillRequest) error {

	return nil
}

func (m *MockDatabase) GetResp(reqs *models.FetchBillRequest) (*models.FetchBillResponse, error) {

	var resp models.FetchBillResponse

	if reqs.IsRealTimeFetch {
		resp = models.FetchBillResponse{
			RespCode: "1",
			Status:   "SUCCESS",
			Response: models.RespData{
				ChID:             1,
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

func NewMockDb() IDatabase {
	return &MockDatabase{}
}

func NewMockHandler() *HandlerDb {
	return &HandlerDb{
		db: NewMockDb(),
	}
}

func TestFetchBill(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()

	// Mock the decrypted text and database:
	router.Use(func(c *gin.Context) {
		c.Set("decryptedText", []byte(`{
			"chId": 1,
			"isRealTimeFetch": true,
			"custDetails": {
				"mobileNo": "9004398093",
				"customerTags": [{"name": "EMAIL", "value": "mk.chekuri@gmail.com"}]
			},
			"agentDetails": {
				"agentId": "AM01AM11BNK519046222",
				"deviceTags": [
					{"name": "INITIATING_CHANNEL", "value": "BSC"},
					{"name": "MOBILE", "value": "7878787123"},
					{"name": "GEOCODE", "value": "28.6139,78.5555"},
					{"name": "POSTAL_CODE", "value": "600001"},
					{"name": "TERMINAL_ID", "value": "3451234560"}
				]
			},
			"billDetails": {
				"billerId": "BESCOM000KAR01",
				"customerParams": [
					{"name": "Consumer Mobile No", "value": "7021398105"},
					{"name": "Account No", "value": "8818908000"}
				]
			}
		}`))
		c.Next()
	})

	router.POST("/fetch-bill", NewMockHandler().FetchBill)

	// Mocking a request
	body := bytes.NewBuffer([]byte(`{"key":"value"}`)) // your request body here
	req, err := http.NewRequest(http.MethodPost, "/fetch-bill", body)
	if err != nil {
		t.Fatalf("Failed to make mock request: %v", err)
	}

	// Recording the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check the response status code and body
	if resp.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but got %d", http.StatusOK, resp.Code)
	}

	if resp.Body == nil {
		t.Fatalf("Expected to get response %d but got %d", response_byte, resp.Body)
	}

}

var Req = models.ReqBody{
	AgentDetails: models.AgentDetail{AgentId: "AM01YB41BSC519046456"}, BillDetails: models.BillDetails{
		BillerID: "OUSH40000NAT02",
		CustomerParams: []models.CustomerParam{{
			Name:  "Registered Mobile Number / Viewing Card Number",
			Value: "9987654321",
		},
		},
	},
}

var Resp_byte = []byte(`
	{
		"agentDetails": {
			"agentId": "AM01YB41BSC519046456"
		},
		"billDetails": {
			"billerId": "OUSH40000NAT02",
			"customerParams": [
				{
					"name": "Registered Mobile Number / Viewing Card Number",
					"value": "9987654321"
				}
			]
		}
	}`)

func (m *MockDatabase) ValidateBillValidation(req *models.ReqBody) bool {

	return true
}
func (m *MockDatabase) CreateBillValidation(req *models.ReqBody) error {

	return nil
}

func (m *MockDatabase) GetRespBillValidation(req *models.ReqBody) (*models.BillValidationResponse, error) {

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

func TestBillValidation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()

	// Mock the decrypted text and database:
	router.Use(func(c *gin.Context) {
		c.Set("decryptedText", []byte(`{
			
			"agentDetails": {
				"agentId": "AM01YB41BSC519046456",
				
			},
			"billDetails": {
				"billerId": "OUSH40000NAT02",
				"customerParams": [
					{"name": "Registered Mobile Number / Viewing Card Number", "value": "9987654321"},
				
				]
			}
		}`))
		c.Next()
	})

	router.POST("/bill-validation", NewMockHandler().BillValidation)
	// Mocking a request
	body := bytes.NewBuffer([]byte(`{"key":"value"}`)) // your request body here
	req, err := http.NewRequest(http.MethodPost, "/bill-validation", body)
	if err != nil {
		t.Fatalf("Failed to make mock request: %v", err)
	}

	// Recording the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check the response status code and body
	if resp.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but got %d", http.StatusOK, resp.Code)
	}

	if resp.Body == nil {
		t.Fatalf("Expected to get response %d but got %d", response_byte, resp.Body)
	}

}

var Requests = models.ReqtBody{
	AgentDetails: models.AgentDetails{AgentID: "AM01AM11BNK519046222",
		DeviceTags: []models.DeviceTag{
			{
				Name:  "INITIATING_CHANNEL",
				Value: "BSC",
			},
			{
				Name:  "MOBILE",
				Value: "7878787123",
			},
			{
				Name:  "GEOCODE",
				Value: "28.6139,78.5555",
			},
			{
				Name:  "POSTAL_CODE",
				Value: "600001",
			},
			{
				Name:  "TERMINAL_ID",
				Value: "3451234560",
			},
		}},

	AmountDetails: models.AmountDetails{

		Amount:         "1000",
		Currency:       "356",
		CustConvFee:    "0",
		CouCustConvFee: "0",
	},

	BillDetails: models.BillDetails{
		BillerID: "OUA500000PUN01",
		CustomerParams: []models.CustomerParam{{
			Name:  "Consumer Mobile No",
			Value: "9865778954",
		},
			{
				Name:  "UID",
				Value: "8596547893214111",
			},
			{
				Name:  "Account No",
				Value: "258931",
			},
		},
	},
	ChId: 1,
	CustDetails: models.CustDetails{
		MobileNo: "9004398093",
		CustomerTags: []models.CustomerTag{
			{
				Name:  "EMAIL",
				Value: "mk.chekuri@gmail.com",
			},
		},
	},
	PaymentDetails: models.PaymentDetails{
		PaymentInfo: []models.PaymentInfo{
			{
				Name:  "Remarks",
				Value: "Payment",
			},
		},
		PaymentMode: "Cash",
		QuickPay:    "No",
		SplitPay:    "No",
		OffusPay:    "Yes",
	},
	RefId:           "LZ2EAV1BVVZ0JJQU05MA3RA59IN21301822",
	ClientRequestId: "NUITUFLQRTOU",
}

var Response_byte = []byte(`{
    "agentDetails": {
        "agentId": "AM01AM11BNK519046222",
        "deviceTags": [
            {
                "name": "INITIATING_CHANNEL",
                "value": "BSC"
            },
            {
                "name": "MOBILE",
                "value": "7878787123"
            },
            {
                "name": "GEOCODE",
                "value": "28.6139,78.5555"
            },
            {
                "name": "POSTAL_CODE",
                "value": "600001"
            },
            {
                "name": "TERMINAL_ID",
                "value": "3451234560"
            }
        ]
    },
	"amountDetails":{
		"amount":"1000",
		"currency":"356",
		"custConvFee":"0",
		"couCustConvFee":"0",
	} ,
    
    
    "billDetails": {
        "billerId": "OUA500000PUN01",
        "customerParams": [
            {
                "name": "Consumer Mobile No",
                "value": "9865778954"
            },
			{
				Name:  "UID",
				Value: "8596547893214111",
			},
            {
                "name": "Account No",
                "value": "258931"
            }
        ]
    },
	"custDetails":{
		"customerTags":[
			{
				"name": "Email",
                "value": "manoj.chekuri@npci.org.in",
			}
		]
	},
	"mobileNo":"9876596093",
	"paymentDetails":{
		"paymentInfo":[
			{
				"name": "Remarks",
                "value": "Payment",
			}
		]
	
	"paymentMode":"Cash",
	"quickPay":"No",
	"splitPay":"No",
	"offusPay":"Yes",
		},
		"refId":"LZ2EAV1BVVZ0JJQU05MA3RA59IN21301822",
		"clientRequestId":"NUITUFLQRTOU",
	},
}`)

func (m *MockDatabase) ValidateNonPrepaidPayment(reqs *models.ReqtBody) bool {

	return true
}

func (m *MockDatabase) CreateNonPrepaidPayment(reqs *models.ReqtBody) error {

	return nil
}

func (m *MockDatabase) GetRespNonPrepaidPayment(reqs *models.ReqtBody) (*models.NonPrepaidResponse, error) {

	var resp models.NonPrepaidResponse

	if reqs.AmountDetails.Amount != "" {
		resp = models.NonPrepaidResponse{
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
		resp = models.NonPrepaidResponse{
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

	return &resp, nil
}

func TestNonPrepaid(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()

	// Mock the decrypted text and database:
	router.Use(func(c *gin.Context) {
		c.Set("decryptedText", []byte(`{
			"agentDetails": {
				"agentId": "AM01AM11BNK519046222",
				"deviceTags": [
					{
						"name": "INITIATING_CHANNEL",
						"value": "BSC"
					},
					{
						"name": "MOBILE",
						"value": "7878787123"
					},
					{
						"name": "GEOCODE",
						"value": "28.6139,78.5555"
					},
					{
						"name": "POSTAL_CODE",
						"value": "600001"
					},
					{
						"name": "TERMINAL_ID",
						"value": "3451234560"
					}
				]
			},
			"amountDetails":{
				"amount":"1000",
				"currency":"356",
				"custConvFee":"0",
				"couCustConvFee":"0",
			} ,
			
			
			"billDetails": {
				"billerId": "OUA500000PUN01",
				"customerParams": [
					{
						"name": "Consumer Mobile No",
						"value": "9865778954"
					},
					{
						Name:  "UID",
						Value: "8596547893214111",
					},
					{
						"name": "Account No",
						"value": "258931"
					}
				]
			},
			"custDetails":{
				"customerTags":[
					{
						"name": "Email",
						"value": "manoj.chekuri@npci.org.in",
					}
				]
			},
			"mobileNo":"9876596093",
			"paymentDetails":{
				"paymentInfo":[
					{
						"name": "Remarks",
						"value": "Payment",
					}
				]
			
			"paymentMode":"Cash",
			"quickPay":"No",
			"splitPay":"No",
			"offusPay":"Yes",
				},
				"refId":"LZ2EAV1BVVZ0JJQU05MA3RA59IN21301822",
				"clientRequestId":"NUITUFLQRTOU",
			}
		}`))
		c.Next()
	})

	router.POST("/fetch-nonprepaidbill", NewMockHandler().NonPrepaid)

	// Mocking a request
	body := bytes.NewBuffer([]byte(`{"key":"value"}`)) // your request body here
	req, err := http.NewRequest(http.MethodPost, "/fetch-nonprepaidbill", body)
	if err != nil {
		t.Fatalf("Failed to make mock request: %v", err)
	}

	// Recording the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check the response status code and body
	if resp.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but got %d", http.StatusOK, resp.Code)
	}

	if resp.Body == nil {
		t.Fatalf("Expected to get response %d but got %d", response_byte, resp.Body)
	}

}
