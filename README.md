pkgCaptchaGen
======
Captcha .png image fast generate pkg
  
Install
-------
go get github.com/Cergoo/pkgCaptchaGen

Example
-------
[Image](https://github.com/Cergoo/pkgCaptchaGen/example/1.png)

Usage
-----
    captcha = pkgCaptchaGen.New("../fonts/Number/perpetua_bold1.png", 4, 150)
    img, txt := captcha.Gen()
    
See example/main.go

license
-------
Under terms of ISC license      