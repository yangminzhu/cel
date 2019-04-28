package main

import (
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/yangminzhu/cel/attributes"
	"github.com/yangminzhu/cel/util"
	"log"

	celgo "github.com/google/cel-go/cel"
)

var (
	printAST bool
	spewCfg = spew.NewDefaultConfig()
)

func init() {
	flag.BoolVar(&printAST, "ast", false, "print parsed AST in YAML, skip evaluation")
	spewCfg.SortKeys = true
	spewCfg.DisableCapacities = true
}

func main() {
	flag.Parse()

	expression := flag.Arg(0)
	if expression == "" {
		log.Fatalln("CEL expressions must not be empty")
	}

	celenv, err := celgo.NewEnv(attributes.EnvOpt())
	if err != nil {
		log.Fatalf("cel environment creation error: %s\n", err)
	}

	ast, iss := celenv.Parse(expression)
	if iss != nil {
		log.Printf("AST parse issues:\n%s\n", iss)
	}
	ast, iss = celenv.Check(ast)
	if iss != nil {
		log.Printf("AST check issues:\n%s\n", iss)
	}
	if ast == nil {
		return
	}
	if printAST {
		yml, err := util.ToYAML(ast.Expr())
		if err != nil {
			log.Fatalf("converting AST to yaml error: %v", err)
		} else {
			fmt.Printf("%s\n\n", yml)
		}
		return
	}

	prg, err := celenv.Program(ast)
	if err != nil {
		log.Fatalf("generating program error: %s", err)
	}
	yml, err := util.ToYAML(&attributes.Example)
	if err != nil {
		log.Printf("converting Example attributes to yaml error: %v", err)
	}
	log.Printf("using attributes:\n%s\n", yml)
	val, detail, err := prg.Eval(attributes.MyActivation())
	if detail != nil {
		log.Printf("details: %v", detail)
	}
	if err != nil {
		log.Fatalf("evaluation error: %s", err)
	} else {
		log.Printf("evaluated successfully: %v (%s)", val.Value(), val.Type())
	}
	fmt.Println(val.Value())
}
