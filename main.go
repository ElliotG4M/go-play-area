package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/ElliotG4M/PlayArea/models"
)

var cache = map[int]models.Book{}

// define a new rng based on current unix time
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Uncomment for user system demo
// func main() {
// 	controllers.RegisterControllers()
// 	http.ListenAndServe(":3000", nil)
// }

// Uncomment for concurrent books demo
func main() {

	// Defines a new wait group which will allow us to wait until goroutines have executed
	// We should use a pointer as we are going to be passing this around a lot
	waitGroup := &sync.WaitGroup{}

	// Defines a new mutex which will allow us to prevent goroutines from simultaneously accessing & writing the same data
	// We define a Mutex or a Read Write Mutex (This is an RWM. A Mutex will just allow us to lock/unlock)
	mutex := &sync.RWMutex{}

	for i := 0; i < 10; i++ {
		// Generate a random number between 1 and 10
		id := random.Intn(10) + 1

		// Define how many goroutines the waitgroup should wait to be completed
		waitGroup.Add(2)

		// Define a goroutine. This function will process in it's own thread
		go func(id int, waitGroup *sync.WaitGroup, mutex *sync.RWMutex) {
			if book, ok := queryCache(id, mutex); ok {
				fmt.Println("Found book in cache\n", book.ToString())
			}
			// Tell the wait group that the goroutine has finished
			waitGroup.Done()
		}(id, waitGroup, mutex)

		go func(id int, waitGroup *sync.WaitGroup, mutex *sync.RWMutex) {
			if book, ok := queryDatabase(id, mutex); ok {
				fmt.Println("Found book in database\n", book.ToString())
			}
			waitGroup.Done()
		}(id, waitGroup, mutex)

		// We need to add in some time for the goroutines to execute. Without adding this sleep,
		// the program will just generate a load of goroutines and not do anything with them
		// this should be done with wait groups rather than sleeps so that we wait the right amount of time
		// time.Sleep(300 * time.Millisecond)

		// Wait until all goroutines are finished
		waitGroup.Wait()
	}
}

func queryCache(id int, mutex *sync.RWMutex) (models.Book, bool) {
	// Locks the cache for writing only. Multiple readers will still be able to access the cache
	mutex.RLock()
	book, ok := cache[id]
	mutex.RUnlock()
	return book, ok
}

func queryDatabase(id int, mutex *sync.RWMutex) (models.Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, book := range models.Books {
		if book.ID == id {
			// Lock processing so that queryCache doesn't try to read from the cache, and another queryDatabase goroutine doesn't try to write
			mutex.Lock()
			cache[id] = book
			// We've finished writing to the cache, unlock the mutex
			mutex.Unlock()
			return book, true
		}
	}
	return models.Book{}, false
}

func unbufferedChannelsExample() {
	waitGroup := &sync.WaitGroup{}

	// Channels must be defined using make. We also assign it the type it will hold
	channel := make(chan int)

	waitGroup.Add(2)

	go func(channel chan int, waitGroup *sync.WaitGroup) {
		// Pull an int out of the channel. The program will wait here until a value has been put in the channel
		// With an unbuffered channel, every write must have a matching read. If we try to read twice, the program will pause
		// until a 2nd value has been written. If we try to write twice but read once, the program will crash
		receivedInt := <-channel
		fmt.Println("Received int from channel ", receivedInt)
		waitGroup.Done()
	}(channel, waitGroup)

	go func(channel chan int, waitGroup *sync.WaitGroup) {
		// Put an int into the channel
		channel <- 42
		waitGroup.Done()
	}(channel, waitGroup)

	waitGroup.Wait()
}

func bufferedChannelsExample() {
	waitGroup := &sync.WaitGroup{}

	// Defines a buffer space of 1, i.e. we can have one int stored in the channel without having to read it. An unbuffered channel has buffer space 0
	channel := make(chan int, 1)

	waitGroup.Add(2)

	// Using <- here, we tell the goroutine that this channel is only for receiving data (arrow points out of the channel)
	go func(channel <-chan int, waitGroup *sync.WaitGroup) {
		receivedInt := <-channel
		fmt.Println("Received int from channel ", receivedInt)
		waitGroup.Done()
	}(channel, waitGroup)

	// Using <- here, we tell the goroutine that this channel is only for sending data (arrow points into the channel)
	go func(channel chan<- int, waitGroup *sync.WaitGroup) {
		// Put an int into the channel, this one will be read by the other goroutine
		channel <- 42
		// Put a 2nd int into the channel, this one will sit in the buffer as it is not read
		channel <- 24
		waitGroup.Done()
	}(channel, waitGroup)

	waitGroup.Wait()
}
