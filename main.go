package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func tab(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "  "
	}
	return s
}

func scan(in io.Reader) string {
	s := bufio.NewScanner(in)
	s.Scan()
	return s.Text()
}

func run(in io.Reader, out io.Writer) {
	str := scan(in)

	i, p := 0, 0
	for {
		c := str[p]
		if c == ' ' {
			str = str[:p] + str[p+1:]
		} else if c == ',' {
			fmt.Fprintln(out, tab(i)+str[0:p+1])
			str = str[p+1:]
			p = 0
		} else if c == '(' || c == '[' {
			fmt.Fprintln(out, tab(i)+str[0:p+1])
			str = str[p+1:]
			p = 0
			i++
		} else if c == ')' || c == ']' {
			fmt.Fprintln(out, tab(i)+str[0:p])
			str = str[p:]
			p = 0
			i--
			p++
			if len(str) == 1 {
				// data class なので最後は必ず ')' になるはず
				fmt.Fprintln(out, str)
				break
			}
		} else {
			p++
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
