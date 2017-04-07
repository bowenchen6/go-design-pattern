package visitor

import (
	"fmt"
)

// Visitor visit my computer input device.
type Visitor interface {
	visitMouse(m Mouse)
	visitKeyboard(k Keyboard)
}

// InputDevice represent my computer input device.
type InputDevice struct {
	Mouse
	Keyboard
}

func (input InputDevice) Accept(v Visitor) {
	input.Mouse.accept(v)
	input.Keyboard.accept(v)
}

// Mouse represent my computer mouse.
type Mouse struct {
	name string
}

func (m Mouse) accept(v Visitor) {
	v.visitMouse(m)
}

func (m Mouse) input() string {
	return "my mouse input left, right or middle."
}

func (m Mouse) isLoadDrive() string {
	return "my mouse drive have load."
}

func (m Mouse) price() int {
	return 20
}

// Keyboard represent my computer keyboard.
type Keyboard struct {
	name string
}

func (k Keyboard) accept(v Visitor) {
	v.visitKeyboard(k)
}

func (k Keyboard) input() string {
	return "my Keyboard input a-z or 0-9 or ... ."
}

func (k Keyboard) isLoadDrive() string {
	return "my keyboard drive have not load."
}

func (k Keyboard) price() int {
	return 200
}

// IsLoadDriveVisitor used to check my computer input device is or not load device.
type IsLoadDriveVisitor string

func (vd IsLoadDriveVisitor) visitMouse(m Mouse) {
	fmt.Println("IsLoadDriveVisitor visitor my computer mouse")
	fmt.Println(m.isLoadDrive())
}

func (vd IsLoadDriveVisitor) visitKeyboard(k Keyboard) {
	fmt.Println("IsLoadDriveVisitor visitor my computer keyboard.")
	fmt.Println(k.isLoadDrive())
}

// InputVisitor used to make my computer input device input data.
type InputVisitor string

func (vi InputVisitor) visitMouse(m Mouse) {
	fmt.Println("InputVisitor visitor my computer mouse.")
	fmt.Println(m.input())
}

func (vi InputVisitor) visitKeyboard(k Keyboard) {
	fmt.Println("InputVisitor visitor my computer keyboard.")
	fmt.Println(k.input())
}

// PriceVisitor used to calculation my computer input device total price.
type PriceVisitor int

func (vp *PriceVisitor) visitMouse(m Mouse) {
	myMousePrice := m.price()
	*vp = *vp + (PriceVisitor)(myMousePrice)
}

func (vp *PriceVisitor) visitKeyboard(k Keyboard) {
	myKeyboardPrice := k.price()
	*vp = *vp + (PriceVisitor)(myKeyboardPrice)
}
