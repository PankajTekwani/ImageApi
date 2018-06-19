package main

import(
	"fmt"
	"log"
	"bufio"
	"os"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	//"time"
	"github.com/gorilla/mux"
	//"model"
)

var cnt = 1
var max_arr_len = 11
var search_result = 10

/*
This handle looks for the tag (recieved from the user) in the TAG_MAP data structure and returns the json containing the list of the images that most probably contains that tag.
*/
func searchTagMap(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	tag := params["tag"]
	output := search(tag)
	//fmt.Println(output)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
	return
}


func hitPredictEndPoint(endPoint string, imageUrl string) {
	fmt.Print(cnt,", ")
	cnt = cnt + 1
	client := &http.Client{
	}

	jsonBody := `{
    "inputs": [
      {
        "data": {
          "image": {
            "url": "` + imageUrl + `"
          }
        }
      }
    ]
}`

	var jsonStr = []byte(jsonBody)
	
	request, err := http.NewRequest("Post", endPoint, bytes.NewBuffer(jsonStr))

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Key d4020ea7489a485baaa05e69dcb4264a")
	response, err := client.Do(request)
	
	defer response.Body.Close()
	
	if err != nil {
        fmt.Println("Http Post request failed")
    } else {
		responseData, _ :=ioutil.ReadAll(response.Body)
		var resJson Response
		
		json.Unmarshal([]byte(string(responseData)), &resJson)
		for _, obj := range resJson.Outputs[0].Data.Concepts {
			//fmt.Println(obj.Name, obj.Value)
			addToTagMap(obj.Name, imageUrl, obj.Value)
		}
		//fmt.Println(resJson.Outputs[0].Data.Concepts[0].Name)
	}
}


func main() {

	FILE_PATH := "images.txt"
    file, err := os.Open(FILE_PATH)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	
	endPoint := "https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs"
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		hitPredictEndPoint(endPoint, scanner.Text())
        //fmt.Println(scanner.Text())
    }
	//fmt.Println(TAG_MAP)
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	
	router := mux.NewRouter()
	router.HandleFunc("/{tag}", searchTagMap).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Server Started")
}
