package attributes

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/google/cel-go/checker/decls"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	celgo "github.com/google/cel-go/cel"
)

var (
	spewCfg = spew.NewDefaultConfig()

	stringMapType = &exprpb.Type{
		TypeKind: &exprpb.Type_MapType_{
			MapType: &exprpb.Type_MapType{
				KeyType:   &exprpb.Type{
					TypeKind: &exprpb.Type_Primitive{
						Primitive: exprpb.Type_STRING,
					},
				},
				ValueType: &exprpb.Type{
					TypeKind: &exprpb.Type_Primitive{
						Primitive: exprpb.Type_STRING,
					},
				},
			},
		},
	}
	stringListType = &exprpb.Type{
		TypeKind: &exprpb.Type_ListType_{
			ListType: &exprpb.Type_ListType{
				ElemType: &exprpb.Type{
					TypeKind: &exprpb.Type_Primitive{
						Primitive:  exprpb.Type_STRING,
					},
				},
			},
		},
	}
	intListType = &exprpb.Type{
		TypeKind: &exprpb.Type_ListType_{
			ListType: &exprpb.Type_ListType{
				ElemType: &exprpb.Type{
					TypeKind: &exprpb.Type_Primitive{
						Primitive:  exprpb.Type_INT64,
					},
				},
			},
		},
	}

	Sample = SampleAttributes{
		"a": 10,
		"b": 20,
		"c": true,
		"d": "hello",
		"e": []int{11, 22, 33},
		"f": []string{"f1", "f2"},
		"g": map[string]string{
			"g1": "g11",
			"g2": "g22",
		},
	}
)

func init() {
	spewCfg.SortKeys = true
	spewCfg.DisableCapacities = true
}

type SampleAttributes map[string]interface{}

func (sa SampleAttributes) String() string {
	return spewCfg.Sprintf("%v\n", map[string]interface{}(sa))
}

func EnvOpt() celgo.EnvOption {
	var ret []*exprpb.Decl
	for k, v := range Sample {
		var t *exprpb.Type
		switch v.(type) {
		case string:
			t = decls.String
		case bool:
			t = decls.Bool
		case int, int8, int16, int32, int64:
			t = decls.Int
		case uint, uint8, uint16, uint32, uint64:
			t = decls.Uint
		case float32, float64:
			t = decls.Double
		case map[string]string:
			t = stringMapType
		case []string:
			t = stringListType
		case []int:
			t = intListType
		}
		ret = append(ret, decls.NewIdent(k, t, nil))
	}

	return celgo.Declarations(ret...)
}
