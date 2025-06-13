package main

import(
	"log"
	"time"
	"strconv"
	"strings"
	"context"
	"net/http"
	"io/ioutil"
	
	"golang.org/x/net/html"
	"github.com/andybalholm/cascadia"
)

func search() []string {
	var links []string
	for p:=1; p<2; p++ {
		newLinks := getPage(p)
		links = append(links, newLinks...)
	}
	
	return links
}

func getPage(page int) []string {
	url := "http://insecam.org/en/bynew/?page=" + strconv.Itoa(page)
	
	rawHTML, err := request(url)
	if err != nil {
		log.Println(err)
		return []string{}
	}
	
	doc := parse(rawHTML)
	return GetLinks(doc)
}

func request(url string) (string, error) {
	//Timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second * 10))
	defer cancel()
	
	//Request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	
	//Download
	client := &http.Client{}
	res, reqErr := client.Do(req)
	if reqErr != nil {
		return "", reqErr
	}
	defer res.Body.Close()
	
	//Read
	byt, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return "", readErr
	}
	
	return string(byt[0:]), nil
}

func parse(rawHTML string) *html.Node {
	doc, err := html.Parse(strings.NewReader(rawHTML))
	if err != nil {
		log.Println(err)
		return nil
	}
	return doc
}

func GetLinks(doc *html.Node) []string {
	links := []string{}
	row := QuerySelectorAll(doc, ".row")
	if len(row) == 0 {
		log.Println("No row class")
		return []string{}
	}
	
	
	for n := range row[0].ChildNodes() {
		imgs := QuerySelectorAll(n, "img")
		
		for _, i := range imgs {
			links = append(links, GetAttribute(i, "src"))
		}
	}
	return links
}

func QuerySelectorAll(n *html.Node, query string) []*html.Node {
	selector, err := cascadia.Parse(query)
	if err != nil {
		return []*html.Node{}
	}
	return cascadia.QueryAll(n, selector)
}

func GetAttribute(n *html.Node, attrKey string) string {
	for _, attr := range n.Attr {
		if attr.Key == attrKey {
			return attr.Val
		}
	}
	return ""
}