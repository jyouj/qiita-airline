# QiitaAirline
![For participating in the 7th Annual Charleston Sports Fest Given this day 13th of June in the year 2020](https://user-images.githubusercontent.com/41543603/56892548-f2b63700-6aba-11e9-9413-49e980b2e3b7.png)

This is a high-performance LINE Bot which searches Qiita!

## Usage
You add this LINE Bot to your LINE app as one friend. And you will use this searching function.

Please use [URL](https://line.me/R/ti/p/%40xfz6432g) or this QR to add a friend of LINE.

![line_qr](https://user-images.githubusercontent.com/41543603/56892457-a0751600-6aba-11e9-8e55-f9911f2b7ce4.png)

### command
`-s` is to search three articles from Qiita's top page. This LINE Bot will teach information, such as titles, authors and urls.

`-t <search word>` is to search a article from Qiita's search page. If you want to do multi-word search, the form of `-t <search word>` has to be `-t <word1+word2>`. `+` is necessary.

`-h` is to teach you how to use this LINE Bot.

## Dependencies
- [golang/go](https://github.com/golang/go)
- [line/line-bot-sdk-go](https://github.com/line/line-bot-sdk-go)
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [sclevine/agouti](https://github.com/sclevine/agouti)
- [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)
