package gokalshi_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/jayydoesdev/gokalshi/market"
	"github.com/joho/godotenv"
)

func TestGetSeriesList(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Fatal("failed to load .env file:", err)
	}

	key := os.Getenv("KALSHI_KEY_ID")
	keyPem := os.Getenv("KALSHI_PRIVATE_KEY")

	if key == "" || keyPem == "" {
		t.Fatal("KALSHI_KEY_ID or KALSHI_PRIVATE_KEY not set in environment")
	}

	resp, err := market.GetSeriesList(key, keyPem, market.SeriesQuery{
		Category: "Sports",
	})
	if err != nil {
		t.Fatalf("GetSeries failed: %v", err)
	}

	var result market.SeriesList
	if err := json.Unmarshal(resp, &result); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if len(result.Series) == 0 {
		t.Errorf("Test failed! Series not found! %v", result.Series)
	}

	t.Log(result.Series[0].Title)
}
