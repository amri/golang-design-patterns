package main

import "fmt"

type ILaptop interface {
	setMemory(amount int)
	getMemory() int
	getBrand() string
}

type AbstractLaptopFactory interface {
	createLaptop() ILaptop
}

type laptop struct {
	brand  string
	memory int
}

type AppleLaptop struct {
	laptop
}

func (al *AppleLaptop) setMemory(amount int) {
	al.memory = amount
}

func (al *AppleLaptop) getMemory() int {
	return al.memory
}

func (al *AppleLaptop) getBrand() string {
	return al.brand
}

type AppleLaptopFactory struct {
}

func (a *AppleLaptopFactory) createLaptop() ILaptop {
	return &AppleLaptop{
		laptop{
			brand:  "Apple",
			memory: 128,
		},
	}
}

type WindowsLaptopFactory struct {
}

func (w *WindowsLaptopFactory) createLaptop() ILaptop {
	return &WindowsLaptop{
		laptop{
			brand:  "Windows",
			memory: 256,
		},
	}
}

type WindowsLaptop struct {
	laptop
}

func (al *WindowsLaptop) setMemory(amount int) {
	al.memory = amount
}

func (al *WindowsLaptop) getMemory() int {
	return al.memory
}

func (al *WindowsLaptop) getBrand() string {
	return al.brand
}

func getLaptopFactory(brand string) AbstractLaptopFactory {
	if brand == "apple" {
		return &AppleLaptopFactory{}
	}

	if brand == "windows" {
		return &WindowsLaptopFactory{}
	}
	return nil
}

func main() {
	appleFactory := getLaptopFactory("apple")
	windowsFactory := getLaptopFactory("windows")

	appleLaptop1 := appleFactory.createLaptop()
	appleLaptop2 := appleFactory.createLaptop()
	appleLaptop2.setMemory(200)
	windowsLaptop := windowsFactory.createLaptop()

	fmt.Println(appleLaptop1.getBrand(), appleLaptop1.getMemory(), appleLaptop2.getMemory(), windowsLaptop.getBrand(), windowsLaptop.getMemory())
}
