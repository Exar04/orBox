package expcommand

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Expcommand = &cobra.Command{
    Use:"expo",
    Short:"this is expo",
    Long: "this is experemtnt",
    Run: showExp,
}

func showExp(cmd *cobra.Command, arg []string){
    fmt.Println("ok this is working properly")
}
