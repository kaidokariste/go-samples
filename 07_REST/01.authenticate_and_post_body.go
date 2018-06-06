package main

import (
	"bytes"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

//Function for making POST request to provided API endpoint,
//putting also JSON body as attachment
func PostJSONBody(queryUrl string, bufJson *bytes.Buffer) (responseStatus string, responseDate string, r []byte) {
	//HTTP - Post Request
	req, err := http.NewRequest("POST", queryUrl, bufJson)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer MyToken")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("[http/Client]: %s\n", err)
	}
	//Close the reponse body if done reading
	defer resp.Body.Close()
	responseDate = resp.Header["Date"][0]
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("[ReadAllResponseBody]: %s\n", err)
	}
	return resp.Status, responseDate, body
}

func main() {
	//Prepare key:values for json marshaling
	Message := map[string]string{"from": "God", "color": "purple", "message": "Hello Mortals!", "notify": "true", "message_format": "text"}
	jsonMessage, _ := json.Marshal(Message)
	//Convert json
	bufToken := bytes.NewBuffer(jsonMessage)
	//Print message before sending
	log.Println(string(jsonMessage))
	//Collect the response
	vResponseStatus, vResponseDate, vBody := PostJSONBody("https://chat/v2/room/666/notification", bufToken)
	log.Println(vResponseStatus, vResponseDate, string(vBody))
}
