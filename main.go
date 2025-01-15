package main

import (
	"fmt"
	chain "github.com/lianzhilu/mychain/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "mychain",
		Short: "mychain",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("please enter an option")
		},
	}
	cmd.AddCommand(chain.NewCommand())
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
