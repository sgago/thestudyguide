package vendingmachine

import (
	"fmt"
	"testing"
)

type state string

const (
	Idle     state = "idle"
	Select   state = "select"
	Dispense state = "dispense"
	restart  state = "restart"

	Start = Idle
)

type VendingMachine struct {
	current state
}

func NewVendingMachine() *VendingMachine {
	vm := &VendingMachine{
		current: Start,
	}

	vm.Input("restart")

	return vm
}

func (vm *VendingMachine) Input(input string) {
	switch vm.current {
	case Idle:
		vm.handleIdleInput(input)
	case Select:
		vm.handleSelectInput(input)
	case Dispense:
		vm.handleDispense(input)
	}
}

func (vm *VendingMachine) handleIdleInput(input string) {
	if input == "restart" {
		fmt.Println("\nWelcome!")
		vm.current = Start
		return
	}

	if input == "select" {
		fmt.Println("Selected...")
		vm.current = Select
		return
	}
}

func (vm *VendingMachine) handleSelectInput(input string) {
	if input == "confirm" {
		vm.current = Dispense
		vm.Input("dispense")
		return
	}

	if input == "cancel" {
		fmt.Println("Canceling...")
		vm.current = Start
		return
	}
}

func (vm *VendingMachine) handleDispense(_ string) {
	fmt.Println("Dispensing... Thank you!")
	vm.handleIdleInput("restart")
}

func Test_Start(t *testing.T) {
	NewVendingMachine()
}

func TestHappyPath(t *testing.T) {
	vm := NewVendingMachine()

	vm.Input("select")
	vm.Input("confirm")
}
