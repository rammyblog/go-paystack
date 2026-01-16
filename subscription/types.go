package subscription

import (
	"github.com/rammyblog/go-paystack/internal/types"
	"github.com/rammyblog/go-paystack/plans"
	"github.com/rammyblog/go-paystack/transaction"
)

type SubscriptionListResponse struct {
	types.SubscriptionResponse
	Customer transaction.Customer `json:"customer,omitempty"`
	Plan     plans.PlanResponse   `json:"plan,omitempty"`
}

type SubscriptionSplitList struct {
	Meta types.PaginationMeta       `json:"meta"`
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
