package main

import (
		"github.com/PuerkitoBio/goquery"
		"google.golang.org/grpc"
    "strings"
    "fmt"
)

func main() {
  doc, err := goquery.NewDocument("https://kushi-tanaka.com/news/")
  if err != nil {
		panic(err)
  }

	var target string = ""
	var link string = ""

	doc.Find("div#news-wrap > div > section.news-box span.icon").Each(func(_ int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text == "キャンペーン" {
			ps := s.Parent().Parent().Parent().Find("dt > a")
			target, _ = ps.Html()
			link, _  = ps.Attr("href")
			return
		}
  })

	target = strings.TrimSpace(target)
	link = strings.TrimSpace(link)
	fmt.Print(target)
	fmt.Print(link)
}
