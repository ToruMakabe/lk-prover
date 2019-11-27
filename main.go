package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const inputFormatMsg = "Please input n^2 * n^2 numbers 0 or 1-9 delimitted by conma. 0 is empty as Sudoku cell."

func prover() int {
	fmt.Print("Antecedent? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	as := strings.Split(scanner.Text(), ",")
	var antecedents []string
	for _, a := range as {
		antecedents = append(antecedents, strings.Join(strings.Fields(a), ""))
	}
	fmt.Println("input is", antecedents)

	fmt.Print("Consequent? ")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cs := strings.Split(scanner.Text(), ",")
	var consequents []string
	for _, c := range cs {
		consequents = append(consequents, strings.Join(strings.Fields(c), ""))
	}
	fmt.Println("input is", consequents)

	st := time.Now()
	// 処理時間を表示する.
	et := time.Now()
	fmt.Println("Time: ", et.Sub(st))

	return 0
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, err.Error()+"\n")
}

func main() {
	os.Exit(prover())
}
