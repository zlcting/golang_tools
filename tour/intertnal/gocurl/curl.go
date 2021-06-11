package gocurl

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Get(url string) (string, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)

	//接收服务端返回给客户端的信息
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		str, _ := ioutil.ReadAll(response.Body)
		bodystr := string(str)
		return bodystr, nil
	}
	return "_", errors.New("未知错误")
}

func Post(posturl string, data map[string]string) (string, error) {
	client := &http.Client{}
	postValues := url.Values{}

	for k, v := range data {
		postValues.Add(k, v)
	}
	// postValues.Add("info", "")
	// postValues.Add("message", `DkCxcs0z6Z03uHWOHOASf2xen+7oNoSad+KG2ss0hkE79211GlgjepmMFRW4zLiF51pVYHHOBFDYYJrnokq5d0ceKYY6ONzbBYKCJMzD7guN3qMYf48Cl9g0bDVb1oMbuN2PstzORe800Q72moQaHVRPiqh7VZ6NCXnkLrtnY64=`)

	resp, err := client.PostForm(posturl, postValues)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
	}
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		return string(body), nil
	}

	return "_", errors.New("未知错误")
}
