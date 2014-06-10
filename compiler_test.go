package mutan

import (
	"fmt"
	"strings"
	"testing"
)

func TestCompiler(t *testing.T) {
	asm, err := CompileStage(strings.NewReader(`
	var a
	var b

	if 200 < this.value() {
		a = 20
		b = 30
	}
	`), false)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(asm)
}

func TestArray(t *testing.T) {
	_, err := CompileStage(strings.NewReader(`
	var[10] a
	var[10] b
	var address = 0xa46df28529eb8aa8b8c025b0b413c5f4b688352f
	call(address, 0, 15, a, b)
	// comment int8 b = 20
	`), false)

	if err != nil {
		fmt.Println(err)
	}
}

func TestPlusPlus(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	      var i = 0
	      i++
	`), false)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ast)
}

func TestElse(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	      if 20 < 10 {
		      this.store[1001] = 10^20
	      } else {
		      this.store[1001] = 10^50
	      }
	`), false)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ast)
}

func TestHex(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		var address = 0xa46df28529eb8aa8b8c025b0b413c5f4b688352f
	`), false)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ast)
}

func TestData(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		var t = this.data[2]
		// OR
		asm {
			PUSH 64
			CALLDATALOAD
		}
	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestTransact(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		transact(0xaa1adef765cd, 100, nil)
	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestCreate(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	a := 0xdeadbeef
	res := create(10000, a)
	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestOpt(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		var[2] test
		test[0] = 1000000000000000000000000000000000000000000
	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestBalance(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		var a = this.balance()
		if a > 10 {
			var b = 10
		}
	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestOps(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		var test = 10 >= 20
		var itst2 = 10 <= 20
		var test3 = 10 != 20
		var test4 = 10 % 1 == 0

		test4 = test & test4
		test4 = test4 | test3
		test4 = test4 ^ test3
		test4 = test4 ** 2
	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestBoolean(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		var t = true
		var f = false
	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestNot(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		if !this.data[0] {
		}
	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestShift(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		var test = 8 << 2
		test = 256 >> 3
	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestLambda(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	a := "hello world"
	a = "test"
	b := 20

	var[2] c
	c[0] = 10
	c[2] = 100

	d := "hello"

	exit compile {
		to := this.data[0]
		from := this.origin()
		value := this.data[1]

		if this.store[from] >= value {
			this.store[from] = this.store[from] - value
			this.store[to] = this.store[to] + value
		}
	}
	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestString(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	a := "hello world"
	b := 10
	this.store[1] = "hello"

	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestReturn(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
		return 1000
		return this.store[this.data[0]]
		var a = 10
		var b = 20
		return b
	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestSome(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	var a = this.address()
	var in = "jeff"
	var out
	var res = call(0xa4976648142a1e624c27dca4e9b1a6d8195f660c, 0, 0, in, out)
	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestCc(t *testing.T) {
	instr1 := newIntInstr(intPush1, "")
	instr2 := newIntInstr(intPush2, "")
	instr3 := newIntInstr(intPush3, "")
	instr4 := newIntInstr(intPush4, "")
	instr5 := newIntInstr(intPush5, "")
	instr6 := newIntInstr(intPush6, "")

	cc(nil, nil, instr1, instr2, nil, instr3, instr4, instr5, instr6)
}

func TestFuncDef(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	func two() {
		return 2
	}

	func one() {
		var b = two()

		return b
	}

	var a = one()
	var b = "hello"

	exit a

	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestFuncArgs(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	func fn(var a, var b) {
		var[10] c
		c[2] = b
		return c[1]
	}

	var a = fn("test", 5)
	exit a
	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestPointers(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	var a = 10
	var *b = &a
	*b = 9
	var c = *b
	`), false)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ast)
}

func TestAsm(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	func sha3(var *a, var s) var {
		m_push(a + s)
		m_push(a)
		asm {
			sha3
			push1 9   ; This is a comemnt
			pop       ; This is another comment
		}

		return m_pop() // Returns outcome of the SHA3
	}

	var a = 5
	var b = sha3(&a, 1)
	var *c = &a
	*c = *c * *c


	exit b
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}

}

func TestWhile(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	i := 0
	for i < 10 {
		i++
	}
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}

}

func TestIpmort(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	import "test"
	`), false)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}

}
func TestStatementList(t *testing.T) {
	ast, err := CompileStage(strings.NewReader(`
	var a = 10
	var b = a + 2
	`), true)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(ast)
	}
}
