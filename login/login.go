package login

import (
	"net/http"

	"github.com/comforme/comforme/common"
	"github.com/comforme/comforme/databaseActions"
)

var loginTemplate *template.Template

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	var data map[string]interface{}
	if req.Method == "POST" {
		email := req.PostFormValue("email")
		username := req.PostFormValue("username")
		password := req.PostFormValue("password")
		sessionid, err := databaseActions.Login(email, password)
		if err != nil {

		} else {
			common.SetSessionCookie(req, sessionid)
		}
	}

	// TODO: Add template and compile it.
	common.ExecTemplate(nil, res, data)
}

const loginTemplateText = `<!DOCTYPE html>
<html>
<head>
	<link href="https://cdnjs.cloudflare.com/ajax/libs/foundation/5.5.0/css/normalize.min.css" rel="stylesheet" type="text/css" />
	<link href="https://cdnjs.cloudflare.com/ajax/libs/foundation/5.5.0/css/foundation.min.css" rel="stylesheet" type="text/css" />
	<script src="https://cdnjs.cloudflare.com/ajax/libs/foundation/5.5.0/js/jquery.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/foundation/5.5.0/js/foundation.min.js"></script>
	<meta charset="utf-8" />
	<title>ComFor.Me</title>
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/foundation/5.5.0/css/style.css" />
	<script scr="https://cdnjs.cloudflare.com/ajax/libs/foundation/5.5.0/js/login.js"></script>
</head>
<body>
	<nav class="top-bar" data-topbar>
		<ul class="title-area">
			<li class="name"></li>
			<li class="toggle-topbar menu-icon">
				<a href="#">Menu <span class="icon-menu"></span></a>
			</li>
		</ul>
		<section class="top-bar-section">
			<ul class="left">
				<li>
					<a href="/">Main Page</a>
				</li>
			</ul>
		</section>
	</nav>
	<div class="content">
		<div class="row">
			<div class="large-4 medium-3 small-1 columns">&nbsp;</div>
			<div class="large-4 medium-6 small-10 columns">
				<section class="login-tabs">
					<dl class="tabs" data-tab>
						<dd class="active"><a href="#sign-up-form">Sign Up</a></dd>
						<dd><a href="#log-in-form">Log In</a></dd>
					</dl>
					<div class="tabs-content">
						<div class="content active" id="sign-up-form">
							<form method="post" action="/">
								<div>
									<input type="text" name="username" placeholder="User Name"{{if .username}} value="{{.username}}"{{end}}{{if .registerUsernameError}} class="error"{{end}}>
								</div>
								<div>
									<input type="email" name="email" placeholder="Email"{{if .email}} value="{{.email}}"{{end}}{{if .registerEmailError}} class="error"{{end}}>
								</div>
								<div>
									<button type="submit" class="button">Submit</button>
								</div>
							</form>
						</div>
						<div class="content" id="log-in-form">
							<form method="post" action="/">
								<div>
									<input type="email" name="email" placeholder="Email"{{if .email}} value="{{.email}}"{{end}}{{if .loginEmailError}} class="error"{{end}}>
								</div>
								<div>
									<input type="password" name="password" placeholder="Password">
								</div>
								<div>
									<button type="submit" class="button">Submit</button>
								</div>
							</form>
						</div>
					</div>
				</section>
			</div>
			<div class="large-4 medium-3 small-1 columns">&nbsp;</div>
		</div>
	</div>
	<script>$(document).foundation();</script>
</body>
</html>
`
