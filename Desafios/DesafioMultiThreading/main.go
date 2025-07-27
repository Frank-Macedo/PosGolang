package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type BrasilApiCep struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Não foi informado Cep como argumento")
	}
	cep := os.Args[1]

	chBrasilApi := make(chan []byte)
	chViaCep := make(chan []byte)
	done := make(chan struct{})

	urls := []string{
		fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep),
		fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep),
	}

	go GetAdress(urls[0], chBrasilApi, done)
	go GetAdress(urls[1], chViaCep, done)

	select {
	case msg := <-chBrasilApi:
		close(done)
		var address BrasilApiCep
		err := json.Unmarshal(msg, &address)
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Printf("Recebido de BrasilApi: %v\n", address)

	case msg := <-chViaCep:
		close(done)
		var address ViaCep
		err := json.Unmarshal(msg, &address)
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Printf("Recebido de ViaCep: %v\n", address)

	case <-time.After(time.Second * 1):
		close(done)
		println("Timeout")

	}

}

func GetAdress(url string, ch chan []byte, done chan struct{}) {

	req, err := http.Get(url)
	if err != nil {
		log.Printf("Erro ao acessar Url %s: %v", url, err)
		return
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Erro ao ler o body de %s: %v", url, err)
		return
	}
	select {
	case ch <- res:
	case <-done:
	}
}
