package utils

import (
	"fmt"
	"strings"

	"github.com/gauravshewale/dictionary-command-line-tool/model"
)

func PrintResult(rs *[]model.SearchResults) {
	for _, v := range *rs {
		var desc string
		if len(v.Shortdef) > 0 {
			desc = strings.Join(v.Shortdef, ";")
		}
		fmt.Printf("`%s` (%s): %s\n", v.Hwi.Hw, v.Fl, desc)
		// `ek-sər-sīz` (noun): the act of bringing into play or realizing in action
		fmt.Println("-------------------------------------------------------")

	}
}
