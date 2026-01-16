package transaction_splits

import "github.com/rammyblog/go-paystack/internal/types"

type TransactionSplitSubAccount struct {
	Subaccount string `json:"subaccount,omitempty"`
	Share      int    `json:"share,omitempty"`
}

type TransactionSplitRequest struct {
	Name             string                       `json:"name,omitempty"`
	Type             string                       `json:"type,omitempty"`
	Currency         string                       `json:"currency,omitempty"`
	Subaccounts      []TransactionSplitSubAccount `json:"subaccounts,omitempty"`
	BearerType       string                       `json:"bearer_type,omitempty"`
	BearerSubaccount string                       `json:"bearer_subaccount,omitempty"`
	Active           bool                         `json:"active,omitempty"`
}

type TransactionSplitList struct {
	Meta types.PaginationMeta       `json:"meta"`
	Data []TransactionSplitResponse `json:"data"`
}

type TransactionSplitResponse struct {
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Type             string `json:"type,omitempty"`
	Currency         string `json:"currency,omitempty"`
	Integration      int    `json:"integration,omitempty"`
	Domain           string `json:"domain,omitempty"`
	SplitCode        string `json:"split_code,omitempty"`
	Active           bool   `json:"active,omitempty"`
	BearerType       string `json:"bearer_type,omitempty"`
	BearerSubaccount int    `json:"bearer_subaccount,omitempty"`
	CreatedAt        string `json:"createdAt,omitempty"`
	UpdatedAt        string `json:"updatedAt,omitempty"`
	Subaccounts      []struct {
		Subaccount struct {
			ID                  int         `json:"id,omitempty"`
			SubaccountCode      string      `json:"subaccount_code,omitempty"`
			BusinessName        string      `json:"business_name,omitempty"`
			Description         string      `json:"description,omitempty"`
			PrimaryContactName  interface{} `json:"primary_contact_name,omitempty"`
			PrimaryContactEmail interface{} `json:"primary_contact_email,omitempty"`
			PrimaryContactPhone interface{} `json:"primary_contact_phone,omitempty"`
			Metadata            interface{} `json:"metadata,omitempty"`
			PercentageCharge    int         `json:"percentage_charge,omitempty"`
			SettlementBank      string      `json:"settlement_bank,omitempty"`
			AccountNumber       string      `json:"account_number,omitempty"`
		} `json:"subaccount,omitempty"`
		Share int `json:"share,omitempty"`
	} `json:"subaccounts,omitempty"`
	TotalSubaccounts int `json:"total_subaccounts,omitempty"`
}
