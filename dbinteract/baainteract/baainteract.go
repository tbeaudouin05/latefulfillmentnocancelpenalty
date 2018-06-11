package baainteract

import (
	"database/sql"
	"log"

	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/row/lfncpenaltyexception"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/row/useraccess"
)

// need to write load to for each form + add to two different tables: exception + date exception at the same time each time + add unique ID

// LoadDateExceptionToDb loads purchase request to baa_application.vendor_service.lfnc_penalty_exception
func LoadDateExceptionToDb(lfncPenaltyException *lfncpenaltyexception.LfncPenaltyException, dbBaa *sql.DB) error {

	// prepare statement to insert values into baa_application.vendor_service.lfnc_penalty_exception
	insertDateExceptionStr := `
INSERT INTO baa_application.vendor_service.lfnc_penalty_exception (
	id_lfnc_penalty_exception
	,start_date 
	,end_date
	,initiator
	,amount
	,reason
	,lfnc_penalty_exception_status
	,fk_lfnc_penalty_exception_status
	,seller_name) 
VALUES (@p1,@p2,@p3,@p4,@p5,@p6,'pending',1, 'All');

INSERT INTO baa_application.vendor_service.lfnc_penalty_date_exception (
	fk_lfnc_penalty_exception
	,start_date 
	,end_date
	,fk_lfnc_penalty_exception_status) 
VALUES (@p1,@p2,@p3,1);
`
	insertDateException, err := dbBaa.Prepare(insertDateExceptionStr)

	res, err := insertDateException.Exec(
		lfncPenaltyException.IDLfncPenaltyException,
		lfncPenaltyException.StartDate,
		lfncPenaltyException.EndDate,
		lfncPenaltyException.Initiator,
		lfncPenaltyException.Amount,
		lfncPenaltyException.Reason,
	)

	log.Println(res)

	return err
}

// LoadSellerDateExceptionToDb loads purchase request to baa_application.vendor_service.lfnc_penalty_exception
func LoadSellerDateExceptionToDb(lfncPenaltyException *lfncpenaltyexception.LfncPenaltyException, dbBaa *sql.DB) error {

	// prepare statement to insert values into baa_application.vendor_service.lfnc_penalty_exception
	insertSellerDateExceptionStr := `
INSERT INTO baa_application.vendor_service.lfnc_penalty_exception (
	id_lfnc_penalty_exception
	,seller_name
	,id_seller
	,start_date 
	,end_date
	,initiator
	,amount
	,reason
	,lfnc_penalty_exception_status
	,fk_lfnc_penalty_exception_status) 
VALUES (@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8,'pending',1);

INSERT INTO baa_application.vendor_service.lfnc_penalty_seller_date_exception (
	fk_lfnc_penalty_exception
	,seller_name
	,id_seller
	,start_date 
	,end_date
	,fk_lfnc_penalty_exception_status) 
VALUES (@p1,@p2,@p3,@p4,@p5,1);
`
	insertSellerDateException, err := dbBaa.Prepare(insertSellerDateExceptionStr)

	res, err := insertSellerDateException.Exec(
		lfncPenaltyException.IDLfncPenaltyException,
		lfncPenaltyException.SellerName,
		lfncPenaltyException.IDSeller,
		lfncPenaltyException.StartDate,
		lfncPenaltyException.EndDate,
		lfncPenaltyException.Initiator,
		lfncPenaltyException.Amount,
		lfncPenaltyException.Reason,
	)

	log.Println(res)

	return err
}

func LoadSoiExceptionToDb(lfncPenaltyException *lfncpenaltyexception.LfncPenaltyException, dbBaa *sql.DB) error {

	// prepare statement to insert values into baa_application.vendor_service.lfnc_penalty_soi_exception
	insertSoiExceptionStr := `
INSERT INTO baa_application.vendor_service.lfnc_penalty_exception (
	id_lfnc_penalty_exception
	,bob_id_sales_order_item
	,seller_name
	,id_seller
	,initiator
	,amount
	,reason
	,lfnc_penalty_exception_status
	,fk_lfnc_penalty_exception_status) 
VALUES (@p1,@p2,@p3,@p4,@p5,@p6,@p7,'pending',1);

INSERT INTO baa_application.vendor_service.lfnc_penalty_soi_exception (
	fk_lfnc_penalty_exception
	,bob_id_sales_order_item
	,fk_lfnc_penalty_exception_status) 
VALUES (@p1,@p2,1);
`
	insertSoiException, err := dbBaa.Prepare(insertSoiExceptionStr)

	res, err := insertSoiException.Exec(
		lfncPenaltyException.IDLfncPenaltyException,
		lfncPenaltyException.BobIDSalesOrderNumber,
		lfncPenaltyException.SellerName,
		lfncPenaltyException.IDSeller,
		lfncPenaltyException.Initiator,
		lfncPenaltyException.Amount,
		lfncPenaltyException.Reason,
	)

	log.Println(res)

	return err
}

func CreateNewUser(userFormInput *useraccess.User, dbBaa *sql.DB) error {
	// prepare statement to insert values into baa_application.vendor_service.lfnc_penalty_exception
	insertNewUserStr := `
INSERT INTO baa_application.vendor_service.user_access (
	email 
	,name
	,access) 
VALUES (@p1,@p2,'lfnc_user')`
	insertNewUser, err := dbBaa.Prepare(insertNewUserStr)

	res, err := insertNewUser.Exec(
		userFormInput.Email,
		userFormInput.Name,
	)

	log.Println(res)

	return err
}

