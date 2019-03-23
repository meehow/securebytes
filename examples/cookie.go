package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/meehow/securebytes"
)

const cookieName = "securebytes"

var sb = securebytes.New([]byte(os.Getenv("SECRET")), securebytes.ASN1Serializer{})

// Session is a struct which will be saved in a cookie
type Session struct {
	UserID int
	Name   string
}

func main() {
	http.HandleFunc("/", handler)
	addr := "localhost:8080"
	fmt.Printf("Listening on http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(cookieName)
	if err == nil {
		// Cookie found, let's read it
		var session Session
		err = sb.DecryptBase64(cookie.Value, &session)
		if err != nil {
			fmt.Fprintf(w, "Decryption error: %v", err)
			return
		}
		fmt.Fprintf(w, "Your session cookie: %#v has been encoded to %s",
			session, cookie.Value)
		return
	}
	// Cookie not found, create a new one
	session := Session{
		UserID: 1234567890,
		Name:   "meehow",
	}
	b64, err := sb.EncryptToBase64(session)
	if err != nil {
		fmt.Fprintf(w, "Encryption error: %v", err)
		return
	}
	cookie = &http.Cookie{
		Name:     cookieName,
		Value:    b64,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	fmt.Fprint(w, "The cookie has been set. You can refresh this page to read it.")
}
