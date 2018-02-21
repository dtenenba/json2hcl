package main

import (
	"bytes"
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
)

func main() {
	js.Global.Set("golib", map[string]interface{}{
		"toHCL": toHCL,
	})
}

func toHCL(inputStr string) (string, error) {
	ast, err := jsonParser.Parse([]byte(inputStr))
	if err != nil {
		return "", fmt.Errorf("unable to parse JSON: %s", err)
	}

	var w bytes.Buffer
	err = printer.Fprint(&w, ast)
	if err != nil {
		return "", fmt.Errorf("unable to print HCL: %s", err)
	}

	return w.String(), nil
}
