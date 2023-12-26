package goboringavatars

import (
	"fmt"
	"strings"
)

// generatePixelColors creates a list of colors based on the name and a color palette
func generatePixelColors(name string, colors []string) map[int]string {
	numFromName := hashCode(name)

	colorList := make(map[int]string)

	for i := 0; i < 64; i++ {
		color := getRandomColor(numFromName%(i+1), colors)
		colorList[i] = color
	}

	return colorList
}

// pixel generates an SVG string representing a pixelated avatar
func (a config) pixel() string {

	maskID := "avatar__pixel"
	dsize := 80

	var svg strings.Builder
	a.start(&svg, maskID, dsize)

	pixelColors := generatePixelColors(a.name, a.colors)

	// Generate the SVG rect elements
	// for i := 0; i < 64; i++ {
	// 	x := (i % 8) * 10
	// 	y := (i / 8) * 10
	// 	svg.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="10" height="10" fill="%s" />`, x, y, pixelColors[i]))
	// }

	svg.WriteString(fmt.Sprintf(`<rect width="10" height="10" fill="%s"></rect>`, pixelColors[0]))
	svg.WriteString(fmt.Sprintf(`<rect x="20" width="10" height="10" fill="%s"></rect>`, pixelColors[1]))
	svg.WriteString(fmt.Sprintf(`<rect x="40" width="10" height="10" fill="%s"></rect>`, pixelColors[2]))
	svg.WriteString(fmt.Sprintf(`<rect x="60" width="10" height="10" fill="%s"></rect>`, pixelColors[3]))
	svg.WriteString(fmt.Sprintf(`<rect x="10" width="10" height="10" fill="%s"></rect>`, pixelColors[4]))
	svg.WriteString(fmt.Sprintf(`<rect x="30" width="10" height="10" fill="%s"></rect>`, pixelColors[5]))
	svg.WriteString(fmt.Sprintf(`<rect x="50" width="10" height="10" fill="%s"></rect>`, pixelColors[6]))
	svg.WriteString(fmt.Sprintf(`<rect x="70" width="10" height="10" fill="%s"></rect>`, pixelColors[7]))
	svg.WriteString(fmt.Sprintf(`<rect y="10" width="10" height="10" fill="%s"></rect>`, pixelColors[8]))
	svg.WriteString(fmt.Sprintf(`<rect y="20" width="10" height="10" fill="%s"></rect>`, pixelColors[9]))
	svg.WriteString(fmt.Sprintf(`<rect y="30" width="10" height="10" fill="%s"></rect>`, pixelColors[10]))
	svg.WriteString(fmt.Sprintf(`<rect y="40" width="10" height="10" fill="%s"></rect>`, pixelColors[11]))
	svg.WriteString(fmt.Sprintf(`<rect y="50" width="10" height="10" fill="%s"></rect>`, pixelColors[12]))
	svg.WriteString(fmt.Sprintf(`<rect y="60" width="10" height="10" fill="%s"></rect>`, pixelColors[13]))
	svg.WriteString(fmt.Sprintf(`<rect y="70" width="10" height="10" fill="%s"></rect>`, pixelColors[14]))
	svg.WriteString(fmt.Sprintf(`<rect x="20" y="10" width="10" height="10" fill="%s"></rect>`, pixelColors[15]))
	svg.WriteString(fmt.Sprintf(`<rect x="20" y="20" width="10" height="10" fill="%s"></rect>`, pixelColors[16]))
	svg.WriteString(fmt.Sprintf(`<rect x="20" y="30" width="10" height="10" fill="%s"></rect>`, pixelColors[17]))
	svg.WriteString(fmt.Sprintf(`<rect x="20" y="40" width="10" height="10" fill="%s"></rect>`, pixelColors[18]))
	svg.WriteString(fmt.Sprintf(`<rect x="20" y="50" width="10" height="10" fill="%s"></rect>`, pixelColors[19]))
	svg.WriteString(fmt.Sprintf(`<rect x="20" y="60" width="10" height="10" fill="%s"></rect>`, pixelColors[20]))
	svg.WriteString(fmt.Sprintf(`<rect x="20" y="70" width="10" height="10" fill="%s"></rect>`, pixelColors[21]))
	svg.WriteString(fmt.Sprintf(`<rect x="40" y="10" width="10" height="10" fill="%s"></rect>`, pixelColors[22]))
	svg.WriteString(fmt.Sprintf(`<rect x="40" y="20" width="10" height="10" fill="%s"></rect>`, pixelColors[23]))
	svg.WriteString(fmt.Sprintf(`<rect x="40" y="30" width="10" height="10" fill="%s"></rect>`, pixelColors[24]))
	svg.WriteString(fmt.Sprintf(`<rect x="40" y="40" width="10" height="10" fill="%s"></rect>`, pixelColors[25]))
	svg.WriteString(fmt.Sprintf(`<rect x="40" y="50" width="10" height="10" fill="%s"></rect>`, pixelColors[26]))
	svg.WriteString(fmt.Sprintf(`<rect x="40" y="60" width="10" height="10" fill="%s"></rect>`, pixelColors[27]))
	svg.WriteString(fmt.Sprintf(`<rect x="40" y="70" width="10" height="10" fill="%s"></rect>`, pixelColors[28]))
	svg.WriteString(fmt.Sprintf(`<rect x="60" y="10" width="10" height="10" fill="%s"></rect>`, pixelColors[29]))
	svg.WriteString(fmt.Sprintf(`<rect x="60" y="20" width="10" height="10" fill="%s"></rect>`, pixelColors[30]))
	svg.WriteString(fmt.Sprintf(`<rect x="60" y="30" width="10" height="10" fill="%s"></rect>`, pixelColors[31]))
	svg.WriteString(fmt.Sprintf(`<rect x="60" y="40" width="10" height="10" fill="%s"></rect>`, pixelColors[32]))
	svg.WriteString(fmt.Sprintf(`<rect x="60" y="50" width="10" height="10" fill="%s"></rect>`, pixelColors[33]))
	svg.WriteString(fmt.Sprintf(`<rect x="60" y="60" width="10" height="10" fill="%s"></rect>`, pixelColors[34]))
	svg.WriteString(fmt.Sprintf(`<rect x="60" y="70" width="10" height="10" fill="%s"></rect>`, pixelColors[35]))
	svg.WriteString(fmt.Sprintf(`<rect x="10" y="10" width="10" height="10" fill="%s"></rect>`, pixelColors[36]))
	svg.WriteString(fmt.Sprintf(`<rect x="10" y="20" width="10" height="10" fill="%s"></rect>`, pixelColors[37]))
	svg.WriteString(fmt.Sprintf(`<rect x="10" y="30" width="10" height="10" fill="%s"></rect>`, pixelColors[38]))
	svg.WriteString(fmt.Sprintf(`<rect x="10" y="40" width="10" height="10" fill="%s"></rect>`, pixelColors[39]))
	svg.WriteString(fmt.Sprintf(`<rect x="10" y="50" width="10" height="10" fill="%s"></rect>`, pixelColors[40]))
	svg.WriteString(fmt.Sprintf(`<rect x="10" y="60" width="10" height="10" fill="%s"></rect>`, pixelColors[41]))
	svg.WriteString(fmt.Sprintf(`<rect x="10" y="70" width="10" height="10" fill="%s"></rect>`, pixelColors[42]))
	svg.WriteString(fmt.Sprintf(`<rect x="30" y="10" width="10" height="10" fill="%s"></rect>`, pixelColors[43]))
	svg.WriteString(fmt.Sprintf(`<rect x="30" y="20" width="10" height="10" fill="%s"></rect>`, pixelColors[44]))
	svg.WriteString(fmt.Sprintf(`<rect x="30" y="30" width="10" height="10" fill="%s"></rect>`, pixelColors[45]))
	svg.WriteString(fmt.Sprintf(`<rect x="30" y="40" width="10" height="10" fill="%s"></rect>`, pixelColors[46]))
	svg.WriteString(fmt.Sprintf(`<rect x="30" y="50" width="10" height="10" fill="%s"></rect>`, pixelColors[47]))
	svg.WriteString(fmt.Sprintf(`<rect x="30" y="60" width="10" height="10" fill="%s"></rect>`, pixelColors[48]))
	svg.WriteString(fmt.Sprintf(`<rect x="30" y="70" width="10" height="10" fill="%s"></rect>`, pixelColors[49]))
	svg.WriteString(fmt.Sprintf(`<rect x="50" y="10" width="10" height="10" fill="%s"></rect>`, pixelColors[50]))
	svg.WriteString(fmt.Sprintf(`<rect x="50" y="20" width="10" height="10" fill="%s"></rect>`, pixelColors[51]))
	svg.WriteString(fmt.Sprintf(`<rect x="50" y="30" width="10" height="10" fill="%s"></rect>`, pixelColors[52]))
	svg.WriteString(fmt.Sprintf(`<rect x="50" y="40" width="10" height="10" fill="%s"></rect>`, pixelColors[53]))
	svg.WriteString(fmt.Sprintf(`<rect x="50" y="50" width="10" height="10" fill="%s"></rect>`, pixelColors[54]))
	svg.WriteString(fmt.Sprintf(`<rect x="50" y="60" width="10" height="10" fill="%s"></rect>`, pixelColors[55]))
	svg.WriteString(fmt.Sprintf(`<rect x="50" y="70" width="10" height="10" fill="%s"></rect>`, pixelColors[56]))
	svg.WriteString(fmt.Sprintf(`<rect x="70" y="10" width="10" height="10" fill="%s"></rect>`, pixelColors[57]))
	svg.WriteString(fmt.Sprintf(`<rect x="70" y="20" width="10" height="10" fill="%s"></rect>`, pixelColors[58]))
	svg.WriteString(fmt.Sprintf(`<rect x="70" y="30" width="10" height="10" fill="%s"></rect>`, pixelColors[59]))
	svg.WriteString(fmt.Sprintf(`<rect x="70" y="40" width="10" height="10" fill="%s"></rect>`, pixelColors[60]))
	svg.WriteString(fmt.Sprintf(`<rect x="70" y="50" width="10" height="10" fill="%s"></rect>`, pixelColors[61]))
	svg.WriteString(fmt.Sprintf(`<rect x="70" y="60" width="10" height="10" fill="%s"></rect>`, pixelColors[62]))
	svg.WriteString(fmt.Sprintf(`<rect x="70" y="70" width="10" height="10" fill="%s"></rect>`, pixelColors[63]))

	a.end(&svg)

	return svg.String()
}
