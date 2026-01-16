package main

import (
	"context"
	"log"
	"math/rand"
	"strings"

	"github.com/rammyblog/go-paystack"
	"github.com/rammyblog/go-paystack/transaction"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	var sb strings.Builder
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
func initializeTransaction(ctx context.Context, c *paystack.Client) {
	resp, err := c.Transaction.Initialize(ctx, &transaction.TransactionRequest{
		Amount:      100000,
		Email:       "Onas@gmail.com",
		Currency:    "NGN",
		Reference:   randomString(49),
		CallbackURL: "https://ngrok.com/rammyblof",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n Initialize transaction \n-%+v\n", resp)
}

func verifyTransaction(ctx context.Context, c *paystack.Client) {
	resp, err := c.Transaction.Verify(ctx, "yinmu")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n verify transaction \n-%+v\n", resp.Log)
}

func listTransaction(ctx context.Context, c *paystack.Client) {
	resp, err := c.Transaction.List(ctx, paystack.Query("page", "1"))
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\ntransactions-%+v\n", resp)
}

func fetchTransaction(ctx context.Context, c *paystack.Client) {
	resp, err := c.Transaction.FetchById(ctx, 3267520111)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n fetch transaction \n-%+v\n", resp)
}

func chargeTransaction(ctx context.Context, c *paystack.Client) {

	resp, err := c.Transaction.Charge(ctx, &transaction.TransactionRequest{
		AuthorizationCode: "AUTH_72btv547",
		Amount:            1000000,
		Email:             "onasanyatunde@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n charge transaction \n-%+v\n", resp)
}

func totalTransaction(ctx context.Context, c *paystack.Client) {

	resp, err := c.Transaction.Total(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n total transaction \n-%+v\n", resp)
}

func viewTransactionTimeline(ctx context.Context, c *paystack.Client) {
	resp, err := c.Transaction.FetchById(ctx, 3267520111)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n fetch transaction \n-%+v\n", resp)
}

func exportTransaction(ctx context.Context, c *paystack.Client) {

	resp, err := c.Transaction.Export(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n Export transaction \n-%+v\n", resp)
}

func partialDebit(ctx context.Context, c *paystack.Client) {

	resp, err := c.Transaction.PartialDebit(ctx, &transaction.TransactionRequest{
		AuthorizationCode: "AUTH_72btv547",
		Amount:            1000000,
		Email:             "onas@gmail.com",
		Currency:          "NGN",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n partial debit transaction \n-%+v\n", resp)
}
