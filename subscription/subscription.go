package subscription

import (
	"context"
	"fmt"

	"github.com/rammyblog/go-paystack/internal/helpers"
	"github.com/rammyblog/go-paystack/internal/types"
)

// Subscription is the resource representing a Paystack subscription.
type Subscription struct {
	client types.Requester
}

// newSubscription returns a new Subscription.
func New(client types.Requester) *Subscription {
	return &Subscription{
		client: client,
	}
}

// CreateSubscription creates a new subscription.
// It sends a POST request to the "/plan" endpoint with the provided CreateSubscriptionRequest
// and returns a SubscriptionResponse and an error, if any.
// For more details see https://paystack.com/docs/api/subscription/#create
func (su *Subscription) CreateSubscription(ctx context.Context, sub *CreateSubscriptionRequest) (*types.SubscriptionResponse, error) {
	url := "/subscription"
	resp := &types.SubscriptionResponse{}
	err := su.client.PostResource(ctx, url, sub, resp)
	return resp, err
}

//	List/Search Subscription
//
// For more details see https://paystack.com/docs/api/subscription/#list
func (su *Subscription) List(ctx context.Context, params ...types.QueryType) (*SubscriptionSplitList, error) {
	var url string
	if len(params) > 0 {
		url = helpers.AddQueryToUrl("subscription", params...)
	} else {
		url = "/subscription"
	}
	resp := &SubscriptionSplitList{}

	err := su.client.GetResource(ctx, url, resp)
	return resp, err
}

//	Fetch a Subscription
//
// For more details see https://paystack.com/docs/api/subscription/#fetch
func (su *Subscription) Fetch(ctx context.Context, id string) (*SubscriptionListResponse, error) {

	url := fmt.Sprintf("/subscription/%s", id)

	resp := &SubscriptionListResponse{}

	err := su.client.GetResource(ctx, url, resp)
	return resp, err
}

// Enable a subscription
// For more details see https://paystack.com/docs/subscriptions/#enable
func (su *Subscription) Enable(ctx context.Context, sub *EnableDisableSubscriptionRequest) (*types.SubscriptionResponse, error) {
	url := "/subscription/enable"
	resp := &types.SubscriptionResponse{}
	err := su.client.PostResource(ctx, url, sub, resp)
	return resp, err
}

// Disable a subscription
// For more details see https://paystack.com/docs/subscriptions/#disable
func (su *Subscription) Disable(ctx context.Context, sub *EnableDisableSubscriptionRequest) (*types.SubscriptionResponse, error) {
	url := "/subscription/disable"
	resp := &types.SubscriptionResponse{}
	err := su.client.PostResource(ctx, url, sub, resp)
	return resp, err
}

// Generate Update Subscription Link
// For more details see https://paystack.com/docs/subscriptions/#generate-update-subscription-link
func (su *Subscription) GenerateUpdateSubscriptionLink(ctx context.Context, id string) (*GenerateUpdateSubscriptionLinkResponse, error) {
	url := fmt.Sprintf("/subscription/%s/manage/link", id)
	resp := &GenerateUpdateSubscriptionLinkResponse{}
	err := su.client.GetResource(ctx, url, resp)
	return resp, err
}

// Send Update Subscription Link
// For more details see https://paystack.com/docs/subscriptions/#send-update-subscription-link
func (su *Subscription) SendUpdateSubscriptionLink(ctx context.Context, id string, body *interface{}) (*types.SubscriptionResponse, error) {
	url := fmt.Sprintf("/subscription/%s/manage/email", id)
	resp := &types.SubscriptionResponse{}
	err := su.client.PostResource(ctx, url, body, resp)
	return resp, err
}
