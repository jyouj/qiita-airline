package controller

import (
  "log"
  "time"
  "strings"

  "github.com/PuerkitoBio/goquery"
  "github.com/sclevine/agouti"
)

// Start ChromeDriver
func ChromeController(url string) string {
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

  return content
}

// scraiping for Qiita Top Page
func QiitaScraper(url string) string {
  // Start ChromeDriver
  content := ChromeController(url)

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

// scraiping search result
func SearchScraper(url string) string {
  content := ChromeController(url)

  res := strings.NewReader(content)

  // read HTML
  doc, err := goquery.NewDocumentFromReader(res)
  if err != nil {
    log.Fatal(err)
  }

  var box QiitaBox
  var boxes QiitaBoxes

  doc.Find(".searchResult").EachWithBreak(func(i int, s *goquery.Selection) bool {
    // get title & author
    box.Title = s.Find("h1.searchResult_itemTitle > a").Text()
    box.Author = s.Find("div.searchResult_header > a").Text()

    // get URL
    uncorrectUrl, _ := s.Find("h1.searchResult_itemTitle > a").Attr("href")
    correctUrl := "https://qiita.com" + uncorrectUrl
    box.Url = correctUrl

    boxes = append(boxes, box)

    time.Sleep(1 * time.Second)

    if i == 0 {
      return false
    }
    return true
  })

  // check existance
  if len(boxes) > 0 {
    result := "こちらの記事が見つかりました！\n" + boxes[0].Title + "by" + boxes[0].Author + "\n" + boxes[0].Url
    return result
  } else {
    return "お探しのものは見つかりませんでした"
  }
}
