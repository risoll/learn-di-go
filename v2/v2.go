package v2

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var url string

func init() {
	flag.StringVar(&url, "url", "http://google.com", "Which URL do you want?")
	flag.Parse()
}

func send(link string) error {
	client := http.Client{}
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
	err := send(url)
	if err != nil {
		panic(err)
	}
}
