package process

import (
	"fmt"

	"github.com/rkitamu/gomono/internal/astutil"
)

func Step1ImportAnalysis(mainPath string) error {
	imports, err := astutil.ExtractImports(mainPath)
	if err != nil {
		return err
	}
	fmt.Println("Imports:", imports)
	return nil
}
