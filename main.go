package main

import(
	"os"
	"os/signal"
	
	"log"
	"embed"
	"context"
	"net/http"
	"html/template"
	
	"github.com/julienschmidt/httprouter"
	"github.com/RileySun/Scaled/utils"
)

//Embed
//go:embed html/*
var HTMLFiles embed.FS
var scraper *Scraper

func main() {
	//Scraper
	scraper = NewScraper()

	//Server
	router := httprouter.New()
	router.GET("/", Handle)
	router.ServeFiles("/static/*filepath", http.Dir("html/static"))
		
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//graceful(cancel)
	utils.StartHTTPServer(ctx, "8080", router)
}

func Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tmpl, parseErr := template.ParseFS(HTMLFiles, "html/index.html")
	if parseErr != nil {
		log.Println("Dashboard Template Parse: ", parseErr)
	}
	
	results := scraper.Scrape()
	templateData := struct {
    	Results []*Result
	}{
		results,
	}
	
	
	tmpl.Execute(w, templateData)
}

func graceful(cancel func()) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	
	go func() {
		<-stop
		cancel()
	}()
}