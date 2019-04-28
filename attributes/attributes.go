package attributes

import (
	celgo "github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/interpreter"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"log"
)

var (
	stringMapType = decls.NewMapType(decls.String, decls.String)
	stringListType = decls.NewListType(decls.String)
	intListType = decls.NewListType(decls.Int)

	Example = AttributesContext{
		Headers: map[string]string{
			":path": "/info/v1",
			":host": "httpbin",
			"x-id": "123456",
			":method": "GET",
			"authorization": "bearer 123456",
		},
		Tls: true,
		Sni: "www.httpbin.com",
		Port: 8080,
		Clusters: []string{"cluster-1", "cluster-2", "cluster-3"},
		Weighet: []int32{10, 20, 30, 40},
	}
)

func MyActivation() interpreter.Activation {
	// Workaround for https://github.com/google/cel-go/issues/208
	h :=  map[string]interface{}{}
	for k, v := range Example.GetHeaders() {
		h[k] = v
	}
	// TODO: Use reflection.
	m := map[string]interface{}{
		"headers":  h,
		"tls":      Example.GetTls(),
		"port":     Example.GetPort(),
		"clusters": Example.GetClusters(),
		"weight":   Example.GetWeighet(),
	}
	a, err := interpreter.NewActivation(m)
	if err != nil {
		log.Fatalf("Activation creation error: %v", err)
	}
	return a
}

func EnvOpt() celgo.EnvOption {
	var ret []*exprpb.Decl
	// TODO: Use reflection.
	ret = append(ret, decls.NewIdent("headers", stringMapType, nil))
	ret = append(ret, decls.NewIdent("tls", decls.Bool, nil))
	ret = append(ret, decls.NewIdent("sni", decls.String, nil))
	ret = append(ret, decls.NewIdent("port", decls.Uint, nil))
	ret = append(ret, decls.NewIdent("clusters", stringListType, nil))
	ret = append(ret, decls.NewIdent("weight", intListType, nil))

	//for k, v := range Example {
	//	var t *exprpb.Type
	//	switch v.(type) {
	//	case string:
	//		t = decls.String
	//	case bool:
	//		t = decls.Bool
	//	case int, int8, int16, int32, int64:
	//		t = decls.Int
	//	case uint, uint8, uint16, uint32, uint64:
	//		t = decls.Uint
	//	case float32, float64:
	//		t = decls.Double
	//	case map[string]string:
	//		t = stringMapType
	//	case []string:
	//		t = stringListType
	//	case []int:
	//		t = intListType
	//	}
	//	ret = append(ret, decls.NewIdent(k, t, nil))
	//}

	return celgo.Declarations(ret...)
}
