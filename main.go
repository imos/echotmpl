package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// Convert expands a template.
func Convert(tmpl []byte) ([]byte, error) {
	output := make([]byte, 0, len(tmpl)*2)

	// isText is true iff tmpl[pos] is in text mode (not in code).
	isText := true
	// The code block is in the echo mode (<?="foo"?>).
	isEchoCode := false
	// Whether echo function starts.
	isEchoing := false

	pos := 0
	hasPrefix := func(prefix string) bool {
		return bytes.HasPrefix(tmpl[pos:], []byte(prefix))
	}
	for ; pos < len(tmpl); pos++ {
		if isText {
			leftEchoBracket := hasPrefix("<?=")
			leftBracket := hasPrefix("<?")

			if leftEchoBracket {
				pos += 2
			} else if leftBracket {
				pos += 1
			}

			if leftEchoBracket || leftBracket {
				if isEchoing {
					output = append(output, []byte("\");\n")...)
					isEchoing = false
				}
				isText = false
				isEchoCode = leftEchoBracket
				if leftEchoBracket {
					output = append(output, []byte("echo(")...)
				}
				continue
			}

			if !isEchoing {
				output = append(output, []byte("echo(\"")...)
				isEchoing = true
			}
			if strconv.IsPrint(rune(tmpl[pos])) {
				if tmpl[pos] == '\\' || tmpl[pos] == '"' || tmpl[pos] == '?' {
					output = append(output, '\\')
				}
				output = append(output, tmpl[pos])
			} else {
				output = append(output, []byte(fmt.Sprintf("\\x%02x", tmpl[pos]))...)
			}
		} else if hasPrefix("?>") {
			pos += 1
			if isEchoCode {
				output = append(output, []byte(");\n")...)
			}
			isText = true
			isEchoing = false
			isEchoCode = false
		} else {
			output = append(output, tmpl[pos])
		}
	}

	return output, nil
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read: %+v", err)
	}
	output, err := Convert(input)
	if err != nil {
		log.Fatalf("failed to convert: %+v", err)
	}
	fmt.Printf("%s", output)
}
