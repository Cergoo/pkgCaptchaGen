// (c) 2014 Cergoo
// under terms of ISC license

// Package pkgCaptchaGen it's .png image generator implementation
package pkgCaptchaGen

import (
	"bytes"
	"github.com/Cergoo/gol/err"
	"github.com/Cergoo/pkgCaptchaGen/imageUtil"
	"image"
	"image/png"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
)

type (
  // TCaptchaGen main struct 
	TCaptchaGen struct {
		alphabetChars    string             // chars
		alphabetImages   []*image.RGBA      // images
		alphabetLength   uint8              // length
		resultLength     uint8              // length of a result
		resultWidthHight [2]int             // width and hight of a result
		distort          imageUtil.IDistort //
	}
)

// New constructor of a new capcha generator
func New(pathtofile string, length uint8, width int) *TCaptchaGen {
	var (
		e          error
		file       *os.File
		imgdecod   image.Image
		filesource []byte
	)
	captchaGen := new(TCaptchaGen)
	captchaGen.resultLength = length
	captchaGen.resultWidthHight[0] = width
	file, e = os.Open(pathtofile)
	err.Panic(e)
	imgdecod, e = png.Decode(file)
	file.Close()
	err.Panic(e)
	captchaGen.alphabetImages, captchaGen.resultWidthHight[1] = imageUtil.Split(imgdecod)
	dir := path.Dir(pathtofile)
	filesource, e = ioutil.ReadFile(dir + "/_chars.txt")
	err.Panic(e)
	captchaGen.alphabetChars = string(filesource)
	captchaGen.distort = new(imageUtil.TDistort1)

	return captchaGen
}

// Gen gen new captcha
func (t *TCaptchaGen) Gen() (bytes.Buffer, string) {
	var (
		r        int
		currentx int
		b        bytes.Buffer
	)

	img := image.NewRGBA(image.Rect(0, 0, t.resultWidthHight[0], t.resultWidthHight[1]))
	rezult := ""
	t.distort.Init()
	for i := uint8(0); i < t.resultLength; i++ {
		r = rand.Intn(len(t.alphabetImages))
		rezult += string(t.alphabetChars[r])
		t.distort.Distort(t.alphabetImages[r], img, currentx, 0)
		currentx += t.alphabetImages[r].Bounds().Dx()
	}
	png.Encode(&b, img)
	return b, rezult
}
