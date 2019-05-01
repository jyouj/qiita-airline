package controller

import (
  "strings"
)

func QiitaController(message string) string {
  slice := strings.Split(message, " ")
  result := "`-h`と入力して操作方法を確認してください"

  // command checker
  switch slice[0] {
  case "-s":
    url := "https://qiita.com"
    result := QiitaScraper(url)
    return result
  case "-t":
    if len(slice) > 1 {
      url := "https://qiita.com/search?q=" + slice[1]
      result := SearchScraper(url)
      return result
    } else {
      result := "検索したい言葉を入力してください"
      return result
    }
  case "-a":
    url := "https://qiita.com"
    result := AuthorScraper(url)
    return result
  case "-m":
    url := "https://qiita.com/milestones"
    result := MilestonesScraper(url)
    return result
  case "-c":
    result := "お問い合わせはこちらのTwitterアカウントまでお寄せください\nhttps://twitter.com/jyouj__"
    return result
  case "-h":
    word1 := "QiitaAirline利用案内\n"
    word2 := "-s: Qiitaの人気記事を3つ検索します(少し遅いかも。ティータイムをお楽しみください)\n"
    word3 := "-t <検索ワード>: 検索ワードにヒットしたQiita記事を1つお届けします\n"
    word4 := "-a: 今週の人気ユーザーをお伝えします！\n"
    word5 := "-m: マイルストーンを検索します\n"
    word6 := "-h: 使い方ガイドを表示します\n"
    word7 := "-c: お問い合わせ先を表示します"
    result := word1 + word2 + word3 + word4 + word5 + word6 + word7
    return result
  }
  return result
}
