package ntdll

import "golang.org/x/sys/windows"

var dll = windows.NewLazyDLL("ntdll.dll")

type NtStatus uint32

const (
	StatusInfoLengthMismatch NtStatus = 0xC0000004
)

func (s NtStatus) IsSuccess() bool {
	return s >= 0
}
