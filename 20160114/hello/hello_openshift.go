package main

import (
	"fmt"
	"net"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	intfs, err := net.Interfaces()
	if err != nil {
		fmt.Print(err)
		return
	}
	for _, i := range intfs {
		if (i.Flags&net.FlagUp != 0) && (i.Flags&(net.FlagLoopback|net.FlagPointToPoint) == 0) {
			addr, err := i.Addrs()
			if err != nil {
				fmt.Print(err)
			}
			fmt.Fprint(w, "My IP is : ", addr, "\n")
		}
	}
}

func main() {
	http.HandleFunc("/", helloHandler)

	go func() {
		fmt.Println("serving on 8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}()

	go func() {
		fmt.Println("serving on 8888")
		err := http.ListenAndServe(":8888", nil)
		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}()
	select {}
}
