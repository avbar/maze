package menu

import (
	"image/color"

	"github.com/avbar/maze/internal/assets"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"
)

func (m *Menu) addButtons() {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Spacing(20),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top: 30,
			}),
		)),
	)
	m.ui.Container.AddChild(c)

	c.AddChild(newButton("OK", func() {
		m.prevSettings = m.settings
		m.close()
	}))

	c.AddChild(newButton("Cancel", func() {
		if m.settings != m.prevSettings {
			m.settings = m.prevSettings
			*m.sliderValues[sliderNameCols] = m.prevSettings.Cols
			*m.sliderValues[sliderNameRows] = m.prevSettings.Rows
		}
		m.close()
	}))
}

func newButton(text string, handler func()) *widget.Button {
	buttonRes := newButtonResources()

	return widget.NewButton(
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		}), widget.WidgetOpts.MinSize(150, 6)),
		widget.ButtonOpts.Image(buttonRes.image),
		widget.ButtonOpts.Text(text, buttonRes.face, buttonRes.color),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			handler()
		}),
	)
}

type buttonResources struct {
	image *widget.ButtonImage
	color *widget.ButtonTextColor
	face  font.Face
}

func newButtonResources() *buttonResources {
	return &buttonResources{
		image: &widget.ButtonImage{
			Idle:    image.NewNineSliceColor(color.RGBA{90, 125, 147, 255}),
			Hover:   image.NewNineSliceColor(color.RGBA{50, 85, 117, 255}),
			Pressed: image.NewNineSliceColor(color.RGBA{20, 55, 87, 255}),
		},

		color: &widget.ButtonTextColor{
			Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
		},

		face: assets.MenuFont,
	}
}
