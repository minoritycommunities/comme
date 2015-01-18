package pages

import (
	"html/template"
	"net/http"

	"github.com/comforme/comforme/common"
	"github.com/comforme/comforme/databaseActions"
	"github.com/comforme/comforme/templates"
)

var newPageTemplate *template.Template

func init() {
	newPageTemplate = template.Must(template.New("siteLayout").Parse(templates.SiteLayout))
	template.Must(newPageTemplate.New("nav").Parse(templates.NavBar))
	template.Must(newPageTemplate.New("content").Parse(newPageTemplateText))
}

func NewPageHandler(res http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{}
	data["formAction"] = req.URL.Path
	if req.Method == "POST" {
		cookie, err := req.Cookie("sessionid")
		sessionId := cookie.Value
		if err == nil {
			title := req.PostFormValue("title")
			description := req.PostFormValue("description")
			address := req.PostFormValue("address")
			categories := req.PostFormValue("categories")
			err := databaseActions.CreatePage(sessionId, title, description, address, int(categories[0]))
			if err == nil {
				data["successMsg"] = "Created " + title + "!"
			} else {
				data["errorMsg"] = "Failed to create page!"
			}
		}
	}

	common.ExecTemplate(newPageTemplate, res, data)
}

const newPageTemplateText = `
<div class="row">
	<div class="large-centered medium-centered large-8 medium-8 columns">
	<div class="content" id="add-page-form">
        {{if .successMsg}}<div class="alert-box success">{{.successMsg}}</div>{{end}}
        {{if .errorMsg}}<div class="alert-box alert">{{.errorMsg}}</div>{{end}}
		<form method="POST" action="{{.formAction}}" align="center">
            <fieldset>
            <legend>Create a New Page</legend>
			<div>
				<input type="text" name="title" placeholder="page title" align="center">
			</div>
			<div>
				<textarea name="description" placeholder="description" rows="15"></textarea>
			</div>
			<div>
				<input type="text" name="address/location" placeholder="address">
			</div>
			<div>
				<input type="text" name="categories" placeholder="categories">
			</div>
			<div style="text-align:center">
				<button type="submit" class="button" name="sign-up" value="true">Submit</button>
			</div>
            </fieldset>
		</form>
	</div>
	</div>
</div>		
`