package audit

import (
	"errors"
	"time"
)

const(
	ENTITY_USER = "USER"
	ENTITY_LIST = "LIST"
	ENTITY_ITEM = "ITEM"

	ACTION_CREATE = "CREATE"
	ACTION_UPDATE = "UPDATE"
	ACTION_DELETE = "DELETE"
	ACTION_GET      = "GET"
	ACTION_LOGIN = "LOGIN"
	ACTION_REGISTER = "REGISTER"
)

var (
	entities = map[string]LogRequest_Entities{
		ENTITY_USER: LogRequest_USER,
		ENTITY_LIST: LogRequest_LIST,
	}

	actions = map[string]LogRequest_Actions{
		ACTION_CREATE:   LogRequest_CREATE,
		ACTION_UPDATE:   LogRequest_UPDATE,
		ACTION_GET:      LogRequest_GET,
		ACTION_DELETE:   LogRequest_DELETE,
		ACTION_REGISTER: LogRequest_REGISTER,
		ACTION_LOGIN:    LogRequest_LOGIN,
	}
)

type LogItem struct {
	Entity    string    `bson:"entity"`
	Action    string    `bson:"action"`
	EntityID  int64     `bson:"entity_id"`
	Timestamp time.Time `bson:"timestamp"`
}

func ToPbEntity(entity string) (LogRequest_Entities, error) {
	val, ex := entities[entity]
	if !ex {
		return 0, errors.New("invalid entity")
	}

	return val, nil
}

func ToPbAction(action string) (LogRequest_Actions, error) {
	val, ex := actions[action]
	if !ex {
		return 0, errors.New("invalid action")
	}

	return val, nil
}

