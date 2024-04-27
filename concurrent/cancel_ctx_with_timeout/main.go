package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// 基于时间的取消
func main() {

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("status code:", res.StatusCode)
}
