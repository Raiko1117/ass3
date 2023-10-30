package main

import "fmt"

func SingletonExample() {
	singleton := getInstance()
	fmt.Printf("Data from Singleton: %s\n", singleton.data)

	singleton.SetData("New data")

	fmt.Printf("Updated data from Singleton: %s\n", singleton.data)
}

func CommandExample() {
	light := &Light{}
	turnOn := &TurnOnCommand{device: light}
	turnOff := &TurnOffCommand{device: light}

	remote := &RemoteControl{}
	remote.addCommand(turnOn)
	remote.addCommand(turnOff)

	remote.pressButton(0) // Turns the light on
	remote.pressButton(1) // Turns the light off
}

func main() {

	//SingletonExample()
	CommandExample()
}
