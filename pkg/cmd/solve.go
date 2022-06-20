package cmd

import (
	"fmt"

	"github.com/k-avy/CalGo/pkg/postfix"
	"github.com/k-avy/CalGo/pkg/prefix"
	"github.com/spf13/cobra"
)

var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve an expression",
}

// calgo solve prefix "expression"
func solvePrefix() *cobra.Command {
	prefixCmd := &cobra.Command{
		Use:   "prefix",
		Short: "solve a prefix expression",
		Args:  cobra.ExactArgs(1),
		//DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := prefix.SolvePrefixExpression(args[0])
			if err != nil {
				return err
			}
			fmt.Printf("The result of prefix expression, %s is: %d", args[0], res)
			return nil
		},
	}
	return prefixCmd
}

func solvePostfix() *cobra.Command {
	postfixCmd := &cobra.Command{
		Use:   "postfix",
		Short: "solve a postfix expression",
		Args:  cobra.ExactArgs(1),
		//DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := postfix.SolvePostfixExpression(args[0])
			if err != nil {
				return err
			}
			fmt.Printf("The result of postfix expression, %s is: %d", args[0], res)
			return nil
		},
	}
	return postfixCmd
}

func Command() *cobra.Command {
	solveCmd.AddCommand(solvePrefix(), solvePostfix())
	return solveCmd
}
