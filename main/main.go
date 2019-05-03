package main

import (
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
	"github.com/yangminzhu/cel/attributes"
	"github.com/yangminzhu/cel/util"
	"log"

	celgo "github.com/google/cel-go/cel"
)

var (
	printAST  bool
	printText bool
	spewCfg                 = spew.NewDefaultConfig()
)

func init() {
	flag.BoolVar(&printAST, "ast", false, "print parsed AST in YAML, skip evaluation")
	flag.BoolVar(&printText, "text", false, "print parsed AST in JSON, skip evaluation")
	spewCfg.SortKeys = true
	spewCfg.DisableCapacities = true
}

func getProgram(expression string) celgo.Program {
	celenv, err := celgo.NewEnv(attributes.EnvOpt())
	if err != nil {
		log.Fatalf("cel environment creation error: %s\n", err)
	}

	a, iss := celenv.Parse(expression)
	if iss != nil {
		log.Printf("AST parse issues:\n%s\n", iss)
	}
	ast, iss := celenv.Check(a)
	if iss != nil {
		log.Printf("AST check issues:\n%s\n", iss)
	}
	if ast == nil {
		return nil
	}
	if printAST {
		var out string
		var err error
		if printText {
			out = proto.MarshalTextString(ast.Expr())
		} else {
			out, err = util.ToYAML(ast.Expr())
		}
		if err != nil {
			log.Fatalf("converting AST to yaml error: %v", err)
		} else {
			fmt.Printf("%s\n\n", out)
		}
		return nil
	}

	prg, err := celenv.Program(ast, celgo.EvalOptions(celgo.OptFoldConstants))
	if err != nil {
		log.Fatalf("generating program error: %s", err)
	}
	return prg
}

func main() {
	flag.Parse()

	expression := flag.Arg(0)
	if expression == "" {
		log.Fatalln("CEL expressions must not be empty")
	}

	yml, err := util.ToYAML(&attributes.Example)
	if err != nil {
		log.Printf("converting Example attributes to yaml error: %v", err)
	}
	log.Printf("using attributes:\n%s\n", yml)

	prg := getProgram(expression)
	if prg == nil {
		return
	}
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
