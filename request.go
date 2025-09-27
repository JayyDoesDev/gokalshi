package gokalshi

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

const baseURL = "https://api.elections.kalshi.com/trade-api/v2"

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("invalid key file")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key.(*rsa.PrivateKey), nil
}

func SignPSSText(key *rsa.PrivateKey, msg string) (string, error) {
	hash := sha256.Sum256([]byte(msg))
	sig, err := rsa.SignPSS(rand.Reader, key, crypto.SHA256, hash[:], nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sig), nil
}

func Request[T any](route, method, keyID, keyPath string, params ...interface{}) (T, error) {
	var zero T

	data, err := json.Marshal(params)
	if err != nil {
		return zero, err
	}

	ts := time.Now().UnixMilli()
	timestamp := []byte(string(rune(ts)))
	msg := string(timestamp) + method + route

	priv, err := LoadPrivateKey(keyPath)
	if err != nil {
		return zero, err
	}

	sig, err := SignPSSText(priv, msg)
	if err != nil {
		return zero, err
	}

	req, err := http.NewRequest(method, baseURL+route, bytes.NewBuffer(data))
	if err != nil {
		return zero, err
	}

	req.Header.Set("KALSHI-ACCESS-KEY", keyID)
	req.Header.Set("KALSHI-ACCESS-SIGNATURE", sig)
	req.Header.Set("KALSHI-ACCESS-TIMESTAMP", string(timestamp))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return zero, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return zero, errors.New("non-2xx status: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return zero, err
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return zero, err
	}

	return result, nil
}
