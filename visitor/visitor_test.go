package visitor

import (
	"fmt"
	"testing"
)

func TestIsLoadDriveVisitor(t *testing.T) {
	input := InputDevice{}
	var v IsLoadDriveVisitor
	input.Accept(v)
}

func TestInputVisitor(t *testing.T) {
	input := InputDevice{}
	var v InputVisitor
	input.Accept(v)
}

func TestPriceVisitor(t *testing.T) {
	input := InputDevice{}
	var v PriceVisitor
	input.Accept(&v)
	fmt.Println(v)
}
