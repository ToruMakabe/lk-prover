package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/ToruMakabe/lk-prover/pfparser"
)

const inputFormatMsg = "Please input LK sequent as (assumptions) |- (conclutions)\nNagation:~, And:&, Or:|, Implication:>\nYou can specify multiple assumtions/conclutions delimitted by comma\nSample: A&B,C |- A,B\n"

type node struct {
	parent      *node
	assumptions []string
	conclutions []string
	child       []*node
	valid       bool
}

func walk(n node) {
	if n.parent == nil {
		fmt.Printf("%v |- %v\n", n.assumptions, n.conclutions)
	} else {
		fmt.Printf("%v |- %v  (Parent: %v |- %v)\n", n.assumptions, n.conclutions, n.parent.assumptions, n.parent.conclutions)
	}

	if n.child == nil {
		fmt.Println("** End of branch **")
	}

	for _, c := range n.child {
		walk(*c)
	}
}

func isValid(a []string, c []string) bool {
	m := make(map[string]bool)
	var (
		s []string
		u []string
	)

	re := regexp.MustCompile(`^[A-Z]$`)

	s = append(a)
	for _, e := range s {
		if re.FindStringSubmatch(e) == nil {
			return false
		}
		if !m[e] {
			m[e] = true
			u = append(u, e)
		}
	}

	v := false
	for _, i := range u {
		for _, j := range c {
			if re.FindStringSubmatch(j) == nil {
				return false
			}
			if i == j {
				v = true
			}
		}
	}
	if v == true {
		return true
	}
	return false
}

func decompose(l string, p string, a []string, c []string) (string, [][]string, [][]string) {
	pf := pfparser.PfParse(l)

	conn := pf[1]
	v1 := pf[0]
	v2 := pf[2]

	var (
		rv1 [][]string
		rv2 [][]string
	)

	if pf == nil {
		return "", nil, nil
	}

	if conn == "" {
		if strings.HasPrefix(v1, "~") {
			if p == "a" {
				rv1 = append(rv1, a)
				rv1 = append(rv1, append(c, strings.TrimLeft(v1, "~")))
				return "~L", rv1, rv2
			}
			rv1 = append(rv1, append(a, strings.TrimLeft(v1, "~")))
			rv1 = append(rv1, c)
			return "~R", rv1, rv2
		}
		rv1 = append(rv1, []string{v1})
		return "", rv1, nil
	}

	switch conn {
	case ">":
		if p == "a" {
			rv1 = append(rv1, a)
			rv1 = append(rv1, append(c, v1))
			rv2 = append(rv2, append(a, v2))
			rv2 = append(rv2, c)
			return ">L", rv1, rv2
		}
		rv1 = append(rv1, append(a, v1))
		rv1 = append(rv1, append(c, v2))
		return ">R", rv1, nil
	case "&":
		if p == "a" {
			rv1 = append(rv1, append(a, v1, v2))
			rv1 = append(rv1, c)
			return "&L", rv1, nil
		}
		rv1 = append(rv1, a)
		rv1 = append(rv1, append(c, v1))
		rv2 = append(rv2, a)
		rv2 = append(rv2, append(c, v2))
		return "&R", rv1, rv2
	case "|":
		if p == "a" {
			rv1 = append(rv1, append(a, v1))
			rv1 = append(rv1, c)
			rv2 = append(rv2, append(c, v1))
			rv2 = append(rv2, c)
			return "|L", rv1, rv2
		}
		rv1 = append(rv1, a)
		rv1 = append(rv1, append(c, v1, v2))
		return "|R", rv1, nil
	}
	return "", nil, nil
}

func evalPf(n *node) bool {
	a := n.assumptions
	c := n.conclutions
	if isValid(a, c) {
		n.valid = true
		return true
	}

	for i, s := range a {
		var t []string
		t = append(t, a[:i]...)
		t = append(t, a[i+1:]...)
		conn, d1, d2 := decompose(s, "a", t, c)
		if conn != "" {
			//			n.conn = conn
			child := node{n, d1[0], d1[1], nil, false}
			n.child = append(n.child, &child)
			if d2 != nil {
				child := node{n, d2[0], d2[1], nil, false}
				n.child = append(n.child, &child)
			}
			return true
		}
	}

	for i, s := range c {
		var t []string
		t = append(t, c[:i]...)
		t = append(t, c[i+1:]...)
		conn, d1, d2 := decompose(s, "c", a, t)
		if conn != "" {
			child := node{n, d1[0], d1[1], nil, false}
			n.child = append(n.child, &child)
			if d2 != nil {
				child := node{n, d2[0], d2[1], nil, false}
				n.child = append(n.child, &child)
			}
			return true
		}
	}

	return false
}

func parseSeq(r *node, n *node) int {
	e := evalPf(n)
	if !e {
		r.valid = false
	}

	for _, c := range n.child {
		parseSeq(r, c)
	}

	return 0
}

func prove() int {

	fmt.Println(inputFormatMsg)
	fmt.Print("Sequent? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := strings.Split(strings.Join(strings.Fields(scanner.Text()), ""), "|-")

	if len(s) != 2 {
		fmt.Println(inputFormatMsg)
		return 1
	}

	as := strings.Split(s[0], ",")
	var assumptions []string
	if as[0] != "" {
		for _, a := range as {
			assumptions = append(assumptions, a)
		}
	}
	fmt.Println("assumptions: ", assumptions)

	cs := strings.Split(s[1], ",")
	var conclutions []string
	if cs[0] != "" {
		for _, c := range cs {
			conclutions = append(conclutions, c)
		}
	}
	fmt.Println("conclutions: ", conclutions)

	/* for debug
	assumptions := []string{}
	conclutions := []string{"A>(B>A)"}
	*/

	st := time.Now()

	root := node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)
	if root.valid == true {
		fmt.Println("Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	} else {
		fmt.Println("Unprovable")
	}

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
