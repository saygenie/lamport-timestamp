package main

import (
	"fmt"
)

type Process struct {
	id   int
	time int
}

func (p *Process) send(msg string, dst *Process) {
	fmt.Printf("p%d send %s to p%d\n", p.id, msg, dst.id)
	p.time = p.time + 1
	timestamp := p.time
	dst.receive(msg, timestamp+1)
}

func (p *Process) receive(msg string, timestamp int) {
	if p.time = p.time + 1; timestamp > p.time {
		p.time = timestamp
	}
	fmt.Println(msg)
}

func main() {
	p1 := Process{id: 1}
	p2 := Process{id: 2}

	fmt.Printf("p1.time: %d\n", p1.time)
	fmt.Printf("p2.time: %d\n", p2.time)
	p1.send("Hello", &p2)
	fmt.Printf("p1.time: %d\n", p1.time)
	fmt.Printf("p2.time: %d\n", p2.time)
}
