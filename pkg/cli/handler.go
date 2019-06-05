package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	commandNAme = "gogen"
)

type Func func(arg []string)

type CommandHandler interface {
	AddRoutes()
	SetDefault(fn func(arg []string))
	Run()
}

type commandHandler struct {
	rootCmd *cobra.Command
	cmds    []*cobra.Command
	fn      []Func
}

func NewCommandHandler() CommandHandler {
	return &commandHandler{
		cmds: []*cobra.Command{},
	}
}

func (h *commandHandler) AddRoutes() {
	h.rootCmd = &cobra.Command{
		Use: commandNAme,
		Run: h.run,
	}
	h.cmds = []*cobra.Command{}

	h.cmds = append(h.cmds, &cobra.Command{
		Use:   "config [show , auth , run , all]",
		Short: "Wizard to configure the run",
		Long:  `config is used to prpmpt question needed for the basic configuration.`,
		Args:  cobra.MaximumNArgs(1),
		Run:   h.config,
	})


	//h.cmds = append(h.cmds, &cobra.Command{
	//	Use:   "version",
	//	Short: "Print the version number of " + commandNAme,
	//	Long:  `All software has versions. This is Hugo's`,
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("v1.0")
	//	},
	//})

	h.rootCmd.AddCommand(h.cmds...)
}

func (h *commandHandler) SetDefault(fn func(arg []string)) {
	h.fn = fn
}
func (h *commandHandler) Run() {
	h.rootCmd.Execute()
}

/////////////////////////////////////////////
// Handler

func (h *commandHandler) config(cmd *cobra.Command, args []string) {
	if h.validateFirstTime() {
		return
	}

	var confType string
	if len(args) == 0 {
		confType = "run"
	} else {
		confType = args[0]
	}

	switch confType {
	case "all":
		h.confSrv.ConfigAll()
	case "auth":
		h.confSrv.ConfigToken()
	case "run":
		h.confSrv.ConfigRun()
	case "show":
		h.confSrv.ShowConfig()
	default:
		fmt.Println("Available optinos are: auth , run , all , show")
	}
}

func (h *commandHandler) run(cmd *cobra.Command, args []string) {

}
