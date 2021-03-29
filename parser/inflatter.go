package parser

import (
	"errors"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/onigura/paich/utils"
)

func Flat(g GenerateCmdArgToken) error {
	defer g.Suicide()

	paichTemplateDir, _ := os.Getwd()
	paichTemplateDir += "/template/"
	if !utils.FolderExists(paichTemplateDir) {
		return errors.New("Template directory missing on your current directory.")
	}

	if &g.Template == nil {
		return errors.New("Template name invalid, use --template=[name].")
	}

	paichTemplateDir += (*g.Template + ".paich")
	if !utils.FileExists(paichTemplateDir) {
		return errors.New("Template unavailable, make sure you template name same with the name provide on flags template." + paichTemplateDir)
	}

	file, _ := os.Open(paichTemplateDir)
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	tmpl := template.Must(template.New(*g.Template).Funcs(template.FuncMap{
		"deref": func(i *FieldToken) FieldToken { return *i },
	}).Parse(string(fileBytes)),
	)

	return tmpl.Execute(os.Stdout, g)
}
