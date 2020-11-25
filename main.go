package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type cringeWriter struct {
	target io.Writer
}

func newCringer(dst io.Writer) *cringeWriter {
	return &cringeWriter{
		target: dst,
	}
}

func (c *cringeWriter) Write(b []byte) (n int, err error) {
	s := string(b)
	for i, v := range s {
		var nt int
		if i%2 == 0 {
			nt, err = c.target.Write([]byte(strings.ToUpper(string(v))))
		} else {
			nt, err = c.target.Write([]byte(strings.ToLower(string(v))))
		}
		n += nt
	}
	return
}

func main() {
	args := os.Args[1:]

	c := newCringer(os.Stdout)
	if len(args) > 0 {
		in := ""
		for _, v := range args {
			in += v + " "
		}
		_, err := io.Copy(c, strings.NewReader(in))
		checkErr(err)
		fmt.Println()
	} else {
		in := bufio.NewReader(os.Stdin)
		_, err := io.Copy(c, in)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
