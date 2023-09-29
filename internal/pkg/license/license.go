package license

import (
	"fmt"

	"github.com/hyperboloide/lk"
)

type License interface {
	Encode() (string, error)
}

type defaultLicense struct {
	licenseData *lk.License
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
