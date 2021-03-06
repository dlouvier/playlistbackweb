package spotify

import (
	"net/http"
	"os"
)

// GetRequests ssssss
func GetRequests(url string) string {

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN"))
	resp := httpClient(req)

	if err != nil {
		panic(err)
	}

	return string(resp)
}
