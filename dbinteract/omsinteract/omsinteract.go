package omsinteract

import (
	"database/sql"
	"log"
)

func GetBobIDSupplierFromSellerName(dbOms *sql.DB, supplierName string) (iDSupplier string) {

	// store bobIDSupplierFromSellerName in a string
	bobIDSupplierFromSellerName := `
	SELECT 
	is1.bob_id_supplier 
	FROM ims_supplier is1
	WHERE is1.name_en = ?`

	err := dbOms.QueryRow(bobIDSupplierFromSellerName, supplierName).Scan(&iDSupplier)
	if err != nil {
		if err.Error() == `sql: no rows in result set` {
			return ""
		} else {
			log.Fatal(err.Error())
		}

	}

	return iDSupplier

}

func GetSupplierInfoFromBobIDSalesOrderItem(dbOms *sql.DB, iDBobSalesOrderItem string) []string {

	supplierInfo := []string{"", ""}

	// store iDBobSalesOrderItemQuery in a string
	supplierInfoFromBobIDSalesOrderItemQuery := `
		SELECT 
		is1.name_en 
		,is1.bob_id_supplier
		FROM ims_sales_order_item isoi
		LEFT JOIN ims_supplier is1
		ON isoi.bob_id_supplier = is1.bob_id_supplier
		WHERE isoi.bob_id_sales_order_item = ?`

	err := dbOms.QueryRow(supplierInfoFromBobIDSalesOrderItemQuery, iDBobSalesOrderItem).Scan(&supplierInfo[0], &supplierInfo[1])
	if err != nil {
		if err.Error() == `sql: no rows in result set` {
			log.Println(`sql: no rows in result set!`)
			return []string{"", ""}
		} else {
			log.Fatal(err.Error())
		}

	}

	return supplierInfo

}

/*

SELECT
 SI.bob_id_sales_order_item
 ,SI.canceled_at
 ,SI.fk_supplier AS id_supplier
 ,SC.name supplier_name
 ,SC.city_type
 ,SI.handled_by_marketplace_at
 ,SI.item_received_at

FROM StagingDB_Replica.Gathering.tblSalesItem SI

JOIN StagingDB_Replica.Gathering.tblSupplierCatalog SC
ON SC.id_supplier = SI.fk_supplier

WHERE SI.handled_by_marketplace_at IS NOT NULL
AND SI.item_received_at IS NOT NULL
AND SI.handled_by_marketplace_at >= DATEADD(DAY, -40, GETDATE()))

*/

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
