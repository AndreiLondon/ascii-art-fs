package main

import "fmt"

func main() {
	text, banner := getArguments()
	// If it's an empty string it will exit the program
	if text == "" {
		return
	}
	if text == "\\n" {
		fmt.Println()
		return
	}

	result := convText(text, banner)
	fmt.Println(result)
}

func convText(text string, banner string) string {
	path := banner + ".txt"
	output := readFile(path)
	myMap := parseTemplate(output)
	return printMessage(text, myMap)
}
