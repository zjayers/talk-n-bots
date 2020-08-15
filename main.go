package main

import "fmt"

// Define base bot interface
// If a struct has a method of getGreeting that returns a string -> structs are automatically a member of type bot
type bot interface {
    getGreeting() string
}

// Create the englishBot and spanishBot
type englishBot struct{}
type spanishBot struct{}

func main() {

    eb := englishBot{}
    sb := spanishBot{}

    printGreeting(eb)
    printGreeting(sb)

}

func printGreeting(b bot) {
    fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
    //  Custom logic for generating an english greeting
    return "Hi there"
}

func (sb spanishBot) getGreeting() string {
    //  Custom logic for generating a spanish greeting
    return "Hola"
}
