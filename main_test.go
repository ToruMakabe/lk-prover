package main

import (
	"testing"
)

func TestParse(t *testing.T) {

	var antecedents []string
	var consequents []string
	var root node

	// Valid
	antecedents = []string{"A", "B"}
	consequents = []string{"A"}
	root = node{nil, antecedents, consequents, nil, false}

	parse(&root, &root)

	if !root.valid {
		t.Errorf("Failed (Valid): The root is %v", root)
	}

	// ->L
	antecedents = []string{"A", "A->B"}
	consequents = []string{"B"}
	root = node{nil, antecedents, consequents, nil, false}

	parse(&root, &root)

	if !root.valid {
		t.Errorf("Failed (->L): The root is %v", root)
	}

	// ->R
	antecedents = []string{"A"}
	consequents = []string{"B", "A->B"}
	root = node{nil, antecedents, consequents, nil, false}

	parse(&root, &root)

	if !root.valid {
		t.Errorf("Failed (->R): The root is %v", root)
	}

}

func TestDecompose(t *testing.T) {

	var l string
	var p string
	var a []string
	var c []string
	var lc string
	var d1 [][]string
	var d2 [][]string

	// ->L
	l = "A->B"
	p = "l"
	a = []string{"A"}
	c = []string{"B"}

	lc, d1, d2 = decompose(l, p, a, c)

	if lc == "->L" {
		t.Logf("OK (->L): The propositonal formula is %v %v%v", lc, d1, d2)
	} else {
		t.Errorf("Failed (->L): The propositonal formula is %v %v%v", lc, d1, d2)
	}

	// ->R
	l = "A->B"
	p = "r"
	a = []string{"A"}
	c = []string{"B"}

	lc, d1, d2 = decompose(l, p, a, c)

	if lc == "->R" {
		t.Logf("OK (->L): The propositonal formula is %v %v%v", lc, d1, d2)
	} else {
		t.Errorf("Failed (->L): The propositonal formula is %v %v%v", lc, d1, d2)
	}

}
