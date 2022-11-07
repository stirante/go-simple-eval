package main

import (
	"bufio"
	"fmt"
	"github.com/stirante/go-simple-query-language/query"
	"github.com/stirante/go-simple-query-language/query/utils"
	"go.uber.org/zap"
	"os"
	"strings"
)

func main() {
	utils.InitLogging(zap.DebugLevel)
	query.Init()
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
		eval, err := query.Eval(text, scope)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(utils.ToString(eval))
		}
		fmt.Print("> ")
	}
}
