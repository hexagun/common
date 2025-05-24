package common

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestDecodeJoin(t *testing.T) {

	gameId := 111
	playerId := 123
	md := MessageDecoder{}

	msg := IncomingMessage{
		Type:     "join",
		GameID:   strconv.Itoa(gameId),
		PlayerID: strconv.Itoa(playerId),
		// GameID   string          `json:"gameId,omitempty"`
		// PlayerID string          `json:"playerId,omitempty"`
		// Payload  json.RawMessage `json:"payload,omitempty"`
	}

	action := md.Decode(&msg)
	expectedAction := NewJoinAction(gameId, playerId)
	assert.Equal(t, action, expectedAction, "Failed to decode join message.")
}

func TestDecodeStart(t *testing.T) {

	gameId := 111
	playerId := 123
	opponentId := 456
	md := MessageDecoder{}

	payload := StartPayload{
		YourToken:  "x",
		OpponentID: strconv.Itoa(opponentId),
		FirstTurn:  strconv.Itoa(playerId),
	}
	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	msg := IncomingMessage{
		Type:     "start",
		GameID:   strconv.Itoa(gameId),
		PlayerID: strconv.Itoa(playerId),
		Payload:  data,
	}
	action := md.Decode(&msg)
	expectedAction := NewStartAction(gameId, playerId, payload)
	assert.Equal(t, action, expectedAction, "Failed to decode start message.")
}

func TestDecodeMove(t *testing.T) {

	gameId := 111
	playerId := 123
	md := MessageDecoder{}

	payload := PlayerMovePayload{
		Row: 1,
		Col: 1,
	}
	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	msg := IncomingMessage{
		Type:     "move",
		GameID:   strconv.Itoa(gameId),
		PlayerID: strconv.Itoa(playerId),
		Payload:  data,
	}
	action := md.Decode(&msg)
	expectedAction := NewPlayerMoveAction(gameId, playerId, payload)
	assert.Equal(t, action, expectedAction, "Failed to decode start message.")
}

func TestGameStateUpdateMove(t *testing.T) {

	gameId := 111
	playerId := 123
	md := MessageDecoder{}

	payload := GameStateUpdatePayload{
		Board:    [3][3]string{},
		NextTurn: strconv.Itoa(playerId),
		Winner:   strconv.Itoa(playerId),
	}
	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	msg := IncomingMessage{
		Type:     "gamestateupdate",
		GameID:   strconv.Itoa(gameId),
		PlayerID: strconv.Itoa(playerId),
		Payload:  data,
	}
	action := md.Decode(&msg)
	expectedAction := NewGameStateUpdateAction(gameId, playerId, payload)
	assert.Equal(t, action, expectedAction, "Failed to decode start message.")
}
