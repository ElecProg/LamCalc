package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/ElecProg/LamCalc"
)

type cLCStatement struct {
	command    string
	parameters []interface{}
}

func executeStatement(stmnt cLCStatement) {
	switch stmnt.command {
	case "exit":
		os.Exit(0)

	case "clear":
		var clearCmd *exec.Cmd

		if runtime.GOOS == "windows" {
			clearCmd = exec.Command("cls")
		} else {
			clearCmd = exec.Command("clear")
		}

		clearCmd.Stdout = os.Stdout
		clearCmd.Run()

	case "let":
		// TODO: Is there an elegant way to stop computation?
		fmt.Print("\n" + stmnt.parameters[0].(string) + " =\n\n")

		lf := stmnt.parameters[1].(LamCalc.LamTerm).Expand()
		globals[stmnt.parameters[0].(string)] = lf

		fmt.Print("    " + lf.String() + "\n\n")

	case "fold":
		// TODO: Is there an elegant way to stop computation?
		expression := stmnt.parameters[0].(LamCalc.LamTerm).Expand()
		fmt.Print("\n" + expression.String() + " =\n\n")

		for _, term := range stmnt.parameters[1].([]string) {
			if globals[term].Equals(expression) {
				fmt.Print("    " + term + "\n\n")
			}
		}

	case "show":
		// TODO: Is there an elegant way to stop computation?
		fmt.Print("\n" + stmnt.parameters[0].(LamCalc.LamTerm).String() + " =\n\n")
		fmt.Print("    " + stmnt.parameters[0].(LamCalc.LamTerm).Expand().String() + "\n\n")
	}
}