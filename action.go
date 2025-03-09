package common

import "fmt"

type ActionType int

const (
	GameStateUpdate ActionType = iota
	GameOver
	Join
	Leave
	Start
	Move
	Error
)

func (a ActionType) String() string {
	switch a {
	case GameStateUpdate:
		return "gamestateupdate"
	case GameOver:
		return "gameover"
	case Join:
		return "join"
	case Leave:
		return "leave"
	case Start:
		return "start"
	case Move:
		return "move"
	case Error:
		return "error"
	default:
		return fmt.Sprintf("%d", int(a))
	}
}

type ActionHeader struct {
	Type     ActionType
	GameId   int
	PlayerId int
}

type Action interface {
	GetHeader() ActionHeader
	GetPayload() interface{}
}

type JoinAction struct {
	header ActionHeader
}

// functional
func NewJoinAction(gameId, playerId int) *JoinAction {
	return &JoinAction{
		ActionHeader{
			Type:     Join,
			GameId:   gameId,
			PlayerId: playerId,
		},
	}
}

func (a *JoinAction) GetType() ActionHeader {
	return a.header
}

func (a *JoinAction) GetPayload() interface{} {
	return nil
}

type StartAction struct {
	header ActionHeader
}

func NewStartAction(gameId, playerId int) *StartAction {
	return &StartAction{
		ActionHeader{
			Type:     Join,
			GameId:   gameId,
			PlayerId: playerId,
		},
	}
}

func (a *StartAction) GetType() ActionHeader {
	return a.header
}

func (a *StartAction) GetPayload() interface{} {
	return nil
}

type PlayerMovePayload struct {
	Row, Col int
}

type PlayerMoveAction struct {
	header  ActionHeader
	Payload PlayerMovePayload
}

func NewPlayerMoveAction(gameId, playerId int, payload PlayerMovePayload) *PlayerMoveAction {
	return &PlayerMoveAction{
		header: ActionHeader{
			Type:     Join,
			GameId:   gameId,
			PlayerId: playerId,
		},

		Payload: payload,
	}
}

func (a *PlayerMoveAction) GetType() ActionHeader {
	return a.header
}

func (a *PlayerMoveAction) GetPayload() interface{} {
	return a.Payload
}

type GameStateUpdatePayload struct {
	Board    [3][3]string
	NextTurn string
	Winner   string
}

type GameStateUpdateAction struct {
	header  ActionHeader
	Payload GameStateUpdatePayload
}

func NewGameStateUpdateAction(gameId, playerId int, payload GameStateUpdatePayload) *GameStateUpdateAction {
	return &GameStateUpdateAction{
		header: ActionHeader{
			Type:     Join,
			GameId:   gameId,
			PlayerId: playerId,
		},
		Payload: payload,
	}
}

func (a *GameStateUpdateAction) GetType() ActionHeader {
	return a.header
}

func (a *GameStateUpdateAction) GetPayload() interface{} {
	return a.Payload
}

type GameOverPayload struct {
	Winner string
	Board  [3][3]string
}

type GameOverAction struct {
	header  ActionHeader
	Payload GameOverPayload
}

func NewGameOverAction(gameId int, payload GameOverPayload) *GameOverAction {
	return &GameOverAction{
		header: ActionHeader{
			Type:     GameOver,
			GameId:   gameId,
			PlayerId: 0,
		},
		Payload: payload,
	}
}

func (a *GameOverAction) GetType() ActionHeader {
	return a.header
}

func (a *GameOverAction) GetPayload() interface{} {
	return a.Payload
}

type ErrorPayload struct {
	Reason  string
	Message string
}

type ErrorAction struct {
	header  ActionHeader
	Payload ErrorPayload
}

func NewErrorAction(message string, reason string) *ErrorAction {
	return &ErrorAction{
		header: ActionHeader{
			Type:     Error,
			GameId:   0,
			PlayerId: 0,
		},
		Payload: ErrorPayload{
			Reason:  reason,
			Message: message,
		},
	}
}

func (a *ErrorAction) GetType() ActionHeader {
	return a.header
}

func (a *ErrorAction) GetPayload() interface{} {
	return a.Payload
}
