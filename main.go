package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const inputFormatMsg = "Please input n^2 * n^2 numbers 0 or 1-9 delimitted by conma. 0 is empty as Sudoku cell."

type node struct {
	parent *node
	left   []string
	right  []string
	child  []node
	valid  bool
}

func evalProp(n node) bool {
	return true
}

func parse(r node, n node) int {
	e := evalProp(r)
	if e == false {
		r.valid = false
	}
	for _, c := range n.child {
		parse(r, c)
	}

	return 0
}

func prove() int {
	fmt.Print("Sequent? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := strings.Split(scanner.Text(), "|-")

	as := strings.Split(s[0], ",")
	var antecedents []string
	for _, a := range as {
		antecedents = append(antecedents, strings.Join(strings.Fields(a), ""))
	}
	fmt.Println("Antecedents: ", antecedents)

	cs := strings.Split(s[1], ",")
	var consequents []string
	for _, c := range cs {
		consequents = append(consequents, strings.Join(strings.Fields(c), ""))
	}
	fmt.Println("Consequents: ", consequents)

	st := time.Now()

	root := node{nil, antecedents, consequents, nil, false}
	fmt.Println("Root: ", root)

	parse(root, root)

	// 処理時間を表示する.
	et := time.Now()
	fmt.Println("Time: ", et.Sub(st))

	return 0
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, err.Error()+"\n")
}

func main() {
	os.Exit(prove())
}
