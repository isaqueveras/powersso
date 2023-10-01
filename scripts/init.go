package scripts

import "github.com/isaqueveras/powersso/utils"

func Init(logg *utils.Logger) {
	go CreateUserAdmin(logg)
}
