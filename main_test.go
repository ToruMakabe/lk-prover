package main

import (
	"fmt"
	"testing"
)

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
		fmt.Println("Info (Valid): Provable")
		fmt.Println("*** Root of sequent ***")
		walk(root)
		fmt.Println("*** End of sequent ***")
		fmt.Println()
	}

	// ~L
	assumptions = []string{"~A", "A"}
	conclutions = []string{"B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (~L): The root is %v", root)
	} else {
		fmt.Println("Info (~L): Provable")
		fmt.Println("*** Root of sequent ***")
		walk(root)
		fmt.Println("*** End of sequent ***")
		fmt.Println()
	}

	// ~R
	assumptions = []string{"A"}
	conclutions = []string{"~B", "B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (~R): The root is %v", root)
	} else {
		fmt.Println("Info (~R): Provable")
		fmt.Println("*** Root of sequent ***")
		walk(root)
		fmt.Println("*** End of sequent ***")
		fmt.Println()
	}

	// >L
	assumptions = []string{"A", "A>B"}
	conclutions = []string{"B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (>L): The root is %v", root)
	} else {
		fmt.Println("Info (>L): Provable")
		fmt.Println("*** Root of sequent ***")
		walk(root)
		fmt.Println("*** End of sequent ***")
		fmt.Println()
	}

	// >R
	assumptions = []string{"A", "B"}
	conclutions = []string{"B", "A>B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (>R): The root is %v", root)
	} else {
		fmt.Println("Info (>R): Provable")
		fmt.Println("*** Root of sequent ***")
		walk(root)
		fmt.Println("*** End of sequent ***")
		fmt.Println()
	}

	// &L
	assumptions = []string{"A", "A&B"}
	conclutions = []string{"B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (&L): The root is %v", root)
	} else {
		fmt.Println("Info (&L): Provable")
		fmt.Println("*** Root of sequent ***")
		walk(root)
		fmt.Println("*** End of sequent ***")
		fmt.Println()
	}

	// &R
	assumptions = []string{"A"}
	conclutions = []string{"B", "A&B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (&R): The root is %v", root)
	} else {
		fmt.Println("Info (&R): Provable")
		fmt.Println("*** Root of sequent ***")
		walk(root)
		fmt.Println("*** End of sequent ***")
		fmt.Println()
	}

	// |L
	assumptions = []string{"A", "A|B"}
	conclutions = []string{"B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (|L): The root is %v", root)
	} else {
		fmt.Println("Info (|L): Provable")
		fmt.Println("*** Root of sequent ***")
		walk(root)
		fmt.Println("*** End of sequent ***")
		fmt.Println()
	}

	// |R
	assumptions = []string{"A"}
	conclutions = []string{"B", "A|B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (|R): The root is %v", root)
	} else {
		fmt.Println("Info (|R): Provable")
		fmt.Println("*** Root of sequent ***")
		walk(root)
		fmt.Println("*** End of sequent ***")
		fmt.Println()
	}

	// Complex #1
	assumptions = []string{"~A&B", "A&B"}
	conclutions = []string{"B", "C|D"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (Complex #1): The root is %v", root)
	} else {
		fmt.Println("Info (Complex #1): Provable")
		fmt.Println("*** Root of sequent ***")
		walk(root)
		fmt.Println("*** End of sequent ***")
		fmt.Println()
	}

	// Complex #2
	assumptions = []string{"~A>B", "A&B", "C"}
	conclutions = []string{"A>B", "~A"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (Complex #2): The root is %v", root)
	} else {
		fmt.Println("Info (Complex #2): Provable")
		fmt.Println("*** Root of sequent ***")
		walk(root)
		fmt.Println("*** End of sequent ***")
		fmt.Println()
	}

	// Complex #3
	assumptions = []string{"~(A>B)", "A&B", "C"}
	conclutions = []string{"A>B", "~A"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("Failed (Complex #3): The root is %v", root)
	} else {
		fmt.Println("Info (Complex #3): Provable")
		fmt.Println("*** Root of sequent ***")
		walk(root)
		fmt.Println("*** End of sequent ***")
		fmt.Println()
	}

}
