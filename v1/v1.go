package v1

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Run()  {
	client := http.Client{}
	res, err := client.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	if res == nil {
		fmt.Println("Empty response!")
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Couldn't read body!")
		return
	}

	fmt.Println(string(body))

}
