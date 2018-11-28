package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server is running")
	})

	// BEGIN shutdown
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)
		fmt.Println("Wait for 2 seconds to finish processing")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()

	// END shutdown

	fmt.Println("Hello World!")

	http.ListenAndServe(":8080", nil)

	/*
		a := []int{5, 4, 1, 2, 8}
		b := challenges.Digits(a)
		fmt.Println("---")
		fmt.Println(b)
	*/
}
