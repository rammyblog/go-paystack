package transaction

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

type mockRequester struct {
	responseData interface{}
	err          error
	capturedURL  string
	capturedBody interface{}
}

func NewTestTransaction(overrides ...func(*TransactionResponse)) *TransactionResponse {
	txn := &TransactionResponse{
		ID:        123,
		Status:    "success",
		Reference: "ref_123",
	}
	for _, override := range overrides {
		override(txn)
	}
	return txn
}

func (m *mockRequester) GetResource(ctx context.Context, url string, res interface{}) error {
	m.capturedURL = url
	if m.err != nil {
		return m.err
	}
	return copyResponse(m.responseData, res)
}

func (m *mockRequester) PostResource(ctx context.Context, url string, body, res interface{}) error {
	m.capturedURL = url
	m.capturedBody = body
	if m.err != nil {
		return m.err
	}
	return copyResponse(m.responseData, res)
}

func copyResponse(src, dst interface{}) error {
	if src == nil {
		return nil
	}
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dst)
}

func (m *mockRequester) DeleteResource(ctx context.Context, url string, res interface{}) error {
	m.capturedURL = url
	return m.err
}

func (m *mockRequester) PutResource(ctx context.Context, url string, body, res interface{}) error {
	m.capturedURL = url
	return m.err
}

func TestTransaction_Verify(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		mockData := NewTestTransaction(func(tr *TransactionResponse) {
			gofakeit.Struct(&tr)
		})
		mock := &mockRequester{
			responseData: mockData,
		}

		tn := New(mock)
		trRef := "ref_123"
		resp, err := tn.Verify(context.Background(), trRef)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedURL := fmt.Sprintf(VERIFY_TRANSACTION, trRef)
		if mock.capturedURL != expectedURL {
			t.Errorf("URL = %q, want %q", mock.capturedURL, expectedURL)
		}

		if resp.ID != mockData.ID {
			t.Errorf("ID = %d, want %d", resp.ID, mockData.ID)
		}
		if resp.Status != mockData.Status {
			t.Errorf("Status = %q, want %q", resp.Status, mockData.Status)
		}
	})

	t.Run("error", func(t *testing.T) {
		mock := &mockRequester{err: errors.New("API error")}
		tn := New(mock)

		_, err := tn.Verify(context.Background(), "ref_123")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("empty reference", func(t *testing.T) {
		mockData := NewTestTransaction()
		mock := &mockRequester{responseData: mockData}
		tn := New(mock)

		_, err := tn.Verify(context.Background(), "")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expectedURL := fmt.Sprintf(VERIFY_TRANSACTION, "")
		if mock.capturedURL != expectedURL {
			t.Errorf("URL = %q, want %q", mock.capturedURL, expectedURL)
		}
	})

	t.Run("special characters in reference", func(t *testing.T) {
		mockData := NewTestTransaction()
		mock := &mockRequester{responseData: mockData}
		tn := New(mock)
		trRef := "ref_123&test=value"

		_, err := tn.Verify(context.Background(), trRef)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expectedURL := fmt.Sprintf(VERIFY_TRANSACTION, trRef)
		if mock.capturedURL != expectedURL {
			t.Errorf("URL = %q, want %q", mock.capturedURL, expectedURL)
		}
	})
}
