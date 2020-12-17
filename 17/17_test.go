package main

import "testing"

func TestExample1(t *testing.T) {
	b := parse(example)

	for i := 0; i < 6; i++ {
		b = b.iterate()
	}

	count := b.count()
	if count != 112 {
		t.Fatalf("P1 Bad %v != 112", count)
	}
}

func TestExample2(t *testing.T) {
	b4 := parse4(example)

	for i := 0; i < 6; i++ {
		b4 = b4.iterate()
	}

	count := b4.count()
	if count != 848 {
		t.Fatalf("P2 Bad %v != 848", count)
	}
}
