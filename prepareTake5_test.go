package main

import(
	"fmt"
	"os"
	"testing"
)

func TestContentTake5(t *testing.T) {
	fmt.Fprintf(os.Stdout, "Testing contents of Take Five pattern...\n")
	tempo := 100
	ptrn := prepareTake5(tempo)
	expected := "Take Five"
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
	expectInt = 5
	if ptrn.timeSig[0] != expectInt {
		fmt.Fprintf(os.Stdout, "Top value of time signature set incorrectly.\n")
		fmt.Fprintf(os.Stdout, "Expected: %d\nActual: %d\n", expectInt, ptrn.timeSig[0])
		t.Error()
	}
	expectInt = 4
	if ptrn.timeSig[1] != expectInt {
		fmt.Fprintf(os.Stdout, "Bottom value of time signature set incorrectly.\n")
		fmt.Fprintf(os.Stdout, "Expected: %d\nActual: %d\n", expectInt, ptrn.timeSig[1])
		t.Error()
	}
	expectInt = 100
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
	expectInt = 10
	if ptrn.divsPerBar != expectInt {
		fmt.Fprintf(os.Stdout, "Number of divisions per bar set incorrectly.\n")
		fmt.Fprintf(os.Stdout, "Expected: %d\nActual: %d\n", expectInt, ptrn.divsPerBar)
		t.Error()
	}

	sounds := [][]string{
		{"Bass Drum"},
		{"Snare", "Accent"},
		{},
		{"Bass Drum"},
		{"Snare", "Accent"},
		{},
		{"Snare", "Bass Drum", "Cymbal"},
		{},
		{"Snare", "Bass Drum", "Cymbal"},
		{},
	}

	if !compareInstrumentation(ptrn.instPerBeat, sounds) {
		fmt.Fprintf(os.Stdout, "Unexpect instrumentation. ")
		fmt.Fprintf(os.Stdout, "Expected: %v\nActual: %v\n", sounds, ptrn.instPerBeat)
		t.Error()
	}
}