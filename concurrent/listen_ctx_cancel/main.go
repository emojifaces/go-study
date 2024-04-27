package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// 监听取消事件
func main() {

	err := http.ListenAndServe(":8080", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		ctx := request.Context()

		select {
		case <-ctx.Done():
			fmt.Fprintln(os.Stdout, "request canceled")
		case <-time.After(3 * time.Second):
			writer.Write([]byte("hello world!"))
		}

	}))
	if err != nil {
		fmt.Println("error:", err)
	}
}
