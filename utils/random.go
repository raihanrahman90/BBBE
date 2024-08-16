package utils

import (
	"encoding/json"
	"log"
	"math/rand"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func LogObject(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("error")
	}
	log.Printf("response data " + string(jsonData))
}
