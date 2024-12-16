package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(solve(os.Args[1], os.Args[2:]))
}

func solve(a string, wrds []string) string {
	quizMatrix := matrix(a)
	for i, row := range quizMatrix {
		for x, col := range row {
			for wx, w := range wrds {
				if col == "2" {
					if across(quizMatrix, i, x, w) {
						quizMatrix = placeWord(quizMatrix, w, "hrz", i, x)
						wrds = trimBox(wrds, w)
					} else {
						wx++
					}
					if down(quizMatrix, i, x, w) {
						quizMatrix = placeWord(quizMatrix, w, "vtc", i, x)
						wrds = trimBox(wrds, w)
					} else {
						wx++
					}
				} else if col == "1" {
					if across(quizMatrix, i, x, w) {
						quizMatrix = placeWord(quizMatrix, w, "hrz", i, x)
						wrds = trimBox(wrds, w)
					} else if down(quizMatrix, i, x, w) {
						quizMatrix = placeWord(quizMatrix, w, "vtc", i, x)
						wrds = trimBox(wrds, w)
					}
				} else if col == "." || col == "0" {
					continue
				}
			}
		}
	}
	return output(quizMatrix)
}

func output(a [][]string) string {
	var output string
	for i, row := range a {
		for _, col := range row {
			output += col
		}
		if i != len(a)-1 {
			output += "\n"
		}
	}
	return output
}

func down(a [][]string, i1, x int, wrd string) bool {
	n := len(wrd)
	for i := 0; i < n; i++ {
		if !checkDown(a, i1, (x + i)) {
			return false
		}
	}
	return true
}

func across(a [][]string, i1, x int, wrd string) bool {
	n := len(wrd)
	for i := 0; i < n; i++ {
		if !checkAcross(a, (i + i1), x) {
			return false
		}
	}
	return true
}

func trimBox(a []string, word string) []string {
	var final []string
	for _, w := range a {
		if w != word {
			final = append(final, w)
		}
	}
	return final
}

func placeWord(a [][]string, word string, option string, i, x int) [][]string {
	if option == "hz" {
		for i1, row := range a {
			for x1 := range row {
				if x1 == x {
					if i1 == i {
						for _, c := range word {
							a[x1][i1] = string(c)
							i1++
						}
					}
					break
				}
			}
		}
	} else if option == "vtc" {
		for i1, row := range a {
			for x1 := range row {
				if x1 == x {
					if i1 == i {
						for _, c := range word {
							a[x1][i1] = string(c)
							x1++
						}
						break
					}
				}
			}
		}
	}
	return a
}

func checkDown(a [][]string, i, x int) bool {
	for i1, row := range a {
		for x1, col := range row {
			if x1 == x {
				if i1 == i {
					if col[i1+1] == '.' {
						return false
					}
				}
			}
		}
	}
	return true
}

func checkAcross(a [][]string, i, x int) bool {
	for i1, row := range a {
		for x1, col := range row {
			if x1 == x {
				if i1 == i {
					if col[x1+1] == '.' {
						return false
					}
				}
			}
		}
	}
	return true
}

func matrix(a string) [][]string {
	var (
		temp   []string
		matrix [][]string
	)
	for i, c := range a {
		if (i+1)%4 == 0 || string(c) == "\n" {
			matrix = append(matrix, temp)
		} else {
			temp = append(temp, string(c))
		}
	}
	return matrix
}
