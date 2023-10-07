package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"BBT/models"

	"github.com/gin-gonic/gin"
)

type IHandlers interface {
	FetchBill(c *gin.Context)
	BillValidation(c *gin.Context)
	NonPrepaid(c *gin.Context)

	// Prepaid(c *gin.Context)
}

type HandlerDb struct {
	db IDatabase
}

func (d *HandlerDb) FetchBill(c *gin.Context) {
	var reqs models.FetchBillRequest
	data, _ := c.Get("decryptedText")
	json.Unmarshal(data.([]byte), &reqs)

	// db := NewDataBase()
	sa := d.db.Validate(&reqs)
	fmt.Print(sa)
	if !d.db.Validate(&reqs) {

		c.JSON(http.StatusBadRequest, "invalid data")
	} else {
		d.db.Create(&reqs)

		resp, err := d.db.GetResp(&reqs)
		if err != nil {
			log.Fatal(err)
			return
		}
		c.JSON(http.StatusOK, resp)

	}
}

func (d *HandlerDb) BillValidation(c *gin.Context) {
	var req models.ReqBody
	data, _ := c.Get("decryptedText")
	json.Unmarshal(data.([]byte), &req)

	// db := NewDataBase()
	sa := d.db.ValidateBillValidation(&req)
	fmt.Print(sa)
	if !d.db.ValidateBillValidation(&req) {

		c.JSON(http.StatusBadRequest, "invalid data")
	} else {
		d.db.CreateBillValidation(&req)

		resp, err := d.db.GetRespBillValidation(&req)
		if err != nil {
			log.Fatal(err)
			return
		}
		c.JSON(http.StatusOK, resp)

	}
}

func (d *HandlerDb) NonPrepaid(c *gin.Context) {
	var reqs models.ReqtBody
	data, _ := c.Get("decryptedText")
	json.Unmarshal(data.([]byte), &reqs)

	// db := NewDataBase()
	sa := d.db.ValidateNonPrepaidPayment(&reqs)
	fmt.Print(sa)
	if !d.db.ValidateNonPrepaidPayment(&reqs) {

		c.JSON(http.StatusBadRequest, "invalid data")
	} else {
		d.db.CreateNonPrepaidPayment(&reqs)

		resp, err := d.db.GetRespNonPrepaidPayment(&reqs)
		if err != nil {
			log.Fatal(err)
			return
		}
		c.JSON(http.StatusOK, resp)

	}
}
func NewHandlerDb() *HandlerDb {
	return &HandlerDb{
		db: NewDataBase(),
	}
}

// func (d *HandlerDb) Prepaid(c *gin.Context) {
// 	var reqs models.PrepaidBillRequest
// 	data, _ := c.Get("decryptedText")
// 	json.Unmarshal(data.([]byte), &reqs)

// 	// db := NewDataBase()
// 	sa := d.db.ValidatePrepaid(&reqs)
// 	fmt.Print(sa)
// 	if !d.db.ValidatePrepaid(&reqs) {

// 		c.JSON(http.StatusBadRequest, "invalid data")
// 	} else {
// 		d.db.CreatePrepaid(&reqs)

// 		resp, err := d.db.GetRespPrepaid(&reqs)
// 		if err != nil {
// 			log.Fatal(err)
// 			return
// 		}
// 		c.JSON(http.StatusOK, resp)

// 	}
// }
