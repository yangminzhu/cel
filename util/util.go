package util

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

func ToYAML(msg proto.Message) (string, error) {
	if msg == nil {
		return "", fmt.Errorf("nil message")
	}

	m := jsonpb.Marshaler{}
	js, err := m.MarshalToString(msg)
	if err != nil {
		return "", err
	}

	ret, err := yaml.JSONToYAML([]byte(js))
	return string(ret), err
}
