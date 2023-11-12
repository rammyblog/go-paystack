package main

import (
	"context"

	"github.com/rammyblog/go-paystack"
)

const APIKey = "sk_test_0000000000000000000000"

func main() {
	ctx := context.Background()
	newClient := paystack.NewClient(APIKey)
	removeSubaccountToTransactionSplit(ctx, newClient)
}
