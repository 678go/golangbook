package main

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main() {

	ctx, _ := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	//defer cancel()

	//执行任务
	url := "http://dl.xzg01.com:83/OpRoot/MemberScoreList.aspx?uid=0&op=0&uname=003008"
	err := chromedp.Run(ctx, VisitWeb(url,
		"wordpress_logged_in_74564337e1d2f71045d66f5a1af3d1b1", "1",
	))
	if err != nil {
		log.Fatal(err)
	}
	//var res string
	//for i := 1; i < 27170; i++ {
	//	//执行
	//	err = chromedp.Run(ctx, DoCrawler(&res)) //执行爬虫任务
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	if err := chromedp.Run(ctx, DoCrawler()); err != nil {
		log.Println(err)
	}
}

func VisitWeb(url string, cookies ...string) chromedp.Tasks {
	//创建一个chrome任务
	return chromedp.Tasks{
		//ActionFunc是一个适配器，允许使用普通函数作为操作。
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 设置Cookie存活时间
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			// 添加Cookie到chrome
			for i := 0; i < len(cookies); i += 2 {
				//SetCookie使用给定的cookie数据设置一个cookie； 如果存在，可能会覆盖等效的cookie。
				err := network.SetCookie(cookies[i], cookies[i+1]).
					// 设置cookie到期时间
					WithExpires(&expr).
					// 设置cookie作用的站点
					WithDomain("dl.xzg01.com:83"). //访问网站主体
					// 设置httponly,防止XSS攻击
					WithHTTPOnly(true).
					//Do根据提供的上下文执行Network.setCookie。
					Do(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		}),
		// 跳转指定的url地址
		chromedp.Navigate(url),
	}
}
func DoCrawler() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate("https://666java.com/12141.html"),
	}
}
