package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s \"first phrase\" \"second phrase\"\n", os.Args[0])
		os.Exit(1)
	}

	first := url.QueryEscape(os.Args[1])
	second := url.QueryEscape(os.Args[2])

	request := fmt.Sprintf("http://www.google.com/trends/fetchComponent?q=%s,%s&cid=TIMESERIES_GRAPH_0&export=3", first, second)
	fmt.Println(request)
	resp, err := http.Get(request)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))

	fmt.Printf("\"%s\" - %d results\n", first, 10000)
	fmt.Printf("\"%s\" - %d results\n", second, 100000)
	fmt.Println("Nothing to say")
}
