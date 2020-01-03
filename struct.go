package main

/*
 * name: A string representing the name of the pattern.
 * timeSig: An array representation of the time signature with the index 0 being the top number and 1 being the bottom
 * tempo: An in representation of the tempo (to be set by the user). Note: Tempo input is validated in main.go
 * bars: An integer representation of the number of bars the pattern should play
 * divsPerBar: An integer representation of the number of divisions per bar (i.e., If dividing 4/4 into eighth notes, 8)
 * instPerBeat: A slice of slices of strings to represent which instruments are play on each division.
 */
type Pattern struct {
	name        string // String name of the patter (like "Four on the Floor")
	timeSig     [2]int
	tempo       int
	bars        int
	divsPerBar  int        // Number of subdivision to be counted in the bar
	instPerBeat [][]string // Which instruments are playing on the beat. Can be none.
}

func prepare4OTF(tempo int) Pattern {
	var ptrn Pattern

	// Construct pattern struct
	ptrn.name = "Four on the Floor"
	ptrn.timeSig[0] = 4
	ptrn.timeSig[1] = 4
	ptrn.tempo = tempo
	ptrn.bars = 4
	ptrn.divsPerBar = 8
	sounds := make([][]string, ptrn.divsPerBar)

	// add the sounds
	sounds[0] = append(sounds[0], "Bass Drum")
	sounds[2] = append(sounds[2], "HiHat")
	sounds[4] = append(sounds[4], "Bass Drum")
	sounds[4] = append(sounds[4], "Snare")
	sounds[6] = append(sounds[6], "HiHat")
	ptrn.instPerBeat = sounds
	return ptrn
}

func prepareGravity(tempo int) Pattern {
	var ptrn Pattern
	// Construct pattern struct
	ptrn.name = "Gravity"
	ptrn.timeSig[0] = 6
	ptrn.timeSig[1] = 8
	ptrn.tempo = tempo
	ptrn.bars = 4
	ptrn.divsPerBar = 12
	sounds := make([][]string, ptrn.divsPerBar)

	// add the sounds
	sounds[0] = append(sounds[0], "Bass Drum")
	sounds[5] = append(sounds[5], "Bass Drum")
	sounds[6] = append(sounds[6], "Snare")
	sounds[6] = append(sounds[6], "HiHat")
	ptrn.instPerBeat = sounds
	return ptrn
}

func prepareTake5(tempo int) Pattern {
	var ptrn Pattern
	// Construct pattern struct
	ptrn.name = "Take Five"
	ptrn.timeSig[0] = 5
	ptrn.timeSig[1] = 4
	ptrn.tempo = tempo
	ptrn.bars = 4
	ptrn.divsPerBar = 10
	sounds := make([][]string, ptrn.divsPerBar)

	// add the sounds
	sounds[0] = append(sounds[0], "Bass Drum")
	sounds[1] = append(sounds[1], "Snare")
	sounds[1] = append(sounds[1], "Accent")
	sounds[3] = append(sounds[3], "Bass Drum")
	sounds[4] = append(sounds[4], "Snare")
	sounds[4] = append(sounds[4], "Accent")
	sounds[6] = append(sounds[6], "Snare")
	sounds[6] = append(sounds[6], "Bass Drum")
	sounds[6] = append(sounds[6], "Cymbal")
	sounds[8] = append(sounds[8], "Snare")
	sounds[8] = append(sounds[8], "Bass Drum")
	sounds[8] = append(sounds[8], "Cymbal")
	ptrn.instPerBeat = sounds
	return ptrn
}

func prepareCowbell(tempo int) Pattern {
	var ptrn Pattern

	// Construct pattern struct
	ptrn.name = "We Need More Cowbell"
	ptrn.timeSig[0] = 4
	ptrn.timeSig[1] = 4
	ptrn.tempo = tempo
	ptrn.bars = 4
	ptrn.divsPerBar = 8
	sounds := make([][]string, ptrn.divsPerBar)

	// add the sounds
	sounds[0] = append(sounds[0], "Bass Drum")
	sounds[0] = append(sounds[0], "Cowbell")
	sounds[2] = append(sounds[2], "HiHat")
	sounds[2] = append(sounds[2], "Cowbell")
	sounds[4] = append(sounds[4], "Bass Drum")
	sounds[4] = append(sounds[4], "Snare")
	sounds[4] = append(sounds[4], "Cowbell")
	sounds[6] = append(sounds[6], "HiHat")
	sounds[6] = append(sounds[6], "Cowbell")
	ptrn.instPerBeat = sounds
	return ptrn
}
