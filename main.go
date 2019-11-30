package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

const inputFormatMsg = "Please input n^2 * n^2 numbers 0 or 1-9 delimitted by conma. 0 is empty as Sudoku cell."

type node struct {
	parent      *node
	antecedents []string
	consequents []string
	conn        string
	child       []*node
	valid       bool
}

func isValid(a []string, c []string) bool {
	m := make(map[string]bool)
	var s []string
	var u []string

	s = append(a)
	for _, e := range s {
		if !m[e] {
			m[e] = true
			u = append(u, e)
		}
	}

	for _, i := range u {
		for _, j := range c {
			if i == j {
				return true
			}
		}
	}
	return false
}

func decompose(l string, p string, a []string, c []string) (string, [][]string, [][]string) {
	// ToDo: Nagation case

	re := regexp.MustCompile(`^\s*(~?[A-Z!])\s*(->|\|{2}|&&)\s*(~?[A-Z!])\s*$`)
	pf := re.FindStringSubmatch(l)
	if pf == nil {
		return "", nil, nil
	}

	conn := pf[2]
	v1 := pf[1]
	v2 := pf[3]

	var rv1 [][]string
	var rv2 [][]string

	switch conn {
	case "->":
		if p == "l" {
			rv1 = append(rv1, a)
			rv1 = append(rv1, append(c, v1))
			rv2 = append(rv2, append(a, v2))
			rv2 = append(rv2, c)
			return "->L", rv1, rv2
		}
		rv1 = append(rv1, append(a, v1))
		rv1 = append(rv1, append(c, v2))
		return "->R", rv1, nil
	case "&&":
		if p == "l" {
			rv1 = append(rv1, append(a, v1, v2))
			rv1 = append(rv1, c)
			return "&&L", rv1, nil
		}
		rv1 = append(rv1, a)
		rv1 = append(rv1, []string{v1})
		rv2 = append(rv2, c)
		rv2 = append(rv2, []string{v2})
		return "&&R", rv1, rv2
	case "||":
		if p == "l" {
			rv1 = append(rv1, a)
			rv1 = append(rv1, []string{v1})
			rv2 = append(rv2, c)
			rv2 = append(rv2, []string{v2})
			return "||L", rv1, rv2
		}
		rv1 = append(rv1, a)
		rv1 = append(rv1, append(c, v1, v2))
		return "||R", rv1, nil
	}
	return "", nil, nil
}

func evalProp(n *node) bool {
	a := n.antecedents
	c := n.consequents
	if isValid(a, c) {
		return true
	}

	for i, s := range a {
		var t []string
		t = append(t, a[:i]...)
		t = append(t, a[i+1:]...)
		conn, d1, d2 := decompose(s, "l", t, c)
		if conn != "" {
			n.conn = conn
			fmt.Printf("Proposotional Formula: %s %s %s\n", conn, d1, d2)
			if d1 != nil {
				child := node{n, d1[0], d1[1], "", nil, false}
				n.child = append(n.child, &child)
			}
			if d2 != nil {
				child := node{n, d2[0], d2[1], "", nil, false}
				n.child = append(n.child, &child)
			}
			return true
		}
	}

	for i, s := range c {
		var t []string
		t = append(t, c[:i]...)
		t = append(t, c[i+1:]...)
		conn, d1, d2 := decompose(s, "r", a, t)
		if conn != "" {
			n.conn = conn
			fmt.Printf("Proposotional Formula: %s %s %s\n", conn, d1, d2)
			if d1 != nil {
				child := node{n, d1[0], d1[1], "", nil, false}
				n.child = append(n.child, &child)
			}
			if d2 != nil {
				child := node{n, d2[0], d2[1], "", nil, false}
				n.child = append(n.child, &child)
			}
			return true
		}
	}

	return false
}

func parse(r *node, n *node) int {
	e := evalProp(r)
	if e == true {
		r.valid = true
	}
	/*
		for _, c := range n.child {
			parse(r, &c)
		}
	*/

	return 0
}

func prove() int {

	fmt.Print("Sequent? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := strings.Split(strings.Join(strings.Fields(scanner.Text()), ""), "|-")

	as := strings.Split(s[0], ",")
	var antecedents []string
	for _, a := range as {
		antecedents = append(antecedents, a)
	}
	fmt.Println("Antecedents: ", antecedents)

	cs := strings.Split(s[1], ",")
	var consequents []string
	for _, c := range cs {
		consequents = append(consequents, c)
	}
	fmt.Println("Consequents: ", consequents)

	// Debug
	// antecedents := []string{"A", "A->B", "A"}
	// consequents := []string{"A"}

	st := time.Now()

	root := node{nil, antecedents, consequents, "", nil, false}
	fmt.Println("Root: ", root)

	parse(&root, &root)
	fmt.Println("Valid: ", root.valid)
	fmt.Println("Decomposition Tree: ", root)

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
