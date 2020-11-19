package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	//size of Universe
	width  = 80
	height = 15
)

//Universe is a 2D field of cells
type Universe [][]bool

//NewUniverse generates an empty Universe
func NewUniverse() Universe {
	universe := make(Universe, height)
	for i := range universe {
		universe[i] = make([]bool, width)
	}
	return universe
}

//Show cleans the screen and shows the Universe
func (u Universe) Show() {
	//fmt.Print("\x0c")//maybe is for Windows system
	fmt.Print("\033[2J") //cleans the screen for Linux/Darwin systems
	var b byte
	//to show the Universe it needs to be printed whole at once, as String for example
	buf := make([]byte, 0, (width+1)*height)
	for _, h := range u {
		for _, w := range h {
			if w { //there is an alive cell
				b = '*'
			} else {
				b = ' '
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}
	fmt.Print(string(buf))
}

//Seed is for initial generating of Universe with alive cells
func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ { //it will be 25% of alive cells
		h := rand.Intn(height)
		w := rand.Intn(width)
		if u[h][w] == false {
			u[h][w] = true
		} else {
			i-- //this place is already with cell, one more circle added
		}
	}
}

//Alive is a checks, wheather the cell is alive or not
//There is a problem of processing a border cells,
//so the solution is to loop the Universe in every direction
func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

//Neighbors is counting alive cells near the desired point
func (u Universe) Neighbors(x, y int) int {
	var totAlive int = 0
	//need to check all 8 cells near and count alives
	for yi := -1; yi <= 1; yi++ {
		for xi := -1; xi <= 1; xi++ {
			if !(xi == 0 && yi == 0) && u.Alive(x+xi, y+yi) {
				totAlive++
			}
		}
	}
	return totAlive
}

//Next is defines the next status of cell
func (u Universe) Next(x, y int) bool {
	//there are defined rules of next status of cell based on alive cells near
	if u.Alive(x, y) {
		switch u.Neighbors(x, y) {
		case 0, 1:
			return false
		case 2, 3:
			return true
		case 4, 5, 6, 7, 8:
			return false
		}
	} else if u.Neighbors(x, y) == 3 {
		return true
	}
	return false
}

//Step is generates new Universe (of next step) based on old Universe
func Step(a, b Universe) {
	for ya, yaLine := range a {
		for xa := range yaLine {
			b[ya][xa] = a.Next(xa, ya)
		}
	}
}
func main() {
	a, b := NewUniverse(), NewUniverse() //old and new Universes, as previous and next step
	a.Seed()                             //initial filling the Universe with cells
	for i := 0; i < 160; i++ {
		a.Show()   //displays one screen of current Universe (one frame)
		Step(a, b) //generates next step
		a, b = b, a
		time.Sleep(time.Second / 16)
	}
}
