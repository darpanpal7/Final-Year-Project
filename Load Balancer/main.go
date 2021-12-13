package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

const TYPE = "http"
const HOST = "127.0.0.1"
const PORT = "8080"
const LIM = 10

type Balancing interface {
	allocateServer() int
}

type RoundRobin struct {
	currServer      int
	numberOfServers int
}

func (r *RoundRobin) allocateServer() int {
	var serverNumber int = r.currServer
	r.currServer = (r.currServer + 1) % r.numberOfServers
	return serverNumber
}

type Random struct {
	numberOfServers int
}

func (r *Random) allocateServer() int {
	var serverNumber int = rand.Intn(r.numberOfServers)
	return serverNumber
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func getRandomServer() int {
	var serverNumber int = rand.Intn(LIM)
	return serverNumber
}

func getRequestURL(URIpath string) string {
	var serverNumber int = getRandomServer()
	var serverPort string = strconv.Itoa(3000 + serverNumber)
	var requestURL string = TYPE + "://" + HOST + ":" + serverPort + "/" + URIpath
	return requestURL
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Request on " + PORT)
	fmt.Fprintf(w, "Hello World!")
}

func makeRequest(w http.ResponseWriter, r *http.Request) {

	var URIpath string = r.URL.Path
	URIpath = URIpath[1:]
	fmt.Println("Request with id = " + URIpath)

	var requestURL string = getRequestURL(URIpath)

	resp, err := http.Get(requestURL)
	handleError(err)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var URIpath string = r.URL.Path

	switch URIpath {
	case "/favicon.ico":
		break
	case "/":
		sayHello(w, r)
		break
	default:
		makeRequest(w, r)
	}
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+PORT, nil)
}
