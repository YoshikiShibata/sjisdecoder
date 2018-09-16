// Copyright © 2018 Yoshiki Shibata. All rights reserved.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/encoding/japanese"
)

func main() {
	flag.Parse()

	var sjis = make([]byte, 4096)
	var sjisBuf bytes.Buffer

	for {
		n, err := os.Stdin.Read(sjis)
		if n > 0 {
			sjisBuf.Write(sjis[:n])
		}
		if err != nil {
			break
		}
	}

	sjisBytes := sjisBuf.Bytes()
	utf8Bytes := make([]byte, len(sjisBytes)*2)

	decoder := japanese.ShiftJIS.NewDecoder()
	nDst, _, err := decoder.Transform(utf8Bytes, sjisBytes, true)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Transform failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(strings.Replace(
		string(utf8Bytes[:nDst]),
		"\r", "", -1))
}
