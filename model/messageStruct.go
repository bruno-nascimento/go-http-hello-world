package model

import (
	"math/rand"
	"time"
)

// MessageStruct lero lero
type MessageStruct struct {
	// Message lero lero
	Message string `json:"message"`
	// EncriptedMessage lero lero
	EncriptedMessage []byte `json:"encryptedMessage"`
	// DecodedMessage lero lero
	ReencodedMessage []byte `json:"reencodedMessage"`
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

// From lero lero
func (messageStruct MessageStruct) From(message string) MessageStruct {
	messageStruct.Message = message
	return messageStruct
}

// NewRandomMessage lero lero
func (messageStruct MessageStruct) NewRandomMessage(stringLength ...int) MessageStruct {
	if stringLength != nil && stringLength[0] != 0 {
		messageStruct.Message = randomString(stringLength[0])
	} else {
		messageStruct.Message = randomString(32)
	}
	return messageStruct
}
