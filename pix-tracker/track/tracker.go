package track

import (
	"fmt"
	"io"
	"strings"
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

func (trk *tracker) Track(msg ...interface{}) error {
	var prefixMsg = "Message to print: %v"
	if trk.w == nil{
		return ErrNoWriter
	}

	var msgParsed []string
	for _,s := range msg{
		msgParsed = append(msgParsed, fmt.Sprintf("%v",s))
	}
	_, e := trk.w.Write([]byte(fmt.Sprintf(prefixMsg, strings.Join(msgParsed, " "))))
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
