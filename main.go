package main

import (
	"net/http"
	"bytes"
	"encoding/json"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/base64"
	"strings"
	"os"
	"flag"
	"time"
)

var url string = "https://texttospeech.googleapis.com/v1beta1/text:synthesize"
var googleapikey string = ""

func main() {

	sentencesFile := flag.String("sentences", "en-US-taranis.csv", "Path to csv files containing all the sentences")
	outputFolder := flag.String("output", ".", "Output folder for wav files")

	flag.Parse()

	if strings.HasSuffix(*outputFolder, "/") {
		o := *outputFolder
		b := o[:len(o)-1]
		outputFolder = &b
	}

	if !strings.HasPrefix(*outputFolder, "/") {
		o := *outputFolder
		b := "./" + o
		outputFolder = &b
	}

	sentencesContent, err := ioutil.ReadFile(*sentencesFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, sentence := range strings.Split(string(sentencesContent), "\n") {
		values := strings.Split(sentence, ";")
		folder := values[0]
		file := values[1]
		ttsInput := values[2]
		os.MkdirAll(*outputFolder+"/"+folder, os.ModePerm)
		err := ioutil.WriteFile(*outputFolder+"/"+folder+"/"+file, synthesize(ttsInput), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	ioutil.WriteFile("output.wav", synthesize("Hello, this is a test."), os.ModePerm)
}

func synthesize(input string) []byte {

	r := Request{
		Input:       SynthesisInput{Text: input},
		Voice:       VoiceSelectionParams{LanguageCode: "en-US", Name: "en-US-Wavenet-D", SsmlGender: MALE},
		AudioConfig: AudioConfig{AudioEncoding: LINEAR16}}

	jsonString, err := json.Marshal(r)

	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("charset", "utf-8")
	request.Header.Set("Authorization", "Bearer "+googleapikey)

	client := &http.Client{}
	httpResponse, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer httpResponse.Body.Close()

	fmt.Println("response Status:", httpResponse.Status)
	fmt.Println("response Headers:", httpResponse.Header)
	body, _ := ioutil.ReadAll(httpResponse.Body)
	fmt.Println("response Body:", string(body))

	var ttsResponse Response
	json.Unmarshal(body, &ttsResponse)

	reader := strings.NewReader(ttsResponse.AudioContent)
	b := base64.NewDecoder(base64.StdEncoding, reader)
	bytesWAV, err := ioutil.ReadAll(b)

	return bytesWAV
}
