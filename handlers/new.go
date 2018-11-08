package handlers

import (
	"fmt"

	"encoding/base64"

	"github.com/bruno-nascimento/go-http-hello-world/crypto"
	"github.com/bruno-nascimento/go-http-hello-world/model"
	"github.com/bruno-nascimento/go-http-hello-world/server"
	"github.com/google/uuid"
	"github.com/kataras/iris"
)

// Ping lero lero
func Ping(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "pong",
	})
	result, _ := crypto.Encrypt([]byte("ping"), server.SessionMap["lero"].New.Key)
	fmt.Println(base64.StdEncoding.EncodeToString(result))
}

// New lero lero
func New(ctx iris.Context) {
	var newStruct model.NewStruct
	if err := ctx.ReadJSON(&newStruct); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	newStruct.Name = ctx.Params().Get("name")
	server.NewSession(newStruct, model.KeyStruct{uuid.New().String()})
	ctx.JSON(iris.Map{
		"key":  server.SessionMap[newStruct.Name].Key.UUID,
		"name": model.OwnName,
	})
}

// Message lero lero
func Message(ctx iris.Context) {
	name := ctx.Params().Get("name")
	var reqMap = make(map[string]string)
	if err := ctx.ReadJSON(&reqMap); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	fmt.Println(reqMap)
	var message model.MessageStruct
	message.EncriptedMessage = []byte(reqMap["message"])
	plain, _ := crypto.Decrypt([]byte(reqMap["message"]), server.SessionMap[name].New.Key)
	fmt.Println(">>> " + string(plain[:]))
	message.Message = string(plain)
	enc, _ := crypto.Encrypt([]byte(message.Message), server.SessionMap[name].New.Key)
	message.ReencodedMessage = enc
	session := server.SessionMap[name]
	session.Message = message
	ctx.JSON(iris.Map{
		"message": message.EncriptedMessage,
	})
	fmt.Println(server.SessionMap)
}
