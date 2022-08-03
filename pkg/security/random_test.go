// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package security_test

import (
	"regexp"
	"testing"

	"github.com/isaqueveras/power-sso/pkg/security"
)

func TestRandomString(t *testing.T) {
	generated := []string{}
	reg := regexp.MustCompile(`[a-zA-Z0-9]+`)

	for i := 0; i < 30; i++ {
		length := 5 + i
		result := security.RandomString(length)

		if len(result) != length {
			t.Errorf("(%d) Expected the length of the string to be %d, got %d", i, length, len(result))
		}

		if match := reg.MatchString(result); !match {
			t.Errorf("(%d) The generated strings should have only [a-zA-Z0-9]+ characters, got %q", i, result)
		}

		for _, str := range generated {
			if str == result {
				t.Errorf("(%d) Repeating random string - found %q in \n%v", i, result, generated)
			}
		}

		generated = append(generated, result)
	}
}