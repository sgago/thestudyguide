package broadcast

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestBroadcast_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer server.Close()

	client := resty.New()
	broadcast := Get(client.R(), server.URL)

	resp := broadcast.Send()
	assert.NoError(t, resp.Err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
}

func TestBroadcast_Post(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "created"}`))
	}))
	defer server.Close()

	client := resty.New()
	broadcast := Post(client.R(), server.URL)

	resp := broadcast.Send()
	assert.NoError(t, resp.Err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode())
}

func TestBroadcast_SendC(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer server.Close()

	client := resty.New()
	broadcast := Get(client.R(), server.URL, server.URL)

	responses := broadcast.SendC()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		resp := <-responses
		assert.NoError(t, resp.Err)
		assert.Equal(t, http.StatusOK, resp.StatusCode())
	}()

	go func() {
		defer wg.Done()
		resp := <-responses
		assert.NoError(t, resp.Err)
		assert.Equal(t, http.StatusOK, resp.StatusCode())
	}()

	wg.Wait()
}
