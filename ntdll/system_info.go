package ntdll

import (
	"fmt"
	"unsafe"
)

var procNtQuerySystemInformation = dll.NewProc("NtQuerySystemInformation")

type SystemInformationClass interface {
	ID() uint32
	UnmarshalNT(b []byte) error
}

func NtQuerySystemInformation(class SystemInformationClass) error {
	len := uint32(0)
	if status := sysNtQuerySystemInformation(class.ID(), nil, 0, &len); status != StatusInfoLengthMismatch {
		return fmt.Errorf("ntdll: NtQuerySystemInformation with status %d", status)
	}

	b := make([]byte, len)
	if status := sysNtQuerySystemInformation(class.ID(), &b[0], len, nil); !status.IsSuccess() {
		return fmt.Errorf("ntdll: NtQuerySystemInformation with status %d", status)
	}

	return class.UnmarshalNT(b)
}

func sysNtQuerySystemInformation(sysInfoClass uint32, sysInfo *byte, sysInfoLen uint32, retLen *uint32) NtStatus {
	r, _, _ := procNtQuerySystemInformation.Call(
		uintptr(sysInfoClass),
		uintptr(unsafe.Pointer(sysInfo)),
		uintptr(sysInfoLen),
		uintptr(unsafe.Pointer(retLen)),
	)

	return NtStatus(r)
}
