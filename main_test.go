package main

import (
	"testing"
)

/*
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
*/

func TestEvalProp(t *testing.T) {

	var antecedents []string
	var consequents []string
	var root node

	// ->L
	antecedents = []string{"A", "A->B"}
	consequents = []string{"B"}
	root = node{nil, antecedents, consequents, "", nil, false}

	if !evalProp(&root) {
		t.Errorf("Failed (->L): The root is %v", root)
	} else {
		t.Logf("Info (->L): The root is %v", root)
	}

	// ->R
	antecedents = []string{"A"}
	consequents = []string{"B", "A->B"}
	root = node{nil, antecedents, consequents, "", nil, false}

	if !evalProp(&root) {
		t.Errorf("Failed (->R): The root is %v", root)
	} else {
		t.Logf("Info (->R): The root is %v", root)
	}

	// &&L
	antecedents = []string{"A", "A&&B"}
	consequents = []string{"B"}
	root = node{nil, antecedents, consequents, "", nil, false}

	if !evalProp(&root) {
		t.Errorf("Failed (&&L): The root is %v", root)
	} else {
		t.Logf("Info (&&L): The root is %v", root)
	}

	// &&R
	antecedents = []string{"A"}
	consequents = []string{"B", "A&&B"}
	root = node{nil, antecedents, consequents, "", nil, false}

	if !evalProp(&root) {
		t.Errorf("Failed (&&R): The root is %v", root)
	} else {
		t.Logf("Info (&&R): The root is %v", root)
	}

	// ||L
	antecedents = []string{"A", "A||B"}
	consequents = []string{"B"}
	root = node{nil, antecedents, consequents, "", nil, false}

	if !evalProp(&root) {
		t.Errorf("Failed (||L): The root is %v", root)
	} else {
		t.Logf("Info (||L): The root is %v", root)
	}

	// ||R
	antecedents = []string{"A"}
	consequents = []string{"B", "A||B"}
	root = node{nil, antecedents, consequents, "", nil, false}

	if !evalProp(&root) {
		t.Errorf("Failed (||R): The root is %v", root)
	} else {
		t.Logf("Info (||R): The root is %v", root)
	}

}
