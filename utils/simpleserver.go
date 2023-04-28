package utils

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"
)

const (
	serverHost = "localhost:9999"
	dataPath   = "/api/v1/data&id=1"
	countPath  = "/api/v1/count"
)

var jsonData = []byte(`{
	"key": "test-key",
	"value": 123,
	"obj": {
		"prop1": null,
		"prop2": [1,2,3]
	}
}`)

func RunServer() {
	server1 := startServerWithRandomPort()
	defer server1.Close()

	server2 := startServerWithSpecifiedPort()
	defer server2.Close()

	for {

		fmt.Println("--- server1 ---")
		getData(server1.URL, dataPath)
		getData(server1.URL, countPath)

		fmt.Println("--- server2 ---")
		getData(server2.URL, dataPath)
		getData(server2.URL, countPath)

		<-time.After(3 * time.Second)
	}
}

func getData(baseUrl string, apiPath string) {
	urlPath, err := url.JoinPath(baseUrl, apiPath)
	if err != nil {
		fmt.Printf("failed to join path: %s\n", err)
	}

	response, err := http.Get(urlPath)
	if err != nil {
		fmt.Printf("failed to get data: %s\n", err)
	}
	defer response.Body.Close()

	fmt.Printf("status: %s\n", response.Status)
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("failed to read data: %s\n", err)
		}
		fmt.Println(string(bodyBytes))
	}

	p := make([]byte, 100000)
	dataLength, err := response.Body.Read(p)
	if err != nil {
		fmt.Printf("read error: %s\n", err)
	}
	fmt.Printf("body(%d) : %s\n", dataLength, p)
}

func startServerWithRandomPort() *httptest.Server {
	server := httptest.NewServer(handler())
	fmt.Printf("with random port: %s\n", server.URL)

	return server
}

func startServerWithSpecifiedPort() *httptest.Server {
	listener, err := net.Listen("tcp", serverHost)
	if err != nil {
		fmt.Println(err)
	}

	server := httptest.NewUnstartedServer(handler())
	server.Listener.Close()
	server.Listener = listener
	server.Start()

	fmt.Printf("with specified port: %s\n", server.URL)
	return server
}

func handler() http.Handler {
	count := 0

	return http.HandlerFunc(func(response http.ResponseWriter, req *http.Request) {
		// fmt.Printf("Host      : %s\n", req.Host)
		// fmt.Printf("URL       : %s\n", req.URL)
		// fmt.Printf("RequestURI: %s\n", req.RequestURI)
		// fmt.Printf("Method    : %s\n", req.Method)

		var responseBytes []byte

		if strings.Contains(req.URL.String(), dataPath) {
			responseBytes = jsonData
		} else if strings.Contains(req.URL.String(), countPath) {
			data := fmt.Sprintf("{ count: %d }", count)
			responseBytes = []byte(data)
		} else {
			response.WriteHeader(http.StatusBadRequest)
			return
		}

		response.WriteHeader(http.StatusOK)
		_, err := response.Write(responseBytes)
		if err != nil {
			fmt.Println(err)
		}
	})
}
