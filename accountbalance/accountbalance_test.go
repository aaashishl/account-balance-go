package accountbalance

import (	
	"encoding/json"
	"net/http"
	"testing"
	"net/http/httptest"
)

func TestAddressIsRequired(t *testing.T) {
	baseURL := "https://backend.testnet.alephium.org"
	response, err := GetAccountBalance(baseURL,"")
	want := "Address is required"
	if err != nil && response == (AccountBalance{}) && err.Error() != want {
		t.Fatalf("Got %v, want %v", err.Error(), want)
	}
}

func TestResponseReturn(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        response := AccountBalance{
            Balance:       "999",
            LockedBalance: "0",
        }
        json.NewEncoder(w).Encode(response)
    }))
    defer server.Close()

    originalTransport := http.DefaultTransport
    defer func() { http.DefaultTransport = originalTransport }()
    http.DefaultTransport = server.Client().Transport
	
    response, err := GetAccountBalance(server.URL,"1DrDyTr9RpRsQnDnXo2YRiPzPW4ooHX5LLoqXrqfMrpQH")
    if err != nil {
        t.Fatalf("Got %v, want nil", err)
    }
    expectedResponse := AccountBalance{Balance: "999", LockedBalance: "0"}
    if response != expectedResponse {
        t.Fatalf("Got %v, want %v", response, expectedResponse)
    }
}
func TestRequestFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        response := map[string]string{
			"error": "Internal Server Error",
		}
        json.NewEncoder(w).Encode(response)
    }))
    defer server.Close()

    originalTransport := http.DefaultTransport
    defer func() { http.DefaultTransport = originalTransport }()
    http.DefaultTransport = server.Client().Transport
	
    response, err := GetAccountBalance(server.URL,"1DrDyTr9RpRsQnDnXo2YRiPzPW4ooHX5LLoqXrqfMrpQH")
    if err != nil {
        t.Fatalf("Got %v, want nil", err)
    }
    expectedResponse := AccountBalance{}
    if response != expectedResponse {
        t.Fatalf("Got %v, want %v", response, expectedResponse)
    }
}