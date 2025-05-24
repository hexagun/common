package common

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type MessageDecoder struct {
}

func (g MessageDecoder) Decode(msg *IncomingMessage) Action {
	var action Action
	var gameId int
	var playerId int

	if msg.GameID != "" {
		id, err := strconv.Atoi(msg.GameID)
		if err != nil {
			// ... handle error
			panic(err)
		}
		gameId = id
	}

	if msg.PlayerID != "" {
		id, err := strconv.Atoi(msg.PlayerID)
		if err != nil {
			// ... handle error
			panic(err)
		}
		playerId = id
	}

	switch msg.Type {
	case "gamestateupdate":
		var gameStateUpdatePayload struct {
			Board    [3][3]string //`json:"Board"`
			NextTurn string       //`json:"NextTurn"`
			Winner   string       //`json:"Winner"`
		}

		if err := json.Unmarshal(msg.Payload, &gameStateUpdatePayload); err == nil {
			// fmt.Printf("Player made a move at row %d, col %d\n", movePayload.Row, movePayload.Col)
			action = NewGameStateUpdateAction(gameId, playerId, GameStateUpdatePayload{
				Board:    gameStateUpdatePayload.Board,
				NextTurn: gameStateUpdatePayload.NextTurn,
				Winner:   gameStateUpdatePayload.Winner,
			})
		}
	case "gameover":
		var gameOverPayload struct {
			Board  [3][3]string //`json:"Board"`
			Winner string       //`json:"Winner"`
		}

		if err := json.Unmarshal(msg.Payload, &gameOverPayload); err == nil {
			// fmt.Printf("Player made a move at row %d, col %d\n", movePayload.Row, movePayload.Col)
			action = NewGameOverAction(gameId, GameOverPayload{
				Board:  gameOverPayload.Board,
				Winner: gameOverPayload.Winner,
			})
		}
	//case "error":
	case "start":
		var startPayload struct {
			YourToken  string //`json:"yourToken"`
			OpponentID string //`json:"opponentId"`
			FirstTurn  string //`json:"firstTurn"` // could also be a player ID
		}

		if err := json.Unmarshal(msg.Payload, &startPayload); err == nil {
			// fmt.Printf("Player made a move at row %d, col %d\n", movePayload.Row, movePayload.Col)
			action = NewStartAction(gameId, playerId, StartPayload{
				YourToken:  startPayload.YourToken,
				OpponentID: startPayload.OpponentID,
				FirstTurn:  startPayload.FirstTurn,
			})
		}
	case "join":
		action = NewJoinAction(gameId, playerId)
	// case "leave":
	// 	action = NewLeaveAction(gameId, playerId)
	case "move":
		var movePayload struct {
			Row int `json:"row"`
			Col int `json:"col"`
		}
		if err := json.Unmarshal(msg.Payload, &movePayload); err == nil {
			fmt.Printf("Player made a move at row %d, col %d\n", movePayload.Row, movePayload.Col)
			action = NewPlayerMoveAction(gameId, playerId, PlayerMovePayload{
				Row: movePayload.Row,
				Col: movePayload.Col,
			})
		}

	default:
		fmt.Println("Not Decoding stuff")
	}

	// game.Dispatch(action)
	return action
}
