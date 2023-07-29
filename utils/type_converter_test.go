// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package utils_test

import (
	"testing"

	"github.com/isaqueveras/powersso/utils"
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
		data, err := utils.TypeConverter[domain](&scenario.in)
		if err != nil {
			t.Errorf("(%d) Expected nil, got %v", i, err)
			continue
		}

		if data.FullName != scenario.in.Name {
			t.Error("Expected nil, got", err)
		}
	}
}
