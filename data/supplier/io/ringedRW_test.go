package io

import (
	"testing"
	"fmt"
)

func TestRandomRead(t *testing.T) {
	writer, data, cursor := NewRingedWriter(100, nil)
	reader := NewRingedReader(data, cursor)

	cash := make([]byte, 1)
	for i := 0; i < 200; i++ {
		n, err := reader.Read(cash)
		//t.Log(fmt.Sprintf("Random number read: %v", cash[0]))
		if n != 1 {
			t.Error("Can't read []byte elem of size 1")
		}
		if err != nil {
			t.Error("Error accured while random ringed read;", err.Error())
		}
	}

	cash = []byte{byte(1), byte(2), byte(3), byte(4), byte(5)}
	n, err := writer.Write(cash)
	if n != 5 {
		t.Error("Buffered writer wrote not all buffer")
	}
	if err != nil {
		t.Error("Error acured on ringed wright: ", err.Error())
	}
	readBuffer := make([]byte, 5)
	n, err = reader.Read(readBuffer)
	if len(cash) != n {
		t.Error("Read wrong size of bytes: (exp/got)", n, len(cash))
	}
	if err != nil {
		t.Error("Error acured on ringed read: ", err.Error())
	}

	for i := range cash {
		if cash[i] != readBuffer[i] {
			t.Error(fmt.Sprintf("read and wrote values are not equal: expected %d, have: %d", cash[i], readBuffer[i]))
		}
	}
}
