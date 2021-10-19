package communication

import (
	"fmt"
	"time"
)

var AllMessages = []Message{
	{From: "A", To: "B", Text: "Text A to B", Datetime: time.Now()},
	{From: "B", To: "A", Text: "Text B to A", Datetime: time.Now().AddDate(0,0,1)},
	{From: "A", To: "C", Text: "Text A to C", Datetime: time.Now().AddDate(0,0,2)},
	{From: "B", To: "C", Text: "Text B to C", Datetime: time.Now().AddDate(0,0,3)},
	{From: "C", To: "A", Text: "Text C to A", Datetime: time.Now().AddDate(0,0,4)},
	{From: "C", To: "B", Text: "Text C to B", Datetime: time.Now().AddDate(0,0,5)},
}

type Message struct {
	From string
	To string
	Text string
	Datetime time.Time
}

func (m Message) String() string {
	return fmt.Sprintf("Message: \nFrom: %v \nTo: %v \nText: %v \nDate: %v", m.From, m.To, m.Text, m.Datetime.Format("2006-01-02"))
}

