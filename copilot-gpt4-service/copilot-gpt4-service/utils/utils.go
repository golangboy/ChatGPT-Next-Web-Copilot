package utils

import (
	"copilot-gpt4-service/cache"
	"copilot-gpt4-service/config"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Set the Authorization in the cache.
func setAuthorizationToCache(copilotToken string, authorization cache.Authorization) {
	cache.CacheInstance.Set(copilotToken, authorization)
}

// Obtain the Authorization from the cache.
func getAuthorizationFromCache(copilotToken string) *cache.Authorization {
	extraTime := rand.Intn(600) + 300
	if authorization, ok := cache.CacheInstance.Get(copilotToken); ok {
		if authorization.ExpiresAt > time.Now().Unix()+int64(extraTime) {
			return &authorization
		}
	}
	return &cache.Authorization{}
}

// When obtaining the Authorization, first attempt to retrieve it from the cache. If it is not available in the cache, retrieve it through an HTTP request and then set it in the cache.
func GetAuthorizationFromToken(copilotToken string) (string, int, string) {
	authorization := getAuthorizationFromCache(copilotToken)
	if authorization.Token == "" {
		getAuthorizationUrl := "https://api.github.com/copilot_internal/v2/token"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", getAuthorizationUrl, nil)
		req.Header.Set("Authorization", "token "+copilotToken)
		response, err := client.Do(req)
		if err != nil {
			return "", http.StatusInternalServerError, err.Error()
		}
		if response.StatusCode != 200 {
			body, _ := io.ReadAll(response.Body)
			return "", response.StatusCode, string(body)
		}
		defer response.Body.Close()

		body, _ := io.ReadAll(response.Body)

		newAuthorization := &cache.Authorization{}
		if err = json.Unmarshal(body, &newAuthorization); err != nil {
			fmt.Println("err", err)
		}
		authorization.Token = newAuthorization.Token
		setAuthorizationToCache(copilotToken, *newAuthorization)
	}
	return authorization.Token, http.StatusOK, ""
}

// Retrieve the GitHub Copilot Plugin Token from the request header.
func GetAuthorization(c *gin.Context) (string, bool) {
	if config.ConfigInstance.CopilotToken == "" {
		copilotToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if copilotToken == "" {
			return "", false
		} else {
			return copilotToken, true
		}
	} else {
		return config.ConfigInstance.CopilotToken, true
	}
}
