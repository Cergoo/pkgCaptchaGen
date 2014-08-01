package main

import (
	"encoding/base64"
	"fmt"
	"github.com/Cergoo/pkgCaptchaGen"
	"net/http"
)

var (
	captcha *pkgCaptchaGen.TCaptchaGen
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store, no-cache")
	img, txt := captcha.Gen()
	pngbase64 := base64.StdEncoding.EncodeToString(img.Bytes())

	fmt.Fprint(w, `<!doctype html>
<meta charset=utf-8>
<html lang=ru><head></head><body><div>`, "<img src=\"data:image/png;base64,", pngbase64, "\">", txt, "</div></body></html>")
}

func main() {
	captcha = pkgCaptchaGen.New("../fonts/Number/perpetua_bold1.png", 4, 150)
	http.HandleFunc("/", get)
	http.ListenAndServe("localhost:9999", nil)
}
