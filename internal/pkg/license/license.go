package license

import (
	"fmt"

	swissknife "github.com/Sagleft/swiss-knife"
	"github.com/hyperboloide/lk"
)

type License interface {
	Encode() (string, error)
	Save(filepath string) error
	Load(filepath string) (License, error)
}

type defaultLicense struct {
	licenseData *lk.License
}

func (l *defaultLicense) Save(filepath string) error {
	licData, err := l.Encode()
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	if err := swissknife.SaveStringToFile(filepath, licData); err != nil {
		return fmt.Errorf("save: %w", err)
	}
	return nil
}

func (l *defaultLicense) Load(filepath string) (License, error) {
	//lk.LicenseFromB32String()

	return nil, nil // TODO
}

func (l *defaultLicense) Encode() (string, error) {
	// the b32 representation of our license
	// this is what you give to your customer.
	licenseB32, err := l.licenseData.ToB32String()
	if err != nil {
		return "", fmt.Errorf("encode license file: %w", err)
	}

	return licenseB32, nil
}
