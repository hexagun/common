package common

import (
	"encoding/json"
	"strconv"
)

type OutgoingMessage struct {
	Type     string      `json:"type"`
	GameID   string      `json:"gameId,omitempty"`
	PlayerID string      `json:"playerId,omitempty"`
	Payload  interface{} `json:"payload,omitempty"`
}

type IncomingMessage struct {
	Type     string          `json:"type"`
	GameID   string          `json:"gameId,omitempty"`
	PlayerID string          `json:"playerId,omitempty"`
	Payload  json.RawMessage `json:"payload,omitempty"`
}

func NewOutgoingMessage(aType ActionType, gameId, playerId int, payload interface{}) *JoinAction {
	return &OutgoingMessage{
		Type:     aType.String(),
		GameID:   strconv.Itoa(gameId),
		PlayerID: strconv.Itoa(playerId),
		Payload:  payload,
	}
}
