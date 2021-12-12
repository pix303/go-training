package track

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var trackBuffer bytes.Buffer
	tracker := NewTracker(&trackBuffer)
	assert.NotNil(t, tracker)

	tracker.Track("test","test","test",5)
	assert.Contains(t, trackBuffer.String(), "test 5", "Should return true if it write something")
	tracker.Track("")
	assert.Contains(t, trackBuffer.String(), "test 5", "Should return true if it write something")

	tracker = NewTracker(nil)
	err := tracker.Track("return some error")
	if assert.Error(t, err){
		assert.Equal(t, ErrNoWriter,err)
	}

	assert.Contains(t, ErrNoWriter.Error(), "writer","Should contain the word \"writer\"" )

	var testError trackerError = "Error test"
	assert.Equal(t, testError.Error(), "Error test", "Should instantiate error") 
}
