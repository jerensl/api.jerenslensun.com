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

func (h HTTPClient) NotSubscriberStatus(t *testing.T, token string) {
	var subscriber client.Status 
	response, err := h.client.SubscriberStatus(context.Background(), client.SubscriberStatusJSONRequestBody{
		Token: token,
	})
	require.NoError(t, err)

	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&subscriber)
	require.Equal(t, false, subscriber.Status)
	require.Equal(t, http.StatusOK, response.StatusCode)
}

func (h HTTPClient) AlreadySubscriberStatus(t *testing.T, token string) {
	var subscriber client.Status 
	response, err := h.client.SubscriberStatus(context.Background(), client.SubscriberStatusJSONRequestBody{
		Token: token,
	})
	require.NoError(t, err)

	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&subscriber)
	require.Equal(t, true, subscriber.Status)
	require.Equal(t, http.StatusOK, response.StatusCode)
}

func (h HTTPClient) SubscibeNotification(t *testing.T, token string) {
	response, err := h.client.SubscribeNotification(context.Background(), client.SubscribeNotificationJSONRequestBody{
		Token: token,
	})
	require.NoError(t, err)

	require.Equal(t, http.StatusCreated, response.StatusCode)
}

func (h HTTPClient) AlreadySubscibeNotification(t *testing.T, token string) {
	response, err := h.client.SubscribeNotification(context.Background(), client.SubscribeNotificationJSONRequestBody{
		Token: token,
	})
	require.NoError(t, err)

	require.Equal(t, http.StatusInternalServerError, response.StatusCode)
}


func (h HTTPClient) UnsubscibeNotification(t *testing.T, token string) {
	response, err := h.client.UnsubscribeNotification(context.Background(), client.UnsubscribeNotificationJSONRequestBody{
		Token: token,
	})
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, response.StatusCode)
}

func (h HTTPClient) UnsubsciberFromSubscriberNotExist(t *testing.T, token string) {
	response, err := h.client.UnsubscribeNotification(context.Background(), client.UnsubscribeNotificationJSONRequestBody{
		Token: token,
	})
	require.NoError(t, err)

	require.Equal(t, http.StatusInternalServerError, response.StatusCode)
}

func (h HTTPClient) SendNotification(t *testing.T, title, message string) {
	response, err := h.client.SendNotification(context.Background(), client.SendNotificationJSONRequestBody{
		Title: title,
		Message: message,
	}, client.RequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("X-API-KEY", os.Getenv("API_KEY"))
		return nil
	}))
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, response.StatusCode)
}

func (h HTTPClient) SendNotificationWithoutAuthz(t *testing.T, title, message string) {
	response, err := h.client.SendNotification(context.Background(), client.SendNotificationJSONRequestBody{
		Title: title,
		Message: message,
	})
	require.NoError(t, err)

	require.Equal(t, http.StatusUnauthorized, response.StatusCode)
}