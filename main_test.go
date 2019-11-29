package main

import (
	"testing"
)

func TestParse(t *testing.T) {

	antecedents := []string{"A", "A->B", "A"}
	consequents := []string{"A"}
	root := node{nil, antecedents, consequents, nil, false}

	parse(&root, &root)

	if !root.valid {
		t.Fatalf("The root is %v", root)
	}

	t.Logf("The root is %v", root)
}
