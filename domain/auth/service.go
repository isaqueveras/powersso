package auth

import (
	"encoding/base32"

	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/utils"
)

// Service structure with repositories
type Service struct {
	repoFlag IFlag
	repoOTP  IOTP
}

// NewAuthService init new service
func NewAuthService(repoFlag IFlag, repoOTP IOTP) IAuthService {
	return &Service{repoFlag: repoFlag, repoOTP: repoOTP}
}

// Configure2FA add the flags to the configured 2fa user and generates the 2fa token
func (s *Service) Configure2FA(userID *uuid.UUID) (err error) {
	if err = s.repoFlag.Set(userID, FlagOTPSetup); err != nil {
		return err
	}

	if err = s.repoFlag.Set(userID, FlagOTPEnable); err != nil {
		return err
	}

	data := []byte(utils.RandomString(26))
	dst := make([]byte, base32.StdEncoding.EncodedLen(len(data)))
	base32.StdEncoding.Encode(dst, data)
	return s.repoOTP.SetToken(utils.Pointer(string(dst)))
}

func (*Service) GenerateQrCode2FA() {}
