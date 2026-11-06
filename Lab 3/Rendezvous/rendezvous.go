//rendezvous.go Template Code
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
// Created: 14/09/2026
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

//Global variables shared between functions --A BAD IDEA

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, barrier chan bool, release chan bool) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	barrier <- true
	//Rendezvous here
	<-release
	fmt.Println("Part B", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	barrier := make(chan bool)
	release := make(chan bool)
	threadCount := 10 // Changed to 10 because of the comment under

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, barrier, release)
	}
	for range threadCount {
		<-barrier
	}
	for range threadCount {
		release <- true
	}
	wg.Wait() //wait here until everyone (10 go routines) is done

}
