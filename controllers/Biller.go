package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"BBT/config"
	"BBT/models"

	"github.com/gin-gonic/gin"
)

//IV         = "1461618689689168"(x-iv)
//passphrase = "noenonrgkgneroiw"(x-key)

func CreateABiller(c *gin.Context) {
	var req models.Biller
	data, _ := c.Get("decryptedText")
	json.Unmarshal(data.([]byte), &req)
	db := config.Database.ConnectToDB()
	defer db.Close()

	c.JSON(http.StatusOK, "Biller created Successfully.....")

}

// GetBillers is used to define a function in Golang which is used to get all the Billers from a database.
func GetBiller(c *gin.Context) {
	//GetTodos function in Golang that handles a GET request for a list of Billers.
	var billers []models.Biller
	// db := Config.ConnectToDB()
	db := config.Database.ConnectToDB()
	defer db.Close()
	//Query a database table called "Biller".
	//db.Query() func to query the table and stores the results in a variable called "row".
	//checks if an error occurred while querying the table. If an error occurred,
	//the code will print the error message to the console and then return.
	row, err := db.Query("SELECT * FROM billers")
	if err != nil {
		fmt.Fprint(c.Writer, err)
		return
	}
	//Iterates over a collection of rows from a SQL query and stores each row into the "biller" variable which is of type Models.Biller.
	//It does this by scanning each row and assigning the values to the ID, Title, and Description fields of the biller variable.
	// If an error is encountered, the error is printed to the writer.
	for row.Next() {
		var biller models.Biller
		// if err := row.Scan(&biller.ID, &biller.BillerID, &biller.BillerName, &biller.BillerCategory, &biller.BillerChannel, &biller.BillerSubCategoryName, &biller.BillerCustomerInfo, &biller.BillerMinAmount, &biller.BillerFetchBill, &biller.BillerCategoryKey, &biller.BillerCategoryName, &biller.CoverageCity, &biller.CoverageState, &biller.CoveragePincode, &biller.BillerUpdatedDate, &biller.Status, &biller.IsAvailable); err != nil {
		if err := row.Scan(&biller.ID, &biller.BillerID, &biller.BillerName, &biller.BillerAliasName, &biller.BillerCategoryName, &biller.BillerSubCategoryName, &biller.BillerMode, &biller.BillerAcceptsAdhoc, &biller.ParentBiller, &biller.ParentBillerID, &biller.BillerCoverage, &biller.FetchRequirement, &biller.SupportBillValidation, &biller.BillerEffctvFrom, &biller.BillerPymtModes, &biller.IntChngFee, &biller.Status, &biller.AdditonalInfo, &biller.PlanAdditionalInfo); err != nil {

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
func BillFetch(c *gin.Context) {
	//assign the value of the "id" parameter from the "c" object to a var called "id"."c" object is assumed to be an instance of a type that provides access to the "Params" object.
	//The "Params" object is assumed to have a method called "ByName" which takes a parameter and returns the value of the corresponding parameter from the "c" object.

	id := c.Params.ByName("id")
	var biller models.Biller
	db := config.Database.ConnectToDB()
	defer db.Close()
	err := db.QueryRow("SELECT * FROM billers where ID=?", id).Scan(&biller.ID, &biller.BillerID, &biller.BillerName, &biller.BillerAliasName, &biller.BillerCategoryName, &biller.BillerSubCategoryName, &biller.BillerMode, &biller.BillerAcceptsAdhoc, &biller.ParentBiller, &biller.ParentBillerID, &biller.BillerCoverage, &biller.FetchRequirement, &biller.SupportBillValidation, &biller.BillerEffctvFrom, &biller.BillerPymtModes, &biller.IntChngFee, &biller.Status, &biller.AdditonalInfo, &biller.PlanAdditionalInfo)
	if err != nil {
		fmt.Fprint(c.Writer, errors.New("biller data not found"))
		return
	}

	data, _ := json.Marshal(biller)
	fmt.Println(string(data))
	c.JSON(http.StatusOK, AESEncrypt(string(data), []byte(c.Request.Header.Get("x-key")), c.Request.Header.Get("x-iv")))
}

func ValidateABill(c *gin.Context) {
	var input models.Bill
	decryptedData, exists := c.Get("decryptedText")
	if !exists {
		c.AbortWithError(http.StatusBadRequest, errors.New("decrypted data not found"))
		return
	}
	json.Unmarshal(decryptedData.([]byte), &input)
	db := config.Database.ConnectToDB()
	defer db.Close()

	// Retrieve data from the database
	query := "SELECT id, billerid FROM bills WHERE id = ? AND billerid = ? "
	var result models.Bill

	err := db.QueryRow(query, input.ID, input.BillerID).Scan(&result.ID, &result.BillerID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Bill not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, AESEncrypt("bill validate created Successfully.....", []byte(c.Request.Header.Get("x-key")), c.Request.Header.Get("x-iv")))
}
