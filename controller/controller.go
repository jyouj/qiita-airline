package controller

import (
  "log"
  "net/http"
  "strings"

  "github.com/PuerkitoBio/goquery"
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
    result := `QiitaAirlineのご利用方法のご案内\n
    -s: Qiitaの人気記事を3つ検索します\n
    -t <検索ワード>: 検索ワードにヒットしたQiita記事を3つお届けします\n
    -h: 使い方ガイドを表示します`
    return result
  }
  return result
}

// scraiping for Qiita
func QiitaSearch(url string) string {
  // request url
  res, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }

  // read HTML
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  // QiitaBox
  var box QiitaBox

  // QiitaBoxes
  var boxes QiitaBoxes

  // scraipe
  doc.Find("div.p-home_main").EachWithBreak(func(i int, s *goquery.Selection) bool {
    // get title & author
    box.Title = s.Find("a.tr-Item_title").Text()
    box.Author = s.Find("a.tr-Item_author").Text()
    // get URL
    uncorrectUrl, _ := s.Find("a.tr-Item_title").Attr("href")
    correctUrl := "https://qiita.com" + uncorrectUrl
    box.Url = correctUrl

    boxes = append(boxes, box)

    if i == 3 {
      return false
    }
    return true
  })

  // result
  result0 := "Qiitaを検索してます......\n"
  result1 := boxes[0].Title + "by" + boxes[0].Author + "\n" + boxes[0].Url + "\n"
  result2 := boxes[1].Title + "by" + boxes[1].Author + "\n" + boxes[1].Url + "\n"
  result3 := boxes[2].Title + "by" + boxes[2].Author + "\n" + boxes[2].Url
  result := result0 + result1 + result2 + result3

  return result
}
