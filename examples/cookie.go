package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/meehow/securebytes"
)

const cookieName = "securebytes"

var sb = securebytes.New([]byte(os.Getenv("SECRET")))

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
		var secret map[string]string
		err = sb.DecryptBase64(cookie.Value, &secret)
		if err != nil {
			fmt.Fprintf(w, "Decryption error: %v", err)
			return
		}
		fmt.Fprintf(w, "Your secret cookie: %#v", secret)
		return
	}
	// Cookie not found, create a new one
	secret := map[string]string{"Hi": "Hello"}
	b64, err := sb.EncryptToBase64(&secret)
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
