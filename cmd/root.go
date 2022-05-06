package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "Danh sách tool Nhập môn mật mã học",
		Long:  "Danh sách tool Nhập môn mật mã học",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
