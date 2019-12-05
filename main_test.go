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
		t.Errorf("(Already valid test): Failed. The root is %v", root)
	} else {
		fmt.Println("(Already valid test): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// ~L
	assumptions = []string{"~A", "A"}
	conclutions = []string{"B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(~L test): Failed. The root is %v", root)
	} else {
		fmt.Println("(~L test): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// ~R
	assumptions = []string{"A"}
	conclutions = []string{"~B", "B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(~R test): Failed. The root is %v", root)
	} else {
		fmt.Println("(~R test): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// >L
	assumptions = []string{"A", "A>B"}
	conclutions = []string{"B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(>L test): Failed. The root is %v", root)
	} else {
		fmt.Println("(>L test): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// >R
	assumptions = []string{"A", "B"}
	conclutions = []string{"B", "A>B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(>R test): Failed. The root is %v", root)
	} else {
		fmt.Println("(>R test): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// &L
	assumptions = []string{"A", "A&B"}
	conclutions = []string{"B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(&L test): Failed. The root is %v", root)
	} else {
		fmt.Println("(&L test): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// &R
	assumptions = []string{"A"}
	conclutions = []string{"A", "A&B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(&R test): Failed. The root is %v", root)
	} else {
		fmt.Println("(&R test): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// |L
	assumptions = []string{"A", "A|B"}
	conclutions = []string{"A"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(|L test): Failed. The root is %v", root)
	} else {
		fmt.Println("(|L test): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// |R
	assumptions = []string{"A"}
	conclutions = []string{"B", "A|B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(|R test): Failed. The root is %v", root)
	} else {
		fmt.Println("(|R test): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// Complex #1
	assumptions = []string{"~A&B", "A&B"}
	conclutions = []string{"B", "C|D"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(Complex test #1): Failed. The root is %v", root)
	} else {
		fmt.Println("(Complex test #1): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// Complex #2
	assumptions = []string{"~A>B", "A&B", "C"}
	conclutions = []string{"A>B", "~A"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(Complex test #2): Failed. The root is %v", root)
	} else {
		fmt.Println("(Complex test #2): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// Complex #3
	assumptions = []string{"~(A>B)", "A&B", "C"}
	conclutions = []string{"A>B", "~A"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(Complex test #3): Failed. The root is %v", root)
	} else {
		fmt.Println("(Complex test #3): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// Hilbert-style deduction system axiom #1.1
	assumptions = []string{}
	conclutions = []string{"A>(B>A)"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(Hilbert axiom test #1.1): Failed. The root is %v", root)
	} else {
		fmt.Println("(Hilbert axiom test #1.1): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// Hilbert-style deduction system axiom #1.2
	assumptions = []string{}
	conclutions = []string{"(A>(B>C))>((A>B)>(A>C))"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(Hilbert axiom test #1.2): Failed. The root is %v", root)
	} else {
		fmt.Println("(Hilbert axiom test #1.2): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// Hilbert-style deduction system axiom #2.1
	assumptions = []string{}
	conclutions = []string{"A>(A|B)"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(Hilbert axiom test #2.1): Failed. The root is %v", root)
	} else {
		fmt.Println("(Hilbert axiom test #2.1): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// Hilbert-style deduction system axiom #2.2
	assumptions = []string{}
	conclutions = []string{"B>(A|B)"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(Hilbert axiom test #2.2): Failed. The root is %v", root)
	} else {
		fmt.Println("(Hilbert axiom test #2.2): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// Hilbert-style deduction system axiom #2.3
	assumptions = []string{}
	conclutions = []string{"(A>C)>((B>C)>((A|B)>C))"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(Hilbert axiom test #2.3): Failed. The root is %v", root)
	} else {
		fmt.Println("(Hilbert axiom test #2.3): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// Hilbert-style deduction system axiom #3.1
	assumptions = []string{}
	conclutions = []string{"(A&B)>A"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(Hilbert axiom test #3.1): Failed. The root is %v", root)
	} else {
		fmt.Println("(Hilbert axiom test #3.1): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// Hilbert-style deduction system axiom #3.2
	assumptions = []string{}
	conclutions = []string{"(A&B)>B"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(Hilbert axiom test #3.2): Failed. The root is %v", root)
	} else {
		fmt.Println("(Hilbert axiom test #3.2): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// Hilbert-style deduction system axiom #4.1
	assumptions = []string{}
	conclutions = []string{"(A>B)>((A>~B)>~A)"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(Hilbert axiom test #4.1): Failed. The root is %v", root)
	} else {
		fmt.Println("(Hilbert axiom test #4.1): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// Hilbert-style deduction system axiom #4.2
	assumptions = []string{}
	conclutions = []string{"A>(~A>B)"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(Hilbert axiom test #4.2): Failed. The root is %v", root)
	} else {
		fmt.Println("(Hilbert axiom test #4.2): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

	// Hilbert-style deduction system axiom #4.3
	assumptions = []string{}
	conclutions = []string{"A|~A"}
	root = node{nil, assumptions, conclutions, nil, true}

	parseSeq(&root, &root)

	if !root.valid {
		t.Errorf("(Hilbert axiom test #4.3): Failed. The root is %v", root)
	} else {
		fmt.Println("(Hilbert axiom test #4.3): Provable")
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	}

}
