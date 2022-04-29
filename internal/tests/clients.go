package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/jerensl/api.jerenslensun.com/internal/tests/client"
	"github.com/stretchr/testify/require"
)


type HTTPClient struct {
	client *client.ClientWithResponses
}

func NewHttpClient(t *testing.T) HTTPClient {
	addr := os.Getenv("HTTP_ADDR")
	ok := WaitForPort(addr)
	require.True(t, ok, "App HTTP Timed out")

	url := fmt.Sprintf("http://%v/api", addr)

	clients, err := client.NewClientWithResponses(
		url,
	)
	require.NoError(t, err)

	return HTTPClient{
		client: clients,
	}
}

func (a HTTPClient) NotSubscriberStatus(t *testing.T, token string) {
	var subscriber client.Status 
	response, err := a.client.SubscriberStatus(context.Background(), client.SubscriberStatusJSONRequestBody{
		Token: token,
	})
	require.NoError(t, err)

	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&subscriber)
	require.Equal(t, false, subscriber.Status)
	require.Equal(t, http.StatusOK, response.StatusCode)
}

func (a HTTPClient) AlreadySubscriberStatus(t *testing.T, token string) {
	var subscriber client.Status 
	response, err := a.client.SubscriberStatus(context.Background(), client.SubscriberStatusJSONRequestBody{
		Token: token,
	})
	require.NoError(t, err)

	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&subscriber)
	require.Equal(t, true, subscriber.Status)
	require.Equal(t, http.StatusOK, response.StatusCode)
}

func (a HTTPClient) SubscibeNotification(t *testing.T, token string) {
	response, err := a.client.SubscribeNotification(context.Background(), client.SubscribeNotificationJSONRequestBody{
		Token: token,
	})
	require.NoError(t, err)

	require.Equal(t, http.StatusCreated, response.StatusCode)
}

func (a HTTPClient) AlreadySubscibeNotification(t *testing.T, token string) {
	response, err := a.client.SubscribeNotification(context.Background(), client.SubscribeNotificationJSONRequestBody{
		Token: token,
	})
	require.NoError(t, err)

	require.Equal(t, http.StatusInternalServerError, response.StatusCode)
}


func (a HTTPClient) UnsubscibeNotification(t *testing.T, token string) {
	response, err := a.client.UnsubscribeNotification(context.Background(), client.UnsubscribeNotificationJSONRequestBody{
		Token: token,
	})
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, response.StatusCode)
}

func (a HTTPClient) UnsubsciberFromSubscriberNotExist(t *testing.T, token string) {
	response, err := a.client.UnsubscribeNotification(context.Background(), client.UnsubscribeNotificationJSONRequestBody{
		Token: token,
	})
	require.NoError(t, err)

	require.Equal(t, http.StatusInternalServerError, response.StatusCode)
}