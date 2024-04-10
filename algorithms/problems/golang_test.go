package problems

import (
	"cmp"
	"container/heap"
	"fmt"
	"slices"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestArrays(t *testing.T) {
	// An array of 5 empty elements
	var a [5]int
	fmt.Println("emp:", a)

	// Declare and init an array in one line
	b := [5]int{0, 1, 2, 3, 4}
	fmt.Println("dcl:", b)

	// A 2D array declare and init'ed in one line
	c := [3][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	fmt.Println("2d:", c)
}

func TestSlices(t *testing.T) {
	// Make a slice
	s := make([]int, 5)
	fmt.Println("emp", s)

	// Copy a slice
	copy(s, []int{0, 1, 2, 3, 4})
	fmt.Println("cpy", s)

	// Has a slice operator slice[low:high]
	// where low and high are inclusive and exclusive, respectively.
	// So s[2:5] would return s[2], s[3], and s[4]. Not s[5]!
	fmt.Println("s[2:5]", s[2:5])
	fmt.Println("s[2:]", s[2:])
	fmt.Println("s[:3]", s[:3])
}

func TestSorting(t *testing.T) {
	nums := []int{5, 3, 1, 4, 2}
	strs := []string{"def", "xyz", "abc"}

	// Method 1: The slices.Sort (newer), a generic func
	slices.Sort(nums)
	slices.Sort(strs)

	fmt.Println("nums:", nums)
	fmt.Println("strs:", strs)

	// Method 2: For sorting other things with a compare func
	cmp := func(a, b int) int {
		if a == b {
			return 0
		} else if a > b {
			return 1
		}

		return -1
	}

	slices.SortFunc(nums, cmp)

	// Method 1: The sort.Ints (old)
	sort.Ints(nums)
}

type dog struct {
	name string
}

func TestParameters(t *testing.T) {
	passIntByVal := func(x int) { x++ }

	num := 99
	passIntByVal(num)
	fmt.Println("int pass by value:", num)

	passIntByRef := func(x *int) { *x++ }
	passIntByRef(&num)
	fmt.Println("int pass by ref:", num)

	rex := dog{
		name: "rex",
	}

	passStructByVal := func(d dog) {
		d.name += ", the good boy"
	}

	passStructByVal(rex)
	fmt.Println("struct pass by value:", rex)

	passStructByRef := func(d *dog) {
		d.name += ", the good boy"
	}

	passStructByRef(&rex)
	fmt.Println("struct pass by ref:", rex)

	arr := []int{1, 2, 3}

	passSliceByVal := func(a []int) {
		a = append(a, 4, 5, 6) // Expected warning, this doesn't do anything
		a[0] = 99
	}

	passSliceByVal(arr)
	fmt.Println("slice pass by val:", arr)

	passSliceByRef := func(a *[]int) {
		*a = append(*a, 4, 5, 6) // Expected warning, that's the whole point of doing this
		(*a)[0] = 99
	}

	passSliceByRef(&arr)
	fmt.Println("slice pass by ref:", arr)
}

func TestChannels(t *testing.T) {
	messages := make(chan string)

	go func() {
		// Sends "ping" into messages channel
		messages <- "ping" // sender blocks until the receiver is ready
	}()

	// Receive a channel value into messages. Or like Send any channel values into this var.
	msg := <-messages // receiver blocks until sender is ready

	// Again, sends and recieves block, allowing us to wait at the end of the func for the value.
	fmt.Println("rcvd:", msg)

	// This'll panic with "all goroutines are asleep - deadlock!"
	//bad := <-messages
	//fmt.Println("bad", bad)

	// Unbuffered channels are strict synchronization point where both sender
	// and receiver must be ready. Use unbuffered channels when you need strong
	// synchronization between sender and receiver, the ensure data are shared.
}

func TestChannelBuffer(t *testing.T) {
	// Channels are unbuffered, meaning they will only accept sends (chan <-)
	// if there is a corresponding receive (<- chan).\
	// Buffered channels will receive a limited number of channels without a receiver.
	messages := make(chan string, 2) // This is a buffered channel, it has a capacity > 0

	go func() {
		// The sender chan send data to the channel without blocking, until it is full
		messages <- "buffered"
		messages <- "channel"
		//messages <- "this'll panic cause the buffer is full"
	}()

	// Similarily, we can receive from the channel without blocking until it's empty
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	//fmt.Println(<-messages) // This'll panic too

	// We do get a degree of async freedom with buffered channels.
	// Useful where small delays or async is acceptable. Useful where small delays
	// or bursts of data are expected.
}

func TestChannelDirections(t *testing.T) {
	// When using channels as params, you can increase the type safety by specifying directions.

	// sendFunc will only accept a channel for sending values
	sendFunc := func(sendChan chan<- string, msg string) {
		sendChan <- msg

		// Uncommeting will yield a "invalid operation: cannot receive from send-only channel sendChan"
		//myVal := <-sendChan
	}

	recvFunc := func(recvChan <-chan string) string {
		//recvChan <- "uncommenting creates a 'invalid operation: cannot send to receive-only channel recvChan' error"
		return <-recvChan
	}

	ch := make(chan string, 1)

	sendFunc(ch, "hello")

	fmt.Println(recvFunc(ch))
}

func TestSelect(t *testing.T) {
	chan1 := make(chan string)
	chan2 := make(chan string)

	fmt.Println("starting")

	go func() {
		time.Sleep(time.Second * 5)
		chan1 <- "one done!"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		chan2 <- "two done!"
	}()

	for i := 0; i < 2; i++ {
		// Select will await both values simultaneously
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
			//default:
			// If we uncomment the default case, this loop'll exit prematurely
		}
	}
}

func TestTimeouts(t *testing.T) {
	chan1 := make(chan string, 1) // Buffered chan

	fmt.Println("starting chan1")

	go func() {
		time.Sleep(time.Second * 5)
		chan1 <- "one done!" // Buffered chan, non-blocking to prevent goroutine leaks in case chan is never read
	}()

	select {
	case msg := <-chan1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout on chan 1!")
	}

	chan2 := make(chan string, 1) // Buffered chan

	fmt.Println("starting chan2")

	go func() {
		time.Sleep(time.Second * 1)
		chan2 <- "one done!" // Buffered chan, non-blocking to prevent goroutine leaks in case chan is never read
	}()

	select {
	case msg := <-chan2:
		fmt.Println(msg)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout on chan 2!")
	}
}

func TestChannelRange(t *testing.T) {
	ch := make(chan string, 5)
	ch <- "one"
	ch <- "two"
	ch <- "three"

	close(ch) // Closes the channel, indicates no more values will be sent

	for val := range ch {
		fmt.Println(val)
	}
}

func TestWaitGroups(t *testing.T) {
	longTask := func(id int) {
		fmt.Println("Worker starting: ", id)
		time.Sleep(1 * time.Second)
	}

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1) // Increment the wait group counter for each goroutine launched

		go func(id int) {
			defer wg.Done()
			longTask(id)
		}(i)
	}

	wg.Wait() // Block until all goroutines are done

	// Note: Use https://pkg.go.dev/golang.org/x/sync/errgroup or similar for error handling
}

