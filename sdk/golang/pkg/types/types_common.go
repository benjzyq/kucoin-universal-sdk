package types

import (
	json "github.com/goccy/go-json"
	"fmt"
)

// RestRateLimit represents the rate limiting information for a REST API.
type RestRateLimit struct {
	Limit     int64 // Total resource pool quota
	Remaining int64 // Resource pool remaining quota
	Reset     int64 // Resource pool quota reset countdown (milliseconds)
}

// RestResponse represents a generic response from the REST API.
type RestResponse struct {
	RateLimit *RestRateLimit
	Code      string          `json:"code"`          // Response code
	Data      json.RawMessage `json:"data"`          // Response data(raw json payload)
	Message   string          `json:"msg,omitempty"` // Response message
}

func (resp *RestResponse) Error() error {
	if resp.Code == CodeSuccess {
		return nil
	}
	return fmt.Errorf("server return error, code:{%s}, msg: {%s}", resp.Code, resp.Message)
}

// WsMessage represents a message between the WebSocket client and server.
type WsMessage struct {
	Id             string          `json:"id"`             // A unique identifier for the message.
	Type           string          `json:"type"`           // The type of the message (e.g., WelcomeMessage).
	Sn             int64           `json:"sn"`             // Sequence number to track the order of messages.
	Topic          string          `json:"topic"`          // The topic or channel the message is associated with.
	Subject        string          `json:"subject"`        // The subject of the message, providing additional context.
	PrivateChannel bool            `json:"privateChannel"` // Indicates if the message belongs to a private channel.
	Response       bool            `json:"response"`       // Specifies whether the message is a response.
	RawData        json.RawMessage `json:"data"`           // Raw JSON payload containing additional message data.
}
