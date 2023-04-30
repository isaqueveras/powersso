// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package otp_test

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/otp"
)

var tokens = []string{
	"J5WGCTLVNZSG6II=",
	"JEQGW3TPO4QHS33VEB3W65LMMQQGIZLDN5SGKIDUNBUXGIDDN5SGKLBAPFXXKIDDOVZGS33VOMQQ====",
	"IRXSA4LVMUQHM33DYOVCAZ3PON2GCPZB",
}

func TestOTP(t *testing.T) {
	t.Run("GenerateAndValidateToken", func(t *testing.T) {
		for i := range tokens {
			code, err := otp.GenerateToken(tokens[i], time.Now().Unix()/30)
			if err != nil {
				t.Error(err)
				continue
			}

			if code == "" {
				t.Error("Empty otp code")
			}

			if err := otp.ValidateToken(&tokens[i], &code); err != nil {
				t.Error(err)
				continue
			}
		}
	})

	t.Run("GetURLQRCode", func(t *testing.T) {
		config.LoadConfig()

		for i := range tokens {
			userUUID := uuid.New()
			url := otp.GetUrlQrCode(tokens[i], userUUID.String())
			urlCorrect := otp.QrCodeURL + "otpauth://totp/" + config.Get().Meta.ProjectName + " " + userUUID.String() + "%3Fsecret%3D" + tokens[i]

			if urlCorrect != url {
				t.Error("url not equal")
			}
		}
	})
}

func BenchmarkOTP(b *testing.B) {
	b.Run("GenerateAndValidateToken", func(b *testing.B) {
		for i := range tokens {
			code, err := otp.GenerateToken(tokens[i], time.Now().Unix()/30)
			if err != nil {
				b.Error(err)
				continue
			}

			if err := otp.ValidateToken(&tokens[i], &code); err != nil {
				b.Error(err)
				continue
			}
		}
	})

	b.Run("GetURLQRCode", func(b *testing.B) {
		config.LoadConfig()

		for i := range tokens {
			userUUID := uuid.New()
			url := otp.GetUrlQrCode(tokens[i], userUUID.String())
			urlCorrect := otp.QrCodeURL + "otpauth://totp/" + config.Get().Meta.ProjectName + " " + userUUID.String() + "%3Fsecret%3D" + tokens[i]

			if urlCorrect != url {
				b.Error("url not equal")
			}
		}
	})
}
