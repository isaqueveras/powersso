// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package otp

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/isaqueveras/powersso/config"
	"github.com/isaqueveras/powersso/pkg/oops"
)

const (
	windowSize = 5
	stepSize   = 30

	QrCodeURL = "https://chart.googleapis.com/chart?chs=200x200&cht=qr&chl=200x200&chld=M|0&cht=qr&chl="
)

// ValidateToken validates if the otp is valid
func ValidateToken(token, otp *string) (err error) {
	if otp == nil {
		return oops.Err(errors.New("the OTP must be sent"))
	}

	for _, value := range []int64{
		(time.Now().Unix() / stepSize),
		(time.Now().Unix() - windowSize) / stepSize,
	} {
		var generated string
		if generated, err = GenerateToken(*token, value); err != nil {
			return oops.Err(err)
		}

		if generated == *otp {
			return nil
		}
	}

	return errors.New("invalid OTP code")
}

// GenerateToken generate an OTP token
func GenerateToken(secret string, ts int64) (otp string, err error) {
	// Converts secret to base32 Encoding. Base32 encoding desires a 32-character
	// subset of the twenty-six letters A–Z and ten digits 0–9
	key, err := base32.StdEncoding.DecodeString(strings.ToUpper(secret))
	if err != nil {
		return otp, err
	}

	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(ts))

	// Signing the value using HMAC-SHA1 Algorithm
	hash := hmac.New(sha1.New, key)
	hash.Write(bs)
	h := hash.Sum(nil)

	// We're going to use a subset of the generated hash.
	// Using the last nibble (half-byte) to choose the index to start from.
	// This number is always appropriate as it's maximum decimal 15, the hash will
	// have the maximum index 19 (20 bytes of SHA1) and we need 4 bytes.
	o := h[19] & 15

	var header uint32
	// Get 32 bit chunk from hash starting at the o
	r := bytes.NewReader(h[o : o+4])
	if err = binary.Read(r, binary.BigEndian, &header); err != nil {
		return otp, err
	}

	// Ignore most significant bits as per RFC 4226.
	// Takes division from one million to generate a remainder less than < 7 digits
	h12 := (int(header) & 0x7fffffff) % 1000000

	// Converts number as a string
	otp = strconv.Itoa(int(h12))

	// Add left pad
	if len(otp) < 6 {
		for i := 0; i < 6-len(otp); i++ {
			otp = "0" + otp
		}
	}

	return
}

// GetUrlQrCode returns the url of qr code to configure the otp
func GetUrlQrCode(otpToken string, userName string) (url string) {
	return QrCodeURL + "otpauth://totp/" + config.Get().Meta.ProjectName + " " + userName + "%3Fsecret%3D" + otpToken
}
