package colorimage

import (
	"image"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/disintegration/imaging"
	"github.com/lucasb-eyer/go-colorful"
)

// Model is a struct that contains all the properties of colorimage.
type Model struct {
	Image   image.Image
	Content string
	Width   int
}

// ImageToString converts an image to a string representation of an image.
func ImageToString(width int, img image.Image) string {
	img = imaging.Resize(img, width, 0, imaging.Lanczos)
	b := img.Bounds()
	w := b.Max.X
	h := b.Max.Y
	str := strings.Builder{}

	for y := 0; y < h; y += 2 {
		for x := w; x < width; x += 2 {
			str.WriteString(" ")
		}

		for x := 0; x < w; x++ {
			c1, _ := colorful.MakeColor(img.At(x, y))
			color1 := lipgloss.Color(c1.Hex())
			c2, _ := colorful.MakeColor(img.At(x, y+1))
			color2 := lipgloss.Color(c2.Hex())
			str.WriteString(lipgloss.NewStyle().Foreground(color1).
				Background(color2).Render("▀"))
		}

		str.WriteString("\n")
	}

	return str.String()
}

// SetSize sets the size of the colorimage.
func (m *Model) SetSize(width int) {
	m.Width = width
}

// GetImage returns the currently set image.
func (m Model) GetImage() image.Image {
	return m.Image
}

// SetContent sets the content of the colorimage.
func (m *Model) SetContent(content string) {
	m.Content = content
}

// SetImage sets the image of the colorimage.
func (m *Model) SetImage(img image.Image) {
	m.Image = img
}

// View returns a string representation of a colorimage.
func (m Model) View() string {
	return lipgloss.NewStyle().Width(m.Width).Render(m.Content)
}
