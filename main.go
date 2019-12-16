package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/ToruMakabe/lk-prover/formula"
)

const inputFormatMsg = "Please input LK sequent as (assumptions) |- (conclutions)\nPropositional variables: A-Z\nNagation: ~, And: &, Or: |, Imply: >\nYou can specify multiple assumtions/conclutions delimitted by comma\nSample: A&B,C |- A,B\n"

// nodeはシーケントを格納する構造体である.
type node struct {
	parent      *node
	assumptions []string
	conclutions []string
	child       []*node
	valid       bool
}

// proveは実質的な主処理である.
func prove() int {

	// シーケントの入力を促す.
	fmt.Println(inputFormatMsg)
	fmt.Print("Sequent? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// シーケントを "|-" を区切り記号として, 前提と結論に分解する.
	s := strings.Split(strings.Join(strings.Fields(scanner.Text()), ""), "|-")
	if err := scanner.Err(); err != nil {
		printError(fmt.Errorf("scanner error"))
		return 1
	}

	if len(s) != 2 {
		fmt.Println()
		printError(fmt.Errorf("syntax error"))
		fmt.Println()
		fmt.Println(inputFormatMsg)
		return 1
	}

	// 前提の命題論理式を "," を区切り記号としてスライスに格納する.
	as := strings.Split(s[0], ",")
	var assumptions []string
	if as[0] != "" {
		for _, a := range as {
			assumptions = append(assumptions, a)
		}
	}
	fmt.Println("assumptions: ", assumptions)

	// 結論の命題論理式を "," を区切り記号としてスライスに格納する.
	cs := strings.Split(s[1], ",")
	var conclutions []string
	if cs[0] != "" {
		for _, c := range cs {
			conclutions = append(conclutions, c)
		}
	}
	fmt.Println("conclutions: ", conclutions)

	// 証明可能かを判定するのに要した時間を計測するため, 開始時間を取得する.
	st := time.Now()

	root := node{nil, assumptions, conclutions, nil, true}

	// シーケントの構文解析と評価を行う.
	err := parseSeq(&root, &root)
	if err != nil {
		fmt.Println()
		printError(err)
		fmt.Println()
		fmt.Println(inputFormatMsg)
		return 1
	}
	// 構文解析と評価の結果,恒真なシーケントへ分解できたら, その結果を出力する.できなかった場合は "Unprovable" を出力する.
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

// parseSeqはシーケントを構文解析と評価する.
func parseSeq(r /* root */ *node, n /* node */ *node) error {
	res, err := parsePfs(n)
	if err != nil {
		return err
	}
	// 解析の結果,この時点で分解しきれていない, 恒真でないと判定できる場合は根ノードのvalidフラグを偽にする.
	if !res {
		r.valid = false
	}

	// 再帰的に子ノードのシーケントを構文解析と評価する.
	for _, c := range n.child {
		parseSeq(r, c)
	}

	return nil
}

// parsePfsは命題論理式(列)を構文解析と評価する.
func parsePfs(n /* node */ *node) (bool, error) {
	a := n.assumptions
	c := n.conclutions

	// すでに恒真かを判定する.
	if isValid(a, c) {
		n.valid = true
		return true, nil
	}

	// シーケントの前提を構成する命題論理式を解析する.
	for i, f := range a {
		var t []string
		t = append(t, a[:i]...)
		t = append(t, a[i+1:]...)
		// 分解できるかを判定する.
		conn, d1, d2, err := decomposeSeq(f, "a", t, c)
		if err != nil {
			return false, err
		}
		// 分解できれば子として追加する.
		if conn != "" {
			child := node{n, d1[0], d1[1], nil, false}
			n.child = append(n.child, &child)
			// 2式に分解された場合.
			if d2 != nil {
				child := node{n, d2[0], d2[1], nil, false}
				n.child = append(n.child, &child)
			}
			return true, nil
		}
	}

	// シーケントの結論を構成する命題論理式を解析する.
	for i, f := range c {
		var t []string
		t = append(t, c[:i]...)
		t = append(t, c[i+1:]...)
		// 分解できるかを判定する.
		conn, d1, d2, err := decomposeSeq(f, "c", a, t)
		if err != nil {
			return false, err
		}
		// 分解できれば子として追加する.
		if conn != "" {
			child := node{n, d1[0], d1[1], nil, false}
			n.child = append(n.child, &child)
			// 2式に分解された場合.
			if d2 != nil {
				child := node{n, d2[0], d2[1], nil, false}
				n.child = append(n.child, &child)
			}
			return true, nil
		}
	}

	return false, nil
}

// isValidは前提と結論がリテラルのみで, かつ前提と結論に同じリテラルが含まれる, つまり恒真で証明可能なシーケントかを判定する.
func isValid(a /* assumptions */ []string, c /* conclutions */ []string) bool {
	m := make(map[string]bool)
	var (
		f []string
		u []string
	)

	// リテラルはAからZまでの1文字であるかで判定する.
	re := regexp.MustCompile(`^[A-Z]$`)

	// 前提にあるリテラルの重複を削除する.
	f = append(a)
	for _, key := range f {
		if re.FindStringSubmatch(key) == nil {
			return false
		}
		if !m[key] {
			m[key] = true
			u = append(u, key)
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

// decomposeSeqは規則に従ってシーケントを分解する.
func decomposeSeq(f /* formula */ string, p /* position */ string, a /* assumptions */ []string, c /* conclutions */ []string) (string, [][]string, [][]string, error) {

	// formula.Evalは命題論理式を評価し, 根に論理結合子があれば [(否定)v1] [論理結合子] [v2]の形式で返す. 論理結合子がなければ [(否定)v1]で返す.
	r, err := formula.Eval(f)
	if err != nil {
		return "", nil, nil, err
	}

	conn := r[1]
	v1 := r[0]
	v2 := r[2]

	var (
		rv1 [][]string
		rv2 [][]string
	)

	if r == nil {
		return "", nil, nil, nil
	}

	// シーケントを分解する(否定のみ).
	if conn == "" {
		if strings.HasPrefix(v1, "~") {
			if p == "a" {
				rv1 = append(rv1, a)
				rv1 = append(rv1, append(c, strings.TrimLeft(v1, "~")))
				return "~L", rv1, rv2, nil
			}
			rv1 = append(rv1, append(a, strings.TrimLeft(v1, "~")))
			rv1 = append(rv1, c)
			return "~R", rv1, rv2, nil
		}
		rv1 = append(rv1, []string{v1})
		return "", rv1, nil, nil
	}

	// シーケントを分解する(否定の他).
	switch conn {
	case ">":
		if p == "a" {
			rv1 = append(rv1, a)
			rv1 = append(rv1, append(c, v1))
			rv2 = append(rv2, append(a, v2))
			rv2 = append(rv2, c)
			return ">L", rv1, rv2, nil
		}
		rv1 = append(rv1, append(a, v1))
		rv1 = append(rv1, append(c, v2))
		return ">R", rv1, nil, nil
	case "&":
		if p == "a" {
			rv1 = append(rv1, append(a, v1, v2))
			rv1 = append(rv1, c)
			return "&L", rv1, nil, nil
		}
		rv1 = append(rv1, a)
		rv1 = append(rv1, append(c, v1))
		rv2 = append(rv2, a)
		rv2 = append(rv2, append(c, v2))
		return "&R", rv1, rv2, nil
	case "|":
		if p == "a" {
			rv1 = append(rv1, append(a, v1))
			rv1 = append(rv1, c)
			rv2 = append(rv2, append(c, v1))
			rv2 = append(rv2, c)
			return "|L", rv1, rv2, nil
		}
		rv1 = append(rv1, a)
		rv1 = append(rv1, append(c, v1, v2))
		return "|R", rv1, nil, nil
	}
	return "", nil, nil, nil
}

// walkはシーケントが格納されたツリーを深さ優先で探索し, ノードの前提と結論を, 親ノードが存在すればその前提と結論も表示する.
func walk(n /* node */ node) {
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

// printErrorはエラーメッセージ出力を統一する.
func printError(err /* error */ error) {
	fmt.Fprintf(os.Stderr, err.Error()+"\n")
}

// mainはエントリーポイントと終了コードを返す役割のみとする.
func main() {
	os.Exit(prove())
}
