package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
)

func main(){
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond * 100)

	// create HTTP request
	req, err := http.NewRequestWithContext(http.MethodGet, "http://placehold.it/2000x2000", nil)
	if err != nil {
		panic(err)
	}

	// perform HTTP request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	imageData, err := ioutil.ReadAll(res.Body)
	fmt.Printf("downloaded image size: %d\n", len(imageData))
	
}
