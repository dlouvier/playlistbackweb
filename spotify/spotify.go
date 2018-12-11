package spotify

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dlouvier/playlistback/models"
)

// AppConfig --
type AppConfig struct {
	ClientID     string `json:"ClientID"`
	ClientSecret string `json:"ClientSecret"`
	URLGetToken  string `json:"URLGetToken"`
}

// ConfigParser --
func ConfigParser() AppConfig {
	var appconfig AppConfig
	file, _ := ioutil.ReadFile("conf/spotify.json")
	err := json.Unmarshal(file, &appconfig)
	if err != nil {
		panic(err)
	}

	return appconfig
}

// CalculateSecret giving the clientid and clientsecret it calculates the token
func CalculateSecret(clientid string, clientsecret string) string {
	return base64.StdEncoding.EncodeToString([]byte(clientid + ":" + clientsecret))
}

// GetToken ssssss
func GetToken(url string, secret string) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte("grant_type=client_credentials")))
	req.Header.Set("Authorization", "Basic "+secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp := httpClient(req)

	var ans models.BodyResp

	err := json.Unmarshal(resp, &ans)
	if err != nil {
		panic(err)
	}

	os.Setenv("TOKEN", ans.AccessToken)
}

// LoginAPI --
func LoginAPI() {
	config := ConfigParser()
	secret := CalculateSecret(config.ClientID, config.ClientSecret)
	GetToken(config.URLGetToken, secret)
}

// GetTracks --
func GetTracks(playlistid string) string {
	return GetRequests("https://api.spotify.com/v1/playlists/" + playlistid + "/tracks")
}
