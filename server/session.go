package server

import (
	"github.com/bruno-nascimento/go-http-hello-world/model"
)

// SessionStruct lero-lero
type SessionStruct struct {
	New     model.NewStruct     `json:"new"`
	Key     model.KeyStruct     `json:"key"`
	Message model.MessageStruct `json:"message"`
}

// SessionMap lero-lero
var SessionMap = make(map[string]SessionStruct)

// NewSession lero-lero
func NewSession(newStruct model.NewStruct, keyStruct model.KeyStruct) {
	SessionMap[newStruct.Name] = SessionStruct{newStruct, keyStruct, model.MessageStruct{}}
}
