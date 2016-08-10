package controllers

import (
	"math/rand"
	"time"

	"github.com/astaxie/beego"
	"github.com/fogleman/gg"
)

type CheckCodeController struct {
	Width, Height, Mixed, Length int
	CodeTmpl                     string
	Randhandler                  *rand.Rand
}

func NewCheckCode() *CheckCodeController {
	ccc := CheckCodeController{
		Width:    100,
		Height:   42,
		Mixed:    5,
		Length:   4,
		CodeTmpl: "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz0123456789",
	}
	ccc.Randhandler = rand.New(rand.NewSource(time.Now().UnixNano()))
	return &ccc
}
func (this *CheckCodeController) GetCode() (dc *gg.Context, code string) {
	dc = gg.NewContext(this.Width, this.Height)
	dc.Clear()
	this.randomBack(dc)
	s := float64(this.Width / this.Mixed)
	for j := 0; j < this.Mixed; j++ {
		x := float64(j) * s
		y := float64(this.Randhandler.Int63n(int64(this.Height)))
		dc.Push()
		dc.Translate(x, y)
		dc.Scale(s/2, s/2)
		this.randomQuadratic(dc)
		dc.Pop()
	}
	c := this.randCode()
	this.writeCode(c, dc)
	return dc, c
}
func (this *CheckCodeController) randCode() string {
	var v []byte
	for i := 0; i < this.Length; i++ {
		j := int32(len(this.CodeTmpl))
		k := this.Randhandler.Int31n(j)
		v = append(v, this.CodeTmpl[k])
	}
	return string(v)
}
func (this *CheckCodeController) writeCode(c string, dc *gg.Context) {
	sdir := beego.AppConfig.String("StaticDir")
	if sdir == "" {
		sdir = "static"
	}
	fdir := sdir + "/css/fonts"
	fsize := float64(this.Width/len(c)) * 0.85
	fontlist, _ := DirList(fdir)
	for i := 0; i < len(c); i++ {
		fn := fontlist[int(this.Randhandler.Int63n(int64(len(fontlist))))]
		err := dc.LoadFontFace(fn, fsize)
		beego.Debug(fn)
		if err != nil {
			beego.Debug(err)
		}
		r, g, b, _ := this.randRgba()
		dc.SetRGB255(r, g, b)
		s := beego.Substr(c, i, 1)
		x := fsize*float64(i) + fsize*0.6
		y := fsize * 1.3
		dc.DrawString(s, x, y)
		r, g, b, _ = this.randRgba()
		dc.SetRGB255(r, g, b)
		dc.DrawString(s, x-1, y-1)
	}
}
func (this *CheckCodeController) random() float64 {
	return this.Randhandler.Float64()
}

func (this *CheckCodeController) point() (x, y float64) {
	return this.random(), this.random()
}
func (this *CheckCodeController) randRgba() (r, g, b, a int) {
	r = int(this.Randhandler.Int31n(256))
	g = int(this.Randhandler.Int31n(256))
	b = int(this.Randhandler.Int31n(256))
	a = int(this.Randhandler.Int31n(100))
	return
}
func (this *CheckCodeController) drawCurve(dc *gg.Context) {
	dc.SetRGBA255(this.randRgba())
	dc.FillPreserve()
	dc.SetRGBA255(this.randRgba())
	dc.SetLineWidth(this.random() * float64(this.Height))
	dc.Stroke()
}

func (this *CheckCodeController) drawPoints(dc *gg.Context) {
	dc.SetRGBA255(this.randRgba())
	dc.SetLineWidth(this.random() * float64(this.Height/10))
	dc.Stroke()
}

func (this *CheckCodeController) randomQuadratic(dc *gg.Context) {
	x0, y0 := this.point()
	x1, y1 := this.point()
	x2, y2 := this.point()
	dc.MoveTo(x0, y0)
	dc.QuadraticTo(x1, y1, x2, y2)
	this.drawCurve(dc)
	this.drawPoints(dc)
}

func (this *CheckCodeController) randomBack(dc *gg.Context) {
	r1, g1, b1, a1 := this.randRgba()
	r2, g2, b2, a2 := this.randRgba()
	dc.SetRGBA255(r1, g1, b1, a1)
	dc.SetLineWidth(1)
	for i := 0; i < this.Height; i++ {
		y := float64(i)
		dc.DrawLine(0, y, float64(this.Width), y)
		dc.Stroke()
		r3 := (r2-r1)*i/this.Height - 1 + r1
		g3 := (g2-g1)*i/this.Height - 1 + g1
		b3 := (b2-b1)*i/this.Height - 1 + b1
		a3 := (a2-a1)*i/this.Height - 1 + a1
		dc.SetRGBA255(r3, g3, b3, a3)
	}
}
