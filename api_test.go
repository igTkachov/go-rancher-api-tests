package tests

import (
    "net/http"
    "testing"
    "strings"
	"crypto/tls"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
    url := "https://localhost/v3-public/localProviders/local?action=login"
	payload := `{"description":"UI session","responseType":"cookie","username":"admin","password":"Test!2345678"}`
    
	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")
    req.Header.Set("X-Api-Csrf", "88dca93b21a18e6194fb5b3f69eba38d")

    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        },
    }
    resp, err := client.Do(req)
    assert.NoError(t, err, "Request error")
    defer resp.Body.Close()

	// Check status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)	
	// Check for presence of R_SESS cookie and that the value starts with "token-"
	var sessionCookie *http.Cookie
    for _, cookie := range resp.Cookies() {
        if cookie.Name == "R_SESS" {
            sessionCookie = cookie
            break
        }
    }
    assert.NotNil(t, sessionCookie, "R_SESS cookie isn't present in response")
    assert.True(t, strings.HasPrefix(sessionCookie.Value, "token-"), "Wrong R_SESS cookie value")
}