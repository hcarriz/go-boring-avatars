package goboringavatars

import (
	"fmt"
	"strings"
)

const (
	beamSize = 36
)

type beamData struct {
	wrapperColor      string
	faceColor         string
	backgroundColor   string
	wrapperTranslateX float64
	wrapperTranslateY float64
	wrapperRotate     float64
	wrapperScale      float64
	isMouthOpen       bool
	isCircle          bool
	eyeSpread         float64
	mouthSpread       float64
	faceRotate        float64
	faceTranslateX    float64
	faceTranslateY    float64
}

func generateData(name string, colors []string) (beamData, error) {
	numFromName := hashCode(name)
	wrapperColor := getRandomColor(numFromName, colors)
	preTranslateX := getUnit(numFromName, 10, 1)
	wrapperTranslateX := preTranslateX
	if preTranslateX < 5 {
		wrapperTranslateX = preTranslateX + beamSize/9
	}
	preTranslateY := getUnit(numFromName, 10, 2)
	wrapperTranslateY := preTranslateY
	if preTranslateY < 5 {
		wrapperTranslateY = preTranslateY + beamSize/9
	}

	wtx := getUnit(numFromName, 8, 1)
	if wrapperTranslateX > (beamSize / 6) {
		wtx = wrapperTranslateX / 2
	}
	wty := getUnit(numFromName, 7, 2)
	if wrapperTranslateY > (beamSize / 6) {
		wty = wrapperTranslateY / 2
	}

	ct, err := getContrast(wrapperColor)
	if err != nil {
		return beamData{}, err
	}

	return beamData{
		wrapperColor:      wrapperColor,
		faceColor:         ct,
		backgroundColor:   getRandomColor(numFromName+13, colors),
		wrapperTranslateX: wrapperTranslateX,
		wrapperTranslateY: wrapperTranslateY,
		wrapperRotate:     getUnit(numFromName, 360, 0),
		wrapperScale:      1 + getUnit(numFromName, beamSize/12, 0)/10,
		isMouthOpen:       getBoolean(numFromName, 2),
		isCircle:          getBoolean(numFromName, 1),
		eyeSpread:         getUnit(numFromName, 5, 0),
		mouthSpread:       getUnit(numFromName, 3, 0),
		faceRotate:        getUnit(numFromName, 10, 3),
		faceTranslateX:    wtx,
		faceTranslateY:    wty,
	}, nil

}

func (c config) beam() (string, error) {

	data, err := generateData(c.name, c.colors)
	if err != nil {
		return "", err
	}

	svg := strings.Builder{}

	c.start(&svg, "beam", beamSize)

	rx := beamSize / 6

	if data.isCircle {
		rx = beamSize
	}

	svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="%s"></rect>`, beamSize, beamSize, data.backgroundColor))

	scale := fmt.Sprintf("%.1f", data.wrapperScale)

	scale = strings.TrimSuffix(scale, ".0")

	svg.WriteString(fmt.Sprintf(`<rect x="0" y="0" width="%d" height="%d" transform="translate(%.0f %.0f) rotate(%.0f %d %d) scale(%s)" fill="%s" rx="%d"></rect>`, beamSize, beamSize, data.wrapperTranslateX, data.wrapperTranslateY, data.wrapperRotate, beamSize/2, beamSize/2, scale, data.wrapperColor, rx))

	svg.WriteString(fmt.Sprintf(`<g transform="translate(%.0f %.0f) rotate(%.0f %d %d)">`, data.faceTranslateX, data.faceTranslateY, data.faceRotate, beamSize/2, beamSize/2))

	if data.isMouthOpen {
		svg.WriteString(fmt.Sprintf(`<path d="M15 %.0fc2 1 4 1 6 0" stroke="%s" fill="none" stroke-linecap="round"></path>`, 19+data.mouthSpread, data.faceColor))
	} else {
		svg.WriteString(fmt.Sprintf(`<path d="M13,%.0f a1,0.75 0 0,0 10,0" fill="%s"></path>`, 19+data.mouthSpread, data.faceColor))
	}

	svg.WriteString(fmt.Sprintf(`<rect x="%.0f" y="14" width="1.5" height="2" rx="1" stroke="none" fill="%s"></rect>`, 14-data.eyeSpread, data.faceColor))
	svg.WriteString(fmt.Sprintf(`<rect x="%.0f" y="14" width="1.5" height="2" rx="1" stroke="none" fill="%s"></rect>`, 20+data.eyeSpread, data.faceColor))

	svg.WriteString(`</g>`)
	c.end(&svg)

	return svg.String(), nil

}
