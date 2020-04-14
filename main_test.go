package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	for _, test := range tests {
		in := bytes.NewBufferString(test.src)
		out := &bytes.Buffer{}
		run(in, out)
		if test.want != out.String() {
			t.Errorf("\nwant:\n%s\ngot:\n%s", test.want, out.String())
		}
	}
}

var tests = []struct {
	src  string
	want string
}{
	{
		`foo(aaa=123, bbb="bbbstr", ccc=[10, 20, 30], ddd=bar(eee=null, fff=false))`,
		`foo(
  aaa=123,
  bbb="bbbstr",
  ccc=[
    10,
    20,
    30
  ],
  ddd=bar(
    eee=null,
    fff=false
  )
)
`,
	},
	{
		`foo(aaa="hoge(fuga=123.45, piyo=234.56)", bbb=null)`,
		`foo(
  aaa="hoge(fuga=123.45, piyo=234.56)",
  bbb=null
)
`,
	},
}
