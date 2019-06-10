package integration

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
	"weeny/apiserver"
	"weeny/application"
	"weeny/cache"
	"weeny/hasher"
)

var s *apiserver.ApiServer

func startServer() {

	s = apiserver.NewServer(
		application.New(
			cache.NewRedis("redis", 6739, "", 0),
			&hasher.Md5{},
		),
	)

	if err := s.Start(4444); err != nil {
		// nice
	}
}

func waitForServer() {
	started := time.Now()
	for {
		if started.After(time.Now().Add(5 * time.Second)) {
			panic("Time out waiting for server to start")
		}
		resp, err := http.Get("http://localhost:4444/ping")
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		if string(body) == "Pong" {
			break // up and running!
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func TestMain(m *testing.M) {

	go startServer()

	waitForServer()

	r := m.Run()

	_ = s.Shutdown()

	os.Exit(r)
}

func TestPing(t *testing.T) {
	resp, err := http.Get("http://localhost:4444/ping")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if string(body) != "Pong" {
		t.Fatalf("%s : %s", body, err)
	}
}

// resp, err = http.PostForm("http://duckduckgo.com",
// 		url.Values{"q": {"github"}})
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err = ioutil.ReadAll(resp.Body)
// 	fmt.Println("post:\n", keepLines(string(body), 3))
