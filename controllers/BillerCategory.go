package controllers

import (
	"BBT/config"
	"BBT/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var err error

func BillerCategory(c *gin.Context) {

	var biller models.BillerCategory
	data, _ := c.Get("decryptedText")
	json.Unmarshal(data.([]byte), &biller)
	db := config.Database.ConnectToDB()
	defer db.Close()

	insertSQL := `INSERT INTO billercategories (billerId,billerName,billerAliasName,billerCategoryName,billerSubCategoryName,billerMode,billerAcceptsAdhoc,
		parentBiller,parentBillerId,billerCoverage,fetchRequirement,paymentAmountExactness,supportBillValidation,billerEffctvFrom,status,
		billerResponseType,planMdmRequirement) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(insertSQL, biller.BillerID, biller.BillerName, biller.BillerAliasName, biller.BillerCategoryName, biller.BillerSubCategoryName, biller.BillerMode, biller.BillerAcceptsAdhoc,
		biller.ParentBiller, biller.ParentBillerID, biller.BillerCoverage, biller.FetchRequirement, biller.PaymentAmountExactness, biller.SupportBillValidation,
		biller.BillerEffctvFrom, biller.Status, biller.BillerResponseType, biller.PlanMdmRequirement)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, "BillerCategory created successfully")
}
func GetBillerCategory(c *gin.Context) {
	//GetBillerCategory function in Golang that handles a GET request for a list of Billers.
	var billers []models.BillerCategory
	// db := Config.ConnectToDB()
	db := config.Database.ConnectToDB()
	defer db.Close()
	//Query a database table called "Biller".//db.Query() func to query the table and stores the results in a variable called "row".
	//checks if an error occurred while querying the table. If an error occurred,
	//the code will print the error message to the console and then return.
	row, err := db.Query("SELECT * FROM billercategories")
	if err != nil {
		fmt.Fprint(c.Writer, err)
		return
	}
	//Iterates over a collection of rows from a SQL query and stores each row into the "biller" variable which is of type Models.Biller.
	//It does this by scanning each row and assigning the values to the ID, Title, and Description fields of the biller variable.
	// If an error is encountered, the error is printed to the writer.
	for row.Next() {
		var biller models.BillerCategory

		if err := row.Scan(&biller.ID, &biller.BillerID, &biller.BillerName, &biller.BillerAliasName,
			&biller.BillerCategoryName, &biller.BillerSubCategoryName, &biller.BillerMode,
			&biller.BillerAcceptsAdhoc, &biller.ParentBiller, &biller.ParentBillerID, &biller.BillerCoverage,
			&biller.FetchRequirement, &biller.PaymentAmountExactness, &biller.SupportBillValidation,
			&biller.BillerEffctvFrom, &biller.Status, &biller.BillerResponseType, &biller.PlanMdmRequirement); err != nil {

			fmt.Fprint(c.Writer, err)
			return
		}

		//Adds a "biller" item to the list of "billers".
		//aappend func takes 2 arguments: the list of existing billers and the new biller item that is to be added to the list.
		//Func then adds the new biller item to the end of the existing list and returns the new list.
		billers = append(billers, biller)
	}
	data, _ := json.Marshal(billers)
	fmt.Println(string(data))
	//Send an HTTP response with an array of "billers" as the body of the response,and
	//a status code of 200 (OK). func c.JSON() is used to respond with JSON and the "billers" is the data which will be sent in the response body.
	//The HTTP status code of 200 indicates that the request was successful.
	c.JSON(http.StatusOK, AESEncrypt(string(data), []byte(c.Request.Header.Get("x-key")), c.Request.Header.Get("x-iv")))
}

// BillFetch is a function that retrieves a biller item from a database
func BillerCategoryById(c *gin.Context) {
	//assign the value of the "id" parameter from the "c" object to a var called "id"."c" object is assumed to be an instance of a type that provides access to the "Params" object.
	//The "Params" object is assumed to have a method called "ByName" which takes a parameter and returns the value of the corresponding parameter from the "c" object.

	id := c.Params.ByName("id")
	var biller models.BillerCategory
	db := config.Database.ConnectToDB()
	defer db.Close()

	err := db.QueryRow("SELECT * FROM billercategories where ID=?", id).Scan(&biller.ID, &biller.BillerID, &biller.BillerName, &biller.BillerAliasName,
		&biller.BillerCategoryName, &biller.BillerSubCategoryName, &biller.BillerMode,
		&biller.BillerAcceptsAdhoc, &biller.ParentBiller, &biller.ParentBillerID, &biller.BillerCoverage,
		&biller.FetchRequirement, &biller.PaymentAmountExactness, &biller.SupportBillValidation,
		&biller.BillerEffctvFrom, &biller.Status, &biller.BillerResponseType, &biller.PlanMdmRequirement)

	if err != nil {
		fmt.Fprint(c.Writer, errors.New("biller data not found"))
		return
	}

	data, _ := json.Marshal(biller)
	fmt.Println(string(data))
	c.JSON(http.StatusOK, AESEncrypt(string(data), []byte(c.Request.Header.Get("x-key")), c.Request.Header.Get("x-iv")))
}
