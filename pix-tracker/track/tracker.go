package track

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type trackerError string
func (te trackerError) Error() string{
	return string(te)
}

const(
	ErrNoWriter = trackerError("Tracker has no writer initialized")
)


type Tracker interface {
	Track(...interface{}) error
}

type tracker struct {
	w io.Writer
}

func (trk *tracker) Track(messages ...interface{}) error {
	if trk.w == nil{
		return ErrNoWriter
	}
	
	var parsedMessages []string
	for _,s := range messages{
		parsedMessages = append(parsedMessages, fmt.Sprintf("%v",s))
	}

	t := time.Now()
	_, e := trk.w.Write([]byte(fmt.Sprintf("%s: %v", t.Format("2006-01-02 15:04:05"), strings.Join(parsedMessages, " "))))
	if e != nil {
		return e
	}
	trk.w.Write([]byte("\n"))
	return nil
}

func NewTracker(writer io.Writer) Tracker {
	return &tracker{
		w: writer,
	}
}
