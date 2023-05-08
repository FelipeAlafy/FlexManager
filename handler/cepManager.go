package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CEP struct {
	CEP 		string	`json:"CEP"`
	Logradouro 	string	`json:"Logradouro"`
	Complemento string	`json:"Complemento"`
	Bairro 		string	`json:"Bairro"`
	Localidade 	string	`json:"Localidade"`
	UF 			string	`json:"UF"`
	Ibge 		string	`json:"Ibge"`
	Gia 		string	`json:"Gia"`
	DDD 		string	`json:"DDD"`
	Siafi 		string	`json:"Siafi"`
}

func GetAddress(cep string) CEP {
	var localCep CEP

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	response, err := http.Get(url)
	Error("controller/cepManager.go >> response, http.Get() in json return: ", err)
	
	data, err := ioutil.ReadAll(response.Body)
	Error("controller/cepManager.go >> data, ioutil.ReadAll() while trying to read data: ", err)
	defer response.Body.Close()

	json.Unmarshal(data, &localCep)

	return localCep
}