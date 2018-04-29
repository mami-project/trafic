package mixer

import (
	"fmt"
	"io/ioutil"
	"log"
	"text/template"
)

func setupTemplate(c FlowDesc, defaultTmpl string) (*template.Template, error) {
	var tmplStr string

	if c.Template == "" {
		log.Println("using default template")

		tmplStr = defaultTmpl
	} else {
		var err error

		log.Println("using custom template from", c.Template)

		tmplStr, err = loadTemplate(c.Template)
		if err != nil {
			return nil, err
		}
	}

	// fail hard if a key is missing
	tmpl, err := template.New(c.Kind).Option("missingkey=error").Parse(tmplStr)
	if err != nil {
		return tmpl, fmt.Errorf("cannot parse template: %v", err)
	}

	return tmpl, nil
}

func loadTemplate(f string) (string, error) {
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
