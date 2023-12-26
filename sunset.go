package goboringavatars

import (
	"fmt"
	"strings"
)

const (
	sunsetElements = 4
	sunsetSize     = 80
)

func genSunsetColors(name string, colors []string) map[int]string {

	n := hashCode(name)

	list := make(map[int]string, sunsetElements)

	for i := 0; i < sunsetElements; i++ {
		list[i] = getRandomColor(n+i, colors)
	}

	return list

}

func (c config) sunset() string {

	colors := genSunsetColors(c.name, c.colors)

	svg := strings.Builder{}

	name := strings.ReplaceAll(c.name, " ", "")

	c.start(&svg, "ring", sunsetSize)

	svg.WriteString(fmt.Sprintf(`<path fill="url(#gradient_paint0_linear_%s)" d="M0 0h80v40H0z"></path><path fill="url(#gradient_paint1_linear_%s)" d="M0 40h80v40H0z"></path>`, name, name))

	c.end(&svg,
		fmt.Sprintf(`<linearGradient id="gradient_paint0_linear_%s" x1="%d" y1="0" x2="%d" y2="%d" gradientUnits="userSpaceOnUse"><stop stop-color="%s" /><stop offset="1" stop-color="%s" /></linearGradient>`, name, sunsetSize/2, sunsetSize/2, sunsetSize/2, colors[0], colors[1]),
		fmt.Sprintf(`<linearGradient id="gradient_paint1_linear_%s" x1="%d" y1="%d" x2="%d" y2="%d" gradientUnits="userSpaceOnUse"><stop stop-color="%s" /><stop offset="1" stop-color="%s" /></linearGradient>`, name, sunsetSize/2, sunsetSize/2, sunsetSize/2, sunsetSize, colors[2], colors[3]),
	)

	return svg.String()

}
