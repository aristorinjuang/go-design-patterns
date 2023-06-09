package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "Hello, World!"}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()
	if s1 == s2 {
		fmt.Println("s1 and s2 are the same instance")
	} else {
		fmt.Println("s1 and s2 are different instances")
	}
	fmt.Printf("s1 data: %s\n", s1.data)
	fmt.Printf("s2 data: %s\n", s2.data)
}
