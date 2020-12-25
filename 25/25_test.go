package main

import "testing"

func TestCardTransform(t *testing.T) {
	cardPublic := transform(7, 8)
	if cardPublic != exCardPublic {
		t.Fail()
	}
}

func TestDoorTransform(t *testing.T) {
	doorPublic := transform(7, 11)
	if doorPublic != exDoorPublic {
		t.Fail()
	}
}

func TestFindLoopCard(t *testing.T) {
	cardPublic := exCardPublic
	cardSecret := findLoop(cardPublic)
	if cardSecret != 8 {
		t.Fail()
	}
}

func TestFindLoopDoor(t *testing.T) {
	doorPublic := exDoorPublic
	doorSecret := findLoop(doorPublic)
	if doorSecret != 11 {
		t.Fail()
	}
}

func TestKeysEqual(t *testing.T) {
	exCardLoop := findLoop(exCardPublic)
	exDoorLoop := findLoop(exDoorPublic)

	exKey1 := transform(exCardPublic, exDoorLoop)
	exKey2 := transform(exDoorPublic, exCardLoop)

	if exKey1 != exKey2 {
		t.Logf("%v != %v", exKey1, exKey2)
		t.Fail()
	}
}
