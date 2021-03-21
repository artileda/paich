package parser

import (
	"errors"
	"regexp"
	"strings"
)

type FieldToken struct {
	FieldName string
	FieldType string
	FieldAttr []string
}

type GenerateCmdArgToken struct {
	FieldName string
	Fields    []*FieldToken
	Template  *string
}

func (g *GenerateCmdArgToken) Parse(lines *[]string, template string) error {
	noSpecialCharacter, _ := regexp.Compile("^[a-zA-Z0-9]{1,69}$")

	if noSpecialCharacter.MatchString((*lines)[0]) {
		g.FieldName = (*lines)[0]
	} else {
		return errors.New("Special character on field name not allowed.")
	}

	var dataDescs []*FieldToken
	for _, e := range (*lines)[1:] {
		dataDesc := parseDataDesc(e)
		dataDescs = append(dataDescs, dataDesc)
	}
	g.Fields = dataDescs

	g.Template = &template
	return nil
}

func (g *GenerateCmdArgToken) Suicide() {
	g = nil
}

func parseDataDesc(line string) *FieldToken {
	stemmed := strings.Split(line, ":")

	if len(stemmed) < 4 {
		return &FieldToken{
			FieldName: stemmed[0],
			FieldType: stemmed[1],
			FieldAttr: stemmed[2:],
		}
	} else {
		return nil
	}
}
