package types

import "context"

type QueryType struct {
	Key   string
	Value string
}

// PaginationMeta is pagination metadata for paginated responses from the Paystack API
type PaginationMeta struct {
	Total     int `json:"total"`
	Skipped   int `json:"skipped"`
	PerPage   int `json:"perPage"`
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}

type Requester interface {
	PostResource(ctx context.Context, url string, body, res interface{}) error
	GetResource(ctx context.Context, url string, res interface{}) error
	PutResource(ctx context.Context, url string, body, res interface{}) error
	DeleteResource(ctx context.Context, url string, res interface{}) error
}

type APIResponse map[string]interface{}

type GenericResponse map[string]interface{}

type Metadata map[string]interface{}

type SubscriptionResponse struct {
	Customer         int         `json:"customer,omitempty"`
	Plan             int         `json:"plan,omitempty"`
	Integration      int         `json:"integration,omitempty"`
	Domain           string      `json:"domain,omitempty"`
	Start            int         `json:"start,omitempty"`
	Status           string      `json:"status,omitempty"`
	Quantity         int         `json:"quantity,omitempty"`
	Amount           int         `json:"amount,omitempty"`
	SubscriptionCode string      `json:"subscription_code,omitempty"`
	EmailToken       string      `json:"email_token,omitempty"`
	Authorization    interface{} `json:"authorization,omitempty"`
	EasyCronID       interface{} `json:"easy_cron_id,omitempty"`
	CronExpression   string      `json:"cron_expression,omitempty"`
	NextPaymentDate  string      `json:"next_payment_date,omitempty"`
	OpenInvoice      interface{} `json:"open_invoice,omitempty"`
	ID               int         `json:"id,omitempty"`
	CreatedAt        string      `json:"createdAt,omitempty"`
	UpdatedAt        string      `json:"updatedAt,omitempty"`
}
