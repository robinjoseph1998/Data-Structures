package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://github.com:8000/learn?tutorial=githublearn&gitid=84894399hjj"

func main() {
	fmt.Println(myurl)

	//parsing the url

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("type is : %T\n", qparams)
	fmt.Println(qparams["tutorial"])
	fmt.Println(qparams["gitid"])

	for _, val := range qparams {
		fmt.Println("Params is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "github.com",
		Path:    "/tutorial",
		RawPath: "gitid=84894399hjj",
	}
	otherrURL := partsOfUrl.String()
	fmt.Println(otherrURL)

}
