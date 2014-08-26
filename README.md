pkgCaptchaGen
======
Captcha .png image fast generate pkg
http://godoc.org/github.com/Cergoo/pkgCaptchaGen
  
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