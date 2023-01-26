package main

import (
	"oop_packages/computer"
)

func main() {
	cpu := computer.CPU{}
	cpu.SetArchitecture("64-bit")

	ram := computer.RAM{}
	ram.SetSize(8)

	mboard := computer.Motherboard{}
	mboard.SetCategory("Micro ATX")

	c1 := computer.Computer{}
	c1.SetSpecification(cpu, ram, mboard)
	c1.ShowSpecification()
}