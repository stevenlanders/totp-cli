package command

import (
	"fmt"
	"os"
	"time"

	"github.com/yitsushi/go-commander"
	"github.com/yitsushi/totp-cli/security"
)

// Generate structure is the representation of the generate command
type Generate struct {
}

// Execute is the main function. It will be called on generate command
func (c *Generate) Execute(opts *commander.CommandHelper) {
	token := os.Getenv("TOTP_TOKEN")
	if token == "" {
		panic("TOTP_TOKEN not defined")
	}
	fmt.Println(security.GenerateOTPCode(token, time.Now()))
}

// NewGenerate creates a new Generate command
func NewGenerate(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Generate{},
		Help: &commander.CommandDescriptor{
			Name:             "generate",
			ShortDescription: "Generate a specific OTP",
		},
	}
}
