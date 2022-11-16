package main

import (
	convert "erdira.com/date-converter/convert"
	"github.com/thatisuday/commando"
)

func main() {
	commando.
		SetExecutableName("excelToJson").
		SetVersion("1.0.0").
		SetDescription("CLI tool to create json files from given excel file")

	commando.Register("convert").
		SetDescription("Cmd to convert excel table to json").
		AddFlag("input,i", "path to input excel file", commando.String, "").
		AddFlag("output,o", "path to output directory", commando.String, "").
		SetAction(convert.Run)

	commando.Parse(nil)
}
