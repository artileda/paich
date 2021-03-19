package parser

import (
	"errors"
	"regexp"
	"strings"
)

type DataDescToken struct {
	DataName string
	DataType string
	DataAtrr []string
}

type GenerateCmdArgToken struct {
	FieldName string
	DataDesc  []*DataDescToken
	Template  *string
}

func (g *GenerateCmdArgToken) Parse(lines *[]string, template string) error {
	noSpecialCharacter, _ := regexp.Compile("^[a-zA-Z0-9]{1,69}$")

	if noSpecialCharacter.MatchString((*lines)[0]) {
		g.FieldName = (*lines)[0]
	} else {
		return errors.New("Special character on field name not allowed.")
	}

	var dataDescs []*DataDescToken
	for _, e := range (*lines)[1:] {
		dataDesc := parseDataDesc(e)
		dataDescs = append(dataDescs, dataDesc)
	}

	g.Template = &template
	return nil
}

func (g *GenerateCmdArgToken) Suicide() {
	g = nil
}

func parseDataDesc(line string) *DataDescToken {
	stemmed := strings.Split(line, ":")

	if len(stemmed) < 3 {
		return &DataDescToken{
			DataName: stemmed[0],
			DataType: stemmed[1],
			DataAtrr: stemmed[2:],
		}
	} else {
		return nil
	}
}
