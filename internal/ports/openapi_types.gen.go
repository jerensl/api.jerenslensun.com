// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package ports

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

// Message defines model for Message.
type Message struct {
	// Notification message
	Message string `json:"message"`

	// Notification title
	Title string `json:"title"`
}

// Status defines model for Status.
type Status struct {
	// Subscriber status
	IsActive bool `json:"isActive"`

	// Last updated date
	UpdatedAt int64 `json:"updatedAt"`
}

// Subscriber defines model for Subscriber.
type Subscriber struct {
	// Client Token
	TokenID string `json:"tokenID"`

	// Last updated date
	UpdatedAt int64 `json:"updatedAt"`
}

// SendNotificationJSONBody defines parameters for SendNotification.
type SendNotificationJSONBody Message

// SubscriberStatusJSONBody defines parameters for SubscriberStatus.
type SubscriberStatusJSONBody Subscriber

// SubscribeNotificationJSONBody defines parameters for SubscribeNotification.
type SubscribeNotificationJSONBody Subscriber

// UnsubscribeNotificationJSONBody defines parameters for UnsubscribeNotification.
type UnsubscribeNotificationJSONBody Subscriber

// SendNotificationJSONRequestBody defines body for SendNotification for application/json ContentType.
type SendNotificationJSONRequestBody SendNotificationJSONBody

// SubscriberStatusJSONRequestBody defines body for SubscriberStatus for application/json ContentType.
type SubscriberStatusJSONRequestBody SubscriberStatusJSONBody

// SubscribeNotificationJSONRequestBody defines body for SubscribeNotification for application/json ContentType.
type SubscribeNotificationJSONRequestBody SubscribeNotificationJSONBody

// UnsubscribeNotificationJSONRequestBody defines body for UnsubscribeNotification for application/json ContentType.
type UnsubscribeNotificationJSONRequestBody UnsubscribeNotificationJSONBody
