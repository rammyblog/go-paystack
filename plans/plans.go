package plans

import (
	"context"
	"fmt"

	"github.com/rammyblog/go-paystack/internal/helpers"
	"github.com/rammyblog/go-paystack/internal/types"
)

type Plans struct {
	client types.Requester
}

func New(client types.Requester) *Plans {
	return &Plans{
		client: client,
	}
}

// Create creates a new plan.
// It sends a POST request to the "/plan" endpoint with the provided CreatePlanRequest
// and returns a CreatePlanResponse and an error, if any.
// For more details see https://paystack.com/docs/api/plan/#create
func (pn *Plans) Create(ctx context.Context, pln *CreatePlanRequest) (*PlanResponse, error) {
	url := "/plan"
	resp := &PlanResponse{}
	err := pn.client.PostResource(ctx, url, pln, resp)
	return resp, err
}

//	List/Search Plans
//
// For more details see https://paystack.com/docs/api/plan/#list
func (pn *Plans) List(ctx context.Context, params ...types.QueryType) (*PlanSplitList, error) {
	var url string
	if len(params) > 0 {
		url = helpers.AddQueryToUrl("plan", params...)
	} else {
		url = "/plan"
	}
	resp := &PlanSplitList{}

	err := pn.client.GetResource(ctx, url, resp)
	return resp, err
}

//	Fetch a Plan
//
// For more details see https://paystack.com/docs/api/plan/#fetch
func (pn *Plans) Fetch(ctx context.Context, id string) (*PlanResponse, error) {

	url := fmt.Sprintf("/plan/%s", id)

	resp := &PlanResponse{}

	err := pn.client.GetResource(ctx, url, resp)
	return resp, err
}

// Update a plan
// For more details see https://paystack.com/docs/api/plan/#update
func (pn *Plans) Update(ctx context.Context, id string, txn *CreatePlanRequest) (*PlanResponse, error) {
	url := fmt.Sprintf("/plan/%s", id)
	resp := &PlanResponse{}
	err := pn.client.PutResource(ctx, url, txn, resp)
	return resp, err
}
