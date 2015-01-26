package main

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

func resultJs(firstArg string, secondArg string) string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	_, err := os.Open(dir + "/fake_resp.html")
	result := "{blabal}"

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		request := fmt.Sprintf("http://www.google.com/trends/fetchComponent?q=%s,%s&cid=TIMESERIES_GRAPH_0&export=3", firstArg, secondArg)
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

		result = fmt.Sprintf("%s\n", string(contents))
	}

	return result
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s \"first phrase\" \"second phrase\"\n", os.Args[0])
		os.Exit(1)
	}

	first := url.QueryEscape(os.Args[1])
	second := url.QueryEscape(os.Args[2])
	result := resultJs(first, second)

	fmt.Println(result)

	re := regexp.MustCompile("{.*}")
	fmt.Printf("%q\n", re.FindString(result))

	fmt.Printf("\"%s\" - %d results\n", first, 10000)
	fmt.Printf("\"%s\" - %d results\n", second, 100000)
	fmt.Println("Nothing to say")
}
