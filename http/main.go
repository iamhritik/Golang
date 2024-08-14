package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://homelane.com")
	if err != nil {
		fmt.Println("Error is:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}
