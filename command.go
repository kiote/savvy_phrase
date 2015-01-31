package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

/*
 * get response from google trends (as string)
 */
func resultJs(firstArg string, secondArg string) string {
	request := fmt.Sprintf("http://www.google.com/trends/fetchComponent?q=%s,%s&cid=TIMESERIES_GRAPH_0&export=3", firstArg, secondArg)
	// request := fmt.Sprintf("https://raw.githubusercontent.com/kiote/savvy_phrase/master/fake_resp.html?a=%s,%s", firstArg, secondArg)
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

	result := fmt.Sprintf("%s\n", string(contents))

	return result
}

/*
 * returns two phrases
 */
func getArgs() (string, string) {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s \"first phrase\" \"second phrase\"\n", os.Args[0])
		os.Exit(1)
	}

	return url.QueryEscape(os.Args[1]), url.QueryEscape(os.Args[2])
}

/*
 * converts response from google trends to json-string
 */
func resultJson(result string) string {
	re := regexp.MustCompile("{.*}")
	jsonString := re.FindString(result)
	jsonString = strings.Replace(jsonString, "\\", "", -1)

	return jsonString
}

func main() {
	result := resultJs(getArgs())
	result = resultJson(result)

	result = result[len(result)-100:]

	re := regexp.MustCompile(`"f":"(\d+)"`)
	fs := re.FindAllStringSubmatch(result, -1)

	fmt.Printf("\"%s\" - %s results\n", os.Args[1], fs[0][1])
	fmt.Printf("\"%s\" - %s results\n", os.Args[2], fs[1][1])
}
