package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func second(){
	res, err := http.Get("http://nayon.net")
	if err !=nil{
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode !=200{
		log.Fatalf("Status code error : %d , %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil{
		log.Fatal(err)
	}

	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection){
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d : %s - %s", i, band, title)
	})

}