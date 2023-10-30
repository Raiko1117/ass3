package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Singleton struct {
	data string
}

var singleInstance *Singleton

func getInstance() *Singleton {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating a new Singleton instance")
			singleInstance = &Singleton{data: "Initial data"}
		} else {
			fmt.Println("Singleton instance already created.")
		}
	} else {
		fmt.Println("Singleton instance already created.")
	}

	return singleInstance
}

func (s *Singleton) SetData(data string) {
	s.data = data
}
