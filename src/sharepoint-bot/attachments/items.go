package attachments

import (
	"time"
	"github.com/xrash/smetrics"
)

type Announcements []Announcement

func (as Announcements) GetNewEntries(oldAnnouncements Announcements) Announcements  {
	var newAnnouncements Announcements
	for _, a := range as {
		if a.isNew(oldAnnouncements){
			newAnnouncements = append(newAnnouncements, a)
		}
	}
	return newAnnouncements
}

type Announcement struct {
	Timestamp time.Time `json:"timestamp"`
	Category  string    `json:"category"`
	Heading   string    `json:"heading"`
	Content   string    `json:"-"`
	Created   string    `json:"created"`
}

func (a Announcement) isNew(oldAnnouncements Announcements) bool {
	for _, oldA := range oldAnnouncements {
		if a.Created != oldA.Created { continue	}
		similarity := smetrics.JaroWinkler(a.Heading, oldA.Heading, 0.7, 4)
		if similarity > 0.9 {
			return false
		}
	}
	return true
}