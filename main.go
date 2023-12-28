package goboringavatars

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Consts

const (
	svgSize = 80
)

// Custom errors
var (
	ErrNegativePixels = errors.New("pixels can not be negative")
	ErrInvalidVariant = errors.New("invalid variant")
	ErrEmptyName      = errors.New("name is empty")
	defaultColors     = []string{"#0A0310", "#49007E", "#FF005B", "#FF7D10", "#FFB238"}
)

// Name limits the styles used.
type Name struct {
	name string
}

func (v Name) String() string {
	return v.name
}

func ValidateName(variant Name) bool {
	switch variant.String() {
	case "", "pixel", "bauhaus", "ring", "sunset", "beam":
		return true
	default:
		return false
	}
}

var (
	Bauhaus = Name{"bauhaus"}
	Beam    = Name{"beam"}
	Marble  = Name{""} // Marble is the default Name.
	Pixel   = Name{"pixel"}
	Ring    = Name{"ring"}
	Sunset  = Name{"sunset"}
)

// Config
type config struct {
	size    int
	square  bool
	title   bool
	name    string
	variant Name
	colors  []string
	classes []string
}

type option func(*config) error

func (o option) apply(c *config) error {
	return o(c)
}

// Options
type Option interface {
	apply(*config) error
}

// Size sets the size of the Avatar in pixels.
func Size(pixels int) Option {
	return option(func(c *config) error {

		if pixels < 0 {
			return ErrNegativePixels
		}

		c.size = pixels

		return nil
	})
}

// Square makes the Avatar square.
func Square() Option {
	return option(func(c *config) error {
		c.square = true
		return nil
	})
}

// Title add a title element with the name to the SVG output.
func Title() Option {
	return option(func(c *config) error {
		c.title = true
		return nil
	})
}

// Variant sets the specific variant to be used for the Avatar.
func Variant(variant Name) Option {
	return option(func(c *config) error {

		if !ValidateName(variant) {
			return ErrInvalidVariant
		}

		c.variant = variant

		return nil
	})
}

// Colors sets the five (5) colors that will be used to generate the Avatar.
func Colors(one, two, three, four, five string) Option {
	return option(func(c *config) error {

		c.colors = []string{one, two, three, four, five}

		return nil
	})
}

// Classes adds classes to the svg.
func Classes(list ...string) Option {
	return option(func(c *config) error {
		c.classes = append(c.classes, list...)
		return nil
	})
}

// New generates an avatar for the given name.
func New(name string, opts ...Option) (string, error) {

	if name == "" {
		return "", ErrEmptyName
	}

	var (
		c = config{
			name:   name,
			size:   40,
			colors: defaultColors,
		}
		err error
	)

	for _, opt := range opts {
		err = errors.Join(err, opt.apply(&c))
	}

	if err != nil {
		return "", err
	}

	switch c.variant {
	case Beam:
		return c.beam()
	case Sunset:
		return c.sunset(), nil
	case Ring:
		return c.ring(), nil
	case Bauhaus:
		return c.bauhaus(), nil
	case Pixel:
		return c.pixel(), nil
	default:
		return c.marble(), nil
	}

}

type Avatar string

// Render lets the Avatar be rendered in templ.
// https://github.com/a-h/templ
func (a Avatar) Render(_ context.Context, w io.Writer) error {
	r := string(a)

	_, err := io.WriteString(w, r)

	return err
}

// String reverts the Avatar back into a string.
func (a Avatar) String() string {
	return string(a)
}

func Render(name string, opts ...Option) Avatar {

	result, err := New(name, opts...)
	if err != nil {
		return Avatar("")
	}

	return Avatar(result)

}

func (a config) start(svg *strings.Builder, maskID string, size int) {

	svg.WriteString(fmt.Sprintf(`<svg viewBox="0 0 %d %d" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="%d" height="%d"`, size, size, a.size, a.size))

	if len(a.classes) > 0 {
		svg.WriteString(fmt.Sprintf(` class="%s"`, strings.Join(a.classes, " ")))
	}

	svg.WriteString(`>`)

	// Add the title
	if a.title {
		svg.WriteString(fmt.Sprintf(`<title>%s</title>`, a.name))
	}

	// Add the mask
	svg.WriteString(fmt.Sprintf(`<mask id="%s" maskUnits="userSpaceOnUse" x="0" y="0" width="%d" height="%d"><rect width="%d" height="%d" `, maskID, size, size, size, size))

	if !a.square {
		svg.WriteString(fmt.Sprintf(`rx="%d" `, size*2))
	}

	svg.WriteString(fmt.Sprintf(`fill="#FFFFFF"></rect></mask><g mask="url(#%s)">`, maskID))

}

func (a config) end(svg *strings.Builder, filters ...string) {

	svg.WriteString(`</g>`)

	if len(filters) > 0 {

		svg.WriteString(`<defs>`)

		for _, line := range filters {
			svg.WriteString(line)
		}

		svg.WriteString(`</defs>`)

	}

	svg.WriteString(`</svg>`)

}
