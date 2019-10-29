package main

import (
	"fmt"
	"math/rand"
	"net/url"

	signaler "github.com/y00rb/pylon-signaler/server"
)

// MySignalerServer is server to run as signaler in webrtc
type MySignalerServer struct {
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// AuthenticateRequest is set auth in signaler server
func (m *MySignalerServer) AuthenticateRequest(params url.Values) (apiKey, room, sessionKey string, ok bool) {
	return "ABC", "ABC", randSeq(16), true
}

// OnClientMessage is handle message for client
func (m *MySignalerServer) OnClientMessage(ApiKey, Room, SessionKey string, raw []byte) {
}

func main() {

	fmt.Println(signaler.Start(&MySignalerServer{}, "8383"))

}
