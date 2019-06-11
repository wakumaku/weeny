package integration

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
	"weeny/apiserver"
	"weeny/application"
	"weeny/cache"
	"weeny/hasher"
)

type response struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

var s *apiserver.ApiServer

const (
	serverPort = 4444
	baseURL    = "http://localhost:4444"
)

func startServer() {

	s = apiserver.NewServer(
		application.New(
			cache.NewRedis("redis", 6379, "", 0),
			&hasher.Md5{},
		),
	)

	if err := s.Start(serverPort); err != nil {
		// nice
	}
}

func waitForServer() {
	started := time.Now()
	for {
		if started.After(time.Now().Add(5 * time.Second)) {
			panic("Time out waiting for server to start")
		}
		resp, err := http.Get(baseURL + "/ping")
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

func TestGetNonExistingHash(t *testing.T) {

	expectedResponse := response{Message: "Failure"}

	resp, err := http.Get(baseURL + "/here_an_unexisting_hash")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response body: %s", err)
	}

	var r response
	if err := json.Unmarshal(body, &r); err != nil {
		t.Fatalf("unmarshalling response body: %s", err)
	}

	if !reflect.DeepEqual(expectedResponse, r) {
		t.Fatalf("%s : %s", body, err)
	}
}

func TestPostMalformedURL(t *testing.T) {

	bodyPost := `{"url":"d u ckduckgo.com"}`
	expectedPostResponse := response{
		Message: "Not a valid URL",
		Data:    "",
	}

	resp, err := http.Post(baseURL+"/shortern", "application/json", strings.NewReader(bodyPost))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response body: %s", err)
	}

	var r response
	if err := json.Unmarshal(body, &r); err != nil {
		t.Fatalf("unmarshalling response body: %s", err)
	}

	if !reflect.DeepEqual(expectedPostResponse, r) {
		t.Fatalf("%s : %s", body, err)
	}
}
func TestPostURLAndLookupHash(t *testing.T) {

	// POST an url
	bodyPost := `{"url":"http://duckduckgo.com"}`
	expectedPostResponse := response{
		Message: "Success",
		Data:    "d4faf77f2107085ad92c39bc47530014",
	}

	resp, err := http.Post(baseURL+"/shortern", "application/json", strings.NewReader(bodyPost))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response body: %s", err)
	}

	var r response
	if err := json.Unmarshal(body, &r); err != nil {
		t.Fatalf("unmarshalling response body: %s", err)
	}

	if !reflect.DeepEqual(expectedPostResponse, r) {
		t.Fatalf("%s : %s", body, err)
	}

	// GET a hash
	expectedGetResponse := response{
		Message: "Success",
		Data:    "http://duckduckgo.com",
	}

	resp, err = http.Get(baseURL + "/lookup/d4faf77f2107085ad92c39bc47530014")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response bodyx: %s", err)
	}

	if err := json.Unmarshal(body, &r); err != nil {
		t.Fatalf("unmarshalling response body: %s", err)
	}

	if !reflect.DeepEqual(expectedGetResponse, r) {
		t.Fatalf("%s : %s", body, err)
	}
}
