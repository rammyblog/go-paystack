package main

import (
	"context"
	"log"

	"github.com/rammyblog/go-paystack"
)

func CreateSubscription(ctx context.Context, c *paystack.Client) {
	// Create a subscription
	resp, err := c.Subscription.CreateSubscription(ctx, &paystack.CreateSubscriptionRequest{
		Customer: "CUS_potop1jwzmml5ng  ",
		Plan:     "PLN_48mmlpyngprj1fk",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n Create Subscription \n-%+v\n", resp)
}

func ListSubscription(ctx context.Context, c *paystack.Client) {
	// List subscriptions
	resp, err := c.Subscription.List(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n List Subscriptions \n-%+v\n", resp)
}

func FetchSubscription(ctx context.Context, c *paystack.Client) {
	// Fetch a subscription
	resp, err := c.Subscription.Fetch(ctx, "SUB_369x89yaz8ifiym")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n Fetch Subscription \n-%+v\n", resp)
}

func EnableSubscription(ctx context.Context, c *paystack.Client) {
	// Enable a subscription
	resp, err := c.Subscription.Enable(ctx, &paystack.EnableDisableSubscriptionRequest{
		Code: "SUB_369x89yaz8ifiym",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n Enable Subscription \n-%+v\n", resp)
}

func DisableSubscription(ctx context.Context, c *paystack.Client) {
	// Disable a subscription
	resp, err := c.Subscription.Disable(ctx, &paystack.EnableDisableSubscriptionRequest{
		Code: "SUB_369x89yaz8ifiym",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n Disable Subscription \n-%+v\n", resp)
}

func GenerateUpdateSubscriptionLink(ctx context.Context, c *paystack.Client) {
	// Generate update subscription link
	resp, err := c.Subscription.GenerateUpdateSubscriptionLink(ctx, "SUB_369x89yaz8ifiym")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n Generate Update Subscription Link \n-%+v\n", resp)
}

func SendUpdateSubscriptionLink(ctx context.Context, c *paystack.Client) {
	// Send update subscription link
	resp, err := c.Subscription.SendUpdateSubscriptionLink(ctx, "SUB_369x89yaz8ifiym", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n Send Update Subscription Link \n-%+v\n", resp)
}
