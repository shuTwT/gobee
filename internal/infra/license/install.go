package license

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const licenseFileName = "license.json"

type LicenseInfo struct {
	MachineCode  string `json:"machine_code"`
	LicenseKey   string `json:"license_key"`
	CustomerName string `json:"customer_name"`
	ExpireDate   string `json:"expire_date"`
}

func InstallLicense(machineCode, licenseKey, customerName, expireDate string) error {
	if machineCode == "" || licenseKey == "" {
		return errors.New("机器码和授权密钥不能为空")
	}

	currentMachineCode, err := GenerateMachineCode()
	if err != nil {
		return err
	}

	if machineCode != currentMachineCode {
		return errors.New("机器码不匹配")
	}

	licenseInfo := LicenseInfo{
		MachineCode:  machineCode,
		LicenseKey:   licenseKey,
		CustomerName: customerName,
		ExpireDate:   expireDate,
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	licenseDir := filepath.Join(homeDir, ".gobee")
	if err := os.MkdirAll(licenseDir, 0755); err != nil {
		return err
	}

	licenseFile := filepath.Join(licenseDir, licenseFileName)

	data, err := json.MarshalIndent(licenseInfo, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(licenseFile, data, 0644); err != nil {
		return err
	}

	return nil
}
