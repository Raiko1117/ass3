package main

import "fmt"

type Device interface {
	on()
	off()
}

type Light struct {
	isOn bool
}

func (l *Light) on() {
	l.isOn = true
	fmt.Println("Turning the light on")
}

func (l *Light) off() {
	l.isOn = false
	fmt.Println("Turning the light off")
}

type Command interface {
	execute()
}

type TurnOnCommand struct {
	device Device
}

func (c *TurnOnCommand) execute() {
	c.device.on()
}

type TurnOffCommand struct {
	device Device
}

func (c *TurnOffCommand) execute() {
	c.device.off()
}

type RemoteControl struct {
	buttons []Command
}

func (rc *RemoteControl) addCommand(command Command) {
	rc.buttons = append(rc.buttons, command)
}

func (rc *RemoteControl) pressButton(index int) {
	if index >= 0 && index < len(rc.buttons) {
		rc.buttons[index].execute()
	}
}
