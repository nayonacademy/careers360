package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func collectLinks(link string, count int) []string {
	// Golang all links are colllected in this function
	linklist := []string{}
	for i := 1; i < count; i++ {
		url := link + strconv.Itoa(i)
		// fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".blockHeading").Each(func(i int, s *goquery.Selection) {
			band, _ := s.Find("a").Attr("href")
			linklist = append(linklist, band)
		})

	}
	return linklist
}

// func csvdata() []string {
// 	data := []string{}
// 	link := "https://it.careers360.com/colleges/list-of-bca-mca-colleges-in-india?sort=popularity&page="
// 	count := 33
// 	mycollection := collectLinks(link, count)
// 	for i := 1; i < len(mycollection); i++ {
// 		var name, address, courses_offered, facilities string
// 		affiliated_list := []string{}
// 		stream_list := []string{}

// 	}

// 	return data
// }

func main() {
	t1 := time.Now()
	link := "https://it.careers360.com/colleges/list-of-bca-mca-colleges-in-india?sort=popularity&page="
	count := 33
	mycollection := collectLinks(link, count)
	// fmt.Println(mycollection)
	fmt.Println(len(mycollection))

	// data := []string{}
	for i := 1; i < len(mycollection); i++ {
		// var name, address, courses_offered, facilities string
		// affiliated_list := []string{}
		// stream_list := []string{}
		res, err := http.Get(mycollection[i])
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		
		fmt.Println(mycollection[i])

	}

	t2 := time.Now()
	fmt.Println(t2.Sub(t1).Seconds())
}
