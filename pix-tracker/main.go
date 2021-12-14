package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/pix303/go-training/pix-tracker/track"
)

func main() {
	trk := track.NewTracker(os.Stdout)
	trk.Track("First message")
	trk.Track("Second message")

	var trackData string
	buffer := bytes.NewBufferString(trackData)
	trk2 := track.NewTracker(buffer)
	for i := 0; i < 10; i++{
		trk2.Track(fmt.Sprintf("track number %d",i) )
	}
	fmt.Println(buffer.String())
}