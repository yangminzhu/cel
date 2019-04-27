package main

import (
	"flag"
	"fmt"
	"github.com/yangminzhu/cel/attributes"
	"github.com/yangminzhu/cel/util"
	"log"

	celgo "github.com/google/cel-go/cel"
)

var (
	printAST            bool
	useSampleAttributes bool
)

func init() {
	flag.BoolVar(&printAST, "ast", false, "also print AST in YAML")
	flag.BoolVar(&useSampleAttributes, "sample", true, "use sample attributes")
}

func main() {
	flag.Parse()

	expression := flag.Arg(0)
	if expression == "" {
		log.Fatalln("CEL expressions must not be empty")
	}

	celenv, err := celgo.NewEnv()
	if useSampleAttributes {
		celenv, err = celgo.NewEnv(attributes.EnvOpt())
	}
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
		log.Printf("using sample attributes: %s\n", attributes.Sample)
		return
	}
	if printAST {
		log.Printf("generating AST for %s\n", expression)
		yml, err := util.ToYAML(ast.Expr())
		if err != nil {
			log.Fatalf("converting AST to yaml error: %v", err)
		} else {
			fmt.Printf("%s\n\n", yml)
		}
	}

	prg, err := celenv.Program(ast)
	if err != nil {
		log.Fatalf("generating program error: %s", err)
	}

	var attr map[string]interface{}
	if useSampleAttributes {
		attr = attributes.Sample
		log.Printf("using sample attributes: %s\n", attributes.Sample)
	}
	log.Printf("evaluated expression: %s", expression)
	val, detail, err := prg.Eval(attr)
	if detail != nil {
		log.Printf("details: %v", detail)
	}
	if err != nil {
		log.Fatalf("evaluation error: %s", err)
	}
	fmt.Println(val.Value())
}
