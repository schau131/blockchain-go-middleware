package main


import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}

func main() {
	// Declare a new router
	r := newRouter()

	http.ListenAndServe(":8080", r)

}


func handler(w http.ResponseWriter, r *http.Request) {
	
//	fmt.Fprintf(w, "Hello World!")
	fmt.Println(" Forwarding the request");
	response, err := http.Get("https://httpbin.org/ip")
	
	//client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
	//}
	//req, err := http.NewRequest("GET", "http://localhost:4000/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn=query&args=%5B%22a%22%5D", nil)
	//req.Header.Add("authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU0NjIzNDEsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1NjU0MjYzNDF9.rehI1nI0AZFV3NrKZAV6wlh9-4j5FF8jdDjoEAgIZzk")
	//req.Header.Add("Content-Type", "application/json")
	//response, err := client.Do(req)
	
	
	if err != nil {
		fmt.Printf("The Http request failed with error %s\n", err);
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		fmt.Fprintf(w, string(data))
	}
}