package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Node struct {
	Urls []*string `json:"urls"`
}

var node Node = Node{}
var client http.Client = http.Client{}

func (n *Node) getNodes(res http.ResponseWriter, req *http.Request) {
	js, _ := json.Marshal(n)
	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}

func (n *Node) updateNodes(res http.ResponseWriter, req *http.Request) {

	var node Node
	json.NewDecoder(req.Body).Decode(&node)
	n.Urls = node.Urls
}

func (n *Node) processNodes(res http.ResponseWriter, req *http.Request) {

	// apiPath, _ := mux.Vars(req)["apiPath"]
	httpMethod := req.Method
	var body interface{}
	json.NewDecoder(req.Body).Decode(&body)
	x, _ := json.Marshal(body)

	//get node using round robin
	nodeToHit := *node.getNode()
	fmt.Println("HEllo i get here")
	fmt.Println(req.URL.Path)
	reqq, _ := http.NewRequest(httpMethod, nodeToHit+req.URL.Path, bytes.NewBuffer(x))
	fmt.Println(req.URL.Path)
	resp, err := client.Do(reqq)
	if err != nil {
		panic(err)
	}

	// defer resp.Body.Close()
	y, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(y)
	bodyStringg := []byte(bodyString)
	res.Write(bodyStringg)
}

func (n *Node) getNode() *string {

	t := time.Now()
	t1 := t.Unix()
	t2 := int(t1)
	return n.Urls[t2%len(n.Urls)]
}

func main() {
	fmt.Println("Hi!!")
	router := mux.NewRouter()
	router.HandleFunc("/getNodes", node.getNodes).Methods("GET")
	router.HandleFunc("/createNodes", node.updateNodes).Methods("POST")
	router.PathPrefix("/").HandlerFunc(node.processNodes)

	http.ListenAndServe(":8000", router)
	fmt.Println("Success!!")

}
