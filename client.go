package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/donjnewman/pokedex/internal/pokecache"
)

type Client struct { //create the client data structure, holds a cache
	cache      pokecache.Cache //this pulls the cache info from pokecache file
	httpClient http.Client     //creates an http client for serving requests
}

func NewClient(cacheInterval time.Duration) *Client { //initializes new client similar to new cache initialization
	//takes the interval to reap and returns a POINTER to the client, ensuring no copies are being primaried
	return &Client{ //return the pointer to this Client
		cache: *pokecache.NewCache(cacheInterval), //the cache is a pointer to the new cache generated for this client
		httpClient: http.Client{
			Timeout: time.Minute, //creates an HTTP client with a 5 minute timeout
		},
	}
}

// ttime to add the method onto the client that will perform the get request, although am I also doing this in main?
func (c *Client) GetRequest(url string) ([]byte, error) { //adds the method to the client struct, takes a url request and returns a slice of bytes and error
	if cachedData, ok := c.cache.Get(url); ok { //attempt to set cachedData to a getURL method, if it's ok:
		fmt.Println("Cache accessed for URL: ", url)
		return cachedData, nil

	}
	//if it's not in the cache, make the request
	resp, err := c.httpClient.Get(url) //resp is the result
	if err != nil {                    //if error return the error
		return nil, err
	}
	defer resp.Body.Close() //make sure this doesn't leak

	data, err := io.ReadAll(resp.Body) //ask about this ioreadall
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, data) //add the data to the cache for the future using the add method
	return data, nil       //return the data requested

}
