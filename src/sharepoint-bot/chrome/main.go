package chrome

import (
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/client"
	"log"
	"context"
	"time"
	"sharepoint-bot/cfg"
)

func LoadAnnouncements() (string, error) {

	// create context
	ctxt, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	// create chrome
	c, err := chromedp.New(ctxt, chromedp.WithTargets(client.New(client.URL(cfg.Get().ChromeUrl)).WatchPageTargets(ctxt)), chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	site := cfg.Get().SharepointUrl
	var res string
	var screenshotBuf []byte
	err = c.Run(ctxt, getAnnouncements(site, &res, &screenshotBuf))
	if err != nil {
		return ``, err
	}
	return res, nil
}

func getAnnouncements(site string, res *string, screenshotBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(site),
		chromedp.WaitReady(`.announcements-list`),
		chromedp.OuterHTML(`.announcements-list`, res, chromedp.NodeReady),
		chromedp.CaptureScreenshot(screenshotBuf),
	}
}