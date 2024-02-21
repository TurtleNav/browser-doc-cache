package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"os"
)

// browser Native messaging API sends and expects JSON
// serialzed messages in the format:
//
// uint32 (4 bytes) for JSON packet length
// n bytes of JSON encoded data

type EncodedMessage struct {
	Length  int    `json:"length"`
	Content string `json:"content"`
}

func getMessage() string {
	rawLength := make([]byte, 4)
	os.Stdin.Read(rawLength)
	messageLength := binary.LittleEndian.Uint32(rawLength)

	if messageLength == 0 {
		panic("error: message length is 0!")
	}

	rawMessage := make([]byte, messageLength)
	os.Stdin.Read(rawMessage)

	return string(rawMessage)
}

func encodeMessage(messageContent string) []byte {
	m := EncodedMessage{
		Length:  len(messageContent),
		Content: messageContent,
	}
	encodedContent, err := json.Marshal(m)
	if err != nil {
		panic("JSON marshal failed")
	}
	return encodedContent
}

func sendMessage(message string) {

	encodedContent := encodeMessage(message)

	buf := new(bytes.Buffer)
	msgLen := uint32(len(encodedContent))
	err := binary.Write(buf, binary.LittleEndian, msgLen)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(buf.Bytes())
	os.Stdout.Write(encodedContent)

}

func main() {
	for {
		receivedMessage := getMessage()
		if receivedMessage == "\"ping\"" {
			sendMessage("pong-go")
		}
	}
}
