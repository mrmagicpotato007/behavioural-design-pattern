package command

import "log"

type tv struct {
     isRunning bool
}

func (t *tv) on(){
	t.isRunning = true
	log.Println("Tv started")

}

func (t *tv) off(){
	t.isRunning = false
	log.Println("Tv stopped")
}