// GetExceptionRequest fetches pending purchase requests from baa_application.vendor_service.lfnc_penalty_exception
func GetExceptionRequest(dbBaa *sql.DB) []*lfncpenaltyexception.LfncPenaltyException {
	log.Println(`GetExceptionRequest`)
	rows, err := dbBaa.Query(`SELECT 
		COALESCE(lpe.uid_lfnc_penalty_exception,'') 'id_lfnc_penalty_exception'
		,COALESCE(lpe.initiator,'')  'initiator'
		,COALESCE(lpe.seller_name,'')  'seller_name'
		,COALESCE(CONVERT(VARCHAR(50), lpe.start_date, 101),'') 'start_date'
		,COALESCE(CONVERT(VARCHAR(50), lpe.end_date, 101),'') 'end_date' 
		,COALESCE(lpe.bob_id_sales_order_item,'')  'bob_id_sales_order_item'
    	,COALESCE(CAST(ROUND(lpe.amount ,2) as numeric(36,2)),'') 'amount'
		,COALESCE(lpe.reason,'') 'reason'
		,COALESCE(lpe.lfnc_penalty_exception_status,'')  'lfnc_penalty_exception_status'
		FROM baa_application.vendor_service.lfnc_penalty_exception lpe
		WHERE lpe.fk_lfnc_penalty_exception_status = 1`)

	checkError(err)
	defer rows.Close()

	// Create the data structure that is returned from the function.
	lfncPenaltyExceptionTable := []*lfncpenaltyexception.LfncPenaltyException{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a lfncPenaltyException,
		lfncPenaltyException := &lfncpenaltyexception.LfncPenaltyException{}
		// Populate the attributes of the lfncPenaltyException,
		// and return incase of an error
		err := rows.Scan(
			&lfncPenaltyException.IDLfncPenaltyException,
			&lfncPenaltyException.Initiator,
			&lfncPenaltyException.SellerName,
			&lfncPenaltyException.StartDate,
			&lfncPenaltyException.EndDate,
			&lfncPenaltyException.BobIDSalesOrderNumber,
			&lfncPenaltyException.Amount,
			&lfncPenaltyException.Reason,
			&lfncPenaltyException.LfncPenaltyExceptionStatus,
		)
		checkError(err)
		// Finally, append the result to the returned array, and repeat for
		// the next row
		lfncPenaltyExceptionTable = append(lfncPenaltyExceptionTable, lfncPenaltyException)
	}

	return lfncPenaltyExceptionTable

}

// GetIDExceptionRequest fetches iDExceptionRequestTable from baa_application.vendor_service.lfnc_penalty_exception representing all pending requests
func GetIDExceptionRequest(dbBaa *sql.DB) []string {

	// store iDExceptionRequestQuery in a string
	iDExceptionRequestQuery := `SELECT 
	lpe.uid_lfnc_penalty_exception 
	FROM baa_application.vendor_service.lfnc_penalty_exception lpe
	WHERE lpe.lfnc_penalty_exception_status = 'pending'`

	// write iDExceptionRequestQuery result to an array of fields.InputChoice , this array of rows represents iDExceptionRequestTable
	var iDExceptionRequest string
	var iDExceptionRequestTable []string
	rows, err := dbBaa.Query(iDExceptionRequestQuery)
	checkError(err)
	for rows.Next() {
		err := rows.Scan(&iDExceptionRequest)
		checkError(err)
		iDExceptionRequestTable = append(iDExceptionRequestTable, iDExceptionRequest)
	}
	return iDExceptionRequestTable
}

func AcceptException(iDExceptionRequest string, dbBaa *sql.DB) error {

	acceptExceptionStr := `
	UPDATE baa_application.vendor_service.lfnc_penalty_exception
	SET baa_application.vendor_service.lfnc_penalty_exception.lfnc_penalty_exception_status = 'approved'
      ,baa_application.vendor_service.lfnc_penalty_exception.fk_lfnc_penalty_exception_status = '2' 
	WHERE baa_application.vendor_service.lfnc_penalty_exception.uid_lfnc_penalty_exception = @p1`

	acceptException, err := dbBaa.Prepare(acceptExceptionStr)

	res, err := acceptException.Exec(iDExceptionRequest)

	log.Println(res)

	return err
}

func RejectException(iDExceptionRequest string, dbBaa *sql.DB) error {

	rejectExceptionStr := `
	UPDATE baa_application.vendor_service.lfnc_penalty_exception
	SET baa_application.vendor_service.lfnc_penalty_exception.lfnc_penalty_exception_status = 'rejected'
	,baa_application.vendor_service.lfnc_penalty_exception.fk_lfnc_penalty_exception_status = '3' 
	WHERE baa_application.vendor_service.lfnc_penalty_exception.uid_lfnc_penalty_exception = @p1`

	rejectException, err := dbBaa.Prepare(rejectExceptionStr)

	res, err := rejectException.Exec(iDExceptionRequest)

	log.Println(res)

	return err
}

func GetUserInfo(user *useraccess.User, dbBaa *sql.DB) {

	// store userQuery in a string
	userQuery := `SELECT 
	ua.name
	,ua.access
	FROM baa_application.vendor_service.user_access ua
	WHERE ua.email = @p1`

	err := dbBaa.QueryRow(userQuery, user.Email).Scan(&user.Name, &user.Access)
	if err != nil {
		if err.Error() != `sql: no rows in result set` {
			log.Fatal(err.Error())
		}

	}

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
