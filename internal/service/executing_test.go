package service

import (
	"bufio"
	"github.com/magiconair/properties/assert"
	"strings"
	"testing"
)

func Test_readOutput(t *testing.T) {

	tests := []struct {
		name string
		text string
	}{
		{
			name: "Test1",
			text: "test\ntest\n",
		},
		{
			name: "Empty",
			text: "",
		},
		{
			name: "Only new line",
			text: "\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			strReader := strings.NewReader(tt.text)
			stdoutReader := bufio.NewReader(strReader)
			var outputBuffer []byte
			ch := make(chan struct{}, 1)

			readOutput(stdoutReader, &outputBuffer, ch)

			assert.Equal(t, tt.text, string(outputBuffer))

			select {
			case <-ch:
				break
			default:
				t.Error("Channel must be closed")
			}

		})
	}
}
