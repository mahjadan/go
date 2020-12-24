package server

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/storage"
	"github.com/ory/fosite/token/jwt"
)

var (
	// Check the api documentation of `compose.Config` for further configuration options.
	config = &compose.Config{
		AccessTokenLifespan: time.Minute * 30,
		// ...
	}
	//Using imMemory store with predefined clients and users.
	store         = storage.NewExampleStore()
	secret        = []byte("some-cool-secret-that-is-32bytes")
	privateKey, _ = rsa.GenerateKey(rand.Reader, 2048)

	sessionStore sessions.Store
	log          = logrus.WithField("cmd", "go-oauth2-demo")
)

var oauth2 = compose.ComposeAllEnabled(config, store, secret, privateKey)

const (
	//SessionName to store session under
	SessionName = "go-oauth2-demo"
)


func InitSessionStore(){
	SESSION_ENCRYPTION_KEY:="00000000000000000000000000000000"
	SESSION_AUTHENTICATION_KEY:="secret-session"

	sessionStore = sessions.NewCookieStore(
		[]byte(SESSION_AUTHENTICATION_KEY),
		[]byte(SESSION_ENCRYPTION_KEY),
	)
}


func RegisterHandlers() {
	// Set up oauth2 endpoints. You could also use gorilla/mux or any other router.

	//to test this authorization_code endPoint
	//use this url
	// http://localhost:8080/oauth2/auth?client_id=my-client&redirect_uri=http://localhost:3846/callback&response_type=code&scope=photos+openid+offline&state=some-random-state-foobar&nonce=some-random-nonce
	//check for the storage.NewExampleStore() client info
	http.HandleFunc("/oauth2/auth", AuthEndpoint)
	http.HandleFunc("/oauth2/login", LoginEndpoint)

	// TODO
	//http.HandleFunc("/oauth2/token", tokenEndpoint)
	// revoke tokens
	// http.HandleFunc("/oauth2/revoke", revokeEndpoint)
	// http.HandleFunc("/oauth2/introspect", introspectionEndpoint)
}

func newSession(user string) *openid.DefaultSession {
	return &openid.DefaultSession{
		Claims: &jwt.IDTokenClaims{
			Issuer:      "https://fosite.my-application.com",
			Subject:     user,
			Audience:    []string{"https://my-client.my-application.com"},
			ExpiresAt:   time.Now().Add(time.Hour * 6),
			IssuedAt:    time.Now(),
			RequestedAt: time.Now(),
			AuthTime:    time.Now(),
		},
		Headers: &jwt.Headers{
			Extra: make(map[string]interface{}),
		},
	}
}
