package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	counter    = 0
)

// Current time + Random number + IP address
func GenerateCode() string {
	t := time.Now().UnixNano() % 10000
	ret := fmt.Sprintf("%c%c%c%c%c%c",
		characters[counter],
		characters[rand.Intn(62)],
		characters[t%100%62],
		characters[rand.Intn(62)],
		characters[t/100%100%62],
		characters[rand.Intn(62)])
	counter = (counter + 1) % 62
	return ret
}
