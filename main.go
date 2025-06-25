package main

import (
	"fmt"
	"github.com/ragolsnagol/ragol-cli/core"
	"github.com/ragolsnagol/ragol-cli/core/action"
	"github.com/ragolsnagol/ragol-cli/core/command"
	"github.com/ragolsnagol/ragol-cli/core/context"
	"github.com/ragolsnagol/ragol-cli/core/flag"
)

func main() {
	f, err := flag.NewFlag("--task", "-t", true, true)
	if err != nil {
		panic(err)
	}
	f2, err := flag.NewFlag("--test", "-t", false, true)
	if err != nil {
		panic(err)
	}

	app := core.NewApp(
		"test cli",
		"0.0.1",
		[]command.BaseCommand{
			*command.NewCommand(
				"test",
				"Test command",
				action.NewAction(testCommand),
				[]flag.Flag{
					*f,
				}),
			*command.NewCommand(
				"test2",
				"Test command 2",
				action.NewAction(testCommand2),
				[]flag.Flag{
					*f2,
				}),
		},
	)
	err = app.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func testCommand(ctx context.Context) error {
	fmt.Println("Testing command runs")
	for _, f := range ctx.Flags {
		fmt.Println(f.Name)
	}
	return nil
}

func testCommand2(ctx context.Context) error {
	fmt.Println("Testing command 2 runs")
	for _, f := range ctx.Flags {
		fmt.Printf("%v: %v\n", f.Name, f.Value)
	}
	return nil
}
