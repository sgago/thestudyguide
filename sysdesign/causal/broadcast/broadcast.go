// Package broadcast provides functionality for sending broadcast
// requests to multiple URLs concurrently.
package broadcast

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/go-resty/resty/v2"
)

// Broadcast represents a broadcast request.
type Broadcast struct {
	proto *resty.Request
	verb  string
	urls  []string

	onSuccessFunc func(resp *resty.Response)
	onErrorFunc   func(err error)
}

// BroadcastResponse represents the response of a broadcast request.
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

// SendC sends the broadcast request concurrently and
// returns a channel to receive the responses.
func (b *Broadcast) SendC() chan *BroadcastResponse {
	out := make(chan *BroadcastResponse, len(b.urls))
	var wg sync.WaitGroup

	for _, url := range b.urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()

			req := deepClone(b.proto)
			req.URL = u

			for resp := range sendC(req) {
				if resp.Err == nil && b.onSuccessFunc != nil {
					b.onSuccessFunc(resp.Response)
				} else if resp.Err != nil && b.onErrorFunc != nil {
					b.onErrorFunc(resp.Err)
				}

				out <- resp
			}
		}(url)
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}

func (b *Broadcast) Send() chan bool {
	done := make(chan bool)

	go func() {
		defer close(done)

		success := true
		resps := b.SendC()

		for resp := range resps {
			if resp.Err != nil {
				success = false
			}
		}

		done <- success
	}()

	return done
}

func (b *Broadcast) SendAndWait() bool {
	return <-b.Send()
}

// OnSuccess sets the function to be executed on successful response.
func (b *Broadcast) OnSuccess(do func(resp *resty.Response)) *Broadcast {
	b.onSuccessFunc = do
	return b
}

// OnError sets the function to be executed on error.
func (b *Broadcast) OnError(do func(err error)) *Broadcast {
	b.onErrorFunc = do
	return b
}

// sendC sends the request and returns a channel to receive the response.
func sendC(req *resty.Request) chan *BroadcastResponse {
	out := make(chan *BroadcastResponse, 1)

	go func() {
		defer close(out)
		resp, err := req.Send()
		out <- &BroadcastResponse{Response: resp, Err: err}
	}()

	return out
}

// shallowClone creates a clone of the request.
func shallowClone(req *resty.Request) *resty.Request {
	clone := *req
	return &clone
}

func deepClone(req *resty.Request) *resty.Request {
	clone := shallowClone(req)

	// Clone basic fields
	clone.Method = req.Method
	clone.URL = req.URL
	clone.QueryParam = cloneValues(req.QueryParam)
	clone.FormData = cloneValues(req.FormData)
	clone.PathParams = cloneMap(req.PathParams)
	clone.Cookies = cloneCookies(req.Cookies)

	if req.Body != nil {
		clone.SetBody(req.Body)
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
