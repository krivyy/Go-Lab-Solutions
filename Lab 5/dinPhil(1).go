//Barrier.go Template Code
//Copyright (C) 2026 Terry Huynh

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// --------------------------------------------
// Author: Terry Huynh (C00300806)
// Created: 5/11/2026
// Modified by: Terry
// Issues:
// Helped by: C00259228, Shadrach (C00298390) and Isabel (C00303465)
// Helped: Shadrach (C00298390) and Isabel (C00303465)
// --------------------------------------------

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was eating")
}

func getForks(index int, forks map[int]chan bool, footman chan bool) {
	footman <- true
	forks[index] <- true
	forks[(index+1)%5] <- true
}

func putForks(index int, forks map[int]chan bool, footman chan bool) {
	<-forks[index]
	<-forks[(index+1)%5]
	<-footman
}

func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool, footman chan bool) {
	for {
		think(index)
		getForks(index, forks, footman)
		eat(index)
		putForks(index, forks, footman)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount)

	forks := make(map[int]chan bool)
	for k := range philCount {
		forks[k] = make(chan bool, 1)
	}
	footman := make(chan bool, philCount-1) //footman semaphore: max 4 philosophers at the table
	for N := range philCount {
		go doPhilStuff(N, &wg, forks, footman)
	} //start philosophers
	wg.Wait() //wait here until everyone (10 go routines) is done

} //main
