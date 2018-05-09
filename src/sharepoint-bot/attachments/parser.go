package attachments

import (
	"github.com/antchfx/htmlquery"
	"strings"
	"regexp"
	"golang.org/x/net/html"
	"time"
)

func Parse(htmlString string) (Announcements, error){
	var announcementArr Announcements
	t := time.Now()
	h := htmlString
	cleanupHtmlString(&h)
	doc, err := htmlquery.Parse(strings.NewReader(h))
	if err != nil {
		return announcementArr, err
	}
	htmlquery.FindEach(doc, `//div[contains(@class, 'announcements-item-overflow')]`, func(_ int, node *html.Node) {
		a := Announcement{
			Timestamp: t,
			Category:  "",
			Heading:   "",
			Content:   "",
		}
		catAttr := htmlquery.FindOne(node, `//div[contains(@class, 'announcements-item-category')]/a/img/@alt`)
		if catAttr != nil {
			a.Category = htmlquery.SelectAttr(catAttr, `alt`)
		}

		if c := htmlquery.FindOne(node, `//div[contains(@class,'announcements-item-created')]/text()`); c != nil {
			a.Created = c.Data
		}

		if h := htmlquery.FindOne(node, `//*[substring(name(), 0, 1) = "h"]/text()`); h != nil {
			a.Heading = h.Data
		}

		a.Content = parseAnnouncementBodyNode(node)
		cleanupContent(&a.Content)
		announcementArr = append(announcementArr, a)
	})
	return announcementArr, nil
}
func cleanupContent(content *string) {
	rx := regexp.MustCompile(`((\*\s*\*)|(_\s*_))`)
	*content = rx.ReplaceAllString(*content, ``)
}

func parseAnnouncementBodyNode(node *html.Node) string {
	content := ``
	expressions := []string{
		`//div[contains(@class,'announcements-item-body')]/div/div`,
		`//div[contains(@class,'announcements-item-body')]/div/p`,
		`//div[contains(@class,'announcements-item-body')]/div`,
	}

	for _, xpath := range expressions {
		itemBodyNodes := htmlquery.Find(node, xpath)
		for i := range itemBodyNodes {
			content = htmlquery.InnerText(itemBodyNodes[i])
			if len(strings.TrimSpace(content)) > 0 {
				return content
			}
		}
	}
	return content
}

func cleanupHtmlString(s *string) {
	rx := regexp.MustCompile(`(?i)<strong[^>]*>([^<]*[^;])</strong>`)
	res := rx.ReplaceAllString(*s, "*${1}*")

	rx = regexp.MustCompile(`<font[^>]*>([^<]*)</font>`)
	res = rx.ReplaceAllString(res, `_${1}_`)

	*s = res
}