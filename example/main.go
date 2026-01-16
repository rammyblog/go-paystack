package main

import (
	"context"
	"log"

	"github.com/rammyblog/go-paystack"
)

const APIKey = "x"

func main() {
	ctx := context.Background()
	newClient := paystack.NewClient(APIKey, &paystack.LogConfig{
		Level:      paystack.LogLevelDebug,
		JSONOutput: true,
	})

	resp, err := newClient.Transaction.Verify(ctx, "TX_Onas")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n Initialize transaction \n-%+v\n", resp)
}
