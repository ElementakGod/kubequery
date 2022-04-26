package terminal

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"kubequery/internal/database"
	"os"
	"strings"
)

// Completer
// @Description: sql提示，用于交互状态下输入提示
// @param d
// @return []prompt.Suggest
func Completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}

	args := strings.Split(d.TextBeforeCursor(), " ")
	w := d.GetWordBeforeCursor()

	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	s := []prompt.Suggest{
		{Text: "show tables", Description: "show all tables"},
		{Text: "exit", Description: "exit the program"},
		{Text: "select", Description: ""},
		{Text: "insert", Description: ""},
		{Text: "into", Description: ""},
		{Text: "values", Description: ""},
		{Text: "update", Description: ""},
		{Text: "delete", Description: ""},
		{Text: "from", Description: ""},
		{Text: "where", Description: ""},
		{Text: "and", Description: ""},
		{Text: "inner", Description: ""},
		{Text: "left", Description: ""},
		{Text: "right", Description: ""},
		{Text: "full", Description: ""},
		{Text: "join", Description: ""},
		{Text: "on", Description: ""},
		{Text: "set", Description: ""},
		{Text: "limit", Description: ""},
		{Text: "order", Description: ""},
		{Text: "asc", Description: ""},
		{Text: "desc", Description: ""},
		{Text: "null", Description: ""},
		{Text: "like", Description: ""},
		{Text: "is", Description: ""},
		{Text: "not", Description: ""},
	}
	return prompt.FilterHasPrefix(s, w, true)
}

func Executor(in string) {
	db, err := database.DB.DB()
	if err != nil {
		os.Exit(0)
	}
	//response := strings.TrimSpace(prompt.Input("kubequery >>> ", Completer))
	response := strings.TrimSpace(in)
	responseArr := strings.Fields(response)
	if len(responseArr) > 0 {
		if response == "show tables" {
			ShowTables(db)
			return
		}

		cmd := strings.ToUpper(responseArr[0])

		if cmd == "EXIT" {
			os.Exit(0)
		}

		if cmd == "SELECT" {
			PrintTable(GetData(db, response))
			return
		}

		affectedRows, err := RunQuery(db, response)
		if err != nil {
			fmt.Printf("Error in %v : %v\n", response, err.Error())
			return
		}

		fmt.Printf("%v row(s) affected...\n", affectedRows)
	}
}