func TestBasicRateLimiting(t *testing.T) {
	// Basic rate limiter
	// This mega basic rate limiter will only serve requests every 200ms.
	// It does this by creating a ticker limiter that only unblocks every 200ms.
	// It obviously doesn't do limiting by IP address, API key, or anything like that.

	// Simulate requests are coming from our HTTP server that's receiving requests
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests) // Close the channel, so we can iterate over it

	// Create a ticker, to periodically serve requests
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("basic req: ", req, time.Now())
	}
}

func TestBurstyRateLimiting(t *testing.T) {
	// Bursty rate limiter
	// This also mega basic rate limiter will allow bursts of 3 requests
	// but then rate limits down to serving requests every 200ms.
	// Unlike the basic limiter above, this one uses a limiter time channel.
	// It obviously doesn't do limiting by IP address, API key, or anything like that.

	limiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		limiter <- time.Now()
	}

	go func() {
		// Kick off the ticker, receivers can acrue up to 3 bursty requests
		// and then only 1 request every 200 ms.
		for tick := range time.Tick(200 * time.Millisecond) {
			limiter <- tick
		}
	}()

	// Simulate requests are coming from our HTTP server that's receiving requests
	requests := make(chan int, 10)
	for i := 0; i < 10; i++ {
		requests <- i
	}
	close(requests) // Close the channel, so we can iterate over it

	// On looping thru requests, you'll immediately get 3 requests to the burst limit,
	// but then the limiter is drained. So, the client'll need to wait for 200ms to get another one.
	for req := range requests {
		<-limiter
		fmt.Println("bursty req: ", req, time.Now())
	}
}

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func TestStack(t *testing.T) {
	// TODO: Author tests and comments
}

type Queue []int

func (s *Queue) Enq(x int) {
	*s = append(*s, x)
}

func (s *Queue) Deq() int {
	result := (*s)[0]
	*s = (*s)[1:len(*s)]
	return result
}

func (s *Queue) Peek() int {
	return (*s)[0]
}

func TestQueue(t *testing.T) {
	// TODO: Author a tests and comments
}

func TestMaps(t *testing.T) {
	graph := map[int][]int{
		1: {2, 3},
		2: {1, 2, 3},
		3: {1, 2},
		4: {2},
	}

	fmt.Println(graph[1])
}

func TestBoundedBinarySearch(t *testing.T) {
	needle := 3
	haystack := []int{1, 3, 5, 7, 9}

	// Searching for a value in a sorted slice
	idx, ok := slices.BinarySearch(haystack, needle)

	fmt.Println("idx:", idx, "ok:", ok)

	cmpFunc := func(a, b int) int {
		return cmp.Compare(a, b)
	}

	idx, ok = slices.BinarySearchFunc(haystack, needle, cmpFunc)

	fmt.Println("idx:", idx, "ok:", ok)
}

// Alias a slice type
type MinHeap []int

// And implement the heap.Interface type.
var _ heap.Interface = (*MinHeap)(nil)

func NewMinHeap(nums ...int) *MinHeap {
	h := &MinHeap{}
	*h = nums

	heap.Init(h)

	return h
}

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i int, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Pop() any {
	result := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return result
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func TestHeap(t *testing.T) {
	// Create an empty MinHeap
	h := &MinHeap{}

	// Push elements onto the heap h, using heap.Push(h, whatever)
	heap.Push(h, 3)
	heap.Push(h, 1)
	heap.Push(h, 4)
	heap.Push(h, 2)

	// Or use this cute ctor which heapifies in-place and is more ergonomic to use
	// h2 := NewMinHeap(3, 1, 4, 2)

	// Pop elements from the heap h (retrieve in sorted order for a min-heap)
	// using the heap.Pop(h)
	for h.Len() > 0 {
		fmt.Printf("%d\v", heap.Pop(h))
	}
}
