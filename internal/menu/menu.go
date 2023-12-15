package menu

import (
	"github.com/avbar/maze/internal/common"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	sliderNameCols = "Columns"
	sliderNameRows = "Rows"
)

type Menu struct {
	settings     common.Settings
	prevSettings common.Settings
	ui           *ebitenui.UI
	sliderValues map[string]*int
	close        func()
}

func NewMenu(settings common.Settings, close func()) *Menu {
	c := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    20,
				Left:   20,
				Right:  20,
				Bottom: 20,
			}),
		)),
	)

	ui := &ebitenui.UI{
		Container: c,
	}

	m := &Menu{
		settings:     settings,
		prevSettings: settings,
		ui:           ui,
		sliderValues: make(map[string]*int, 2),
		close:        close,
	}

	m.addSlider(sliderNameCols, &m.settings.Cols)
	m.addSlider(sliderNameRows, &m.settings.Rows)
	m.addButtons()

	return m
}

func (m *Menu) Settings() common.Settings {
	return m.settings
}

func (m *Menu) isCancelKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEscape)
}

func (m *Menu) Update() {
	if m.isCancelKeyPressed() {
		if m.settings != m.prevSettings {
			m.settings = m.prevSettings
			*m.sliderValues[sliderNameCols] = m.prevSettings.Cols
			*m.sliderValues[sliderNameRows] = m.prevSettings.Rows
		}
		m.close()
		return
	}

	m.ui.Update()
}

func (m *Menu) Draw(screen *ebiten.Image) {
	m.ui.Draw(screen)
}
