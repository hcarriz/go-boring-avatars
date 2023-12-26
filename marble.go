package goboringavatars

import (
	"fmt"
	"strings"
)

type marbleProperties struct {
	color      string
	translateY float64
	translateX float64
	scale      float64
	rotate     float64
}

func generateMarbleColors(name string, colors []string) map[int]marbleProperties {
	numFromName := hashCode(name)

	elementsProperties := map[int]marbleProperties{}

	for i := 0; i < 3; i++ {
		element := marbleProperties{
			color:      getRandomColor(numFromName+i, colors),
			translateX: getUnit(numFromName*(i+1), svgSize/10, 1),
			translateY: getUnit(numFromName*(i+1), svgSize/10, 2),
			scale:      1.2 + float64(getUnit(numFromName*(i+1), svgSize/20, 0))/10,
			rotate:     getUnit(numFromName*(i+1), 360, 1),
		}
		elementsProperties[i] = element
	}

	return elementsProperties
}

func (a config) marble() string {

	dsize := 80

	properties := generateMarbleColors(a.name, a.colors)
	maskID := "mask__marble"

	var svg strings.Builder

	a.start(&svg, maskID, dsize)

	svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="%s"></rect>`, dsize, dsize, properties[0].color))

	svg.WriteString(fmt.Sprintf(`<path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="%s" transform="translate(%.0f %.0f) rotate(%.0f %d %d) scale(%.1f)"></path>`, properties[1].color, properties[1].translateX, properties[1].translateY, properties[1].rotate, dsize/2, dsize/2, properties[2].scale))

	svg.WriteString(fmt.Sprintf(`<path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="%s" transform="translate(%.0f %.0f) rotate(%.0f %d %d) scale(%.1f)"></path>`, properties[2].color, properties[2].translateX, properties[2].translateY, properties[2].rotate, dsize/2, dsize/2, properties[2].scale))

	a.end(&svg, `<filter id="prefix__filter0_f" filterUnits="userSpaceOnUser" colorInterpolationFilters="sRGB"><feFlood flood-opacity="0" result="BackgroundImageFix" /><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape" /><feGaussianBlur stdDeviation="7" result="effect1_foregoundBlur"/></filter>`)

	return svg.String()
}
