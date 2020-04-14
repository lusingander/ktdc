package main

import "fmt"

func main() {
	str := `foo(aaa=123,bbb="bbbstr",ccc=[10,20,30],ddd=bar(eee=null,fff=false))`

	i, p := 0, 0
	for {
		c := str[p]
		if c == ',' {
			fmt.Println(tab(i) + str[0:p+1])
			str = str[p+1:]
			p = 0
		} else if c == '(' || c == '[' {
			fmt.Println(tab(i) + str[0:p+1])
			str = str[p+1:]
			p = 0
			i++
		} else if c == ')' || c == ']' {
			fmt.Println(tab(i) + str[0:p])
			str = str[p:]
			p = 0
			i--
			if len(str) == 1 {
				// data class なので最後は必ず ')' になるはず
				fmt.Println(str)
				break
			}
		}
		p++
	}
}

func tab(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "  "
	}
	return s
}
