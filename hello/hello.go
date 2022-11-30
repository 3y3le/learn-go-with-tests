package main

import (
	"bytes"
	"fmt"
)

const englishPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const frenchHelloPrefix = "Bonjour,"

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func Greet(writer *bytes.Buffer, name string) {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	if err != nil {
		return
	}
}
