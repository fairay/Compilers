package main

import (
	"course/compilers/ast"
	"course/compilers/parser"
	"course/compilers/utils"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/antlr4-go/antlr/v4"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

type TreeShapeListener struct {
	*parser.BasetinycListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (listener *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("[%d]\t %s\n", ctx.GetRuleIndex(), ctx.GetText())
}

func (listener *TreeShapeListener) EnterInteger(ctx *parser.IntegerContext) {
	fmt.Println("INT: ", ctx.GetText())
}

func main() {
	walkTree(utils.Config().SourcePath)
	module, err := minimalIR()
    if err != nil {
        panic(err)
    }

    dumpModule(module)
}

func dumpModule(module *ir.Module) {
	err := ioutil.WriteFile(
        utils.Config().OutputPath,
        []byte(module.String()),
        0644,
    )
	if err != nil {
        log.Panicf("Failed to write LLVM IR to file: %s", err.Error())
	}
}

func walkTree(fileName string) {
	input, _ := antlr.NewFileStream(fileName)
	lexer := parser.NewtinycLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	tinycParser := parser.NewtinycParser(stream)

	tinycParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tinycParser.BuildParseTrees = true
	tree := tinycParser.Program()
	// antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

	visitor := ast.NewVisitor()
	visitor.VisitProgram(tree.(*parser.ProgramContext))
}


func minimalIR() (*ir.Module, error) {
    m := ir.NewModule()
    mainFunc := m.NewFunc("main", types.I32)
    entry := mainFunc.NewBlock("")
    entry.NewRet(constant.NewInt(types.I32, 42))
    return m, nil
}


func createIR() (*ir.Module, error) {
    // Создаём нужные типы и константы
    i32 := types.I32
    zero := constant.NewInt(i32, 0)
    a := constant.NewInt(i32, 0x15A4E35) // умножаем PRNG.
    c := constant.NewInt(i32, 1)         // инкрементируем PRNG.

    // Создаём новый модуль LLVM IR.
    m := ir.NewModule()

    // Создаём объявление внешней функции и добавляем его к модулю.
    //
    //    int abs(int x);
    abs := m.NewFunc("abs", i32, ir.NewParam("x", i32))

    // Создаём определение глобальной переменной и добавляем его к модулю.
    //
    //    int seed = 0;
    seed := m.NewGlobalDef("seed", zero)

    // Создаём определение функции и добавляем его к модулю.
    //
    //    int rand(void) { ... }
    rand := m.NewFunc("rand", i32)

    // Создаём неименованный входной базовый блок и добавляем его к функции `rand`.
    entry := rand.NewBlock("")

    // Создаём команды и добавляем их к входному базовому блоку.
    tmp1 := entry.NewLoad(seed.Type(), seed)
    tmp2 := entry.NewMul(tmp1, a)
    tmp3 := entry.NewAdd(tmp2, c)
    // entry.NewStore(tmp3, seed)
    tmp4 := entry.NewCall(abs, tmp3)
    entry.NewRet(tmp4)

    // Печатаем ассемблер LLVM IR этого модуля.
    fmt.Println(m)
    return m, nil
}


func init() {
    utils.Config(os.Args[1:]...)
}
