package main

import (
	"context"
	"log"

	"github.com/rammyblog/go-paystack"
	transaction_splits "github.com/rammyblog/go-paystack/transaction-splits"
)

func createTransactionSplit(ctx context.Context, c *paystack.Client) {
	subaccountArray := []transaction_splits.TransactionSplitSubAccount{
		{Subaccount: "ACCT_jo822mgh7xbjzqk", Share: 60},
	}
	resp, err := c.TransactionSplit.Create(ctx, &transaction_splits.TransactionSplitRequest{
		Name:             "Percentage Split",
		Type:             "percentage",
		Currency:         "NGN",
		Subaccounts:      subaccountArray,
		BearerType:       "subaccount",
		BearerSubaccount: "ACCT_jo822mgh7xbjzqk",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n Initialize transaction split \n-%+v\n", resp)
}

func listTransactionSplits(ctx context.Context, c *paystack.Client) {

	resp, err := c.TransactionSplit.List(ctx)

	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n  transaction splits \n-%+v\n", resp)
}

func fetchTransactionSplit(ctx context.Context, c *paystack.Client) {

	resp, err := c.TransactionSplit.Fetch(ctx, "1499896")

	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n  transaction splits \n-%+v\n", resp)
}

func updateTransactionSplit(ctx context.Context, c *paystack.Client) {

	resp, err := c.TransactionSplit.Update(ctx, "1499896", &transaction_splits.TransactionSplitRequest{
		Name:   "Updated name",
		Active: false,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n updated transaction split \n-%+v\n", resp)
}

func addSubaccountToTransactionSplit(ctx context.Context, c *paystack.Client) {

	resp, err := c.TransactionSplit.AddSubaccount(ctx, "1499896", &transaction_splits.TransactionSplitSubAccount{
		Subaccount: "ACCT_ch9drs2a9pka9fb",
		Share:      10,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n updated transaction split \n-%+v\n", resp)
}

func removeSubaccountToTransactionSplit(ctx context.Context, c *paystack.Client) {

	resp, err := c.TransactionSplit.RemoveSubaccount(ctx, "1499896", &transaction_splits.TransactionSplitSubAccount{
		Subaccount: "ACCT_ch9drs2a9pka9fb",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n remove subaccount transaction split \n-%+v\n", resp)
}
