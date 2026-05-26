package test

import (
	"github.com/spf13/cobra"
)

func Execute(cmd *cobra.Command, args ...string) (output string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ExecuteC(cmd *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}
