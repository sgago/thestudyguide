package broadcast

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/go-resty/resty/v2"
)

type Broadcast struct {
	proto *resty.Request
	verb  string
	urls  []string

	onSuccessFunc func(resp *resty.Response)
	onErrorFunc   func(err error)
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

			// Send the request
			resp, err := req.Send()

			// Handle success or failure
			if err != nil {
				fmt.Printf("Error sending to %s: %v\n", u, err)
			} else {
				if b.onSuccessFunc != nil && resp.StatusCode() == http.StatusOK {
					b.onSuccessFunc(resp)
				}
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

// OnSuccess sets the function to be executed on a successful response.
func (b *Broadcast) OnSuccess(do func(resp *resty.Response)) *Broadcast {
	b.onSuccessFunc = do
	return b
}

// OnError sets the function to be executed on error.
func (b *Broadcast) OnError(do func(err error)) *Broadcast {
	b.onErrorFunc = do
	return b
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
