package dto

type ImageQueryDto struct {
	Width       int    `query:"width"`
	Height      int    `query:"height"`
	Color       string `query:"color"`
	Text        string `query:"text"`
	FontSize    string `query:"font_size"`
	PaddingTop  int    `query:"padding_top"`
	PaddingLeft int    `query:"padding_left"`
}

func GetImageQueryDto() *ImageQueryDto {
	return new(ImageQueryDto)
}

func (obj *ImageQueryDto) FillDefaults() {
	if obj.Width == 0 {
		obj.Width = 300
	}
	if obj.Height == 0 {
		obj.Height = 300
	}
	if obj.Color == "" {
		obj.Color = "#000000"
	}
	if obj.Text == "" {
		obj.Text = "Default text"
	}
	if obj.PaddingTop == 0 {
		obj.PaddingTop = obj.Height / 5
	}
	if obj.PaddingLeft == 0 {
		obj.PaddingLeft = obj.Width / 10
	}
}
