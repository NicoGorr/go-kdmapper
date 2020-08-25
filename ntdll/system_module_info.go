package ntdll

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

type SystemModuleInformation struct {
	NumberOfModules uint32
	Modules         []SystemModule
}

func (SystemModuleInformation) ID() uint32 {
	return 11
}

func (m *SystemModuleInformation) UnmarshalNT(b []byte) error {
	r := bytes.NewReader(b)

	if err := binary.Read(r, binary.LittleEndian, &m.NumberOfModules); err != nil {
		return err
	}

	// Num of Mods Padding
	if _, err := r.Seek(4, io.SeekCurrent); err != nil {
		return err
	}

	m.Modules = make([]SystemModule, m.NumberOfModules)
	return binary.Read(r, binary.LittleEndian, &m.Modules)
}

func (m SystemModuleInformation) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("Number of Modules: %d", m.NumberOfModules))
	b.WriteString("\nModules:")
	for _, mod := range m.Modules {
		b.WriteString(fmt.Sprintf("\n%v", mod))
	}

	return b.String()
}

type SystemModule struct {
	Section        uint64
	MappedBase     uint64
	ImageBase      uint64
	ImageSize      uint32
	Flags          uint32
	LoadOrderIndex uint16
	InitOrderIndex uint16
	LoadCount      uint16
	NameOffset     uint16
	Path           [256]byte
}

func (m SystemModule) String() string {
	var b strings.Builder

	b.WriteString("System Module:")
	b.WriteString(fmt.Sprintf("\nSection: 0x%08X", m.Section))
	b.WriteString(fmt.Sprintf("\nMappedBase: 0x%08X", m.MappedBase))
	b.WriteString(fmt.Sprintf("\nImageBase: 0x%08X", m.ImageBase))
	b.WriteString(fmt.Sprintf("\nImageSize: %d", m.ImageSize))
	b.WriteString(fmt.Sprintf("\nFlags: %d", m.Flags))
	b.WriteString(fmt.Sprintf("\nLoadOrderIndex: %d", m.LoadOrderIndex))
	b.WriteString(fmt.Sprintf("\nInitOrderIndex: %d", m.InitOrderIndex))
	b.WriteString(fmt.Sprintf("\nLoadCount: %d", m.LoadCount))
	b.WriteString(fmt.Sprintf("\nNameOffset: %d", m.NameOffset))
	b.WriteString(fmt.Sprintf("\nPath: %s", string(m.Path[:])))

	return b.String()
}
