package main

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	execute(t, "hello", "standard", "test01.txt")
	execute(t, "hello world", "shadow", "test02.txt")
	execute(t, "nice 2 meet you", "thinkertoy", "test03.txt")
	execute(t, "you & me", "standard", "test04.txt")
	execute(t, "123", "shadow", "test05.txt")
	execute(t, "/(\")", "thinkertoy", "test06.txt")
	execute(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "shadow", "test07.txt")
	execute(t, "\"#$%&/()*+,-./", "thinkertoy", "test08.txt")
	execute(t, "It's Working", "thinkertoy", "test09.txt")
}

func execute(t *testing.T, text string, banner string, testFile string) {
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf(`Failed to readt test file`)
	}

	Want := string(data)
	Found := convText(text, banner)

	if Want != Found {
		t.Fatalf(`convertAndPrint(path, text") fail.
		Want:
%v
		Found:
%v`, Want, Found)
	}
}
