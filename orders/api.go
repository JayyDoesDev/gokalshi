package orders

import "github.com/jayydoesdev/gokalshi"

func GetOrderGroups(keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/portfolio/order_groups", "GET", keyID, keyPem, true, map[string]string{})
}
