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

// Create a barrier data type
type barrier struct {
	theChan  chan bool
	theChan2 chan bool
	theLock  sync.Mutex
	total    int
	count    int
}

// creates a properly initialised barrier
// N== number of threads (go Routines)
func createBarrier(N int) *barrier {
	return &barrier{
		theChan:  make(chan bool, N),
		theChan2: make(chan bool, N),
		total:    N,
		count:    0,
	}
}

// Method belonging to barrier data type
// Blocks until everyone reaches this point then lets everyone continue
func (b *barrier) wait() {
	b.theLock.Lock()
	b.count++
	if b.count == b.total {
		for range b.total {
			b.theChan <- true
		}
	}
	b.theLock.Unlock()
	<-b.theChan
	b.theLock.Lock()
	b.count--
	if b.count == 0 {
		for range b.total {
			b.theChan2 <- true
		}
	}
	b.theLock.Unlock()
	<-b.theChan2
}

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, theBarrier *barrier) bool {
	for round := 0; round < 3; round++ {
		X := time.Duration(rand.IntN(5))
		time.Sleep(X * time.Second)
		fmt.Println("Part A", Num)
		theBarrier.wait()
		fmt.Println("PartB", Num)
	}
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	barrier := createBarrier(5)
	threadCount := 5
	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, barrier)
	}
	wg.Wait() //wait here until everyone (5 go routines) is done

}
