package broadcast

import (
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/go-resty/resty/v2"
)

type Broadcast struct {
	proto *resty.Request
	verb  string
	urls  []string

	beforeFunc   func(req *resty.Request)
	successFunc  func(resp *resty.Response)
	netErrorFunc func(err error)
	appErrorFunc func(resp *resty.Response)
}

type BroadcastResponse struct {
	*resty.Response
	Err error
}

// Get creates a new GET broadcast request.
func Get(req *resty.Request, urls ...string) *Broadcast {
	return &Broadcast{
		verb:  http.MethodGet,
		proto: req,
		urls:  urls,
	}
}

// Post creates a new POST broadcast request.
func Post(req *resty.Request, urls ...string) *Broadcast {
	return &Broadcast{
		verb:  http.MethodPost,
		proto: req,
		urls:  urls,
	}
}

// SendC sends the broadcast request concurrently to all replicas
// and returns a channel to receive the responses.
func (b *Broadcast) SendC() <-chan *BroadcastResponse {
	out := make(chan *BroadcastResponse, len(b.urls))
	var wg sync.WaitGroup

	for _, url := range b.urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()

			// Clone the original request and set the target URL
			req := deepClone(b.proto)
			req.Method = b.verb
			req.URL = u

			b.maybeBefore(req)

			// Send the request
			resp, err := req.Send()

			b.maybeOnNetError(err)

			if resp != nil {
				b.maybeOnSuccess(resp)
				b.maybeOnAppError(resp)
			}

			// Send the response (or error) back on the channel
			out <- &BroadcastResponse{Response: resp, Err: err}
		}(url)
	}

	// Close the channel once all goroutines have completed
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func (b *Broadcast) Send() *BroadcastResponse {
	return <-b.SendC()
}

func (b *Broadcast) Before(do func(req *resty.Request)) *Broadcast {
	b.beforeFunc = do
	return b
}

func (b *Broadcast) maybeBefore(req *resty.Request) {
	if b.successFunc != nil {
		b.beforeFunc(req)
	}
}

// OnSuccess sets the function to be executed on a successful response.
func (b *Broadcast) OnSuccess(do func(resp *resty.Response)) *Broadcast {
	b.successFunc = do
	return b
}

func (b *Broadcast) maybeOnSuccess(resp *resty.Response) {
	if resp != nil && resp.IsSuccess() {
		log.Printf("Success sending to %s\n", resp.Request.URL)

		if b.successFunc != nil {
			b.successFunc(resp)
		}
	}
}

// OnError sets the function to be executed on a network-level error.
// This is for handling timeouts, connection errors, etc.
func (b *Broadcast) OnNetError(do func(err error)) *Broadcast {
	b.netErrorFunc = do
	return b
}

func (b *Broadcast) maybeOnNetError(err error) {
	if err != nil {
		log.Printf("Network error: %v\n", err)

		if b.netErrorFunc != nil {
			b.netErrorFunc(err)
		}
	}
}

// OnError sets the function to be executed on an application-level error.
// This is for handling 4xx and 5xx responses.
func (b *Broadcast) OnAppError(do func(resp *resty.Response)) *Broadcast {
	b.appErrorFunc = do
	return b
}

func (b *Broadcast) maybeOnAppError(resp *resty.Response) {
	if resp != nil && resp.IsError() {
		log.Printf("Application error: %v\n", resp.Error())

		if b.appErrorFunc != nil {
			b.appErrorFunc(resp)
		}
	}
}

// deepClone manually clones the request, ensuring all headers and settings are copied.
func deepClone(req *resty.Request) *resty.Request {
	// Create a new request based on the original
	clone := resty.New().R()

	// Copy all necessary fields from the original request
	clone.Method = req.Method
	clone.URL = req.URL
	clone.QueryParam = cloneValues(req.QueryParam)
	clone.FormData = cloneValues(req.FormData)
	clone.PathParams = cloneMap(req.PathParams)
	clone.Cookies = cloneCookies(req.Cookies)

	// If the body is set, copy it to the cloned request
	if req.Body != nil {
		clone.SetBody(req.Body)
	}

	// Copy headers from the original request
	clone.Header = make(map[string][]string)
	for k, v := range req.Header {
		clone.Header[k] = append([]string{}, v...)
	}

	return clone
}

func cloneValues(src url.Values) url.Values {
	clone := make(url.Values, len(src))
	for k, v := range src {
		clone[k] = v
	}
	return clone
}

// Utility function to clone map[string]string
func cloneMap(src map[string]string) map[string]string {
	clone := make(map[string]string)
	for k, v := range src {
		clone[k] = v
	}
	return clone
}

// Utility function to deep clone cookies
func cloneCookies(src []*http.Cookie) []*http.Cookie {
	clone := make([]*http.Cookie, len(src))
	for i, cookie := range src {
		// Deep copy individual cookie fields
		clone[i] = &http.Cookie{
			Name:     cookie.Name,
			Value:    cookie.Value,
			Path:     cookie.Path,
			Domain:   cookie.Domain,
			Expires:  cookie.Expires,
			Secure:   cookie.Secure,
			HttpOnly: cookie.HttpOnly,
			SameSite: cookie.SameSite,
		}
	}
	return clone
}
