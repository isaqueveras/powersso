package types_test

import (
	"testing"

	"github.com/isaqueveras/power-sso/internal/types"
)

func TestTypeConverter(t *testing.T) {
	type app struct {
		Name string `json:"name"`
	}

	type domain struct {
		FullName string `json:"name"`
	}

	scenarios := []struct {
		in  app
		out domain
	}{
		{in: app{Name: "Isaque"}, out: domain{FullName: ""}},
		{in: app{Name: ""}, out: domain{FullName: "Isaque"}},
		{in: app{}, out: domain{}},
	}

	for i, scenario := range scenarios {
		data, err := types.TypeConverter[domain](&scenario.in)
		if err != nil {
			t.Errorf("(%d) Expected nil, got %v", i, err)
			continue
		}

		if data.FullName != scenario.in.Name {
			t.Error("Expected nil, got", err)
		}
	}
}
