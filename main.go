package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const BrasilApiBaseUrl = "https://brasilapi.com.br"
const ViaCepBaseUrl = "https://viacep.com.br"

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	chBrasilApi := make(chan string)
	chViaCep := make(chan string)

	if len(os.Args) < 2 {
		fmt.Println("Error: Missing CEP argument!")
		waitGroup.Done()
		return
	}

	inputedCep := os.Args[1]
	go fetchBrasilApiCep(inputedCep, &waitGroup, chBrasilApi)
	go fetchViaCep(inputedCep, &waitGroup, chViaCep)
	printResponse(chBrasilApi, chViaCep, &waitGroup)

	waitGroup.Wait()
}

func printResponse(chBrasilApi <-chan string, chViaCep <-chan string, waitGroup *sync.WaitGroup) {
	select {
	case cepDataBrasilApi := <-chBrasilApi:
		fmt.Printf("Response got from %s: \n%s\n", BrasilApiBaseUrl, cepDataBrasilApi)
	case cepDataViaCep := <-chViaCep:
		fmt.Printf("Response got from %s: \n%s\n", ViaCepBaseUrl, cepDataViaCep)
	case <-time.After(1 * time.Second):
		fmt.Println("Error: Timeout!")
		waitGroup.Done()
	}
}

func fetchBrasilApiCep(cep string, waitGroup *sync.WaitGroup, ch chan<- string) {
	//time.Sleep(2 * time.Second) // Uncomment this line to test the timeout error or the other API
	BrasilApiUrl := BrasilApiBaseUrl + "/api/cep/v1/" + cep
	fetchCep(BrasilApiUrl, waitGroup, ch)
}

func fetchViaCep(cep string, waitGroup *sync.WaitGroup, ch chan<- string) {
	//time.Sleep(2 * time.Second) // Uncomment this line to test the timeout error or the other API
	ViaCepUrl := ViaCepBaseUrl + "/ws/" + cep + "/json/"
	fetchCep(ViaCepUrl, waitGroup, ch)
}

func fetchCep(apiUrl string, waitGroup *sync.WaitGroup, ch chan<- string) {
	defer waitGroup.Done()

	log.Printf("Fetching cep data from %s\n", apiUrl)
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Printf("Error fetching cep data from %s\n", apiUrl)
		return
	}

	defer resp.Body.Close()
	cepData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading cep data response from %s\n", apiUrl)
		return
	}

	ch <- string(cepData)
}
