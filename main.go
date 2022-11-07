package main

import (
	"bufio"
	"fmt"
	"github.com/stirante/go-simple-query-language/eval"
	"github.com/stirante/go-simple-query-language/eval/utils"
	"go.uber.org/zap"
	"os"
	"strings"
)

func main() {
	utils.InitLogging(zap.DebugLevel)
	eval.Init()
	repl(map[string]interface{}{
		"version": utils.Semver{Major: 1, Minor: 17},
	})
}

func repl(scope map[string]interface{}) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	for true {
		read, _ := reader.ReadString('\n')
		text := strings.TrimRight(read, "\n\r")
		if text == "exit" {
			break
		}
		eval, err := eval.Eval(text, scope)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(utils.ToString(eval))
		}
		fmt.Print("> ")
	}
}
