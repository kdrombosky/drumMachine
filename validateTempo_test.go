package main

import(
	"fmt"
	"os"
	"testing"
)

var errorResp = -1


func TestValidateTempoNotInt(t *testing.T) {
	fmt.Fprintf(os.Stdout, "Testing incorrect tempo input type...\n")
	tempostr := "Not an int"
	tempo := validateTempo(tempostr)
	if tempo != -1 {
		fmt.Fprintf(os.Stdout, "Validate Tempo did not catch incorrect input.\n")
		fmt.Fprintf(os.Stdout, "Expected Response: %d\nActual Response: %d\n", errorResp, tempo)
		t.Error()
	}

	tempoFloat := "60.00357"
	tempo = validateTempo(tempoFloat)
	if tempo != -1 {
		fmt.Fprintf(os.Stdout, "Validate Tempo did not catch incorrect input.\n")
		fmt.Fprintf(os.Stdout, "Expected Response: %d\nActual Response: %d\n", errorResp, tempo)
		t.Error()
	}

	tempoBool := "false"
	tempo = validateTempo(tempoBool)
	if tempo != -1 {
		fmt.Fprintf(os.Stdout, "Validate Tempo did not catch incorrect input.\n")
		fmt.Fprintf(os.Stdout, "Expected Response: %d\nActual Response: %d\n", errorResp, tempo)
		t.Error()
	}
}

func TestValidateTempoNotInRange(t *testing.T) {
	fmt.Fprintf(os.Stdout, "Testing incorrect tempo input range...\n")
	tempostr := "12"
	tempo := validateTempo(tempostr)
	if tempo != -1 {
		fmt.Fprintf(os.Stdout, "Validate Tempo did not catch out of range input.\n")
		fmt.Fprintf(os.Stdout, "Expected Response: %d\nActual Response: %d\n", errorResp, tempo)
		t.Error()
	}

	tempostr = "201"
	tempo = validateTempo(tempostr)
	if tempo != -1 {
		fmt.Fprintf(os.Stdout, "Validate Tempo did not catch out of range input.\n")
		fmt.Fprintf(os.Stdout, "Expected Response: %d\nActual Response: %d\n", errorResp, tempo)
		t.Error()
	}

	tempostr = "40"
	tempo = validateTempo(tempostr)
	if tempo == -1 {
		fmt.Fprintf(os.Stdout, "Validate Tempo did not allow in range input.\n")
		fmt.Fprintf(os.Stdout, "Expected Response: %d\nActual Response: %d\n", errorResp, tempo)
		t.Error()
	}

	tempostr = "200"
	tempo = validateTempo(tempostr)
	if tempo == -1 {
		fmt.Fprintf(os.Stdout, "Validate Tempo did not allow in range input.\n")
		fmt.Fprintf(os.Stdout, "Expected Response: %d\nActual Response: %d\n", errorResp, tempo)
		t.Error()
	}
}

func TestValidateTempo(t *testing.T) {
	fmt.Fprintf(os.Stdout, "Testing correct tempo input type and range...\n")
	tempostr := "68"
	tempo := validateTempo(tempostr)
	if tempo != 68 {
		fmt.Fprintf(os.Stdout, "Validate Tempo did not output correct tempo integer.\n")
		fmt.Fprintf(os.Stdout, "Expected Response: %d\nActual Response: %d\n", errorResp, tempo)
		t.Error()
	}
	tempostr = "40"
	tempo = validateTempo(tempostr)
	if tempo != 40 {
		fmt.Fprintf(os.Stdout, "Validate Tempo did not allow in range input.\n")
		fmt.Fprintf(os.Stdout, "Expected Response: %d\nActual Response: %d\n", errorResp, tempo)
		t.Error()
	}
	tempostr = "200"
	tempo = validateTempo(tempostr)
	if tempo != 200 {
		fmt.Fprintf(os.Stdout, "Validate Tempo did not allow in range input.\n")
		fmt.Fprintf(os.Stdout, "Expected Response: %d\nActual Response: %d\n", errorResp, tempo)
		t.Error()
	}
}