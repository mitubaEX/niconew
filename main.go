package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getDistinctSlice(slice []string) []string {
	m := make(map[string]bool)
	uniq := []string{}

	for _, e := range slice {
		if !m[e] {
			m[e] = true
			uniq = append(uniq, e)
		}
	}
	return uniq
}

func nicoScrape(pageNum int) string {
	// Request the HTML page.
	res, err := http.Get(fmt.Sprintf("https://www.nicovideo.jp/newarrival?page=%d", pageNum))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var urlList = []string{}
	var titleList = []string{}
	var imgList = []string{}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		if strings.Contains(url, "watch") {
			urlList = append(urlList, "https://www.nicovideo.jp/"+url)
			titleList = append(titleList, s.Text())
		}
	})

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("src")
		if strings.Contains(url, "smile") {
			imgList = append(imgList, url)
		}
	})

	distinctUrlList := getDistinctSlice(urlList)
	distinctTitleList := getDistinctSlice(titleList)
	distinctImgList := getDistinctSlice(imgList)

	// create result string
	var result = ""
	for i, _ := range distinctUrlList {
		result += fmt.Sprintf("<a href=\"%s\" class=\"watch\" title=\"%s\">\n", distinctUrlList[i], distinctTitleList[i+1])
		result += fmt.Sprintf("<img src=\"%s\">", distinctImgList[i])
	}

	return result
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// create html
	var result = ""
	result += "<html><header/><body>"
	result += "<meta http-equiv=\"content-type\" charset=\"utf-8\">"
	for i := 1; i <= 10; i++ {
		result += nicoScrape(i)
	}
	result += "</body></html>"

	// return html
	fmt.Fprintf(w, result)
}

func main() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
