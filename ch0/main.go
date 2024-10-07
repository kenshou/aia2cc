package main

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		println("have 1 arg: return num")
		return
	}
	numStr := os.Args[1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Printf("invalid num: %v", err)
	}
	m := ir.NewModule()
	funcMain := m.NewFunc(
		"main",
		types.I32,
	)
	mb := funcMain.NewBlock("") // llir/llvm would give correct default name for block without name
	mb.NewRet(constant.NewInt(types.I32, int64(num)))
	fmt.Println(m.String())
}
