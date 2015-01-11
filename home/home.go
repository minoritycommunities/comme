package home

import (
	"github.com/comforme/comforme/common"
	"html/template"
	"net/http"
	// "github.com/comforme/comforme/databaseActions"
	"github.com/comforme/comforme/templates"
)

var homeTemplate *template.Template

func init() {
	homeTemplate = template.Must(template.New("siteLayout").Parse(templates.SiteLayout))
	template.Must(homeTemplate.New("content").Parse(homeTemplateText))
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{}

	// TODO: Add template and compile it.
	tmpl, _ := template.New("test").Parse(rootPageTemplateHtml)
	common.ExecTemplate(tmpl, res, data)
}


const homeTemplateText = `<p>Home Page</p>`
