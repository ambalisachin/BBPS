package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Credentials struct can be used to store credentials in a single data type.
type Credentials struct {
	Username string
	Password string
	Server   string
	Dbname   string
}

// Database creates a variable called Database that is set to a Credentials struct.
var Database = Credentials{
	Username: "root",
	Password: "",
	Server:   "tcp(localhost:3306)",
	Dbname:   "bbtt",
}

// ConnectToDB connects to the database
func (m Credentials) ConnectToDB() *sql.DB {
	dataSourceName := m.Username + ":" + m.Password + "@" + m.Server + "/" + m.Dbname
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour * 1)
	fmt.Println("Connected to DB Successfully....... ")
	return db
}

// NewTable creates new table if the table not exist
func NewTable() {
	db := Database.ConnectToDB()
	defer db.Close()
	//checking for create table for user in db exist or not , if not in db
	// crate a table for user in db
	_, e := db.Query("CREATE TABLE IF NOT EXISTS users(Name varchar(20) UNIQUE NOT NULL, Username varchar(20) NOT NULL, Email varchar(20) NOT NULL, Password varchar(20) NOT NULL)")
	if e != nil {
		fmt.Println(e)
	}

	// Create the MySQL table if it billers doesn't exist
	_, er := db.Query("CREATE TABLE IF NOT EXISTS billers(ID int(5) UNIQUE NOT NULL AUTO_INCREMENT,billerID varchar(20)  NOT NULL,billerName varchar(20) NOT NULL, BillerAliasName varchar(20) NOT NULL,BillerSubCategoryName varchar(20)  NOT NULL,BillerMode varchar(20)  NOT NULL, BillerAcceptsAdhoc varchar(20)  NOT NULL, ParentBiller varchar(20)  NOT NULL, ParentBillerID varchar(20)  NOT NULL, BillerCoverage varchar(20)  NOT NULL, FetchRequirement varchar(20)  NOT NULL, SupportBillValidation varchar(20)  NOT NULL, BillerEffctvFrom varchar(20)  NOT NULL, BillerPymtModes varchar(20)  NOT NULL, IntChngFee varchar(20)  NOT NULL, Status varchar(20)  NOT NULL, AdditonalInfo varchar(20)  NOT NULL, PlanAdditionalInfo varchar(20)  NOT NULL)")

	if er != nil {
		fmt.Println(er)
	}

	// Create the MySQL table if it billfetches doesn't exist
	_, e = db.Exec(`CREATE TABLE IF NOT EXISTS billfetches (ID int(5) UNIQUE NOT NULL AUTO_INCREMENT,chId INT,isRealTimeFetch BOOL,mobileNo VARCHAR(255),name VARCHAR(255),agentId VARCHAR(255),geoCode VARCHAR(255),mobileNo2 VARCHAR(255),postalCode VARCHAR(255),terminalId VARCHAR(255),billerId VARCHAR(255),customerMobileNo VARCHAR(255),accountNo VARCHAR(255))`)
	if e != nil {
		log.Fatal(e)
	}

	// Create the MySQL table if it nonprepaidpayments doesn't exist
	createTableQuery := `CREATE TABLE IF NOT EXISTS nonprepaidpayments (ID int(5) UNIQUE NOT NULL AUTO_INCREMENT,agentId VARCHAR(255),initiatingChannel VARCHAR(255),mobile VARCHAR(255),geocode VARCHAR(255),postalCode VARCHAR(255),terminalId VARCHAR(255),amount VARCHAR(255),currency VARCHAR(255),custConvFee VARCHAR(255),couCustConvFee VARCHAR(255),billerId VARCHAR(255),consumerMobileNo VARCHAR(255),uid VARCHAR(255),accountNo VARCHAR(255),chId INT,email VARCHAR(255),mobileNo VARCHAR(255),remarks VARCHAR(255),paymentMode VARCHAR(255),quickPay VARCHAR(255),splitPay VARCHAR(255),offusPay VARCHAR(255),refId VARCHAR(255),clientRequestId VARCHAR(255));`
	_, e = db.Exec(createTableQuery)
	if e != nil {
		log.Fatal(e)
	}

	// Create the MySQL table if it prepaidpayments doesn't exist
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS prepaidpayments (chId INT,refId VARCHAR(255),clientRequestId VARCHAR(255),agentId VARCHAR(255),initiatingChannel VARCHAR(255),amount VARCHAR(255),currency VARCHAR(255),custConvFee VARCHAR(255),couCustConvFee VARCHAR(255),billerId VARCHAR(255),
	mobileNumber VARCHAR(255),circle VARCHAR(255),id VARCHAR(255),email VARCHAR(255),aadhaar VARCHAR(255),pan VARCHAR(255),mobileNo VARCHAR(255),remarks VARCHAR(255),paymentMode VARCHAR(255),quickPay VARCHAR(255),splitPay VARCHAR(255),offusPay VARCHAR(255),planType VARCHAR(255),planId VARCHAR(255))`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the MySQL table if it fetchbills doesn't exist
	CreatTableQuery := `CREATE TABLE IF NOT EXISTS fetchbills (ID int(5) UNIQUE NOT NULL AUTO_INCREMENT,chId INT,isRealTimeFetch BOOL,mobileNo VARCHAR(255),email VARCHAR(255),agendId VARCHAR(255),initiatingchannel VARCHAR(255), mobile VARCHAR(255),geocode VARCHAR(255),postalcode VARCHAR(255),terminalId VARCHAR(255),billerId VARCHAR(255), consumermobileno VARCHAR(255),uid VARCHAR(255),accountno VARCHAR(255));`

	_, e = db.Exec(CreatTableQuery)
	if e != nil {
		log.Fatal(e)
	}

	// Create the MySQL table if it billvalidations doesn't exist
	const createTableSQL = `
    CREATE TABLE IF NOT EXISTS billvalidations (id INT AUTO_INCREMENT PRIMARY KEY, agentId VARCHAR(255),billerId VARCHAR(255),registeredMobileNumber VARCHAR(255));`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createtableSQL := `
	CREATE TABLE IF NOT EXISTS billercategories (
		ID INT AUTO_INCREMENT PRIMARY KEY,
		billerId VARCHAR(255) ,
		billerName VARCHAR(255),
		billerAliasName VARCHAR(255),
		billerCategoryName VARCHAR(255),
		billerSubCategoryName VARCHAR(255),
		billerMode VARCHAR(255),
		billerAcceptsAdhoc VARCHAR(255),
		parentBiller VARCHAR(255),
		parentBillerId VARCHAR(255),
		billerCoverage VARCHAR(255),
		fetchRequirement VARCHAR(255),
		paymentAmountExactness VARCHAR(255),
		supportBillValidation VARCHAR(255),
		billerEffctvFrom DATE,
		status VARCHAR(255),
		billerResponseType VARCHAR(255),
		planMdmRequirement VARCHAR(255)
	) `

	// Execute the SQL table creation query
	_, err = db.Exec(createtableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
