package service

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	configs "github.com/4lerman/e_com/common/config"
	"github.com/4lerman/e_com/common/utils"
	"github.com/4lerman/e_com/payment/types"
)

func GetPaymentToken() (*types.TokenResponse, error) {
	tokenUrl := configs.Envs.Token_Url

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// taken test fields from website
	writer.WriteField("grant_type", "client_credentials")
	writer.WriteField("scope", "webapi usermanagement email_send verification statement statistics payment")
	writer.WriteField("client_id", "test")
	writer.WriteField("client_secret", "yF587AV9Ms94qN2QShFzVR3vFnWkhjbAK3sG")
	writer.WriteField("invoiceID", "938290483292")
	writer.WriteField("amount", "100")
	writer.WriteField("currency", "KZT")
	writer.WriteField("terminal", "67e34d63-102f-4bd1-898e-370781d0074d")

	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, tokenUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var token types.TokenResponse
	// read response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// check error status
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error when reading response body. Status: " + resp.Status)
	}

	err = json.Unmarshal(data, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func GetPublicKey() (*rsa.PublicKey, error) {
	publicKeyURL := "https://testepay.homebank.kz/api/public.rsa"
	resp, err := http.Get(publicKeyURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(body)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to parse RSA public key")
	}
	return rsaPublicKey, nil
}

func encryptData() (string, error) {
	publicKey, err := GetPublicKey()
	if err != nil {
		return "", err
	}
	data := map[string]string{
		"hpan":       "4405639704015096",
		"expDate":    "0125",
		"cvc":        "815",
		"terminalId": "67e34d63-102f-4bd1-898e-370781d0074d",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, jsonData)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

func MakePayment() (*types.PaymentResponse, error) {
	paymentUrl := configs.Envs.Make_Payment_Url
	paymentToken, err := GetPaymentToken()
	fmt.Println("Payment token", paymentToken.AccessToken)

	if err != nil {
		return nil, fmt.Errorf("failed to get payment token: %v", err)
	}

	encryptedData, err := encryptData()
	if err != nil {
		return nil, fmt.Errorf("error when encrypting key: %v", err)
	}

	body := map[string]interface{}{
		"amount":          100,
		"currency":        "KZT",
		"name":            "JON JONSON",
		"cryptogram":      encryptedData,
		"invoiceID":       "938290483292",
		"invoiceIdAlt":    "8564546",
		"description":     "test payment",
		"accountID":       "uuid000001",
		"email":           "jj@example.com",
		"phone":           "77777777777",
		"cardSave":        true,
		"data":            `{\"statement\":{\"name\":\"Arman     Ali\",\"invoiceID\":\"80000016\"}}`,
		"postLink":        "https://testmerchant/order/1123",
		"failurePostLink": "https://testmerchant/order/1123/fail",
	}

	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, paymentUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+paymentToken.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	var payload types.PaymentResponse
	if err := utils.ResParseJSON(resp, &payload); err != nil {
		return nil, fmt.Errorf("error when parsing json")
	}

	return &payload, nil
}
