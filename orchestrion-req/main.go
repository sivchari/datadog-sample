package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data, err := GetSomeData(request.Context())
		if err != nil {
			http.Error(writer, fmt.Sprintf("Failed to get data: %v", err), http.StatusInternalServerError)
			return
		}
		writer.Write(data)
	})

	http.ListenAndServe(":8080", nil)
}

func GetSomeData(ctx context.Context) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second,
	}
	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost, "http://example.com",
		strings.NewReader("Hello, World!"))
	if err != nil {
		return nil, err
	}
	name := "sivchari"
	return handleRequest(ctx, req, client, name)
}

//dd:span my:tag req.Method name
func handleRequest(ctx context.Context, req *http.Request, client *http.Client, name string) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.Status)
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
