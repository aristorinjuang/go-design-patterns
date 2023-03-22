package main

import "fmt"

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (s *RealSubject) Do() string {
	return "RealSubject: doing something"
}

type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) Do() string {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	result := "Proxy: calling the real subject to execute a task\n"
	result += p.realSubject.Do()
	return result
}

func main() {
	var subject Subject = &Proxy{}
	result := subject.Do()
	fmt.Println(result)
}
