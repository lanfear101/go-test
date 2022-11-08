package image

import (
	"bytes"
	"fmt"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"myapp/internal/dto"
	"myapp/pkg/colors"
)

const (
	// Default setting
	imgColorDefault          = "00FFFFFF"
	msgColorDefault          = "000000"
	imgWDefault              = 300
	imgHDefault              = 300
	fontSizeDefault          = 0
	labelTextDefault         = "default text"
	dpiDefault       float64 = 72

	fontfileDefault = "assets/fonts/CeraPro-Medium.ttf"
)

type Label struct {
	Text     string
	FontSize int
	Color    string
}

type Img struct {
	Width  int
	Height int
	Color  string
	Label  Label
}

// Do - entrypoint.
func Do(params interface{}) (*bytes.Buffer, error) {
	c, _ := params.(*dto.ImageQueryDto)
	fmt.Println(c.Width, c.Height)
	// fetch img params: imgW, imgH, text, etc
	// Just make the Text structure
	label := Label{Text: labelTextDefault, FontSize: fontSizeDefault, Color: msgColorDefault}
	// make the Image structure with the necessary fields - height, width, color, and text
	img := Img{Width: imgWDefault, Height: imgHDefault, Color: imgColorDefault, Label: label}

	// generate our image with the text
	return img.generate()
}

// generate - make the image according to the desired size, color, and text.
func (i Img) generate() (*bytes.Buffer, error) {
	// If there are dimensions and there are no requirements for the Text, we will build the default Text.
	if ((i.Width > 0 || i.Height > 0) && i.Label.Text == "") || i.Label.Text == "" {
		i.Label.Text = fmt.Sprintf("%d x %d", i.Width, i.Height)
	}

	// If there are no parameters for the font size, we will construct it based on the sizes of the image.
	if i.Label.FontSize == 0 {
		i.Label.FontSize = i.Width / 10
		if i.Height < i.Width {
			i.Label.FontSize = i.Height / 5
		}
	}

	// Convert the color from string to color.RGBA.
	//clr, err := colors.ToRGBA(i.Color)
	//if err != nil {
	//	return nil, err
	//}
	clr := color.RGBA{R: 0, G: 0, B: 0, A: 0}
	//clr.A = 0
	//clr.R = 0
	//clr.G = 0
	//clr.B = 0
	fmt.Println(clr)
	fmt.Printf("clr = %T\n", clr)
	// Create an in-memory image with the desired size.
	m := image.NewRGBA(image.Rect(0, 0, i.Width, i.Height))
	//Draw a picture:
	// - in the sizes (Bounds)
	// - with color (Uniform - wrapper above color.Color with Image functions)
	// - based on the point (Point) as the base image
	// - fill with color Uniform (draw.Src)
	draw.Draw(m, m.Bounds(), image.NewUniform(clr), image.Point{}, draw.Src)

	// add a text in the picture.
	if err := i.drawLabel(m); err != nil {
		fmt.Println(err)
		return nil, err
	}

	var im image.Image = m

	// Allocate memory for our data (the bytes of the image)
	buffer := &bytes.Buffer{}
	// Let's encode the image into our allocated memory.
	err := png.Encode(buffer, im)

	return buffer, err
}

// drawLabel - add a text in the picture
func (i *Img) drawLabel(m *image.RGBA) error {
	// Convert string text to RGBA.
	clr, err := colors.ToRGBA(i.Label.Color)
	if err != nil {
		return err
	}
	// Get the font (should work with both latin and cyrillic).
	fontBytes, err := ioutil.ReadFile(fontfileDefault)
	if err != nil {
		return err
	}
	fnt, err := truetype.Parse(fontBytes)
	if err != nil {
		return err
	}
	// Prepare a Drawer for drawing text on the image.
	d := &font.Drawer{
		Dst: m,
		Src: image.NewUniform(clr),
		Face: truetype.NewFace(fnt, &truetype.Options{
			Size:    float64(i.Label.FontSize),
			DPI:     dpiDefault,
			Hinting: font.HintingNone,
		}),
	}
	//Setting the baseline.
	d.Dot = fixed.Point26_6{
		X: (fixed.I(i.Width) - d.MeasureString(i.Label.Text)) / 2,
		Y: fixed.I((i.Height+i.Label.FontSize)/2 - 12),
	}
	// Directly rendering text to our RGBA image.
	d.DrawString(i.Label.Text)

	return nil
}
