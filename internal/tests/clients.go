package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/jerensl/jerens-web-api/internal/tests/client"
	"github.com/stretchr/testify/require"
)


type AppHTTPClient struct {
	client *client.ClientWithResponses
}

func NewAppHttpClient(t *testing.T) AppHTTPClient {
	addr := os.Getenv("APP_HTTP_ADDR")
	ok := WaitForPort(addr)
	require.True(t, ok, "App HTTP Timed out")

	url := fmt.Sprintf("http://%v/api", addr)

	clients, err := client.NewClientWithResponses(
		url,
	)
	require.NoError(t, err)

	return AppHTTPClient{
		client: clients,
	}
}

func (a AppHTTPClient) SubscriberStatus(t *testing.T, token string) {
	var subscriber client.Status 
	response, err := a.client.SubscriberStatus(context.Background(), client.SubscriberStatusJSONRequestBody{
		Token: token,
	})
	defer response.Body.Close()
	
	require.NoError(t, err)
	require.NoError(t, response.Body.Close())

	json.NewDecoder(response.Body).Decode(subscriber)
	require.Equal(t, false, subscriber.Status)
	require.Equal(t, http.StatusOK, response.StatusCode)
}