package goboringavatars

import (
	"fmt"
	"strings"
)

const (
	ringSize   = 90
	ringColors = 5
)

func generateRingColors(name string, colors []string) map[int]string {

	numFromName := hashCode(name)

	shuffle := make(map[int]string)

	for i := 0; i < ringColors; i++ {
		shuffle[i] = getRandomColor(numFromName+i, colors)
	}

	return map[int]string{
		0: shuffle[0],
		1: shuffle[1],
		2: shuffle[1],
		3: shuffle[2],
		4: shuffle[2],
		5: shuffle[3],
		6: shuffle[3],
		7: shuffle[0],
		8: shuffle[4],
	}

}

func (c config) ring() string {

	colors := generateRingColors(c.name, c.colors)

	svg := strings.Builder{}

	c.start(&svg, "ring", ringSize)

	svg.WriteString(fmt.Sprintf(`<path d="M0 0h90v45H0z" fill="%s"></path>`, colors[0]))
	svg.WriteString(fmt.Sprintf(`<path d="M0 45h90v45H0z" fill="%s"></path>`, colors[1]))
	svg.WriteString(fmt.Sprintf(`<path d="M83 45a38 38 0 00-76 0h76z" fill="%s"></path>`, colors[2]))
	svg.WriteString(fmt.Sprintf(`<path d="M83 45a38 38 0 01-76 0h76z" fill="%s"></path>`, colors[3]))
	svg.WriteString(fmt.Sprintf(`<path d="M77 45a32 32 0 10-64 0h64z" fill="%s"></path>`, colors[4]))
	svg.WriteString(fmt.Sprintf(`<path d="M77 45a32 32 0 11-64 0h64z" fill="%s"></path>`, colors[5]))
	svg.WriteString(fmt.Sprintf(`<path d="M71 45a26 26 0 00-52 0h52z" fill="%s"></path>`, colors[6]))
	svg.WriteString(fmt.Sprintf(`<path d="M71 45a26 26 0 01-52 0h52z" fill="%s"></path>`, colors[7]))
	svg.WriteString(fmt.Sprintf(`<circle cx="45" cy="45" r="23" fill="%s"></circle>`, colors[8]))

	c.end(&svg)

	return svg.String()

}
