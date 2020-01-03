package main

func compareInstrumentation(exp [][]string, act [][]string) bool {
	for i := range act {
		for j := range act[i]{
			if act[i][j] != exp[i][j] {
				return false
			}
		}
	}
	return true
}
