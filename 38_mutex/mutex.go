package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
   In the previous example we saw how to manage simple counter state using atomic operations.
   For more complex state we can use a mutex to safely access data across multiple goroutines.
*/

func main() {
	// state is a map
	var state = make(map[int]int)

	// mutex that synchronizes access to state
	var mutex = &sync.Mutex{}

	// monitor how many reads and writes we do
	var readOps uint64
	var writeOps uint64

	// start 100 goroutines to execute repeated reads against state, once per millisecond in each
	// goroutine
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				// for each read, we pick a random key to access, Lock() the mutex to ensure exclusive
				// access to state, read the value at the chosen key, Unlock() the mutex and increment
				// readOps count
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	// start 10 goroutines to simulate writes
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[val] = key
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	// let the goroutines work on the state and mutex for a second.
	time.Sleep(time.Second)

	// take final report
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// with a final lock of state, lets see how it ended up.
	mutex.Lock()
	fmt.Println("state", state)
	mutex.Unlock()
}

/*

   $ go run 38_mutex/mutex.go
   readOps: 84759
   writeOpe: 8477
   state map[0:2 1:3 2:2 3:0 4:2 5:3 6:0 7:3 8:1 9:2 10:2 11:0 12:0 13:2 14:4 15:0 16:4 17:4 18:3
   19:0 20:3 21:4 22:2 23:4 24:1 25:1 26:2 27:0 28:0 29:3 30:2 31:0 32:1 33:2 34:4 35:4 36:2 37:0
   38:2 39:3 40:4 41:3 42:0 43:3 44:0 45:3 46:4 47:3 48:3 49:4 50:4 51:0 52:0 53:2 54:4 55:3 56:0
   57:2 58:1 59:3 60:1 61:3 62:2 63:1 64:0 65:2 66:3 67:0 68:0 69:4 70:1 71:1 72:4 73:1 74:2 75:4
   76:1 77:4 78:4 79:1 80:2 81:0 82:4 83:1 84:0 85:3 86:2 87:0 88:3 89:2 90:1 91:0 92:1 93:1 94:0
   95:1 96:2 97:1 98:4 99:1]


   Running the program shows that we executed about 90,000 total operations against our
   mutex-synchronized state.

*/
