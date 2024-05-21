package systeminfo

import (
	"strings"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func ComputerName() string {
	name, err := windows.ComputerName()
	if err != nil {
		return ""
	}
	return name
}
func CPU() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `HARDWARE\DESCRIPTION\System\CentralProcessor\0`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	name, _, err := k.GetStringValue("ProcessorNameString")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(name), nil
}
