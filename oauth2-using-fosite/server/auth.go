package server

import (
	"fmt"

	"net/http"
)

func AuthEndpoint(rw http.ResponseWriter, req *http.Request) {
	// This context will be passed to all methods.
	ctx := req.Context()

	// Let's create an AuthorizeRequest object!
	// It will analyze the request and extract important information like scopes, response type and others.
	ar, err := oauth2.NewAuthorizeRequest(ctx, req)

	if err != nil {
		log.Printf("Error occurred in NewAuthorizeRequest: %+v", err)
		oauth2.WriteAuthorizeError(rw, ar, err)
		return
	}
	// You have now access to authorizeRequest, Code ResponseTypes, Scopes ...

	var requestedScopes string
	for _, this := range ar.GetRequestedScopes() {
		requestedScopes += fmt.Sprintf(`<li><input type="checkbox" name="scopes" value="%s">%s</li>`, this, this)
	}


	// Normally, this would be the place where you would check if the user is logged in and gives his consent.
	// We're simplifying things and just checking if the request includes a valid username and password

	//check if the session exists, if exists check if the username in that session , if so we continue otherwise we redirect to the login-page
	session, err := sessionStore.Get(req, SessionName)
	if err != nil {
		handleSessionError(rw, err)
		return
	}


	//session.Values["username"]=""
	//session.Save(req, rw)
	username, found := session.Values["username"]
	if !found || username == "" {
		queryParam := req.URL.RawQuery // in order to not lose the original request parameters we pass the queryParameters to the login page
		http.Redirect(rw, req, "/oauth2/login" +"?"+queryParam, http.StatusSeeOther)
		log.WithField("username", username).Info("Username is empty/notfound, redirecting")
		return
	}

	req.ParseForm()
	//if req.PostForm.Get("username") != "peter" {
	//	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	//	rw.Write([]byte(`<h1>Login page</h1>`))
	//	rw.Write([]byte(fmt.Sprintf(`
	//		<p>Howdy! This is the log in page. For this example, it is enough to supply the username.</p>
	//		<form method="post">
	//			<p>
	//				By logging in, you consent to grant these scopes:
	//				<ul>%s</ul>
	//			</p>
	//			<input type="text" name="username" /> <small>try peter</small><br>
	//			<input type="submit">
	//		</form>
	//	`, requestedScopes)))
	//	return
	//}

	// let's see what scopes the user gave consent to
	for _, scope := range req.PostForm["scopes"] {
		ar.GrantScope(scope)
	}

	// Now that the user is authorized, we set up a session:
	mySessionData := newSession(username.(string))

	// When using the HMACSHA strategy you must use something that implements the HMACSessionContainer.
	// It brings you the power of overriding the default values.
	//
	// mySessionData.HMACSession = &strategy.HMACSession{
	//	AccessTokenExpiry: time.Now().Add(time.Day),
	//	AuthorizeCodeExpiry: time.Now().Add(time.Day),
	// }
	//

	// If you're using the JWT strategy, there's currently no distinction between access token and authorize code claims.
	// Therefore, you both access token and authorize code will have the same "exp" claim. If this is something you
	// need let us know on github.
	//
	// mySessionData.JWTClaims.ExpiresAt = time.Now().Add(time.Day)

	// It's also wise to check the requested scopes, e.g.:
	// if ar.GetRequestedScopes().Has("admin") {
	//     http.Error(rw, "you're not allowed to do that", http.StatusForbidden)
	//     return
	// }

	// Now we need to get a response. This is the place where the AuthorizeEndpointHandlers kick in and start processing the request.
	// NewAuthorizeResponse is capable of running multiple response type handlers which in turn enables this library
	// to support open id connect.
	response, err := oauth2.NewAuthorizeResponse(ctx, ar, mySessionData)

	// Catch any errors, e.g.:
	// * unknown client
	// * invalid redirect
	// * ...
	if err != nil {
		log.Printf("Error occurred in NewAuthorizeResponse: %+v", err)
		oauth2.WriteAuthorizeError(rw, ar, err)
		return
	}

	//** we can remove the session of this user as his transaction was successful
	//session.Values["username"]= "" // either delete the value
	session.Options.MaxAge = -1 // or you can make it expired (delete the whole session)
	err = session.Save(req, rw)
	if err := session.Save(req, rw); err != nil {
		handleSessionError(rw, err)
		return
	}
	// Last but not least, send the response!
	oauth2.WriteAuthorizeResponse(rw, ar, response)
}



func handleSessionError(w http.ResponseWriter, err error) {
	log.WithField("err", err).Info("Error handling session.")
	http.Error(w, "Application Error", http.StatusInternalServerError)
}