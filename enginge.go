package freemiumengine

import (
	"time"

	"github.com/Sagleft/freemium-engine/internal/pkg/license"
)

const licenseFilePath = "license.dat"

func Create(secret string) error {
	until := time.Now().Add(time.Hour * time.Duration(24))

	data := license.UserData{
		Email:           "tester@mfc.click",
		RegisteredUntil: until,
	}

	return license.GrantUserAccessAndSave(data, secret, licenseFilePath)
}

func Validate(pubkey string) error {
	return license.ValidateUserLicenseFile(licenseFilePath, pubkey)
}

func CreateSecret() (string, error) {
	return license.EncodeKeyToBase32("test")
}
