package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestOtherFanInAndOut(t *testing.T) {
	fileChan := make(chan string)

	fileMap := &sync.Map{}

	getFile("genFile1", fileMap, fileChan)
	getFile("genFile2", fileMap, fileChan)
	getFile("genFile3", fileMap, fileChan)

	for i := 0; i < 15; i++ {
		fmt.Println(<-fileChan)
	}
}

func getFile(name string, fileMap *sync.Map, fileChan chan<- string) {
	go func() {
		for i := 0; ; i++ {
			// TODO: Use our sync map here! If CAS is too hard, we can use a mutex map.

			fileChan <- fmt.Sprint(name, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
}

func TestFanInAndOut(t *testing.T) {
	// Fan out, generates make outputs
	gen1 := generate("gen1_")
	gen2 := generate("gen2_")
	gen3 := generate("gen3_")

	// Fan in, combined outputs on one chan
	multi := fanIn(gen1, gen2, gen3)

	timeout := time.After(10 * time.Second)

	// Get values from multiplexed chan
	for i := 0; i < 50; i++ {
		select {
		case msg := <-multi:
			fmt.Println(msg)
		case <-timeout:
			fmt.Println("Timeout!")
			return
		}
	}
}

func generate(msg string) chan string {
	out := make(chan string)

	timeout := time.After(1 * time.Second)

	go func() {
		for i := 0; ; i++ {
			select {
			case <-timeout:
				fmt.Println(msg, "timeout!")
				close(out)
				return
			default:
				out <- msg + fmt.Sprint(i)
				time.Sleep(time.Duration(rand.Intn(1_000)) * time.Millisecond)
			}
		}
	}()

	return out
}

func fanIn(gen1, gen2, gen3 <-chan string) <-chan string {
	out := make(chan string)

	timeout := time.After(10 * time.Second)

	go func() {
		for {
			select {
			case msg1 := <-gen1:
				out <- msg1
			case msg2 := <-gen2:
				out <- msg2
			case msg3 := <-gen3:
				out <- msg3
			case <-timeout:
				fmt.Println("fanIn timeout!")
				close(out)
			}
		}
	}()

	return out
}
