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

// nodeはシーケントを格納する構造体である.
type node struct {
	parent      *node
	assumptions []string
	conclutions []string
	child       []*node
	valid       bool
}

// walkはシーケントが格納されたツリーを深さ優先で探索し、ノードの前提と結論を、存在すれば親ノードの前提と結論も表示する.
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

// isValidは前提と結論がリテラルのみで、かつ前提と結論に同じリテラルが含まれるか、つまり証明可能なシーケントかを判定する.
func isValid(a []string, c []string) bool {
	m := make(map[string]bool)
	var (
		s []string
		u []string
	)

	// リテラルはAからZまでの1文字であるかで判定する.
	re := regexp.MustCompile(`^[A-Z]$`)

	// 前提にあるリテラルの重複を削除する.
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

	// 前提と結論のリテラルに同じものがあるかを確認する.
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

// decomposeは規則に従ってシーケントを分解する.
func decompose(l string, p string, a []string, c []string) (string, [][]string, [][]string) {

	// pfparser.PfParseは命題論理式を構文解析し、根に論理結合子があれば [(否定)v1] [論理結合子] [v2]の形式で返す. 論理結合子がなければ [(否定)v1]で返す. yaccベースのプログラムである(コード量が多いため、Goのパッケージは分割).
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

	// シーケントを分解する(否定のみ).
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

	// シーケントを分解する(否定の他).
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

// evalPfは命題論理式の集合を構文解析する.
func evalPf(n *node) bool {
	a := n.assumptions
	c := n.conclutions

	// すでにvalidかを判定する.
	if isValid(a, c) {
		n.valid = true
		return true
	}

	// シーケントの前提を構成する命題論理式の集合を解析する.
	for i, s := range a {
		var t []string
		t = append(t, a[:i]...)
		t = append(t, a[i+1:]...)
		// 分解できるかを判定する.
		conn, d1, d2 := decompose(s, "a", t, c)
		// 分解できれば子として追加する.
		if conn != "" {
			child := node{n, d1[0], d1[1], nil, false}
			n.child = append(n.child, &child)
			// 2式に分解された場合.
			if d2 != nil {
				child := node{n, d2[0], d2[1], nil, false}
				n.child = append(n.child, &child)
			}
			return true
		}
	}

	// シーケントの結論を構成する命題論理式の集合を解析する.
	for i, s := range c {
		var t []string
		t = append(t, c[:i]...)
		t = append(t, c[i+1:]...)
		// 分解できるかを判定する.
		conn, d1, d2 := decompose(s, "c", a, t)
		// 分解できれば子として追加する.
		if conn != "" {
			child := node{n, d1[0], d1[1], nil, false}
			n.child = append(n.child, &child)
			// 2式に分解された場合.
			if d2 != nil {
				child := node{n, d2[0], d2[1], nil, false}
				n.child = append(n.child, &child)
			}
			return true
		}
	}

	return false
}

// parseSeqはシーケントを構文解析する.
func parseSeq(r *node, n *node) {
	e := evalPf(n)
	// 解析の結果、この時点で分解しきれていない、validでないと判定できる場合はルートシーケントノードのvaildフラグを偽にする.
	if !e {
		r.valid = false
	}

	// 再帰的に子ノードのシーケントを構文解析する.
	for _, c := range n.child {
		parseSeq(r, c)
	}
}

// proveは実質的な主処理である.
func prove() int {

	// シーケントの入力を促す.
	fmt.Println(inputFormatMsg)
	fmt.Print("Sequent? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// シーケントを "|-" を区切り記号として、前提と結論に分解する.
	s := strings.Split(strings.Join(strings.Fields(scanner.Text()), ""), "|-")

	if len(s) != 2 {
		fmt.Println()
		fmt.Println("Syntax error!!")
		fmt.Println()
		fmt.Println(inputFormatMsg)
		return 1
	}

	// 前提の命題論理式の集合を "," を区切り記号としてスライスに格納する.
	as := strings.Split(s[0], ",")
	var assumptions []string
	if as[0] != "" {
		for _, a := range as {
			assumptions = append(assumptions, a)
		}
	}
	fmt.Println("assumptions: ", assumptions)

	// 結論の命題論理式の集合を "," を区切り記号としてスライスに格納する.
	cs := strings.Split(s[1], ",")
	var conclutions []string
	if cs[0] != "" {
		for _, c := range cs {
			conclutions = append(conclutions, c)
		}
	}
	fmt.Println("conclutions: ", conclutions)

	// 証明可能かを判定するのに要した時間を計測するため、開始時間を取得する.
	st := time.Now()

	root := node{nil, assumptions, conclutions, nil, true}

	// シーケントの構文解析を行う.
	parseSeq(&root, &root)
	// 構文解析の結果vaildなシーケントへ分解できたら、その結果を出力する.できなかった場合は "Unprovable" を出力する.
	if root.valid == true {
		fmt.Println()
		fmt.Println("Provable")
		fmt.Println()
		fmt.Println("*** Root of sequent tree ***")
		walk(root)
		fmt.Println("*** End of sequent tree ***")
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println("Unprovable")
		fmt.Println()
	}

	// 証明可能かを判定するのに要した時間を表示する.
	et := time.Now()
	fmt.Println("Time: ", et.Sub(st))

	return 0
}

// printErrorはエラーメッセージ出力を統一する.
func printError(err error) {
	fmt.Fprintf(os.Stderr, err.Error()+"\n")
}

// mainはエントリーポイントと終了コードを返却する役割のみとする.
func main() {
	os.Exit(prove())
}
