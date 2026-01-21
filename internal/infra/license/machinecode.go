package license

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func GenerateMachineCode() (string, error) {
	var info []string

	hostname, err := os.Hostname()
	if err == nil {
		info = append(info, hostname)
	}

	info = append(info, runtime.GOOS)
	info = append(info, runtime.GOARCH)

	info = append(info, strconv.Itoa(runtime.NumCPU()))

	info = append(info, runtime.Version())

	combined := strings.Join(info, "|")

	hash := md5.Sum([]byte(combined))
	machineCode := hex.EncodeToString(hash[:])

	return machineCode, nil
}
