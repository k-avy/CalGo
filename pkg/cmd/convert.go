package cmd

import (
	"errors"
	"fmt"

	"github.com/k-avy/CalGo/pkg/postfix"
	"github.com/k-avy/CalGo/pkg/prefix"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type expEnum int

const (
	expEnumPre  expEnum = iota // 0
	expEnumPost expEnum = iota // 1
	expEnumIn   expEnum = iota // 2
)

var _ pflag.Value = (*expEnum)(nil)

func (e *expEnum) String() string {
	switch *e {
	case expEnumPre:
		return "prefix"
	case expEnumPost:
		return "postfix"
	case expEnumIn:
		return "infix"
	default:
		return "nil"
	}
}

func (e *expEnum) Set(s string) error {
	switch s {
	case "prefix":
		*e = expEnumPre
		return nil
	case "postfix":
		*e = expEnumPost
		return nil
	case "infix":
		*e = expEnumIn
		return nil
	default:
		return errors.New(`must be one a "prefix", "postfix", "infix"`)
	}
}

func (e *expEnum) Type() string {
	return "expEnum"
}

/*
pre =0, post=1, in =2,
pre to post 1
post to pre 4
post to in  6
pre to in 2
from<<2 | to
*/

func ConvertExpression() *cobra.Command {
	var fromFlag, toFlag = expEnumPre, expEnumPost
	var convertCmd = &cobra.Command{
		Use:   "convert",
		Short: "Convert an expression",
		RunE: func(cmd *cobra.Command, args []string) error {
			if fromFlag == expEnumIn {
				return errors.New("conversion from infix is not supported yet")
			}
			switch fromFlag<<2 | toFlag {
			case 1: // pre to post
				res, err := prefix.ConvertPretoPost(args[0])
				if err != nil {
					return err
				}
				fmt.Printf("The result of prefix to postfix expression, %s is: %s \n", args[0], res)

			case 2: // pre to in
				res, err := prefix.ConvertPretoIn(args[0])
				if err != nil {
					return err
				}
				fmt.Printf("The result of prefix to infix expression, %s is: %s \n", args[0], res)

			case 4: // post to pre
				res, err := postfix.ConvertPosttoPre(args[0])
				if err != nil {
					return err
				}
				fmt.Printf("The result of postfix to prefix expression, %s is: %s \n", args[0], res)

			case 6: // post to in
				res, err := postfix.ConvertPosttoIn(args[0])
				if err != nil {
					return err
				}
				fmt.Printf("The result of postfix to infix expression, %s is: %s \n", args[0], res)

			default:
				return errors.New("conversion not supported")
			}
			return nil

		},
	}
	convertCmd.Flags().Var(&fromFlag, "from", `Type of exp to convert from. allowed: "prefix", "postfix"`)
	convertCmd.Flags().Var(&toFlag, "to", `Type of exp to convert to. allowed: "prefix", "postfix","infix"`)
	convertCmd.MarkFlagRequired("from")
	convertCmd.MarkFlagRequired("to")
	return convertCmd
}
