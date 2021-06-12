package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Token struct {
	AccessToken string `json:"access_token"`
}

func main() {
	var tok Token

	scopeDriveReadOnly := "https://www.googleapis.com/auth/drive.readonly"
	endpoint := "https://oauth2.googleapis.com/token"

	clientIdPtr := flag.String("client", "", "Client Id")
	clientSecretPtr := flag.String("secret", "", "Client Secret")
	refreshTokenPtr := flag.String("refresh", "", "Refresh token")
	flag.Parse()

	clientId := *clientIdPtr
	clientSecret := *clientSecretPtr
	refreshToken := *refreshTokenPtr

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("refresh_token", refreshToken)
	data.Set("scope", scopeDriveReadOnly)

	body := strings.NewReader(data.Encode())

	client := http.DefaultClient
	req, _ := http.NewRequest("POST", endpoint, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("ERROR Request=%s", err)
		return
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&tok)

	fmt.Print(tok.AccessToken)
}
