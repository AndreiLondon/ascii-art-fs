package main

import (
	"fmt"
	"log"
	"os"
)

var usage = `Usage: go run . [STRING] [BANNER]

EX: go run . something standard`

func getArguments() (string, string) {
	// ignore item 0 and get everytning that follows item 0
	args := os.Args[1:]

	if len(args) == 0 || len(args) > 2 {
		fmt.Println(usage)
		os.Exit(0)
	}
	if len(args) == 1 {
		return args[0], "standard"
	}

	return args[0], args[1]
}

func readFile(s string) string {
	data, err := os.ReadFile(s)
	if err != nil {
		log.Fatal("invalid file")
	}
	return string(data)
}
