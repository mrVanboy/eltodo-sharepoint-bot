package main

import (
	"sharepoint-bot/attachments"
	"sharepoint-bot/webhook"
	"sharepoint-bot/cfg"
	"fmt"
	"os"
	"sharepoint-bot/chrome"
)

const maxSavedArrayLength = 20

func main() {
	err := cfg.Load()
	if err != nil {
		panic(err)
	}

	html, err := chrome.LoadAnnouncements()
	if err != nil {
		webhook.NotifyAboutErrors([]error{err})
		panic(err)
	}

	ann, err := attachments.Parse(html)
	if err != nil {
		webhook.NotifyAboutErrors([]error{err})
		panic(err)
	}

	oldAnn, err := attachments.LoadFromStorage()
	if err != nil {
		webhook.NotifyAboutErrors([]error{err})
		panic(err)
	}

	if oldAnn == nil {
		err := attachments.SaveToStorage(ann)
		if err != nil {
			webhook.NotifyAboutErrors([]error{err})
			panic(err)
		}
		return
	}

	newAnn := ann.GetNewEntries(oldAnn)
	if len(newAnn) == 0{
		fmt.Fprintln(os.Stdout, "Announcements aren't new. Exitting...")
		return
	}

	for _, v := range newAnn {
		a := webhook.Attachment{
			AuthorName: v.Category,
			Fallback:   v.Heading,
			Title:      v.Heading,
			Text:       v.Content,
			Footer:     `Sharepoint @ ` + v.Created,
		}
		webhook.NewAttachment(a)
	}

	err = webhook.Send()
	if err != nil {
		webhook.NotifyAboutErrors([]error{err})
		panic(err)
	}

	annForSave := append(newAnn, oldAnn...)
	if len(annForSave) > maxSavedArrayLength {
		annForSave = annForSave[:maxSavedArrayLength]
	}
	err = attachments.SaveToStorage(annForSave)
	if err != nil {
		webhook.NotifyAboutErrors([]error{err})
		panic(err)
	}

}
