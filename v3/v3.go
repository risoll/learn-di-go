package v3

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var url string

type HttpClient interface {
	Get(string) (*http.Response, error)
}

func init() {
	flag.StringVar(&url, "url", "http://google.com", "Which URL do you want?")
	flag.Parse()
}

func send(client HttpClient, link string) error {
	res, err := client.Get(link)
	if err != nil {
		return err
	}

	if res == nil {
		fmt.Println("Empty response!")
		return errors.New("empty response")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Couldn't read body!")
		return err
	}

	fmt.Println(string(body))
	return nil
}

func Run()  {
	client := &http.Client{}
	err := send(client, url)
	if err != nil {
		panic(err)
	}
}
