package paystack

import (
	"context"
	"fmt"
)

// Subscription is the resource representing a Paystack subscription.
type Subscription struct {
	client *Client
}

// newSubscription returns a new Subscription.
func newSubscription(client *Client) *Subscription {
	return &Subscription{
		client: client,
	}
}

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

type SubscriptionListResponse struct {
	SubscriptionResponse
	Customer Customer     `json:"customer,omitempty"`
	Plan     PlanResponse `json:"plan,omitempty"`
}

type SubscriptionSplitList struct {
	Meta PaginationMeta             `json:"meta"`
	Data []SubscriptionListResponse `json:"data"`
}

type CreateSubscriptionRequest struct {
	Customer      string `json:"customer,omitempty"`
	Plan          string `json:"plan,omitempty"`
	Authorization string `json:"authorization,omitempty"`
	StartDate     string `json:"start_date,omitempty"`
}

type EnableDisableSubscriptionRequest struct {
	Code       string `json:"code,omitempty"`
	EmailToken string `json:"email_token,omitempty"`
}

type GenerateUpdateSubscriptionLinkResponse struct {
	Link string `json:"link,omitempty"`
}

// CreateSubscription creates a new subscription.
// It sends a POST request to the "/plan" endpoint with the provided CreateSubscriptionRequest
// and returns a SubscriptionResponse and an error, if any.
// For more details see https://paystack.com/docs/api/subscription/#create
func (su *Subscription) CreateSubscription(ctx context.Context, sub *CreateSubscriptionRequest) (*SubscriptionResponse, error) {
	url := "/subscription"
	resp := &SubscriptionResponse{}
	err := postResource(ctx, su.client, url, sub, resp)
	return resp, err
}

//	List/Search Subscription
//
// For more details see https://paystack.com/docs/api/subscription/#list
func (su *Subscription) List(ctx context.Context, params ...QueryType) (*SubscriptionSplitList, error) {
	var url string
	if len(params) > 0 {
		url = addQueryToUrl("subscription", params...)
	} else {
		url = "/subscription"
	}
	resp := &SubscriptionSplitList{}

	err := getResource(ctx, su.client, url, resp)
	return resp, err
}

//	Fetch a Subscription
//
// For more details see https://paystack.com/docs/api/subscription/#fetch
func (su *Subscription) Fetch(ctx context.Context, id string) (*SubscriptionListResponse, error) {

	url := fmt.Sprintf("/subscription/%s", id)

	resp := &SubscriptionListResponse{}

	err := getResource(ctx, su.client, url, resp)
	return resp, err
}

// Enable a subscription
// For more details see https://paystack.com/docs/subscriptions/#enable
func (su *Subscription) Enable(ctx context.Context, sub *EnableDisableSubscriptionRequest) (*SubscriptionResponse, error) {
	url := "/subscription/enable"
	resp := &SubscriptionResponse{}
	err := postResource(ctx, su.client, url, sub, resp)
	return resp, err
}

// Disable a subscription
// For more details see https://paystack.com/docs/subscriptions/#disable
func (su *Subscription) Disable(ctx context.Context, sub *EnableDisableSubscriptionRequest) (*SubscriptionResponse, error) {
	url := "/subscription/disable"
	resp := &SubscriptionResponse{}
	err := postResource(ctx, su.client, url, sub, resp)
	return resp, err
}

// Generate Update Subscription Link
// For more details see https://paystack.com/docs/subscriptions/#generate-update-subscription-link
func (su *Subscription) GenerateUpdateSubscriptionLink(ctx context.Context, id string) (*GenerateUpdateSubscriptionLinkResponse, error) {
	url := fmt.Sprintf("/subscription/%s/manage/link", id)
	resp := &GenerateUpdateSubscriptionLinkResponse{}
	err := getResource(ctx, su.client, url, resp)
	return resp, err
}

// Send Update Subscription Link
// For more details see https://paystack.com/docs/subscriptions/#send-update-subscription-link
func (su *Subscription) SendUpdateSubscriptionLink(ctx context.Context, id string, body *interface{}) (*SubscriptionResponse, error) {
	url := fmt.Sprintf("/subscription/%s/manage/email", id)
	resp := &SubscriptionResponse{}
	err := postResource(ctx, su.client, url, body, resp)
	return resp, err
}
