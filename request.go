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
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const baseURL = "https://api.elections.kalshi.com"
const apiPrefix = "/trade-api/v2"

func LoadPrivateKeyFromString(pemString string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemString))
	if block == nil {
		return nil, errors.New("invalid PEM key")
	}
	if k, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		return k.(*rsa.PrivateKey), nil
	}
	if k, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return k, nil
	}
	return nil, errors.New("failed to parse private key")
}

func SignPSSText(key *rsa.PrivateKey, msg string) (string, error) {
	sum := sha256.Sum256([]byte(msg))
	sig, err := rsa.SignPSS(rand.Reader, key, crypto.SHA256, sum[:], nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sig), nil
}

func Request[T any](route, method, keyID, keyPem string, useAuth bool, query map[string]string, params ...interface{}) (T, error) {
	var zero T

	if !strings.HasPrefix(route, "/") {
		route = "/" + route
	}

	u, err := url.Parse(baseURL + apiPrefix + route)
	if err != nil {
		return zero, err
	}

	if len(query) > 0 {
		q := u.Query()
		for k, v := range query {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	fullPath := apiPrefix + route
	if len(query) > 0 {
		fullPath += "?" + u.RawQuery
	}

	method = strings.ToUpper(method)
	tsStr := strconv.FormatInt(time.Now().UnixMilli(), 10)
	msg := tsStr + method + fullPath

	var sig string
	if useAuth {
		priv, err := LoadPrivateKeyFromString(keyPem)
		if err != nil {
			return zero, err
		}
		sig, err = SignPSSText(priv, msg)
		if err != nil {
			return zero, err
		}
	}

	fmt.Println("==== KALSHI DEBUG ====")
	fmt.Println("Base URL:     ", baseURL)
	fmt.Println("Full Path:    ", fullPath)
	fmt.Println("Full URL:     ", u.String())
	fmt.Println("HTTP Method:  ", method)
	fmt.Println("Timestamp:    ", tsStr)
	fmt.Println("Signing Msg:  ", msg)
	fmt.Println("======================")

	var reqBody io.Reader
	if len(params) > 0 {
		data, err := json.Marshal(params[0])
		if err != nil {
			return zero, err
		}
		reqBody = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, u.String(), reqBody)
	if err != nil {
		return zero, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "gokalshi")
	if useAuth {
		req.Header.Set("KALSHI-ACCESS-KEY", keyID)
		req.Header.Set("KALSHI-ACCESS-SIGNATURE", sig)
		req.Header.Set("KALSHI-ACCESS-TIMESTAMP", tsStr)
		req.Header.Set("Content-Type", "application/json")
	}

	fmt.Println("Headers:")
	for k, v := range req.Header {
		fmt.Printf("  %s: %s\n", k, v[0])
	}
	fmt.Println("======================")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return zero, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return zero, err
	}

	if _, ok := any(zero).([]byte); ok {
		return any(body).(T), nil
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return zero, fmt.Errorf("non-2xx status: %s body: %s", resp.Status, string(body))
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return zero, err
	}

	return result, nil
}
