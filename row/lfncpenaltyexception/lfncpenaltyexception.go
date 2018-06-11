package lfncpenaltyexception

// LfncPenaltyException represents one purchase request
// it also includes potential error or form to include into the web page used to collect purchase requests
type LfncPenaltyException struct {
	IDLfncPenaltyException       string `json:"id_lfnc_penalty_exception"`
	Initiator                    string `json:"initiator"`
	SellerName                   string `json:"seller_name"`
	IDSeller                     string `json:"id_seller"`
	StartDate                    string `json:"start_date"`
	EndDate                      string `json:"end_date"`
	OrderNr                      string `json:"order_nr"`
	BobIDSalesOrderNumber        string `json:"bob_id_sales_order_item"`
	Amount                       string `json:"amount"`
	Reason                       string `json:"reason"`
	LfncPenaltyExceptionStatus   string `json:"lfnc_penalty_exception_status"`
	FkLfncPenaltyExceptionStatus string `json:"fk_lfnc_penalty_exception_status"`

	Error string
}

/*// Validate validates the data of the purchase request sent by the user
func (lfncPenaltyException *LfncPenaltyException) Validate() bool {

	lfncPenaltyException.Error = ""

	// define validation of each field of the purchase request
	err := validation.ValidateStruct(lfncPenaltyException,
		validation.Field(&lfncPenaltyException.IDLfncPenaltyException, validation.Required),
		validation.Field(&lfncPenaltyException.Initiator, validation.Required),
		validation.Field(&lfncPenaltyException.SellerName, validation.Required),
		validation.Field(&lfncPenaltyException.IDSeller, validation.Required),
		validation.Field(&lfncPenaltyException.StartDate, validation.Required, validation.Date("1/2/2006")),
		validation.Field(&lfncPenaltyException.EndDate, validation.Required, validation.Date("1/2/2006")),
		validation.Field(&lfncPenaltyException.OrderNr, validation.Required, is.Int),
		validation.Field(&lfncPenaltyException.BobIDSalesOrderNumber, validation.Required, is.Int),
		validation.Field(&lfncPenaltyException.Amount, validation.Required, is.Float),
		validation.Field(&lfncPenaltyException.Reason, validation.Required, validation.Length(0, 400)),
		validation.Field(&lfncPenaltyException.LfncPenaltyExceptionStatus, validation.Required),
		validation.Field(&lfncPenaltyException.FkLfncPenaltyExceptionStatus, validation.Required, is.Int),
	)

	// add potential error text to lfncPenaltyException.Error
	if err != nil {
		lfncPenaltyException.Error = err.Error()
	}

	// return true if no error, false otherwise
	return lfncPenaltyException.Error == ""
}

// Render the web page itself given the html template and the lfncPenaltyException
func (lfncPenaltyException *LfncPenaltyException) Render(c *gin.Context, htmlTemplate string) {
	// fetch the htmlTemplate
	tmpl, err := template.ParseFiles(htmlTemplate)
	handleErr(c, err)
	// render the htmlTemplate given the lfncPenaltyException
	err = tmpl.Execute(c.Writer, map[string]interface{}{
		`IDLfncPenaltyException`:       lfncPenaltyException.IDLfncPenaltyException,
		`Initiator`:                    lfncPenaltyException.Initiator,
		`SellerName`:                   lfncPenaltyException.SellerName,
		`IDSeller`:                     lfncPenaltyException.IDSeller,
		`StartDate`:                    lfncPenaltyException.StartDate,
		`EndDate`:                      lfncPenaltyException.EndDate,
		`OrderNr`:                      lfncPenaltyException.OrderNr,
		`BobIDSalesOrderNumber`:        lfncPenaltyException.BobIDSalesOrderNumber,
		`Amount`:                       lfncPenaltyException.Amount,
		`Reason`:                       lfncPenaltyException.Reason,
		`LfncPenaltyExceptionStatus`:   lfncPenaltyException.LfncPenaltyExceptionStatus,
		`FkLfncPenaltyExceptionStatus`: lfncPenaltyException.FkLfncPenaltyExceptionStatus,
	})
	handleErr(c, err)
}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}*/
