package authentication

import (
	"encoding/base32"

	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/config"
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
	return s.repoOTP.SetToken(userID, utils.Pointer(string(dst)))
}

// GenerateQrCode2FA return the formatted url to configure 2-factor authentication
func (s *Service) GenerateQrCode2FA(userID *uuid.UUID) (url *string, err error) {
	userName, token, err := s.repoOTP.GetToken(userID)
	if err != nil {
		return nil, err
	}

	if config.Get().Server.IsModeDevelopment() {
		*userName += " [DEV]"
	}

	projectName := utils.Pointer(config.Get().ProjectName)
	link := utils.GetUrlQrCode(projectName, token, userName)

	return utils.Pointer(link), nil
}
