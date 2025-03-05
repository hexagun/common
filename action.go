package common

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

type JoinAction struct {
	Type     ActionType
	GameId   int
	PlayerId int
}

// functional
func NewJoinAction(gameId, playerId int) *JoinAction {
	return &JoinAction{
		Type:     Join,
		GameId:   gameId,
		PlayerId: playerId,
	}
}

type StartAction struct {
	Type     ActionType
	GameId   int
	PlayerId int
}

func NewStartAction(gameId, playerId int) *StartAction {
	return &StartAction{
		Type:     Start,
		GameId:   gameId,
		PlayerId: playerId,
	}
}

type PlayerMovePayload struct {
	Row, Col int
}

type PlayerMoveAction struct {
	Type     ActionType
	GameId   int
	PlayerId int
	Payload  PlayerMovePayload
}

func NewPlayerMoveAction(gameId, playerId int, payload PlayerMovePayload) *PlayerMoveAction {
	return &PlayerMoveAction{
		Type:     Move,
		GameId:   gameId,
		PlayerId: playerId,
		Payload:  payload,
	}
}

type GameStateUpdatePayload struct {
	Board    [3][3]string
	NextTurn string
	Winner   string
}

type GameStateUpdateAction struct {
	Type     ActionType
	GameId   int
	PlayerId int
	Payload  GameStateUpdatePayload
}

func NewGameStateUpdateAction(gameId, playerId int, payload GameStateUpdatePayload) *GameStateUpdateAction {
	return &GameStateUpdateAction{
		Type:     GameStateUpdate,
		GameId:   gameId,
		PlayerId: playerId,
		Payload:  payload,
	}
}

type GameOverPayload struct {
	Winner string
	Board  [3][3]string
}

type GameOverAction struct {
	Type    ActionType
	GameId  int
	Payload GameOverPayload
}

func NewGameOverAction(gameId int, payload GameOverPayload) *GameOverAction {
	return &GameOverAction{
		Type:    GameOver,
		GameId:  gameId,
		Payload: payload,
	}
}

type ErrorPayload struct {
	Reason string
}

type ErrorAction struct {
	Type    ActionType
	Message string
	Payload ErrorPayload
}

func NewErrorAction(message string, payload ErrorPayload) *ErrorAction {
	return &ErrorAction{
		Type:    Error,
		Message: message,
		Payload: payload,
	}
}
