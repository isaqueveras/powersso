// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package otp

// QRCodeResponse wraps the data to return the qr code url
type QRCodeResponse struct {
	Url *string `json:"url,omitempty"`
}
