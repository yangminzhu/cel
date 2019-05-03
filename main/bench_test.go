package main

import (
	"github.com/google/cel-go/common/types"
	"github.com/yangminzhu/cel/attributes"
	"testing"
)

func BenchmarkEval(b *testing.B) {
	expressions := []struct{
		name string
		expr string
		expect bool
	}{
		{
			name: "only-headers",
			expect: true,
			expr: `
!(headers.ip in blacklist) &&
  ((headers.path.startsWith("v1") && headers.token in ["v1", "v2", "admin"]) ||
   (headers.path.startsWith("v2") && headers.token in ["v2", "admin"]) ||
   (headers.path.startsWith("/admin") && headers.token == "admin" && headers.ip in whitelist))
`,

		},
		{
			name: "const-list",
			expect: true,
			expr: `
!(headers.ip in ["10.0.1.5", "10.0.1.6", "10.0.1.6"]) &&
  ((headers.path.startsWith("v1") && headers.token in ["v1", "v2", "admin"]) ||
   (headers.path.startsWith("v2") && headers.token in ["v2", "admin"]) ||
   (headers.path.startsWith("/admin") && headers.token == "admin" && headers.ip in ["10.0.1.1", "10.0.1.2", "10.0.1.3"]))
`,
		},
		{
			name: "blacklist",
			expect: true,
			expr: `!(headers.ip in ["10.0.1.5", "10.0.1.6", "10.0.1.6"])`,
		},
		{
			name: "v1",
			expect: false,
			expr: `headers.path.startsWith("v1") && headers.token in ["v1", "v2", "admin"]`,
		},
		{
			name: "v2",
			expect: false,
			expr: `headers.path.startsWith("v2") && headers.token in ["v2", "admin"]`,
		},
		{
			name: "admin",
			expect: true,
			expr: `headers.path.startsWith("/admin") && headers.token == "admin" && headers.ip in ["10.0.1.1", "10.0.1.2", "10.0.1.3"]`,
		},
	}

	attr := attributes.MyActivation()
	for _, exp := range expressions {
		prg := getProgram(exp.expr)
		expect := types.Bool(exp.expect)
		b.Run(exp.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				val, _, _ := prg.Eval(attr)
				if val != expect {
					b.Errorf("bad")
				}
			}
		})
	}
}
