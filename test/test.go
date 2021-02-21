package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetCourseResult struct {
	Course_id uint64
	Fullname  string
	Timeend   uint64
}

func main() {

	url := "http://200.14.84.16/api/v1/moodle/getCourse"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Host = "external-udp.xn--ensea-rta.cl"
	req.Header.Add("Email", "manuel.alba@mail.udp.cl")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewDecoder()
}
