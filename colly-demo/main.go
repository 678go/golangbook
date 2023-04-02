package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"strings"
)

func main() {
	baseColly()
}

func baseColly() {

	c := colly.NewCollector()

	c.DetectCharset = true

	collector := c.Clone()

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("request: ", request.URL.String())
		cookie := "wordpress_74564337e1d2f71045d66f5a1af3d1b1=ylin%7C1680185137%7CQqB6u7pdVirH97pSKi6GPfQ8PP4bHFn0yNiYnFlX1sc%7C028c54a4e88837de9ad77fd39e7e438c4f9ad4a6c2ba3b81bbaa707a964cf843; PHPSESSID=ug7dcqjuhjb49huc9tfulnfv0b; Hm_lvt_1f43835170b4149de053b197f1b2c6bf=1679925841,1679926819; wordpress_test_cookie=WP+Cookie+check; wordpress_logged_in_74564337e1d2f71045d66f5a1af3d1b1=ylin%7C1680185137%7CQqB6u7pdVirH97pSKi6GPfQ8PP4bHFn0yNiYnFlX1sc%7C917a2b729a3df2dc8c09b3161ad70e72557559c5458f96219ba69f064e33e397; Hm_lpvt_1f43835170b4149de053b197f1b2c6bf=1680012374"
		err := c.SetCookies("https://666java.com", setCookieRaw(cookie))
		if err != nil {
			log.Println("登录失败,", err)
		}
		fmt.Println("登录成功")
	})

	// 筛选 dev[class=site] 匹配的
	c.OnHTML("div[class=placeholder]", func(e *colly.HTMLElement) {
		//	fmt.Println(e)
		e.ForEach("div>a", func(_ int, htmlElement *colly.HTMLElement) {
			url := htmlElement.Attr("href")
			htmlElement.ForEach("img", func(_ int, htmlElement *colly.HTMLElement) {
				title := htmlElement.Attr("alt")
				if title == "" {
					return
				}
				img := htmlElement.Attr("data-src")
				fmt.Println(url, title, img)
				collector.Visit(url)
			})
		})

	})

	collector.OnHTML("div[class=\"downinfo pay-box\"]", func(e *colly.HTMLElement) {
		e.ForEach("div>a", func(_ int, htmlElement *colly.HTMLElement) {
			fmt.Println(htmlElement)
		})
	})

	c.Visit("https://666java.com/tag/%E6%9F%90%E9%A9%AC")

}

// set cookies raw
func setCookieRaw(cookieRaw string) []*http.Cookie {
	// 可以添加多个cookie
	var cookies []*http.Cookie
	cookieList := strings.Split(cookieRaw, "; ")
	for _, item := range cookieList {
		keyValue := strings.Split(item, "=")
		// fmt.Println(keyValue)
		name := keyValue[0]
		valueList := keyValue[1:]
		cookieItem := http.Cookie{
			Name:  name,
			Value: strings.Join(valueList, "="),
		}
		cookies = append(cookies, &cookieItem)
	}
	return cookies
}
