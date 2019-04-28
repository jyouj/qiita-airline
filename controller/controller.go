package controller

import (
  "log"
  "time"
  // "net/http"
  "strings"

  "github.com/PuerkitoBio/goquery"
  "github.com/sclevine/agouti"
)

func QiitaController(message string) string {
  slice := strings.Split(message, " ")
  result := "`-h`と入力して操作方法を確認してください"

  // command checker
  switch slice[0] {
  case "-s":
    url := "https://qiita.com"
    result := QiitaSearch(url)
    return result
  case "-h":
    word1 := "QiitaAirline利用案内\n"
    word2 := "-s: Qiitaの人気記事を3つ検索します(少し遅いです。ティータイムをお楽しみください)\n"
    word3 := "-t <検索ワード>: 検索ワードにヒットしたQiita記事を3つお届けします\n"
    word4 := "-h: 使い方ガイドを表示します"
    result := word1 + word2 + word3 + word4
    return result
  }
  return result
}

// scraiping for Qiita
func QiitaSearch(url string) string {
  // Start ChromeDriver
  driver := agouti.ChromeDriver(agouti.Browser("chrome"))
  if err := driver.Start(); err != nil {
    log.Fatalf("Failed to start driver:%v", err)
  }
  defer driver.Stop()

  page, err := driver.NewPage()
  if err != nil {
    log.Fatalf("Failed to open page:%v", err)
  }

  if err := page.Navigate(url); err != nil {
    log.Fatalf("Failed to navigate:%v", err)
  }

  content, err := page.HTML()
  if err != nil {
    log.Printf("Failed to get html: %v", err)
  }

  res := strings.NewReader(content)

  // read HTML
  doc, err := goquery.NewDocumentFromReader(res)
  if err != nil {
    log.Fatal(err)
  }

  // QiitaBox
  var box QiitaBox

  // QiitaBoxes
  var boxes QiitaBoxes

  // scraipe
  doc.Find(".tr-Item").EachWithBreak(func(i int, s *goquery.Selection) bool {
    // get title & author
    box.Title = s.Find("a.tr-Item_title").Text()
    box.Author = s.Find("a.tr-Item_author").Text()

    // get URL
    uncorrectUrl, _ := s.Find("a.tr-Item_title").Attr("href")
    correctUrl := "https://qiita.com" + uncorrectUrl
    box.Url = correctUrl

    boxes = append(boxes, box)

    // sleep not to place stress on the server.
    time.Sleep(1 * time.Second)

    if i == 2 {
      return false
    }
    return true
  })

  // result
  result0 := "Qiitaを検索してます......\n"
  result1 := boxes[0].Title + " by " + boxes[0].Author + "\n" + boxes[0].Url + "\n"
  result2 := boxes[1].Title + " by " + boxes[1].Author + "\n" + boxes[1].Url + "\n"
  result3 := boxes[2].Title + " by " + boxes[2].Author + "\n" + boxes[2].Url
  result := result0 + result1 + result2 + result3

  return result
}
