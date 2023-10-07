package controllers

// type IHandler interface {
// 	Prepaid(c *gin.Context)
// }

// type HandlerDB struct {
// 	db IDataBase
// }

// func (d *HandlerDB) Prepaid(c *gin.Context) {
// 	var reqs models.PrepaidBillRequest
// 	data, _ := c.Get("decryptedText")
// 	json.Unmarshal(data.([]byte), &reqs)

// db := NewDataBase()
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

// func NewHandlerDB() *HandlerDB {
// 	return &HandlerDB{
// 		db: NewDatBases(),
// 	}
// 	return nil
// }
