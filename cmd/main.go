package main

import (
	"fmt"
	"log"
	"os"

	"github.com/linzeyan/qrcode"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "qrcode",
	Short:   "Generate or read QRCode.",
	Long:    "Generate a QRCode for a given message or output a message from a QRCode.",
	Args:    cobra.OnlyValidArgs,
	Run:     rootCmdRun,
	Version: "v0.0.1",
}
var f, m string

func rootCmdRun(_ *cobra.Command, _ []string) {
	if f != "" {
		result, err := qrcode.ReadQRCode(f)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(result)
		return
	}
	err := qrcode.GenerateQRCode(m, 600, "/tmp/qrcode.tmp.png")
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	rootCmd.Flags().StringVarP(&f, "file", "f", "", "Specify file path to read")
	rootCmd.Flags().StringVarP(&m, "message", "m", "", "Input message")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
