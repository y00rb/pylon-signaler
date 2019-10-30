package main

import (
	"fmt"
	"log"
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
	roomId := params.Get("roomID")
	if roomId == "" {
		b := make([]byte, 16)
		_, err := rand.Read(b)
		if err != nil {
			log.Fatal(err)
		}
		roomId = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	}
	return "ABC", roomId, randSeq(16), true
}

// OnClientMessage is handle message for client
func (m *MySignalerServer) OnClientMessage(ApiKey, Room, SessionKey string, raw []byte) {
}

func main() {

	fmt.Println(signaler.Start(&MySignalerServer{}, "8383"))

}
