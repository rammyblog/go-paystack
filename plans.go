package paystack

import (
	"context"
	"fmt"
)

type Plans struct {
	client *Client
}

func newPlans(client *Client) *Plans {
	return &Plans{
		client: client,
	}
}

type CreatePlanRequest struct {
	Name         string  `json:"name,omitempty"`
	Amount       float32 `json:"amount,omitempty"`
	Currency     string  `json:"currency,omitempty"`
	Interval     string  `json:"interval,omitempty"`
	Description  string  `json:"description,omitempty"`
	SendInvoices bool    `json:"send_invoices,omitempty"`
	SendSMS      bool    `json:"send_sms,omitempty"`
	InvoiceLimit bool    `json:"invoice_limit,omitempty"`
}

type PlanResponse struct {
	ID                int                    `json:"id,omitempty"`
	Name              string                 `json:"name,omitempty"`
	Amount            int                    `json:"amount,omitempty"`
	Interval          string                 `json:"interval,omitempty"`
	Integration       int                    `json:"integration,omitempty"`
	Domain            string                 `json:"domain,omitempty"`
	PlanCode          string                 `json:"plan_code,omitempty"`
	SendInvoices      bool                   `json:"send_invoices,omitempty"`
	SendSms           bool                   `json:"send_sms,omitempty"`
	HostedPage        bool                   `json:"hosted_page,omitempty"`
	Currency          string                 `json:"currency,omitempty"`
	CreatedAt         string                 `json:"createdAt,omitempty"`
	UpdatedAt         string                 `json:"updatedAt,omitempty"`
	HostedPageURL     string                 `json:"hosted_page_url,omitempty"`
	HostedPageSummary string                 `json:"hosted_page_summary,omitempty"`
	Subscription      []SubscriptionResponse `json:"subscription,omitempty"`
}

type PlanSplitList struct {
	Meta PaginationMeta      `json:"meta"`
	Data []PlanSplitResponse `json:"data"`
}

type SubscriptionResponse struct {
	Customer         int           `json:"customer,omitempty"`
	Plan             int           `json:"plan,omitempty"`
	Integration      int           `json:"integration,omitempty"`
	Domain           string        `json:"domain,omitempty"`
	Start            int           `json:"start,omitempty"`
	Status           string        `json:"status,omitempty"`
	Quantity         int           `json:"quantity,omitempty"`
	Amount           int           `json:"amount,omitempty"`
	SubscriptionCode string        `json:"subscription_code,omitempty"`
	EmailToken       string        `json:"email_token,omitempty"`
	Authorization    Authorization `json:"authorization,omitempty"`
	EasyCronID       interface{}   `json:"easy_cron_id,omitempty"`
	CronExpression   string        `json:"cron_expression,omitempty"`
	NextPaymentDate  string        `json:"next_payment_date,omitempty"`
	OpenInvoice      interface{}   `json:"open_invoice,omitempty"`
	ID               int           `json:"id,omitempty"`
	CreatedAt        string        `json:"createdAt,omitempty"`
	UpdatedAt        string        `json:"updatedAt,omitempty"`
}

type PlanSplitResponse struct {
	Subscription      []SubscriptionResponse `json:"subscription,omitempty"`
	Integration       int                    `json:"integration,omitempty"`
	Domain            string                 `json:"domain,omitempty"`
	Name              string                 `json:"name,omitempty"`
	PlanCode          string                 `json:"plan_code,omitempty"`
	Description       interface{}            `json:"description,omitempty"`
	Amount            int                    `json:"amount,omitempty"`
	Interval          string                 `json:"interval,omitempty"`
	SendInvoices      bool                   `json:"send_invoices,omitempty"`
	SendSms           bool                   `json:"send_sms,omitempty"`
	HostedPage        bool                   `json:"hosted_page,omitempty"`
	HostedPageURL     interface{}            `json:"hosted_page_url,omitempty"`
	HostedPageSummary interface{}            `json:"hosted_page_summary,omitempty"`
	Currency          string                 `json:"currency,omitempty"`
	ID                int                    `json:"id,omitempty"`
	CreatedAt         string                 `json:"createdAt,omitempty"`
	UpdatedAt         string                 `json:"updatedAt,omitempty"`
}

// Create creates a new plan.
// It sends a POST request to the "/plan" endpoint with the provided CreatePlanRequest
// and returns a CreatePlanResponse and an error, if any.
// For more details see https://paystack.com/docs/api/plan/#create
func (pn *Plans) Create(ctx context.Context, pln *CreatePlanRequest) (*PlanResponse, error) {
	url := "/plan"
	resp := &PlanResponse{}
	err := postResource(ctx, pn.client, url, pln, resp)
	return resp, err
}

//	List/Search Plans
//
// For more details see https://paystack.com/docs/api/plan/#list
func (pn *Plans) List(ctx context.Context, params ...QueryType) (*PlanSplitList, error) {
	var url string
	if len(params) > 0 {
		url = addQueryToUrl("plan", params...)
	} else {
		url = "/plan"
	}
	resp := &PlanSplitList{}

	err := getResource(ctx, pn.client, url, resp)
	return resp, err
}

//	Fetch a Plan
//
// For more details see https://paystack.com/docs/api/plan/#fetch
func (pn *Plans) Fetch(ctx context.Context, id string) (*PlanResponse, error) {

	url := fmt.Sprintf("/plan/%s", id)

	resp := &PlanResponse{}

	err := getResource(ctx, pn.client, url, resp)
	return resp, err
}

// Update a plan
// For more details see https://paystack.com/docs/api/plan/#update
func (pn *Plans) Update(ctx context.Context, id string, txn *CreatePlanRequest) (*PlanResponse, error) {
	url := fmt.Sprintf("/plan/%s", id)
	resp := &PlanResponse{}
	err := putResource(ctx, pn.client, url, txn, resp)
	return resp, err
}
