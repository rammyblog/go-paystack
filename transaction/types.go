package transaction

import (
	"github.com/rammyblog/go-paystack/internal/types"
)

// TransactionList is a list object for transactions.
type TransactionList struct {
	Meta types.PaginationMeta  `json:"meta"`
	Data []TransactionResponse `json:"data"`
}

type Authorization struct {
	AuthorizationCode string      `json:"authorization_code"`
	Bin               string      `json:"bin"`
	Last4             string      `json:"last4"`
	ExpMonth          string      `json:"exp_month"`
	ExpYear           string      `json:"exp_year"`
	Channel           string      `json:"channel"`
	CardType          string      `json:"card_type"`
	Bank              string      `json:"bank"`
	CountryCode       string      `json:"country_code"`
	Brand             string      `json:"brand"`
	Reusable          bool        `json:"reusable"`
	Signature         string      `json:"signature"`
	AccountName       interface{} `json:"account_name"`
}

type Customer struct {
	ID                       int         `json:"id"`
	FirstName                interface{} `json:"first_name"`
	LastName                 interface{} `json:"last_name"`
	Email                    string      `json:"email"`
	CustomerCode             string      `json:"customer_code"`
	Phone                    interface{} `json:"phone"`
	Metadata                 interface{} `json:"metadata"`
	RiskAction               string      `json:"risk_action"`
	InternationalFormatPhone interface{} `json:"international_format_phone"`
}

type TransactionTimeline struct {
	StartTime      int           `json:"start_time"`
	TimeSpent      int           `json:"time_spent,omitempty"`
	Attempts       int           `json:"attempts,omitempty"`
	Authentication interface{}   `json:"authentication,omitempty"`
	Errors         int           `json:"errors,omitempty"`
	Success        bool          `json:"success,omitempty"`
	Mobile         bool          `json:"mobile,omitempty"`
	Input          []interface{} `json:"input,omitempty"`
	Channel        string        `json:"channel,omitempty"`
	History        []struct {
		Type    string `json:"type,omitempty"`
		Message string `json:"message,omitempty"`
		Time    int    `json:"time,omitempty"`
	} `json:"history,omitempty"`
}

// TransactionRequest represents a request to start a transaction.
type TransactionRequest struct {
	Amount            float32        `json:"amount,omitempty"`
	Email             string         `json:"email,omitempty"`
	Currency          string         `json:"currency,omitempty"`
	Reference         string         `json:"reference,omitempty"`
	CallbackURL       string         `json:"callback_url,omitempty"`
	Plan              string         `json:"plan,omitempty"`
	InvoiceLimit      int            `json:"invoice_limit,omitempty"`
	Metadata          types.Metadata `json:"metadata,omitempty"`
	Channels          []string       `json:"channels,omitempty"`
	SplitCode         string         `json:"split_code,omitempty"`
	SubAccount        string         `json:"subaccount,omitempty"`
	TransactionCharge int            `json:"transaction_charge,omitempty"`
	AuthorizationCode string         `json:"authorization_code,omitempty"`
	Bearer            string         `json:"bearer,omitempty"`
	Queue             bool           `json:"queue,omitempty"`
}

type InitializeTransactionResponse struct {
	AuthorizationURL string `json:"authorization_url,omitempty"`
	AccessCode       string `json:"access_code,omitempty"`
	Reference        string `json:"reference,omitempty"`
}
type TransactionResponse struct {
	ID                 int                 `json:"id"`
	Domain             string              `json:"domain"`
	Status             string              `json:"status"`
	Reference          string              `json:"reference"`
	Amount             int                 `json:"amount"`
	Message            interface{}         `json:"message"`
	GatewayResponse    string              `json:"gateway_response"`
	Channel            string              `json:"channel"`
	Currency           string              `json:"currency"`
	IPAddress          string              `json:"ip_address"`
	Metadata           interface{}         `json:"metadata"`
	Log                TransactionTimeline `json:"log"`
	Fees               int                 `json:"fees"`
	FeesSplit          interface{}         `json:"fees_split"`
	Authorization      Authorization       `json:"authorization"`
	Customer           Customer            `json:"customer"`
	Plan               interface{}         `json:"plan"`
	Split              interface{}         `json:"split"`
	OrderID            interface{}         `json:"order_id"`
	PaidAt             string              `json:"paidAt"`
	CreatedAt          string              `json:"createdAt"`
	RequestedAmount    int                 `json:"requested_amount"`
	PosTransactionData interface{}         `json:"pos_transaction_data"`
	Source             interface{}         `json:"source"`
	FeesBreakdown      interface{}         `json:"fees_breakdown"`
	TransactionDate    string              `json:"transaction_date"`
	PlanObject         interface{}         `json:"plan_object"`
	Subaccount         interface{}         `json:"subaccount"`
}

type ExportTransaction struct {
	Path string `json:"path,omitempty"`
}
