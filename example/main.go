package main

import (
	"context"

	"github.com/rammyblog/go-paystack"
)

const APIKey = "sk_test_00000000000000"

func main() {
	ctx := context.Background()
	newClient := paystack.NewClient(APIKey)
	updatePlan(ctx, newClient)
}
