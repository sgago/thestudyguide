package concurrency

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

/*
	Reference: https://go.dev/talks/2012/concurrency.slide#1
*/

func Test_ExitsImmediately(t *testing.T) {
	boring := func(msg string) {
		for i := 0; ; i++ {
			fmt.Println(msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}

	// We can think of the go keyword as launching a cheap thread aka goroutine.
	go boring("the caller (the test driving boring code in this case) won't wait and " +
		"only lets boring run like once; this is dangerous and useless")
}

func Test_DelayBeforeExiting(t *testing.T) {
	boring := func(msg string) {
		for i := 0; ; i++ {
			fmt.Println(msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}

	go boring("Loops:")
	fmt.Println("Sleeping the test for a few seconds")
	time.Sleep(time.Second * 3)
	fmt.Println("Test is exiting now")
}

func Test_BasicChannels(t *testing.T) {
	// This is a channel, it lets us communicate with many goroutines.
	ch := make(chan string)

	// Warning: Don't do stuff like this, you'll get deadlock errors
	//ch <- 123
	//val := <-ch

	go func() {
		fmt.Println(<-ch) // Data out
	}()

	ch <- "hello" // Data in
}

func Test_BasicSynchronizationWithChannels(t *testing.T) {
	ch := make(chan int)

	// Note: You could just share ch here instead of passing
	// it as a param, but a chan param would be more typical.
	go func(msg string, c chan int) {
		for i := 0; ; i++ {
			fmt.Println(msg, <-c) // Boring func waits until i is sent
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}("Value from channel:", ch)

	for i := 0; i < 5; i++ {
		ch <- i // Main will pause until i is sent into the channel.
	}

	// Both sender and receiver wait until the other is ready to communicate
	// and share data. Otherwise, sender and/or receiver will hang until the other is ready.

	// Go's approach is not to share memory like in other languages.
	// Goroutines share memory by communicating.

	fmt.Println("Test is exiting now")
}

func Test_Generator_Pattern(t *testing.T) {
	// Channels are first-class citizens and can be passed around through
	// return values and params.

	generate := func(msg string) chan string {
		// Make a chan string and give it to the caller so it can talk to us
		ch := make(chan string)

		// Fire off an internal goroutine
		go func() {
			for i := 0; ; i++ {
				ch <- fmt.Sprint(msg, i)
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			}
		}()

		return ch // And give the channel back to the caller, so we can talk
	}

	ch := generate("We're talking!")

	for i := 0; i < 5; i++ {
		// Get a value out of the chan 5 times, neat.
		// Will only get a value when the internal goroutine is done waiting
		fmt.Println(<-ch)
	}
}

func Test_GenerateMultiple_Pattern(t *testing.T) {
	generate := func(msg string) chan string {
		// Make a chan string and give it to the caller so it can talk to us
		ch := make(chan string)

		// Fire off an internal goroutine
		go func() {
			for i := 0; ; i++ {
				ch <- fmt.Sprint(msg, i)
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			}
		}()

		return ch // And give the channel back to the caller, so we can talk
	}

	// We can make gen multiple goroutines + chans and communicate
	joe := generate("Hey joe!")
	ann := generate("Hi ann!")

	for i := 0; i < 5; i++ {
		// Get a value out of the chan 5 times, neat.
		// Will only get a value when the internal goroutine is done waiting
		// Note that joe and ann are in lockstep here, first joe then bill every time.
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
}

func Test_BasicFanInOrMultiplexed_Pattern(t *testing.T) {
	generate := func(msg string) chan string {
		// Make a chan string and give it to the caller so it can talk to us
		ch := make(chan string)

		// Fire off an internal goroutine
		go func() {
			for i := 0; ; i++ {
				ch <- fmt.Sprint(msg, i)
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			}
		}()

		return ch // And give the channel back to the caller, so we can talk
	}

	// This fanIn or multiplexed func makes it so joe and bill are no longer
	// in lock step. They go in any order as quick as they can.
	fanIn := func(input1, input2 <-chan string) <-chan string {
		c := make(chan string)

		go func() {
			for {
				c <- <-input1
			}
		}()

		go func() {
			for {
				c <- <-input2
			}
		}()

		return c
	}

	ch := fanIn(generate("joe"), generate("ann"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func Test_FanInOrMultiplexed_Pattern(t *testing.T) {
	// Again, generate has an internal goroutine that effectively runs forever
	// and returns us a channel for synchronization. When this internal
	// goroutine and the caller using the channel are both ready, it'll print a msg.
	generate := func(msg string) <-chan string {
		// Make a chan string and give it to the caller so it can talk to us
		ch := make(chan string)

		// Fire off an internal goroutine
		go func() {
			for i := 0; ; i++ {
				ch <- fmt.Sprint(msg, i)
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			}
		}()

		return ch // And give the channel back to the caller, so we can talk
	}

	// This fanIn multiplexes multiple generates outputs into 1.
	// Multiplexing means joining multiple streams of data or signals into 1.
	// Note that we lose sequencing here. They go as fast as the can in any order.
	fanIn := func(inputs ...<-chan string) <-chan string {
		output := make(chan string)

		// We fire off n goroutines that wait on an input channel and
		// move those messages into a multiplexed output channel.
		// So, we combine all those input channels into one output channel.
		// That's all these are doing. Again, this internal goroutine
		// and generates goroutine must both be willing to communicate.
		for _, input := range inputs {
			go func(in <-chan string) {
				for {
					output <- <-in
				}
			}(input)
		}

		return output
	}

	out := fanIn(generate("joe"), generate("ann"), generate("test"))

	//
	for i := 0; i < 15; i++ {
		fmt.Println(<-out)
	}
}

// https://go.dev/talks/2012/concurrency.slide#29

type Message struct {
	str  string
	wait chan bool // Will be same channel among all messages!
}

func TestRestoringOrder(t *testing.T) {
	// We can boost performance and restore sequencing
	// by making goroutines wait their turn.

	// This is going to feel like the movie Inception where we
	// pass a channel through a channel. Again, channels are
	// first-class citizens like an integer, so we can pass them around
	// in very much the same ways.
	generate := func(msg string, wait chan bool) <-chan Message {
		ch := make(chan Message)

		go func() {
			for i := 0; ; i++ {
				ch <- Message{fmt.Sprintf("%s: %d", msg, i), wait}
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
				<-wait
			}
		}()

		return ch
	}

	// The shared channel for restoring sequence
	wait := make(chan bool)

	joe := generate("joe", wait)
	ann := generate("ann", wait)

	for i := 0; i < 5; i++ {
		msg1 := <-joe // Get the joe message, block until joe is ready
		fmt.Println(msg1.str)

		msg2 := <-ann // Get the ann message, block until ann is ready
		fmt.Println(msg2.str)

		// Now we tell the goroutines that we're ready in joe-ann order
		msg1.wait <- true // Tell joe we're ready
		msg2.wait <- true // Now tell ann we're ready
		// wait <- true // Don't do this, deadlocks
	}
}

// https://go.dev/talks/2012/concurrency.slide#31

func TestSelect(t *testing.T) {
	/*
		Select statements are like switch but for channels and communications.
		- All channels are evaluated
		- Select blocks until a channel is ready
		- If multiple channels are ready, it's psuedo-random order
		- A default statement runs if no channel is ready
		We can simplify our multiplexed pattern with select albeit
		losing order or operations. If need order or an unknown number of channels
		we can use the fan-in pattern from above.
	*/

	// Same generate as before
	generate := func(msg string) chan string {
		ch := make(chan string)

		go func() {
			for i := 0; ; i++ {
				ch <- fmt.Sprint(msg, i)
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			}
		}()

		return ch
	}

	fanIn := func(input1, input2 <-chan string) <-chan string {
		ch := make(chan string)

		// Again, if we don't need guaranteed order
		// and know the number of input channels, we can use select.
		go func() {
			for {
				select {
				case msg1 := <-input1:
					ch <- msg1
				case msg2 := <-input2:
					ch <- msg2
				}
			}
		}()

		return ch
	}

	ch := fanIn(generate("joe"), generate("ann"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func TestSelectTimeout(t *testing.T) {
	// Same generate as before
	generate := func(msg string) chan string {
		ch := make(chan string)

		go func() {
			for i := 0; ; i++ {
				ch <- fmt.Sprint(msg, i)
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			}
		}()

		return ch
	}

	// Basically same fanIn as before but with a custom timeout
	fanIn := func(input1, input2 <-chan string) <-chan string {
		ch := make(chan string)

		go func() {
			for {
				select {
				case msg1 := <-input1:
					ch <- msg1
				case msg2 := <-input2:
					ch <- msg2
				}
			}
		}()

		return ch
	}

	ch := fanIn(generate("joe"), generate("ann"))

	// A timeout func after one second
	timeout := time.After(1 * time.Second)
	stop := false

	for i := 0; i < 10 && !stop; i++ {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case <-timeout:
			fmt.Println("Timeout!")
			// We can't just simply call break here cause it only
			// breaks us out of the select statement and basically doesn't do anything.
			// We either need an anon function or a control var like this
			// Alternatively, we could call return and just leave this func entirely, of course.
			stop = true
		}
	}

	fmt.Println("Done")
}

func TestSelectQuit(t *testing.T) {
	// Same generate as before
	generate := func(msg string) chan string {
		ch := make(chan string)

		go func() {
			for i := 0; ; i++ {
				ch <- fmt.Sprint(msg, i)
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			}
		}()

		return ch
	}

	// Basically same fanIn as before
	fanIn := func(input1, input2 <-chan string) <-chan string {
		ch := make(chan string)

		go func() {
			for {
				select {
				case msg1 := <-input1:
					ch <- msg1
				case msg2 := <-input2:
					ch <- msg2
				}
			}
		}()

		return ch
	}

	ch := fanIn(generate("joe"), generate("ann"))

	quit := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			select {
			case msg := <-ch:
				fmt.Println(msg)
			case <-quit:
				fmt.Println("Quitting!")
				return
			}
		}
	}()

	quit <- true // Quits the select loop almost immediately

	fmt.Println("Done")
}

func TestDaisyChain(t *testing.T) {
	const n = 10000

	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	f := func(left, right chan int) {
		left <- 1 + <-right
	}

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) { c <- 1 }(right)

	fmt.Println(<-leftmost)
}

func TestAtomicCounter(t *testing.T) {
	// Shows an example of a low-level atomic primitives.
	// A typical go-way is to use channels, but these still exist
	// so we might as well cover them.

	// Declare an atomic counter like:
	var cnt atomic.Int32

	// Use the atomic counter cnt like
	cnt.Store(3)
	fmt.Println("cnt:", cnt.Add(3)) // cnt: 6
	fmt.Println("cnt:", cnt.Load()) // cnt: 6
}

func TestAtomicCompareAndSwap(t *testing.T) {
	// Compare-and-swap (CAS) can feel a little goofy
	// if it's the first time you're seeing it.

	// CAS steps:
	// 1. Load cnt via cnt.Load()
	// 2. Compare current cnt value to the old value from cnt.Load()
	// 3. If they're the same, cnt.CompareAndSwap swaps the new value in.
	//    However, if cnt's current value does not match the old value,
	//    CompareAndSwap returns false. In that case, we just try again in a nanosecond.
	var cnt atomic.Int32

	cnt.Store(99)
	new := int32(123) // The value we're going to put into cnt via CAS

	for !cnt.CompareAndSwap(cnt.Load(), new) {
		time.Sleep(1 * time.Nanosecond)
	}

	fmt.Println(cnt.Load())
}
