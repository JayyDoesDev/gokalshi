package incentiveprograms

import "github.com/jayydoesdev/gokalshi"

func GetIncentives(keyID, keyPem string, q IncentiveProgramsQuery) ([]byte, error) {
	return gokalshi.Request[[]byte]("/incentive_programs", "GET", keyID, keyPem, true, q.ToMap())
}
