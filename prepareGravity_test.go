package main

import(
	"fmt"
"os"
"testing"
)

func TestContentGravity(t *testing.T) {
	fmt.Fprintf(os.Stdout, "Testing contents of Gravity pattern...\n")
	tempo := 60
	ptrn := prepareGravity(tempo)
	expected := "Gravity"
	if ptrn.name != expected {
		fmt.Fprintf(os.Stdout, "Pattern name set incorrectly.\n")
		fmt.Fprintf(os.Stdout, "Expected: %s\nActual: %s\n", expected, ptrn.name)
		t.Error()
	}
	expectInt := 2
	if len(ptrn.timeSig) != expectInt {
		fmt.Fprintf(os.Stdout, "Time signature contains incorrect number fo values.\n")
		fmt.Fprintf(os.Stdout, "Expected: %d\nActual: %d\n", expectInt, len(ptrn.timeSig))
		t.Error()
	}
	expectInt = 6
	if ptrn.timeSig[0] != expectInt {
		fmt.Fprintf(os.Stdout, "Top value of time signature set incorrectly.\n")
		fmt.Fprintf(os.Stdout, "Expected: %d\nActual: %d\n", expectInt, ptrn.timeSig[0])
		t.Error()
	}
	expectInt = 8
	if ptrn.timeSig[1] != expectInt {
		fmt.Fprintf(os.Stdout, "Bottom value of time signature set incorrectly.\n")
		fmt.Fprintf(os.Stdout, "Expected: %d\nActual: %d\n", expectInt, ptrn.timeSig[1])
		t.Error()
	}
	expectInt = 60
	if ptrn.tempo != expectInt {
		fmt.Fprintf(os.Stdout, "Tempo set incorrectly.\n")
		fmt.Fprintf(os.Stdout, "Expected: %d\nActual: %d\n", expectInt, ptrn.tempo)
		t.Error()
	}
	expectInt = 4
	if ptrn.bars != expectInt {
		fmt.Fprintf(os.Stdout, "Number of bars set incorrectly.\n")
		fmt.Fprintf(os.Stdout, "Expected: %d\nActual: %d\n", expectInt, ptrn.bars)
		t.Error()
	}
	expectInt = 12
	if ptrn.divsPerBar != expectInt {
		fmt.Fprintf(os.Stdout, "Number of divisions per bar set incorrectly.\n")
		fmt.Fprintf(os.Stdout, "Expected: %d\nActual: %d\n", expectInt, ptrn.divsPerBar)
		t.Error()
	}

	sounds := [][]string{
		{"Bass Drum"},
		{},
		{},
		{},
		{},
		{"Bass Drum"},
		{"Snare", "HiHat"},
		{},
		{},
		{},
		{},
		{},
	}

	if !compareInstrumentation(ptrn.instPerBeat, sounds) {
		fmt.Fprintf(os.Stdout, "Unexpect instrumentation. ")
		fmt.Fprintf(os.Stdout, "Expected: %v\nActual: %v\n", sounds, ptrn.instPerBeat)
		t.Error()
	}
}