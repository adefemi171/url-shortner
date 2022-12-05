package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Tiny URL API consumption")

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
		os.Exit(1)
	}

	baseUrl := "http://tinyurl.com/api-create.php?url="
	urlToShorten := os.Args[1]
	getReqUrl := baseUrl + urlToShorten

	response, err := http.Get(getReqUrl)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		_, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
