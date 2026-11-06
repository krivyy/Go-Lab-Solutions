//barrier2.go Template Code
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

//--------------------------------------------
// Author: Terry Huynh (C00300806)
// Created: 14/09/2026
// Modified by: Terry
// Issues:
// Helped by: C00259228, Shadrach (C00298390) and Isabel (C00303465)
// Helped: Shadrach (C00298390) and Isabel (C00303465)
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func doStuff(goNum int, wg *sync.WaitGroup, theLock *sync.Mutex, sem *semaphore.Weighted, ctx context.Context, count *int) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	//we wait here until everyone has completed part A
	theLock.Lock()
	*count++
	//barrier code here
	if *count == 10 {
		sem.Release(10)
	}
	theLock.Unlock()
	sem.Acquire(ctx, 1)
	fmt.Println("PartB", goNum)
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	ctx := context.TODO()
	var theLock sync.Mutex
	sem := semaphore.NewWeighted(int64(totalRoutines))
	count := 0
	sem.Acquire(ctx, int64(totalRoutines))

	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &wg, &theLock, sem, ctx, &count)
	}
	wg.Wait() //wait for everyone to finish before exiting
}
