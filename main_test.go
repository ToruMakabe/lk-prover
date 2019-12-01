package main

import (
	"fmt"
	"testing"
)

func walk(n node) {
	if n.parent == nil {
		fmt.Println("[Root of sequent]")
		fmt.Printf("%v |- %v\n", n.antecedents, n.consequents)
	} else {
		fmt.Printf("%v |- %v (Parent) %v |- %v\n", n.antecedents, n.consequents, n.parent.antecedents, n.parent.consequents)
	}

	if n.child == nil {
		fmt.Println("[End of branch]")
		fmt.Println()
	}

	for _, c := range n.child {
		walk(*c)
	}
}
func TestParse(t *testing.T) {

	var antecedents []string
	var consequents []string
	var root node

	// Valid
	antecedents = []string{"A", "B"}
	consequents = []string{"A"}
	root = node{nil, antecedents, consequents, "", nil, true}

	parse(&root, &root)

	if !root.valid {
		t.Errorf("Failed (Valid): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (Valid): The root is %v", root)
	}

	// ->L
	antecedents = []string{"A", "A->B"}
	consequents = []string{"B"}
	root = node{nil, antecedents, consequents, "", nil, true}

	parse(&root, &root)

	if !root.valid {
		t.Errorf("Failed (->L): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (->L): The root is %v", root)
	}

	// ->R
	antecedents = []string{"A", "B"}
	consequents = []string{"B", "A->B"}
	root = node{nil, antecedents, consequents, "", nil, true}

	parse(&root, &root)

	if !root.valid {
		t.Errorf("Failed (->R): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (->R): The root is %v", root)
	}

	// &&L
	antecedents = []string{"A", "A&&B"}
	consequents = []string{"B"}
	root = node{nil, antecedents, consequents, "", nil, true}

	parse(&root, &root)

	if !root.valid {
		t.Errorf("Failed (&&L): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (&&L): The root is %v", root)
	}

	// &&R
	antecedents = []string{"A"}
	consequents = []string{"B", "A&&B"}
	root = node{nil, antecedents, consequents, "", nil, true}

	parse(&root, &root)

	if !root.valid {
		t.Errorf("Failed (&&R): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (&&R): The root is %v", root)
	}

	// ||L
	antecedents = []string{"A", "A||B"}
	consequents = []string{"B"}
	root = node{nil, antecedents, consequents, "", nil, true}

	parse(&root, &root)

	if !root.valid {
		t.Errorf("Failed (||L): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (||L): The root is %v", root)
	}

	// ||R
	antecedents = []string{"A"}
	consequents = []string{"B", "A||B"}
	root = node{nil, antecedents, consequents, "", nil, true}

	parse(&root, &root)

	if !root.valid {
		t.Errorf("Failed (||R): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (||R): The root is %v", root)
	}

}

func TestEvalProp(t *testing.T) {

	var antecedents []string
	var consequents []string
	var root node

	// ->L
	antecedents = []string{"A", "A->B"}
	consequents = []string{"B"}
	root = node{nil, antecedents, consequents, "", nil, true}

	if !evalProp(&root) {
		t.Errorf("Failed (->L): The root is %v", root)
	} else {
		t.Logf("Info (->L): The root is %v", root)
	}

	// ->R
	antecedents = []string{"A"}
	consequents = []string{"B", "A->B"}
	root = node{nil, antecedents, consequents, "", nil, true}

	if !evalProp(&root) {
		t.Errorf("Failed (->R): The root is %v", root)
	} else {
		t.Logf("Info (->R): The root is %v", root)
	}

	// &&L
	antecedents = []string{"A", "A&&B"}
	consequents = []string{"B"}
	root = node{nil, antecedents, consequents, "", nil, true}

	if !evalProp(&root) {
		t.Errorf("Failed (&&L): The root is %v", root)
	} else {
		t.Logf("Info (&&L): The root is %v", root)
	}

	// &&R
	antecedents = []string{"A"}
	consequents = []string{"B", "A&&B"}
	root = node{nil, antecedents, consequents, "", nil, true}

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
