package reniecsunatclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPersonByDNI(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/persons/71101328" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"success":true,"statusCode":200,"message":"ok","data":{"dni":"71101328","firstNames":"PATRICK","lastNames":"VIDAL MORI"},"path":"` + r.URL.Path + `","timestamp":"2026-06-11T00:00:00Z"}`))
	}))
	defer server.Close()

	client := NewWithBaseURL(server.URL)
	person, err := client.GetPersonByDNI(context.Background(), "71101328")
	if err != nil {
		t.Fatalf("GetPersonByDNI returned error: %v", err)
	}

	if person.DNI != "71101328" {
		t.Fatalf("unexpected dni: %s", person.DNI)
	}
}

func TestGetCompanyByRUCError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"success":false,"statusCode":404,"message":"company not found","error":{"code":"COMPANY_NOT_FOUND","message":"company not found"},"path":"` + r.URL.Path + `","timestamp":"2026-06-11T00:00:00Z"}`))
	}))
	defer server.Close()

	client := NewWithBaseURL(server.URL)
	_, err := client.GetCompanyByRUC(context.Background(), "20604633070")
	if err == nil {
		t.Fatal("expected error")
	}

	apiErr, ok := err.(*Error)
	if !ok {
		t.Fatalf("expected *Error, got %T", err)
	}
	if apiErr.Code != "COMPANY_NOT_FOUND" {
		t.Fatalf("unexpected error code: %s", apiErr.Code)
	}
}
