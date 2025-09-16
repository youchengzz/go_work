package main

import "fmt"

type Usber interface {
	start()
	stop()
}

type Computer struct {
	Type string
}

func (c Computer) work(u Usber) {
	u.start()
	u.stop()
}

func (c Computer) start() {
	fmt.Println("电脑开启", c.Type)
}

func (c Computer) stop() {
	fmt.Println("电脑关机", c.Type)
}

func main() {
	var computer = Computer{}
	var c = Computer{
		Type: "华为",
	}
	c.work(computer)
}
