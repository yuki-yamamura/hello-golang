package main

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	// return "Hello, %v, name"
	return englishHelloPrefix + name
}