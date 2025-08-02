package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClientCaching(t *testing.T) {
	// Create a test server that counts requests
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Mock response for location-area endpoint
		response := `{
			"count": 1054,
			"next": "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
			"previous": null,
			"results": [
				{
					"name": "test-location-1",
					"url": "https://pokeapi.co/api/v2/location-area/1/"
				},
				{
					"name": "test-location-2",
					"url": "https://pokeapi.co/api/v2/location-area/2/"
				}
			]
		}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with short cache interval for testing
	client := NewClient(5*time.Second, 50*time.Millisecond)

	// Note: In a real implementation, we would make baseURL configurable for testing

	// Test URL that points to our test server
	testURL := server.URL

	// First request - should hit the server
	_, err := client.ListLocations(&testURL)
	if err != nil {
		t.Fatalf("First request failed: %v", err)
	}

	if requestCount != 1 {
		t.Errorf("Expected 1 request after first call, got %d", requestCount)
	}

	// Second request with same URL - should use cache
	_, err = client.ListLocations(&testURL)
	if err != nil {
		t.Fatalf("Second request failed: %v", err)
	}

	if requestCount != 1 {
		t.Errorf("Expected 1 request after second call (cache hit), got %d", requestCount)
	}

	// Wait for cache to expire
	time.Sleep(100 * time.Millisecond)

	// Third request - cache should be expired, so should hit server again
	_, err = client.ListLocations(&testURL)
	if err != nil {
		t.Fatalf("Third request failed: %v", err)
	}

	if requestCount != 2 {
		t.Errorf("Expected 2 requests after cache expiry, got %d", requestCount)
	}
}

func TestClientCacheDifferentURLs(t *testing.T) {
	// Create a test server
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := `{
			"count": 1054,
			"next": null,
			"previous": null,
			"results": []
		}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client
	client := NewClient(5*time.Second, 5*time.Minute)

	// Test with different URLs - each should result in a separate request
	testURL1 := server.URL + "/page1"
	testURL2 := server.URL + "/page2"

	// First URL
	_, err := client.ListLocations(&testURL1)
	if err != nil {
		t.Fatalf("Request to URL1 failed: %v", err)
	}

	// Second URL (different from first)
	_, err = client.ListLocations(&testURL2)
	if err != nil {
		t.Fatalf("Request to URL2 failed: %v", err)
	}

	if requestCount != 2 {
		t.Errorf("Expected 2 requests for different URLs, got %d", requestCount)
	}

	// Repeat first URL - should use cache
	_, err = client.ListLocations(&testURL1)
	if err != nil {
		t.Fatalf("Repeat request to URL1 failed: %v", err)
	}

	if requestCount != 2 {
		t.Errorf("Expected 2 requests after cache hit, got %d", requestCount)
	}
}
