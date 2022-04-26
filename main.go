package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"kubequery/internal/database"
	"kubequery/internal/terminal"
)

func main() {
	var err error
	sqlClient := database.NewSqliteDatabase("kubequery")
	database.DB, err = sqlClient.Open(
		&database.DbOptions{
			DbType: "sqlite",
			DbConn: "kubequery",
		})
	if err != nil {
		panic(err)
	}
	db, err := database.DB.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 进入交互模式
	fmt.Printf("kubequery v1.0\n")
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")
	defer fmt.Println("Bye!")
	p := prompt.New(terminal.Executor,
		terminal.Completer,
		prompt.OptionTitle("kubequery: query k8s resource(s) use sql"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
	)
	p.Run()
}
