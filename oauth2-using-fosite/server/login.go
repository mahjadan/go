package server

import (
	"fmt"
	"net/http"
	"strings"

	"html/template"
)

// oauth/login handler checks if the method is not post it renders the form,
// and if its a POST it validate the username and password and if successful redirect with authorization_code.
// otherwise it rerender the form with validations errors.
func LoginEndpoint(writer http.ResponseWriter, request *http.Request) {

	tpl := template.Must(template.ParseFiles("./static/login.html"))

	if request.Method != http.MethodPost {
		tpl.Execute(writer, nil)
		return
	}
	// when successfully logged in we redirect back to the oauth/auth page to continue with the authorization_code
	// here we are saving the queryParameters in a variable to re-pass it later.
	//the queryParameters that we got initially from the oauth2/auth redirect. (check the authEndPoint)
	refererHeader := request.Header.Get("Referer")
	queryParams := strings.Split(refererHeader,"?")

	username := request.FormValue("username")
	password := request.FormValue("password")

	//username and password check usually done with the database.
	if username == "foo" && password == "secret" {

		// creating session to save the username and we check if its available in the oauth2/auth
		// in this way we know that this user was authenticated and we can generate authorization_code
		// TODO improve this mechanism of session verification (we can save the customerId instead of the username )
		session, err := sessionStore.Get(request, SessionName)
		if err != nil {
			handleSessionError(writer, err)
			return
		}

		session.Values["username"] = username
		if err := session.Save(request, writer); err != nil {
			handleSessionError(writer, err)
			return
		}

		url:="/oauth2/auth"
		if len(queryParams) > 1 {
			url = "/oauth2/auth?" + queryParams[1]
		}
		http.Redirect(writer, request, url, 303)

	}else {

		// Unsuccessful login

		validationErr := struct {
			ErrMsg string
			Success bool
		}{
			ErrMsg: "invalid username or password",
			Success: false,
		}

		err := tpl.Execute(writer, validationErr)
		if err != nil {
			fmt.Println("executing template error: " , err)
		}
	}
}
