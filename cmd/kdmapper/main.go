package main

import (
	"fmt"

	"github.com/nicogorr/go-kdmapper/ntdll"
)

func main() {
	modInfo := new(ntdll.SystemModuleInformation)
	if err := ntdll.NtQuerySystemInformation(modInfo); err != nil {
		panic(err)
	}

	fmt.Println("Number of Modules:", modInfo.NumberOfModules)
	fmt.Println(modInfo.Modules)
}
