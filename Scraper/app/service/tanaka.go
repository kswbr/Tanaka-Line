package service

import (
    "strings"
    "github.com/PuerkitoBio/goquery"
    "context"
    pb "../pb"
)

type MyTanakaService struct {
}

func (s *MyTanakaService) GetNews(ctx context.Context, message *pb.GetLatestNewsMessage) (*pb.LatestNewsResponse, error) {
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

  return &pb.LatestNewsResponse{
      Title: strings.TrimSpace(target),
      Link: strings.TrimSpace(link),
  }, nil
}
