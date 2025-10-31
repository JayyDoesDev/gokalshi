package fcm

import "github.com/jayydoesdev/gokalshi"

func GetFCMOrders(keyID, keyPem string, q FCMOrdersQuery) ([]byte, error) {
	return gokalshi.Request[[]byte]("/fcm/orders", "GET", keyID, keyPem, true, q.ToMap())
}

func GetFCMPositions(keyID, keyPem string, q FCMPositionsQuery) ([]byte, error) {
	return gokalshi.Request[[]byte]("/fcm/positions", "GET", keyID, keyPem, true, q.ToMap())
}
