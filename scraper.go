package main

import(
	"log"
	"time"
	"strconv"
	"strings"
	"context"
	"net/url"
	"net/http"
	"io/ioutil"
	"path/filepath"
	
	"golang.org/x/net/html"
	"github.com/andybalholm/cascadia"
)



type Scraper struct {
	page int
}

type Result struct {
	Src string
	Location string
	Type string
}
//<a class="Control-Prev" href="./?p={{.Prev}}">&lt; Prev</a>

func NewScraper() *Scraper {
	scraper := &Scraper{
		page:1,
	}
	
	return scraper
}


func (s *Scraper) Scrape() []*Result {
	url := "http://insecam.org/en/bynew/?page=" + strconv.Itoa(s.page)
	
	rawHTML, err := s.request(url)
	if err != nil {
		log.Println("Download Err", err)
		return []*Result{}
	}
	
	doc := s.prepare(rawHTML)
	
	return s.parse(doc)
}

func (s *Scraper) request(url string) (string, error) {
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

func (s *Scraper) prepare(rawHTML string) *html.Node {
	doc, err := html.Parse(strings.NewReader(rawHTML))
	if err != nil {
		log.Println(err)
		return nil
	}
	return doc
}

func (s *Scraper) parse(doc *html.Node) []*Result {
	var results []*Result
	row := querySelectorAll(doc, ".row")
	if len(row) == 0 {
		log.Println("No row class")
		return []*Result{}
	}
	children := row[0].ChildNodes()
	
	for n := range children {
		imgs := querySelectorAll(n, "img")
		
		
		for _, i := range imgs {
			src := getAttribute(i, "src")
			urlNoQuery, err := url.Parse(src)
			if err != nil {
				log.Println(err)
			}
			
			res := &Result{
				Src:src,
				Location:getAttribute(i, "title")[20:],
				Type:filepath.Ext(urlNoQuery.Path)[1:],
			}
			results = append(results, res)
		}	
	}
	return results
}

func querySelectorAll(n *html.Node, query string) []*html.Node {
	selector, err := cascadia.Parse(query)
	if err != nil {
		log.Println("querySelectorAll err:", err)
		return []*html.Node{}
	}
	return cascadia.QueryAll(n, selector)
}

func getAttribute(n *html.Node, attrKey string) string {
	for _, attr := range n.Attr {
		if attr.Key == attrKey {
			return attr.Val
		}
	}
	return ""
}