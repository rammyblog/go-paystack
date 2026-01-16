package main

import (
	"context"
	"log"

	"github.com/rammyblog/go-paystack"
	"github.com/rammyblog/go-paystack/plans"
)

func createPlan(ctx context.Context, c *paystack.Client) {
	// Create a plan
	resp, err := c.Plan.Create(ctx, &plans.CreatePlanRequest{
		Name:        "Monthly retainer",
		Amount:      500000,
		Interval:    "monthly",
		Description: "Monthly retainer",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n Create Plan \n-%+v\n", resp)
}

func listPlans(ctx context.Context, c *paystack.Client) {
	// List plans
	resp, err := c.Plan.List(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n List Plans \n-%+v\n", resp)
}

func fetchPlan(ctx context.Context, c *paystack.Client) {
	// Fetch a plan
	resp, err := c.Plan.Fetch(ctx, "PLN_48mmlpyngprj1fk")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n Fetch Plan \n-%+v\n", resp)
}

func updatePlan(ctx context.Context, c *paystack.Client) {
	// Update a plan
	resp, err := c.Plan.Update(ctx, "PLN_48mmlpyngprj1fk", &plans.CreatePlanRequest{
		Name:        "Annual retainer",
		Amount:      60000,
		Interval:    "annually",
		Description: "Annual retainer",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n Update Plan \n-%+v\n", resp)
}
