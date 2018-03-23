package http

import (
	"encoding/json"
	"io"
	"labX-graphql-go-graphq-gophers/pkg/starwars"
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)

type service2 struct {
}

func NewService2() (starwars.Service2, error) {
	s := service2{}
	return &s, nil
}

func httpClient(request string) (io.Reader, error) {
	url := "http://localhost:9002/graphql"
	payload := strings.NewReader(request)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return res.Body, nil
}

func (s *service2) Species(id string) []*starwars.Specie {
	request := "{\"query\":\"query ($id: ID) {\\n  species(id: $id) {\\n    designation\\n    language\\n  }\\n}\\n\",\"variables\":{\"id\":\"" + id + "\"}}"
	response, err := httpClient(request)
	if err != nil {
		return nil
	}
	var ret []*starwars.Specie
	a,_:=ioutil.ReadAll(response)
	fmt.Println(string(a))
	if err := json.NewDecoder(response).Decode(&ret); err != nil {
		return nil
	}
	return ret
}

func (s *service2) Homeworld(id string) *starwars.Homeworld {
	request := "{\"query\":\"query ($id: ID) {\\n  homeworld(id: $id) {\\n    name\\n  }\\n}\\n\",\"variables\":{\"id\":\"" + id + "\"}}"
	response, err := httpClient(request)
	if err != nil {
		return nil
	}
	var ret *starwars.Homeworld
	if err := json.NewDecoder(response).Decode(&ret); err != nil {
		return nil
	}
	return ret
}

func (s *service2) AllHomeworld() []*starwars.Homeworld {
	request := "{\"query\":\"{\\n  allHomeworlds {\\n    name\\n  }\\n}\\n\",\"variables\":{}}"
	response, err := httpClient(request)
	if err != nil {
		return nil
	}
	var ret []*starwars.Homeworld
	if err := json.NewDecoder(response).Decode(&ret); err != nil {
		return nil
	}
	return ret
}
