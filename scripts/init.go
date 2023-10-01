// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package scripts

import "github.com/isaqueveras/powersso/utils"

func Init(logg *utils.Logger) {
	go CreateUserAdmin(logg)
}
