package main

import (
	"fmt"
	"testing"
)

func walk(n node) {
	if n.parent == nil {
		fmt.Println("[Root of sequent]")
		fmt.Printf("%v |- %v\n", n.assumptions, n.conclutions)
	} else {
		fmt.Printf("%v |- %v  (Parent: %v |- %v)\n", n.assumptions, n.conclutions, n.parent.assumptions, n.parent.conclutions)
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

	var assumptions []string
	var conclutions []string
	var root node

	// Valid
	assumptions = []string{"A", "B"}
	conclutions = []string{"A"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (Valid): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (Valid): The root is %v", root)
	}

	// ~L
	assumptions = []string{"~A", "A"}
	conclutions = []string{"B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (~L): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (~L): The root is %v", root)
	}

	// ~R
	assumptions = []string{"A"}
	conclutions = []string{"~B", "B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (~R): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (~R): The root is %v", root)
	}

	// >L
	assumptions = []string{"A", "A>B"}
	conclutions = []string{"B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (>L): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (>L): The root is %v", root)
	}

	// >R
	assumptions = []string{"A", "B"}
	conclutions = []string{"B", "A>B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (>R): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (>R): The root is %v", root)
	}

	// &L
	assumptions = []string{"A", "A&B"}
	conclutions = []string{"B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (&L): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (&L): The root is %v", root)
	}

	// &R
	assumptions = []string{"A"}
	conclutions = []string{"B", "A&B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (&R): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (&R): The root is %v", root)
	}

	// |L
	assumptions = []string{"A", "A|B"}
	conclutions = []string{"B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (|L): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (|L): The root is %v", root)
	}

	// |R
	assumptions = []string{"A"}
	conclutions = []string{"B", "A|B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (|R): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (|R): The root is %v", root)
	}

	// Complex #1
	assumptions = []string{"~A&B", "A&B"}
	conclutions = []string{"B", "C|D"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (Complex #1): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (Complex #1): The root is %v", root)
	}

	// Complex #2
	assumptions = []string{"~A>B", "A&B", "C"}
	conclutions = []string{"A>B", "~A"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (Complex #2): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (Complex #2): The root is %v", root)
	}

	// Complex #3
	assumptions = []string{"~(A>B)", "A&B", "C"}
	conclutions = []string{"A>B", "~A"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (Complex #3): The root is %v", root)
	} else {
		walk(root)
		t.Logf("Info (Complex #3): The root is %v", root)
	}

}
