package main

import (
	"context"
	"log"

	"github.com/rammyblog/go-paystack"
)

const APIKey = "x"

func main() {
	ctx := context.Background()
	newClient := paystack.New(APIKey)
	initializeTransaction(ctx, newClient)
}
