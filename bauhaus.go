package goboringavatars

import (
	"fmt"
	"strings"
)

const (
	bauhausElement = 4
)

type bauhausProps struct {
	color      string
	translateY float64
	translateX float64
	rotate     float64
	isSquare   bool
}

func generateBauhausColors(name string, colors []string) map[int]bauhausProps {

	p := map[int]bauhausProps{}

	n := hashCode(name)

	for i := 0; i < bauhausElement; i++ {

		r := bauhausProps{
			color:      getRandomColor(n+i, colors),
			translateX: getUnit(n*(i+1), svgSize/2-(i+17), 1),
			translateY: getUnit(n*(i+1), svgSize/2-(i+17), 2),
			rotate:     getUnit(n*(i+1), 360, 0),
			isSquare:   getBoolean(n, 2),
		}

		p[i] = r
	}

	return p

}

func (c config) bauhaus() string {

	var svg strings.Builder

	props := generateBauhausColors(c.name, c.colors)

	c.start(&svg, "avatar_bauhaus", svgSize)

	sq := svgSize / 8

	if props[0].isSquare {
		sq = svgSize
	}

	svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="%s"></rect>`, svgSize, svgSize, props[0].color))
	svg.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" fill="%s" transform="translate(%.0f %.0f) rotate(%.0f %d %d)"></rect>`,
		(svgSize-60)/2, (svgSize-20)/2, svgSize, sq, props[1].color, props[1].translateX, props[1].translateY, props[1].rotate, svgSize/2, svgSize/2))
	svg.WriteString(fmt.Sprintf(`<circle cx="%d" cy="%d" fill="%s" r="%d" transform="translate(%.0f %.0f)"></circle>`, svgSize/2, svgSize/2, props[2].color, svgSize/5, props[2].translateX, props[2].translateY))
	svg.WriteString(fmt.Sprintf(`<line x1="0" y1="%d" x2="%d" y2="%d" stroke-width="2" stroke="%s" transform="translate(%.0f %.0f) rotate(%.0f %d %d)"></line>`, svgSize/2, svgSize, svgSize/2, props[3].color, props[3].translateX, props[3].translateY, props[3].rotate, svgSize/2, svgSize/2))

	c.end(&svg)

	return svg.String()

}
