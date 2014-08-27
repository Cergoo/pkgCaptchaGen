pkgCaptchaGen
======
Captcha .png image fast generate pkg [![GoDoc](https://godoc.org/github.com/Cergoo/pkgCaptchaGen/imageUtil?status.svg)](https://godoc.org/github.com/Cergoo/pkgCaptchaGen/imageUtil)  
  
Install
-------
go get github.com/Cergoo/pkgCaptchaGen

Example
-------
![Image](https://raw.githubusercontent.com/Cergoo/pkgCaptchaGen/master/example/1.png)

Usage
-----
    captcha = pkgCaptchaGen.New("../fonts/Number/perpetua_bold1.png", 4, 150)
    img, txt := captcha.Gen()
    
See example/main.go

license
-------
Under terms of ISC license      