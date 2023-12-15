package menu

import (
	"fmt"
	"image/color"

	"github.com/avbar/maze/internal/assets"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"
)

func (m *Menu) addSlider(name string, value *int) {
	sliderRes := newSliderResources()
	labelRes := newLabelResources()

	// Container for slider name
	c := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewRowLayout()))
	m.ui.Container.AddChild(c)

	// Slider name
	labelText := widget.NewLabel(
		widget.LabelOpts.TextOpts(widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
		}))),
		widget.LabelOpts.Text(name, labelRes.face, labelRes.color),
	)
	c.AddChild(labelText)

	// Container for slider
	c = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Spacing(20),
		)),
	)
	m.ui.Container.AddChild(c)

	var labelValue *widget.Label

	// Slider
	slider := widget.NewSlider(
		widget.SliderOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
		}), widget.WidgetOpts.MinSize(300, 6)),
		widget.SliderOpts.MinMax(5, 50),
		widget.SliderOpts.Images(sliderRes.trackImage, sliderRes.handle),
		widget.SliderOpts.FixedHandleSize(sliderRes.handleSize),
		widget.SliderOpts.TrackOffset(5),
		widget.SliderOpts.PageSizeFunc(func() int {
			return 1
		}),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			*value = args.Current
			labelValue.Label = fmt.Sprintf("%d", args.Current)
		}),
	)
	slider.Current = *value
	m.sliderValues[name] = &slider.Current
	c.AddChild(slider)

	// Slider value
	labelValue = widget.NewLabel(
		widget.LabelOpts.TextOpts(widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
		}))),
		widget.LabelOpts.Text(fmt.Sprintf("%d", slider.Current), labelRes.face, labelRes.color),
	)
	c.AddChild(labelValue)
}

type sliderResources struct {
	trackImage *widget.SliderTrackImage
	handle     *widget.ButtonImage
	handleSize int
}

func newSliderResources() *sliderResources {
	return &sliderResources{
		trackImage: &widget.SliderTrackImage{
			Idle:  image.NewNineSlice(assets.SliderTrackIdle, [3]int{0, 19, 0}, [3]int{6, 0, 0}),
			Hover: image.NewNineSlice(assets.SliderTrackIdle, [3]int{0, 19, 0}, [3]int{6, 0, 0}),
		},

		handle: &widget.ButtonImage{
			Idle:    image.NewNineSliceSimple(assets.SliderHandleIdle, 0, 5),
			Hover:   image.NewNineSliceSimple(assets.SliderHandleHover, 0, 5),
			Pressed: image.NewNineSliceSimple(assets.SliderHandleHover, 0, 5),
		},

		handleSize: 5,
	}
}

type labelResources struct {
	color *widget.LabelColor
	face  font.Face
}

func newLabelResources() *labelResources {
	return &labelResources{
		color: &widget.LabelColor{
			Idle: color.White,
		},

		face: assets.MenuFont,
	}
}
