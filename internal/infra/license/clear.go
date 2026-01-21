package license

import (
	"os"
	"path/filepath"
)

func ClearLicense() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	licenseFile := filepath.Join(homeDir, ".gobee", licenseFileName)

	if err := os.Remove(licenseFile); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return nil
}
