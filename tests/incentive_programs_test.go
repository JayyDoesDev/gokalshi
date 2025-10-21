package gokalshi

import (
	"encoding/json"
	"os"
	"testing"

	incentiveprograms "github.com/jayydoesdev/gokalshi/incentive_programs"
	"github.com/joho/godotenv"
)

func TestIncentives(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Fatal("failed to load .env file:", err)
	}

	key := os.Getenv("KALSHI_KEY_ID")
	keyPem := os.Getenv("KALSHI_PRIVATE_KEY")

	if key == "" || keyPem == "" {
		t.Fatal("KALSHI_KEY_ID or KALSHI_PRIVATE_KEY not set in environment")
	}

	resp, err := incentiveprograms.GetIncentives(key, keyPem, incentiveprograms.IncentiveProgramsQuery{Limit: 100})
	if err != nil {
		t.Fatalf("GetIncentives: failedL %v", err)
	}

	var result incentiveprograms.IncentivePrograms
	if err := json.Unmarshal(resp, &result); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if len(result.IncentivePrograms) == 0 {
		t.Errorf("Test failed! Incentive programs not found! %v", result.IncentivePrograms)
	}

	t.Log(result.IncentivePrograms[0].SeriesTicker)
}
