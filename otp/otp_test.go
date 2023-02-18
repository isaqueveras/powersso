// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package otp_test

import (
	"testing"

	"github.com/isaqueveras/power-sso/otp"
)

func TestGenerateToken(t *testing.T) {
	type input struct {
		secret string
		ts     int64
	}

	scenarios := []input{
		{secret: "J5WGCTLVNZSG6II=", ts: 41929800},
		{secret: "JEQGW3TPO4QHS33VEB3W65LMMQQGIZLDN5SGKIDUNBUXGIDDN5SGKLBAPFXXKIDDOVZGS33VOMQQ====", ts: 41929415},
		{secret: "IRXSA4LVMUQHM33DYOVCAZ3PON2GCPZB", ts: 41929799},
	}

	for i := range scenarios {
		code, err := otp.GenerateToken(scenarios[i].secret, scenarios[i].ts)
		if err != nil {
			t.Errorf("%v", err)
			continue
		}

		if code == "" {
			t.Error("Empty otp code")
		}
	}
}
