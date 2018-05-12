package chrome

import (
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/client"
	"log"
	"context"
	"time"
	"sharepoint-bot/cfg"
	"io/ioutil"
	"os"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/cdp"
)

func LoadAnnouncements() (string, error) {

	// create context
	ctxt, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	// create chrome
	c, err := chromedp.New(ctxt, chromedp.WithTargets(client.New(client.URL(
		cfg.Get().ChromeUrl)).WatchPageTargets(ctxt)),
		chromedp.WithErrorf(log.Printf))
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
	saveScreenshot(screenshotBuf)
	return res, nil
}

func saveScreenshot(screenshotBuf []byte) {
	err := os.MkdirAll(`artifacts`, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't create artifacts folder: %s\n", err)
		return
	}
	err = ioutil.WriteFile(`artifacts/output.png`, screenshotBuf, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't create screenshot: %s\n", err)
		return
	}
}

func getAnnouncements(site string, res *string, screenshotBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(site),
		chromedp.WaitVisible(`.announcements-list`),
		chromedp.OuterHTML(`.announcements-list`, res, chromedp.NodeReady),
		setViewportAndScale(1920, 1080),
		chromedp.Sleep(5 * time.Second),
		chromedp.CaptureScreenshot(screenshotBuf),
	}
}

func setViewportAndScale(w, h int64) chromedp.ActionFunc {
	return func(ctxt context.Context, ha cdp.Executor) error {
		err := emulation.SetDeviceMetricsOverride(w, h, 1, false).Do(ctxt, ha)
		if err != nil {
			return err
		}
		return nil
	}
}