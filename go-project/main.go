package main


import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"github.com/gorilla/mux"
	"encoding/json"
	"os"
	"bytes"
	"strings"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", registerUser).Methods("GET")
	r.HandleFunc("/create-channel", createChannel).Methods("GET")
	r.HandleFunc("/join-channel", joinChannel).Methods("GET")
	r.HandleFunc("/install-chaincodes", installChaincodes).Methods("GET")
	r.HandleFunc("/get-balance", getBalance).Methods("GET")
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}

func main() {
	// Declare a new router
	r := newRouter()

	http.ListenAndServe(":8080", r)

}

type RegisterRequest struct{
	Username string `json:"username"`
	Orgname string `json:"orgName"`
}

type RegisterResponse struct{
	Success bool `json:"success"`
	Secret string `json:"secret"`
	Message string `json:"message"`
	Token string `json:"token,omitempty"`
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println(" Forwarding the request")
	
//	register := RegisterRequest{}
//	register.Username = r.URL.Query().Get("username")
//	register.Orgname = r.URL.Query().Get("orgName")
	
//	var jsonStr []byte
//	jsonStr, err := json.Marshal(register)
	
//	fmt.Println(string(jsonStr))
	
	data := url.Values{}
	data.Set("username", r.URL.Query().Get("username"))
	data.Set("orgName", r.URL.Query().Get("orgName"))
	
	fmt.Println(strings.NewReader(data.Encode()))
	
	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
	}
	
	req, err := http.NewRequest("POST", "http://localhost:4000/users", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(req)
	
	
	if err != nil {
		fmt.Printf("The Http request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		
		var resJson RegisterResponse
		json.Unmarshal(data, &resJson)
		
		f, err := os.Create("token.txt")
		
		if err != nil {
			fmt.Printf("The Http request failed with error %s\n", err)
		}
		
		f.WriteString(resJson.Token)
		f.Close()
		
		resJson.Token = ""
		
		resJsonStr, err := json.Marshal(resJson)
		fmt.Fprintf(w, string(resJsonStr))
	}
}

func createChannel(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("Creating Channel")
	
	jsonStr := []byte(`{"channelName":"mychannel","channelConfigPath":"../artifacts/channel/mychannel.tx"}`)
		
	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
	}
	
	file, err := os.Open("token.txt")
	defer file.Close()
	b, _ := ioutil.ReadAll(file)
	authToken := "Bearer " + string(b)

	fmt.Println(authToken)
	
	req, err := http.NewRequest("POST", "http://localhost:4000/channels", bytes.NewReader(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", authToken)
	response, err := client.Do(req)
	
	
	if err != nil {
		fmt.Printf("The Http request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		fmt.Fprintf(w, string(data))
	}
}

func joinChannel(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("Joining Channel")
	
	jsonStr := []byte(`{"peers": ["peer0.org1.example.com","peer1.org1.example.com"]}`)
		
	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
	}
	
	file, err := os.Open("token.txt")
	defer file.Close()
	b, _ := ioutil.ReadAll(file)
	authToken := "Bearer " + string(b)

	fmt.Println(authToken)
	
	req, err := http.NewRequest("POST", "http://localhost:4000/channels/mychannel/peers", bytes.NewReader(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", authToken)
	response, err := client.Do(req)
	
	
	if err != nil {
		fmt.Printf("The Http request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		fmt.Fprintf(w, string(data))
	}
}

func installChaincodes(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("Joining Channel")
	
	jsonStr := []byte(`{"peers": ["peer0.org1.example.com","peer1.org1.example.com"],
						"chaincodeName":"mycc",
						"chaincodePath":"github.com/example_cc/go",
						"chaincodeType": "golang",
						"chaincodeVersion":"v0"}`)
		
	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
	}
	
	file, err := os.Open("token.txt")
	defer file.Close()
	b, _ := ioutil.ReadAll(file)
	authToken := "Bearer " + string(b)

	fmt.Println(authToken)
	
	req, err := http.NewRequest("POST", "http://localhost:4000/chaincodes", bytes.NewReader(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", authToken)
	response, err := client.Do(req)
	
	
	if err != nil {
		fmt.Printf("The Http request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		fmt.Fprintf(w, string(data))
	}
}


func instantiateChaincodes(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("Joining Channel")
	
	jsonStr := []byte(`{"chaincodeName":"mycc",
						"chaincodePath":"github.com/example_cc/go",
						"chaincodeType": "golang",
						"args":["a","100","b","200"]}`)
		
	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
	}
	
	file, err := os.Open("token.txt")
	defer file.Close()
	b, _ := ioutil.ReadAll(file)
	authToken := "Bearer " + string(b)

	fmt.Println(authToken)
	
	req, err := http.NewRequest("POST", "http://localhost:4000/channels/mychannel/chaincodes/mycc", bytes.NewReader(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", authToken)
	response, err := client.Do(req)
	
	
	if err != nil {
		fmt.Printf("The Http request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		fmt.Fprintf(w, string(data))
	}
}


func getBalance(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("fetching balance")
		
	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
	}
	
	file, err := os.Open("token.txt")
	defer file.Close()
	b, _ := ioutil.ReadAll(file)
	authToken := "Bearer " + string(b)

	fmt.Println(authToken)
	
	req, err := http.NewRequest("GET", "http://localhost:4000/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn=query&args=%5B%22a%22%5D")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", authToken)
	response, err := client.Do(req)
	
	
	if err != nil {
		fmt.Printf("The Http request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		fmt.Fprintf(w, string(data))
	}
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