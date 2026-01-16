package plans

import (
	"github.com/rammyblog/go-paystack/internal/types"
)

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
	ID                int                          `json:"id,omitempty"`
	Name              string                       `json:"name,omitempty"`
	Amount            int                          `json:"amount,omitempty"`
	Interval          string                       `json:"interval,omitempty"`
	Integration       int                          `json:"integration,omitempty"`
	Domain            string                       `json:"domain,omitempty"`
	PlanCode          string                       `json:"plan_code,omitempty"`
	SendInvoices      bool                         `json:"send_invoices,omitempty"`
	SendSms           bool                         `json:"send_sms,omitempty"`
	HostedPage        bool                         `json:"hosted_page,omitempty"`
	Currency          string                       `json:"currency,omitempty"`
	CreatedAt         string                       `json:"createdAt,omitempty"`
	UpdatedAt         string                       `json:"updatedAt,omitempty"`
	HostedPageURL     string                       `json:"hosted_page_url,omitempty"`
	HostedPageSummary string                       `json:"hosted_page_summary,omitempty"`
	Subscription      []types.SubscriptionResponse `json:"subscription,omitempty"`
}

type PlanSplitList struct {
	Meta types.PaginationMeta `json:"meta"`
	Data []PlanSplitResponse  `json:"data"`
}

type PlanSplitResponse struct {
	Subscription      []types.SubscriptionResponse `json:"subscription,omitempty"`
	Integration       int                          `json:"integration,omitempty"`
	Domain            string                       `json:"domain,omitempty"`
	Name              string                       `json:"name,omitempty"`
	PlanCode          string                       `json:"plan_code,omitempty"`
	Description       interface{}                  `json:"description,omitempty"`
	Amount            int                          `json:"amount,omitempty"`
	Interval          string                       `json:"interval,omitempty"`
	SendInvoices      bool                         `json:"send_invoices,omitempty"`
	SendSms           bool                         `json:"send_sms,omitempty"`
	HostedPage        bool                         `json:"hosted_page,omitempty"`
	HostedPageURL     interface{}                  `json:"hosted_page_url,omitempty"`
	HostedPageSummary interface{}                  `json:"hosted_page_summary,omitempty"`
	Currency          string                       `json:"currency,omitempty"`
	ID                int                          `json:"id,omitempty"`
	CreatedAt         string                       `json:"createdAt,omitempty"`
	UpdatedAt         string                       `json:"updatedAt,omitempty"`
}
