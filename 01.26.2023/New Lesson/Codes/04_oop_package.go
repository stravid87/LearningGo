//This code is not working for 04_oop.go because you have to create your own package with this code.

package computer

import "fmt"

type CPU struct {
	architecture string
}

func (cpu *CPU) SetArchitecture(arch string) {
	cpu.architecture = arch
}

func (cpu *CPU) GetArchitecture() string {
	return cpu.architecture
}

type RAM struct {
	size int
}

func (ram *RAM) SetSize(size int) {
	ram.size = size
}

func (ram *RAM) GetSize() int {
	return ram.size
}

type Motherboard struct {
	category string
}

func (m *Motherboard) SetCategory(cat string) {
	m.category = cat
}

func (m *Motherboard) GetCategory() string {
	return m.category
}

type Computer struct {
	cpu    CPU
	ram    RAM
	mboard Motherboard
}

func (c *Computer) SetSpecification(cpu CPU, ram RAM, mboard Motherboard) {
	c.cpu.SetArchitecture(cpu.GetArchitecture())
	c.ram.SetSize(ram.GetSize())
	c.mboard.SetCategory(mboard.GetCategory())
}

func (c *Computer) ShowSpecification() {
	fmt.Println("CPU: ", c.cpu.GetArchitecture(), ", RAM: ", c.ram.GetSize(), "GB, Motherboard: ", c.mboard.GetCategory())
}