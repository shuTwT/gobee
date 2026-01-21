package license

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

func VerifyLicense() (*LicenseInfo, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	licenseFile := filepath.Join(homeDir, ".gobee", licenseFileName)

	data, err := os.ReadFile(licenseFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("未找到授权文件")
		}
		return nil, err
	}

	var licenseInfo LicenseInfo
	if err := json.Unmarshal(data, &licenseInfo); err != nil {
		return nil, err
	}

	currentMachineCode, err := GenerateMachineCode()
	if err != nil {
		return nil, err
	}

	if licenseInfo.MachineCode != currentMachineCode {
		return nil, errors.New("机器码不匹配")
	}

	expireDate, err := time.Parse(time.RFC3339, licenseInfo.ExpireDate)
	if err != nil {
		return nil, errors.New("授权过期日期格式错误")
	}

	if time.Now().After(expireDate) {
		return nil, errors.New("授权已过期")
	}

	return &licenseInfo, nil
}
