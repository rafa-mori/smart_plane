package main

import (
	cc "github.com/rafa-mori/smart_plane/cmd/cli"
	gl "github.com/rafa-mori/smart_plane/logger"
	vs "github.com/rafa-mori/smart_plane/version"
	"github.com/spf13/cobra"

	"os"
	"strings"
)

type SmartPlane struct {
	parentCmdName string
	printBanner   bool
}

func (m *SmartPlane) Alias() string { return "" }
func (m *SmartPlane) ShortDescription() string {
	return "SmartPlane is a minimalistic backend service with Go."
}
func (m *SmartPlane) LongDescription() string {
	return `SmartPlane: A minimalistic backend service with Go.`
}
func (m *SmartPlane) Usage() string {
	return "smart_plane [command] [args]"
}
func (m *SmartPlane) Examples() []string {
	return []string{"smart_plane start -p ':8080' -b '0.0.0.0' -n 'MyService' -d"}
}
func (m *SmartPlane) Active() bool {
	return true
}
func (m *SmartPlane) Module() string {
	return "smart_plane"
}
func (m *SmartPlane) Execute() error { return m.Command().Execute() }
func (m *SmartPlane) Command() *cobra.Command {
	gl.Log("debug", "Starting SmartPlane CLI...")

	var rtCmd = &cobra.Command{
		Use:     m.Module(),
		Aliases: []string{m.Alias()},
		Example: m.concatenateExamples(),
		Version: vs.GetVersion(),
		Annotations: cc.GetDescriptions([]string{
			m.LongDescription(),
			m.ShortDescription(),
		}, m.printBanner),
	}

	// rtCmd.AddCommand(cc.CertificatesCmdList())
	// rtCmd.AddCommand(cc.ServiceCmdList()...)

	rtCmd.AddCommand(vs.CliCommand())

	// Set usage definitions for the command and its subcommands
	setUsageDefinition(rtCmd)
	for _, c := range rtCmd.Commands() {
		setUsageDefinition(c)
		if !strings.Contains(strings.Join(os.Args, " "), c.Use) {
			if c.Short == "" {
				c.Short = c.Annotations["description"]
			}
		}
	}

	return rtCmd
}
func (m *SmartPlane) SetParentCmdName(rtCmd string) {
	m.parentCmdName = rtCmd
}
func (m *SmartPlane) concatenateExamples() string {
	examples := ""
	rtCmd := m.parentCmdName
	if rtCmd != "" {
		rtCmd = rtCmd + " "
	}
	for _, example := range m.Examples() {
		examples += rtCmd + example + "\n  "
	}
	return examples
}

func RegX() *SmartPlane {
	var printBannerV = os.Getenv("GOBEMIN_PRINT_BANNER")
	if printBannerV == "" {
		printBannerV = "true"
	}

	return &SmartPlane{
		printBanner: strings.ToLower(printBannerV) == "true",
	}
}
