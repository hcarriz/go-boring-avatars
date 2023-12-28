package goboringavatars

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strings"
	"testing"
)

// used for testing
func getInternals(in string) (string, error) {
	list := regexp.MustCompile(`(?s)<g[^>]*>(.*?)</g>`).FindString(in)

	type G struct {
		Content string `xml:",innerxml"`
	}

	var g G
	if err := xml.Unmarshal([]byte(list), &g); err != nil {
		if err := xml.Unmarshal([]byte(fmt.Sprintf("%s</g>", list)), &g); err != nil {
			return "", err
		}
	}

	return g.Content, nil
}

func Test_getInternals(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "",
			in:   `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="128" height="128"><mask id=":r3ch:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r3ch:)"><rect width="80" height="80" fill="#F5B349"></rect><path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="#E8D5B9" transform="translate(4 -4) rotate(36 40 40) scale(1.4)"></path><path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="#E8D5B7" transform="translate(2 2) rotate(234 40 40) scale(1.4)"></path></g><defs><filter id="prefix__filter0_f" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood flood-opacity="0" result="BackgroundImageFix"></feFlood><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"></feBlend><feGaussianBlur stdDeviation="7" result="effect1_foregroundBlur"></feGaussianBlur></filter></defs></svg>`,
			want: `<rect width="80" height="80" fill="#F5B349"></rect><path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="#E8D5B9" transform="translate(4 -4) rotate(36 40 40) scale(1.4)"></path><path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="#E8D5B7" transform="translate(2 2) rotate(234 40 40) scale(1.4)"></path>`,
		},
		{
			name: "",
			in:   `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r3a:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r3a:)"><rect width="80" height="80" fill="#0A0310"></rect><path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="#49007E" transform="translate(-0 -0) rotate(-320 40 40) scale(1.2)"></path><path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="#FF005B" transform="translate(-4 -4) rotate(-300 40 40) scale(1.2)"></path></g><defs><filter id="prefix__filter0_f" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood flood-opacity="0" result="BackgroundImageFix"></feFlood><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"></feBlend><feGaussianBlur stdDeviation="7" result="effect1_foregroundBlur"></feGaussianBlur></filter></defs></svg>`,
			want: `<rect width="80" height="80" fill="#0A0310"></rect><path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="#49007E" transform="translate(-0 -0) rotate(-320 40 40) scale(1.2)"></path><path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="#FF005B" transform="translate(-4 -4) rotate(-300 40 40) scale(1.2)"></path>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getInternals(tt.in)
			if err != nil {
				t.Errorf("recieved error: %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("getInternals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))

	tests := []struct {
		name    string
		args    []Option
		want    string
		wantErr bool
	}{
		// Marble | Default
		{
			name: "Mary Baker",
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r3a:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r3a:)"><rect width="80" height="80" fill="#0A0310"></rect><path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="#49007E" transform="translate(-0 -0) rotate(-320 40 40) scale(1.2)"></path><path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="#FF005B" transform="translate(-4 -4) rotate(-300 40 40) scale(1.2)"></path></g><defs><filter id="prefix__filter0_f" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood flood-opacity="0" result="BackgroundImageFix"></feFlood><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"></feBlend><feGaussianBlur stdDeviation="7" result="effect1_foregroundBlur"></feGaussianBlur></filter></defs></svg>`,
		},
		{
			name: "Amelia Earhart",
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r3b:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r3b:)"><rect width="80" height="80" fill="#FFB238"></rect><path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="#0A0310" transform="translate(-0 -0) rotate(-288 40 40) scale(1.2)"></path><path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="#49007E" transform="translate(4 -4) rotate(252 40 40) scale(1.2)"></path></g><defs><filter id="prefix__filter0_f" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood flood-opacity="0" result="BackgroundImageFix"></feFlood><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"></feBlend><feGaussianBlur stdDeviation="7" result="effect1_foregroundBlur"></feGaussianBlur></filter></defs></svg>`,
		},
		{
			name: "Mary Roebling",
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r3c:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r3c:)"><rect width="80" height="80" fill="#0A0310"></rect><path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="#49007E" transform="translate(6 6) rotate(310 40 40) scale(1.3)"></path><path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="#FF005B" transform="translate(-1 -1) rotate(-105 40 40) scale(1.3)"></path></g><defs><filter id="prefix__filter0_f" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood flood-opacity="0" result="BackgroundImageFix"></feFlood><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"></feBlend><feGaussianBlur stdDeviation="7" result="effect1_foregroundBlur"></feGaussianBlur></filter></defs></svg>`,
		},
		{
			name: "Sarah Winnemucca",
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r3d:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r3d:)"><rect width="80" height="80" fill="#49007E"></rect><path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="#FF005B" transform="translate(-2 -2) rotate(-242 40 40) scale(1.5)"></path><path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="#FF7D10" transform="translate(-7 7) rotate(-183 40 40) scale(1.5)"></path></g><defs><filter id="prefix__filter0_f" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood flood-opacity="0" result="BackgroundImageFix"></feFlood><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"></feBlend><feGaussianBlur stdDeviation="7" result="effect1_foregroundBlur"></feGaussianBlur></filter></defs></svg>`,
		},
		{
			name: "Margaret Brent",
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r3e:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r3e:)"><rect width="80" height="80" fill="#49007E"></rect><path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="#FF005B" transform="translate(0 0) rotate(352 40 40) scale(1.2)"></path><path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="#FF7D10" transform="translate(-4 -4) rotate(-348 40 40) scale(1.2)"></path></g><defs><filter id="prefix__filter0_f" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood flood-opacity="0" result="BackgroundImageFix"></feFlood><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"></feBlend><feGaussianBlur stdDeviation="7" result="effect1_foregroundBlur"></feGaussianBlur></filter></defs></svg>`,
		},
		// Pixel
		{
			name: "Mary Baker",
			args: []Option{Variant(Pixel)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r6d:" mask-type="alpha" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r6d:)"><rect width="10" height="10" fill="#0A0310"></rect><rect x="20" width="10" height="10" fill="#0A0310"></rect><rect x="40" width="10" height="10" fill="#49007E"></rect><rect x="60" width="10" height="10" fill="#0A0310"></rect><rect x="10" width="10" height="10" fill="#0A0310"></rect><rect x="30" width="10" height="10" fill="#FFB238"></rect><rect x="50" width="10" height="10" fill="#49007E"></rect><rect x="70" width="10" height="10" fill="#FFB238"></rect><rect y="10" width="10" height="10" fill="#FF005B"></rect><rect y="20" width="10" height="10" fill="#0A0310"></rect><rect y="30" width="10" height="10" fill="#FFB238"></rect><rect y="40" width="10" height="10" fill="#FFB238"></rect><rect y="50" width="10" height="10" fill="#0A0310"></rect><rect y="60" width="10" height="10" fill="#FF7D10"></rect><rect y="70" width="10" height="10" fill="#0A0310"></rect><rect x="20" y="10" width="10" height="10" fill="#FFB238"></rect><rect x="20" y="20" width="10" height="10" fill="#49007E"></rect><rect x="20" y="30" width="10" height="10" fill="#49007E"></rect><rect x="20" y="40" width="10" height="10" fill="#FF7D10"></rect><rect x="20" y="50" width="10" height="10" fill="#0A0310"></rect><rect x="20" y="60" width="10" height="10" fill="#49007E"></rect><rect x="20" y="70" width="10" height="10" fill="#FFB238"></rect><rect x="40" y="10" width="10" height="10" fill="#FF005B"></rect><rect x="40" y="20" width="10" height="10" fill="#FFB238"></rect><rect x="40" y="30" width="10" height="10" fill="#0A0310"></rect><rect x="40" y="40" width="10" height="10" fill="#FF7D10"></rect><rect x="40" y="50" width="10" height="10" fill="#FF005B"></rect><rect x="40" y="60" width="10" height="10" fill="#FF7D10"></rect><rect x="40" y="70" width="10" height="10" fill="#0A0310"></rect><rect x="60" y="10" width="10" height="10" fill="#0A0310"></rect><rect x="60" y="20" width="10" height="10" fill="#FF005B"></rect><rect x="60" y="30" width="10" height="10" fill="#0A0310"></rect><rect x="60" y="40" width="10" height="10" fill="#FFB238"></rect><rect x="60" y="50" width="10" height="10" fill="#FF7D10"></rect><rect x="60" y="60" width="10" height="10" fill="#0A0310"></rect><rect x="60" y="70" width="10" height="10" fill="#49007E"></rect><rect x="10" y="10" width="10" height="10" fill="#FFB238"></rect><rect x="10" y="20" width="10" height="10" fill="#FF005B"></rect><rect x="10" y="30" width="10" height="10" fill="#49007E"></rect><rect x="10" y="40" width="10" height="10" fill="#0A0310"></rect><rect x="10" y="50" width="10" height="10" fill="#FF005B"></rect><rect x="10" y="60" width="10" height="10" fill="#FF005B"></rect><rect x="10" y="70" width="10" height="10" fill="#FFB238"></rect><rect x="30" y="10" width="10" height="10" fill="#FFB238"></rect><rect x="30" y="20" width="10" height="10" fill="#0A0310"></rect><rect x="30" y="30" width="10" height="10" fill="#0A0310"></rect><rect x="30" y="40" width="10" height="10" fill="#FFB238"></rect><rect x="30" y="50" width="10" height="10" fill="#FFB238"></rect><rect x="30" y="60" width="10" height="10" fill="#FF005B"></rect><rect x="30" y="70" width="10" height="10" fill="#0A0310"></rect><rect x="50" y="10" width="10" height="10" fill="#49007E"></rect><rect x="50" y="20" width="10" height="10" fill="#FFB238"></rect><rect x="50" y="30" width="10" height="10" fill="#49007E"></rect><rect x="50" y="40" width="10" height="10" fill="#FFB238"></rect><rect x="50" y="50" width="10" height="10" fill="#0A0310"></rect><rect x="50" y="60" width="10" height="10" fill="#49007E"></rect><rect x="50" y="70" width="10" height="10" fill="#FF7D10"></rect><rect x="70" y="10" width="10" height="10" fill="#0A0310"></rect><rect x="70" y="20" width="10" height="10" fill="#0A0310"></rect><rect x="70" y="30" width="10" height="10" fill="#0A0310"></rect><rect x="70" y="40" width="10" height="10" fill="#49007E"></rect><rect x="70" y="50" width="10" height="10" fill="#FF005B"></rect><rect x="70" y="60" width="10" height="10" fill="#FF7D10"></rect><rect x="70" y="70" width="10" height="10" fill="#FF005B"></rect></g></svg>`,
		},
		{
			name: "Amelia Earhart",
			args: []Option{Variant(Pixel)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r6e:" mask-type="alpha" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r6e:)"><rect width="10" height="10" fill="#0A0310"></rect><rect x="20" width="10" height="10" fill="#0A0310"></rect><rect x="40" width="10" height="10" fill="#0A0310"></rect><rect x="60" width="10" height="10" fill="#0A0310"></rect><rect x="10" width="10" height="10" fill="#FFB238"></rect><rect x="30" width="10" height="10" fill="#0A0310"></rect><rect x="50" width="10" height="10" fill="#FFB238"></rect><rect x="70" width="10" height="10" fill="#FFB238"></rect><rect y="10" width="10" height="10" fill="#0A0310"></rect><rect y="20" width="10" height="10" fill="#FFB238"></rect><rect y="30" width="10" height="10" fill="#49007E"></rect><rect y="40" width="10" height="10" fill="#0A0310"></rect><rect y="50" width="10" height="10" fill="#FF005B"></rect><rect y="60" width="10" height="10" fill="#FFB238"></rect><rect y="70" width="10" height="10" fill="#FFB238"></rect><rect x="20" y="10" width="10" height="10" fill="#FF005B"></rect><rect x="20" y="20" width="10" height="10" fill="#49007E"></rect><rect x="20" y="30" width="10" height="10" fill="#0A0310"></rect><rect x="20" y="40" width="10" height="10" fill="#FF7D10"></rect><rect x="20" y="50" width="10" height="10" fill="#FFB238"></rect><rect x="20" y="60" width="10" height="10" fill="#FF7D10"></rect><rect x="20" y="70" width="10" height="10" fill="#FF005B"></rect><rect x="40" y="10" width="10" height="10" fill="#FF005B"></rect><rect x="40" y="20" width="10" height="10" fill="#FF005B"></rect><rect x="40" y="30" width="10" height="10" fill="#FFB238"></rect><rect x="40" y="40" width="10" height="10" fill="#0A0310"></rect><rect x="40" y="50" width="10" height="10" fill="#0A0310"></rect><rect x="40" y="60" width="10" height="10" fill="#FFB238"></rect><rect x="40" y="70" width="10" height="10" fill="#0A0310"></rect><rect x="60" y="10" width="10" height="10" fill="#FFB238"></rect><rect x="60" y="20" width="10" height="10" fill="#FFB238"></rect><rect x="60" y="30" width="10" height="10" fill="#FF7D10"></rect><rect x="60" y="40" width="10" height="10" fill="#FF005B"></rect><rect x="60" y="50" width="10" height="10" fill="#FF7D10"></rect><rect x="60" y="60" width="10" height="10" fill="#FFB238"></rect><rect x="60" y="70" width="10" height="10" fill="#0A0310"></rect><rect x="10" y="10" width="10" height="10" fill="#FF7D10"></rect><rect x="10" y="20" width="10" height="10" fill="#FF7D10"></rect><rect x="10" y="30" width="10" height="10" fill="#FF7D10"></rect><rect x="10" y="40" width="10" height="10" fill="#FFB238"></rect><rect x="10" y="50" width="10" height="10" fill="#FFB238"></rect><rect x="10" y="60" width="10" height="10" fill="#FF7D10"></rect><rect x="10" y="70" width="10" height="10" fill="#49007E"></rect><rect x="30" y="10" width="10" height="10" fill="#FF005B"></rect><rect x="30" y="20" width="10" height="10" fill="#FFB238"></rect><rect x="30" y="30" width="10" height="10" fill="#0A0310"></rect><rect x="30" y="40" width="10" height="10" fill="#0A0310"></rect><rect x="30" y="50" width="10" height="10" fill="#FF005B"></rect><rect x="30" y="60" width="10" height="10" fill="#FF005B"></rect><rect x="30" y="70" width="10" height="10" fill="#FFB238"></rect><rect x="50" y="10" width="10" height="10" fill="#FF7D10"></rect><rect x="50" y="20" width="10" height="10" fill="#0A0310"></rect><rect x="50" y="30" width="10" height="10" fill="#0A0310"></rect><rect x="50" y="40" width="10" height="10" fill="#0A0310"></rect><rect x="50" y="50" width="10" height="10" fill="#FFB238"></rect><rect x="50" y="60" width="10" height="10" fill="#FFB238"></rect><rect x="50" y="70" width="10" height="10" fill="#FF005B"></rect><rect x="70" y="10" width="10" height="10" fill="#FFB238"></rect><rect x="70" y="20" width="10" height="10" fill="#49007E"></rect><rect x="70" y="30" width="10" height="10" fill="#FFB238"></rect><rect x="70" y="40" width="10" height="10" fill="#0A0310"></rect><rect x="70" y="50" width="10" height="10" fill="#FFB238"></rect><rect x="70" y="60" width="10" height="10" fill="#FF7D10"></rect><rect x="70" y="70" width="10" height="10" fill="#FF7D10"></rect></g></svg>`,
		},
		{
			name: "Mary Roebling",
			args: []Option{Variant(Pixel)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r6f:" mask-type="alpha" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r6f:)"><rect width="10" height="10" fill="#0A0310"></rect><rect x="20" width="10" height="10" fill="#49007E"></rect><rect x="40" width="10" height="10" fill="#FF005B"></rect><rect x="60" width="10" height="10" fill="#FF7D10"></rect><rect x="10" width="10" height="10" fill="#0A0310"></rect><rect x="30" width="10" height="10" fill="#0A0310"></rect><rect x="50" width="10" height="10" fill="#FFB238"></rect><rect x="70" width="10" height="10" fill="#FF7D10"></rect><rect y="10" width="10" height="10" fill="#FF005B"></rect><rect y="20" width="10" height="10" fill="#0A0310"></rect><rect y="30" width="10" height="10" fill="#FF005B"></rect><rect y="40" width="10" height="10" fill="#49007E"></rect><rect y="50" width="10" height="10" fill="#FF005B"></rect><rect y="60" width="10" height="10" fill="#49007E"></rect><rect y="70" width="10" height="10" fill="#0A0310"></rect><rect x="20" y="10" width="10" height="10" fill="#49007E"></rect><rect x="20" y="20" width="10" height="10" fill="#FF7D10"></rect><rect x="20" y="30" width="10" height="10" fill="#49007E"></rect><rect x="20" y="40" width="10" height="10" fill="#0A0310"></rect><rect x="20" y="50" width="10" height="10" fill="#0A0310"></rect><rect x="20" y="60" width="10" height="10" fill="#49007E"></rect><rect x="20" y="70" width="10" height="10" fill="#FF7D10"></rect><rect x="40" y="10" width="10" height="10" fill="#FF005B"></rect><rect x="40" y="20" width="10" height="10" fill="#49007E"></rect><rect x="40" y="30" width="10" height="10" fill="#0A0310"></rect><rect x="40" y="40" width="10" height="10" fill="#0A0310"></rect><rect x="40" y="50" width="10" height="10" fill="#FF005B"></rect><rect x="40" y="60" width="10" height="10" fill="#49007E"></rect><rect x="40" y="70" width="10" height="10" fill="#FFB238"></rect><rect x="60" y="10" width="10" height="10" fill="#0A0310"></rect><rect x="60" y="20" width="10" height="10" fill="#49007E"></rect><rect x="60" y="30" width="10" height="10" fill="#49007E"></rect><rect x="60" y="40" width="10" height="10" fill="#FF005B"></rect><rect x="60" y="50" width="10" height="10" fill="#0A0310"></rect><rect x="60" y="60" width="10" height="10" fill="#0A0310"></rect><rect x="60" y="70" width="10" height="10" fill="#49007E"></rect><rect x="10" y="10" width="10" height="10" fill="#FF7D10"></rect><rect x="10" y="20" width="10" height="10" fill="#FFB238"></rect><rect x="10" y="30" width="10" height="10" fill="#FF7D10"></rect><rect x="10" y="40" width="10" height="10" fill="#0A0310"></rect><rect x="10" y="50" width="10" height="10" fill="#0A0310"></rect><rect x="10" y="60" width="10" height="10" fill="#49007E"></rect><rect x="10" y="70" width="10" height="10" fill="#FF7D10"></rect><rect x="30" y="10" width="10" height="10" fill="#0A0310"></rect><rect x="30" y="20" width="10" height="10" fill="#0A0310"></rect><rect x="30" y="30" width="10" height="10" fill="#FF005B"></rect><rect x="30" y="40" width="10" height="10" fill="#FFB238"></rect><rect x="30" y="50" width="10" height="10" fill="#49007E"></rect><rect x="30" y="60" width="10" height="10" fill="#0A0310"></rect><rect x="30" y="70" width="10" height="10" fill="#0A0310"></rect><rect x="50" y="10" width="10" height="10" fill="#FF7D10"></rect><rect x="50" y="20" width="10" height="10" fill="#49007E"></rect><rect x="50" y="30" width="10" height="10" fill="#FF7D10"></rect><rect x="50" y="40" width="10" height="10" fill="#FFB238"></rect><rect x="50" y="50" width="10" height="10" fill="#0A0310"></rect><rect x="50" y="60" width="10" height="10" fill="#49007E"></rect><rect x="50" y="70" width="10" height="10" fill="#FF7D10"></rect><rect x="70" y="10" width="10" height="10" fill="#FFB238"></rect><rect x="70" y="20" width="10" height="10" fill="#49007E"></rect><rect x="70" y="30" width="10" height="10" fill="#0A0310"></rect><rect x="70" y="40" width="10" height="10" fill="#49007E"></rect><rect x="70" y="50" width="10" height="10" fill="#FF005B"></rect><rect x="70" y="60" width="10" height="10" fill="#49007E"></rect><rect x="70" y="70" width="10" height="10" fill="#49007E"></rect></g></svg>`,
		},
		{
			name: "Sarah Winnemucca",
			args: []Option{Variant(Pixel)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r6g:" mask-type="alpha" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r6g:)"><rect width="10" height="10" fill="#0A0310"></rect><rect x="20" width="10" height="10" fill="#49007E"></rect><rect x="40" width="10" height="10" fill="#49007E"></rect><rect x="60" width="10" height="10" fill="#49007E"></rect><rect x="10" width="10" height="10" fill="#49007E"></rect><rect x="30" width="10" height="10" fill="#49007E"></rect><rect x="50" width="10" height="10" fill="#FF7D10"></rect><rect x="70" width="10" height="10" fill="#0A0310"></rect><rect y="10" width="10" height="10" fill="#FFB238"></rect><rect y="20" width="10" height="10" fill="#49007E"></rect><rect y="30" width="10" height="10" fill="#49007E"></rect><rect y="40" width="10" height="10" fill="#49007E"></rect><rect y="50" width="10" height="10" fill="#49007E"></rect><rect y="60" width="10" height="10" fill="#FF7D10"></rect><rect y="70" width="10" height="10" fill="#49007E"></rect><rect x="20" y="10" width="10" height="10" fill="#0A0310"></rect><rect x="20" y="20" width="10" height="10" fill="#FFB238"></rect><rect x="20" y="30" width="10" height="10" fill="#FF7D10"></rect><rect x="20" y="40" width="10" height="10" fill="#FF7D10"></rect><rect x="20" y="50" width="10" height="10" fill="#49007E"></rect><rect x="20" y="60" width="10" height="10" fill="#0A0310"></rect><rect x="20" y="70" width="10" height="10" fill="#FF005B"></rect><rect x="40" y="10" width="10" height="10" fill="#FF7D10"></rect><rect x="40" y="20" width="10" height="10" fill="#FF7D10"></rect><rect x="40" y="30" width="10" height="10" fill="#49007E"></rect><rect x="40" y="40" width="10" height="10" fill="#FFB238"></rect><rect x="40" y="50" width="10" height="10" fill="#FFB238"></rect><rect x="40" y="60" width="10" height="10" fill="#FF005B"></rect><rect x="40" y="70" width="10" height="10" fill="#FFB238"></rect><rect x="60" y="10" width="10" height="10" fill="#49007E"></rect><rect x="60" y="20" width="10" height="10" fill="#FFB238"></rect><rect x="60" y="30" width="10" height="10" fill="#0A0310"></rect><rect x="60" y="40" width="10" height="10" fill="#FF7D10"></rect><rect x="60" y="50" width="10" height="10" fill="#49007E"></rect><rect x="60" y="60" width="10" height="10" fill="#49007E"></rect><rect x="60" y="70" width="10" height="10" fill="#FF7D10"></rect><rect x="10" y="10" width="10" height="10" fill="#FFB238"></rect><rect x="10" y="20" width="10" height="10" fill="#FF005B"></rect><rect x="10" y="30" width="10" height="10" fill="#FFB238"></rect><rect x="10" y="40" width="10" height="10" fill="#49007E"></rect><rect x="10" y="50" width="10" height="10" fill="#0A0310"></rect><rect x="10" y="60" width="10" height="10" fill="#49007E"></rect><rect x="10" y="70" width="10" height="10" fill="#FFB238"></rect><rect x="30" y="10" width="10" height="10" fill="#FF005B"></rect><rect x="30" y="20" width="10" height="10" fill="#49007E"></rect><rect x="30" y="30" width="10" height="10" fill="#49007E"></rect><rect x="30" y="40" width="10" height="10" fill="#FF005B"></rect><rect x="30" y="50" width="10" height="10" fill="#FF005B"></rect><rect x="30" y="60" width="10" height="10" fill="#FF7D10"></rect><rect x="30" y="70" width="10" height="10" fill="#49007E"></rect><rect x="50" y="10" width="10" height="10" fill="#FFB238"></rect><rect x="50" y="20" width="10" height="10" fill="#0A0310"></rect><rect x="50" y="30" width="10" height="10" fill="#FF7D10"></rect><rect x="50" y="40" width="10" height="10" fill="#49007E"></rect><rect x="50" y="50" width="10" height="10" fill="#49007E"></rect><rect x="50" y="60" width="10" height="10" fill="#0A0310"></rect><rect x="50" y="70" width="10" height="10" fill="#49007E"></rect><rect x="70" y="10" width="10" height="10" fill="#FF7D10"></rect><rect x="70" y="20" width="10" height="10" fill="#0A0310"></rect><rect x="70" y="30" width="10" height="10" fill="#49007E"></rect><rect x="70" y="40" width="10" height="10" fill="#FF7D10"></rect><rect x="70" y="50" width="10" height="10" fill="#FFB238"></rect><rect x="70" y="60" width="10" height="10" fill="#49007E"></rect><rect x="70" y="70" width="10" height="10" fill="#FF005B"></rect></g></svg>`,
		},
		{
			name: "Margaret Brent",
			args: []Option{Variant(Pixel)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r6h:" mask-type="alpha" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r6h:)"><rect width="10" height="10" fill="#0A0310"></rect><rect x="20" width="10" height="10" fill="#0A0310"></rect><rect x="40" width="10" height="10" fill="#FF005B"></rect><rect x="60" width="10" height="10" fill="#0A0310"></rect><rect x="10" width="10" height="10" fill="#49007E"></rect><rect x="30" width="10" height="10" fill="#FF005B"></rect><rect x="50" width="10" height="10" fill="#49007E"></rect><rect x="70" width="10" height="10" fill="#FFB238"></rect><rect y="10" width="10" height="10" fill="#0A0310"></rect><rect y="20" width="10" height="10" fill="#49007E"></rect><rect y="30" width="10" height="10" fill="#0A0310"></rect><rect y="40" width="10" height="10" fill="#FF7D10"></rect><rect y="50" width="10" height="10" fill="#FF005B"></rect><rect y="60" width="10" height="10" fill="#FF7D10"></rect><rect y="70" width="10" height="10" fill="#49007E"></rect><rect x="20" y="10" width="10" height="10" fill="#FFB238"></rect><rect x="20" y="20" width="10" height="10" fill="#FFB238"></rect><rect x="20" y="30" width="10" height="10" fill="#FFB238"></rect><rect x="20" y="40" width="10" height="10" fill="#49007E"></rect><rect x="20" y="50" width="10" height="10" fill="#49007E"></rect><rect x="20" y="60" width="10" height="10" fill="#FF7D10"></rect><rect x="20" y="70" width="10" height="10" fill="#0A0310"></rect><rect x="40" y="10" width="10" height="10" fill="#FF7D10"></rect><rect x="40" y="20" width="10" height="10" fill="#0A0310"></rect><rect x="40" y="30" width="10" height="10" fill="#49007E"></rect><rect x="40" y="40" width="10" height="10" fill="#FF005B"></rect><rect x="40" y="50" width="10" height="10" fill="#FF7D10"></rect><rect x="40" y="60" width="10" height="10" fill="#FF7D10"></rect><rect x="40" y="70" width="10" height="10" fill="#FF005B"></rect><rect x="60" y="10" width="10" height="10" fill="#49007E"></rect><rect x="60" y="20" width="10" height="10" fill="#FFB238"></rect><rect x="60" y="30" width="10" height="10" fill="#FFB238"></rect><rect x="60" y="40" width="10" height="10" fill="#49007E"></rect><rect x="60" y="50" width="10" height="10" fill="#FFB238"></rect><rect x="60" y="60" width="10" height="10" fill="#49007E"></rect><rect x="60" y="70" width="10" height="10" fill="#FF005B"></rect><rect x="10" y="10" width="10" height="10" fill="#49007E"></rect><rect x="10" y="20" width="10" height="10" fill="#0A0310"></rect><rect x="10" y="30" width="10" height="10" fill="#FF005B"></rect><rect x="10" y="40" width="10" height="10" fill="#49007E"></rect><rect x="10" y="50" width="10" height="10" fill="#0A0310"></rect><rect x="10" y="60" width="10" height="10" fill="#FF7D10"></rect><rect x="10" y="70" width="10" height="10" fill="#49007E"></rect><rect x="30" y="10" width="10" height="10" fill="#0A0310"></rect><rect x="30" y="20" width="10" height="10" fill="#49007E"></rect><rect x="30" y="30" width="10" height="10" fill="#FF7D10"></rect><rect x="30" y="40" width="10" height="10" fill="#FF005B"></rect><rect x="30" y="50" width="10" height="10" fill="#0A0310"></rect><rect x="30" y="60" width="10" height="10" fill="#FF7D10"></rect><rect x="30" y="70" width="10" height="10" fill="#49007E"></rect><rect x="50" y="10" width="10" height="10" fill="#FFB238"></rect><rect x="50" y="20" width="10" height="10" fill="#FF7D10"></rect><rect x="50" y="30" width="10" height="10" fill="#FF7D10"></rect><rect x="50" y="40" width="10" height="10" fill="#0A0310"></rect><rect x="50" y="50" width="10" height="10" fill="#49007E"></rect><rect x="50" y="60" width="10" height="10" fill="#49007E"></rect><rect x="50" y="70" width="10" height="10" fill="#49007E"></rect><rect x="70" y="10" width="10" height="10" fill="#49007E"></rect><rect x="70" y="20" width="10" height="10" fill="#FF005B"></rect><rect x="70" y="30" width="10" height="10" fill="#49007E"></rect><rect x="70" y="40" width="10" height="10" fill="#0A0310"></rect><rect x="70" y="50" width="10" height="10" fill="#FFB238"></rect><rect x="70" y="60" width="10" height="10" fill="#0A0310"></rect><rect x="70" y="70" width="10" height="10" fill="#49007E"></rect></g></svg>`,
		},
		// Bauhaus
		{
			name: "Mary Baker",
			args: []Option{Variant(Bauhaus)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r9g:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r9g:)"><rect width="80" height="80" fill="#0A0310"></rect><rect x="10" y="30" width="80" height="80" fill="#49007E" transform="translate(-8 -8) rotate(320 40 40)"></rect><circle cx="40" cy="40" fill="#FF005B" r="16" transform="translate(-3 -3)"></circle><line x1="0" y1="40" x2="80" y2="40" stroke-width="2" stroke="#FF7D10" transform="translate(-0 -0) rotate(280 40 40)"></line></g></svg>`,
		},
		{
			name: "Amelia Earhart",
			args: []Option{Variant(Bauhaus)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r9h:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r9h:)"><rect width="80" height="80" fill="#FFB238"></rect><rect x="10" y="30" width="80" height="80" fill="#0A0310" transform="translate(-2 -2) rotate(288 40 40)"></rect><circle cx="40" cy="40" fill="#49007E" r="16" transform="translate(12 -12)"></circle><line x1="0" y1="40" x2="80" y2="40" stroke-width="2" stroke="#FF005B" transform="translate(16 -16) rotate(216 40 40)"></line></g></svg>`,
		},
		{
			name: "Mary Roebling",
			args: []Option{Variant(Bauhaus)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r9i:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r9i:)"><rect width="80" height="80" fill="#0A0310"></rect><rect x="10" y="30" width="80" height="80" fill="#49007E" transform="translate(4 4) rotate(310 40 40)"></rect><circle cx="40" cy="40" fill="#FF005B" r="16" transform="translate(-12 -12)"></circle><line x1="0" y1="40" x2="80" y2="40" stroke-width="2" stroke="#FF7D10" transform="translate(-0 0) rotate(260 40 40)"></line></g></svg>`,
		},
		{
			name: "Sarah Winnemucca",
			args: []Option{Variant(Bauhaus)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r9j:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r9j:)"><rect width="80" height="80" fill="#49007E"></rect><rect x="10" y="30" width="80" height="10" fill="#FF005B" transform="translate(-12 -12) rotate(242 40 40)"></rect><circle cx="40" cy="40" fill="#FF7D10" r="16" transform="translate(-9 9)"></circle><line x1="0" y1="40" x2="80" y2="40" stroke-width="2" stroke="#FFB238" transform="translate(-4 -4) rotate(124 40 40)"></line></g></svg>`,
		},
		{
			name: "Margaret Brent",
			args: []Option{Variant(Bauhaus)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r9k:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r9k:)"><rect width="80" height="80" fill="#49007E"></rect><rect x="10" y="30" width="80" height="10" fill="#FF005B" transform="translate(0 0) rotate(352 40 40)"></rect><circle cx="40" cy="40" fill="#FF7D10" r="16" transform="translate(-3 -3)"></circle><line x1="0" y1="40" x2="80" y2="40" stroke-width="2" stroke="#FFB238" transform="translate(-4 -4) rotate(344 40 40)"></line></g></svg>`,
		},
		// Ring
		{
			name: "Mary Baker",
			args: []Option{Variant(Ring)},
			want: `<svg viewBox="0 0 90 90" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":rip:" maskUnits="userSpaceOnUse" x="0" y="0" width="90" height="90"><rect width="90" height="90" rx="180" fill="#FFFFFF"></rect></mask><g mask="url(#:rip:)"><path d="M0 0h90v45H0z" fill="#0A0310"></path><path d="M0 45h90v45H0z" fill="#49007E"></path><path d="M83 45a38 38 0 00-76 0h76z" fill="#49007E"></path><path d="M83 45a38 38 0 01-76 0h76z" fill="#FF005B"></path><path d="M77 45a32 32 0 10-64 0h64z" fill="#FF005B"></path><path d="M77 45a32 32 0 11-64 0h64z" fill="#FF7D10"></path><path d="M71 45a26 26 0 00-52 0h52z" fill="#FF7D10"></path><path d="M71 45a26 26 0 01-52 0h52z" fill="#0A0310"></path><circle cx="45" cy="45" r="23" fill="#FFB238"></circle></g></svg>`,
		},
		{
			name: "Amelia Earhart",
			args: []Option{Variant(Ring)},
			want: `<svg viewBox="0 0 90 90" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":riq:" maskUnits="userSpaceOnUse" x="0" y="0" width="90" height="90"><rect width="90" height="90" rx="180" fill="#FFFFFF"></rect></mask><g mask="url(#:riq:)"><path d="M0 0h90v45H0z" fill="#FFB238"></path><path d="M0 45h90v45H0z" fill="#0A0310"></path><path d="M83 45a38 38 0 00-76 0h76z" fill="#0A0310"></path><path d="M83 45a38 38 0 01-76 0h76z" fill="#49007E"></path><path d="M77 45a32 32 0 10-64 0h64z" fill="#49007E"></path><path d="M77 45a32 32 0 11-64 0h64z" fill="#FF005B"></path><path d="M71 45a26 26 0 00-52 0h52z" fill="#FF005B"></path><path d="M71 45a26 26 0 01-52 0h52z" fill="#FFB238"></path><circle cx="45" cy="45" r="23" fill="#FF7D10"></circle></g></svg>`,
		},
		{
			name: "Mary Roebling",
			args: []Option{Variant(Ring)},
			want: `<svg viewBox="0 0 90 90" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":rir:" maskUnits="userSpaceOnUse" x="0" y="0" width="90" height="90"><rect width="90" height="90" rx="180" fill="#FFFFFF"></rect></mask><g mask="url(#:rir:)"><path d="M0 0h90v45H0z" fill="#0A0310"></path><path d="M0 45h90v45H0z" fill="#49007E"></path><path d="M83 45a38 38 0 00-76 0h76z" fill="#49007E"></path><path d="M83 45a38 38 0 01-76 0h76z" fill="#FF005B"></path><path d="M77 45a32 32 0 10-64 0h64z" fill="#FF005B"></path><path d="M77 45a32 32 0 11-64 0h64z" fill="#FF7D10"></path><path d="M71 45a26 26 0 00-52 0h52z" fill="#FF7D10"></path><path d="M71 45a26 26 0 01-52 0h52z" fill="#0A0310"></path><circle cx="45" cy="45" r="23" fill="#FFB238"></circle></g></svg>`,
		},
		{
			name: "Sarah Winnemucca",
			args: []Option{Variant(Ring)},
			want: `<svg viewBox="0 0 90 90" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":ris:" maskUnits="userSpaceOnUse" x="0" y="0" width="90" height="90"><rect width="90" height="90" rx="180" fill="#FFFFFF"></rect></mask><g mask="url(#:ris:)"><path d="M0 0h90v45H0z" fill="#49007E"></path><path d="M0 45h90v45H0z" fill="#FF005B"></path><path d="M83 45a38 38 0 00-76 0h76z" fill="#FF005B"></path><path d="M83 45a38 38 0 01-76 0h76z" fill="#FF7D10"></path><path d="M77 45a32 32 0 10-64 0h64z" fill="#FF7D10"></path><path d="M77 45a32 32 0 11-64 0h64z" fill="#FFB238"></path><path d="M71 45a26 26 0 00-52 0h52z" fill="#FFB238"></path><path d="M71 45a26 26 0 01-52 0h52z" fill="#49007E"></path><circle cx="45" cy="45" r="23" fill="#0A0310"></circle></g></svg>`,
		},
		{
			name: "Margaret Brent",
			args: []Option{Variant(Ring)},
			want: `<svg viewBox="0 0 90 90" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":rit:" maskUnits="userSpaceOnUse" x="0" y="0" width="90" height="90"><rect width="90" height="90" rx="180" fill="#FFFFFF"></rect></mask><g mask="url(#:rit:)"><path d="M0 0h90v45H0z" fill="#49007E"></path><path d="M0 45h90v45H0z" fill="#FF005B"></path><path d="M83 45a38 38 0 00-76 0h76z" fill="#FF005B"></path><path d="M83 45a38 38 0 01-76 0h76z" fill="#FF7D10"></path><path d="M77 45a32 32 0 10-64 0h64z" fill="#FF7D10"></path><path d="M77 45a32 32 0 11-64 0h64z" fill="#FFB238"></path><path d="M71 45a26 26 0 00-52 0h52z" fill="#FFB238"></path><path d="M71 45a26 26 0 01-52 0h52z" fill="#49007E"></path><circle cx="45" cy="45" r="23" fill="#0A0310"></circle></g></svg>`,
		},
		// Sunset
		{
			name: "Mary Baker",
			args: []Option{Variant(Sunset)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":rls:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:rls:)"><path fill="url(#gradient_paint0_linear_MaryBaker)" d="M0 0h80v40H0z"></path><path fill="url(#gradient_paint1_linear_MaryBaker)" d="M0 40h80v40H0z"></path></g><defs><linearGradient id="gradient_paint0_linear_MaryBaker" x1="40" y1="0" x2="40" y2="40" gradientUnits="userSpaceOnUse"><stop stop-color="#0A0310"></stop><stop offset="1" stop-color="#49007E"></stop></linearGradient><linearGradient id="gradient_paint1_linear_MaryBaker" x1="40" y1="40" x2="40" y2="80" gradientUnits="userSpaceOnUse"><stop stop-color="#FF005B"></stop><stop offset="1" stop-color="#FF7D10"></stop></linearGradient></defs></svg>`,
		},
		{
			name: "Amelia Earhart",
			args: []Option{Variant(Sunset)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":rlt:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:rlt:)"><path fill="url(#gradient_paint0_linear_AmeliaEarhart)" d="M0 0h80v40H0z"></path><path fill="url(#gradient_paint1_linear_AmeliaEarhart)" d="M0 40h80v40H0z"></path></g><defs><linearGradient id="gradient_paint0_linear_AmeliaEarhart" x1="40" y1="0" x2="40" y2="40" gradientUnits="userSpaceOnUse"><stop stop-color="#FFB238"></stop><stop offset="1" stop-color="#0A0310"></stop></linearGradient><linearGradient id="gradient_paint1_linear_AmeliaEarhart" x1="40" y1="40" x2="40" y2="80" gradientUnits="userSpaceOnUse"><stop stop-color="#49007E"></stop><stop offset="1" stop-color="#FF005B"></stop></linearGradient></defs></svg>`,
		},
		{
			name: "Mary Roebling",
			args: []Option{Variant(Sunset)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":rlu:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:rlu:)"><path fill="url(#gradient_paint0_linear_MaryRoebling)" d="M0 0h80v40H0z"></path><path fill="url(#gradient_paint1_linear_MaryRoebling)" d="M0 40h80v40H0z"></path></g><defs><linearGradient id="gradient_paint0_linear_MaryRoebling" x1="40" y1="0" x2="40" y2="40" gradientUnits="userSpaceOnUse"><stop stop-color="#0A0310"></stop><stop offset="1" stop-color="#49007E"></stop></linearGradient><linearGradient id="gradient_paint1_linear_MaryRoebling" x1="40" y1="40" x2="40" y2="80" gradientUnits="userSpaceOnUse"><stop stop-color="#FF005B"></stop><stop offset="1" stop-color="#FF7D10"></stop></linearGradient></defs></svg>`,
		},
		{
			name: "Sarah Winnemucca",
			args: []Option{Variant(Sunset)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":rlv:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:rlv:)"><path fill="url(#gradient_paint0_linear_SarahWinnemucca)" d="M0 0h80v40H0z"></path><path fill="url(#gradient_paint1_linear_SarahWinnemucca)" d="M0 40h80v40H0z"></path></g><defs><linearGradient id="gradient_paint0_linear_SarahWinnemucca" x1="40" y1="0" x2="40" y2="40" gradientUnits="userSpaceOnUse"><stop stop-color="#49007E"></stop><stop offset="1" stop-color="#FF005B"></stop></linearGradient><linearGradient id="gradient_paint1_linear_SarahWinnemucca" x1="40" y1="40" x2="40" y2="80" gradientUnits="userSpaceOnUse"><stop stop-color="#FF7D10"></stop><stop offset="1" stop-color="#FFB238"></stop></linearGradient></defs></svg>`,
		},
		{
			name: "Margaret Brent",
			args: []Option{Variant(Sunset)},
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":rm0:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:rm0:)"><path fill="url(#gradient_paint0_linear_MargaretBrent)" d="M0 0h80v40H0z"></path><path fill="url(#gradient_paint1_linear_MargaretBrent)" d="M0 40h80v40H0z"></path></g><defs><linearGradient id="gradient_paint0_linear_MargaretBrent" x1="40" y1="0" x2="40" y2="40" gradientUnits="userSpaceOnUse"><stop stop-color="#49007E"></stop><stop offset="1" stop-color="#FF005B"></stop></linearGradient><linearGradient id="gradient_paint1_linear_MargaretBrent" x1="40" y1="40" x2="40" y2="80" gradientUnits="userSpaceOnUse"><stop stop-color="#FF7D10"></stop><stop offset="1" stop-color="#FFB238"></stop></linearGradient></defs></svg>`,
		},
		// Beam
		{
			name: "Mary Baker",
			args: []Option{Variant(Beam)},
			want: `<svg viewBox="0 0 36 36" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":rov:" maskUnits="userSpaceOnUse" x="0" y="0" width="36" height="36"><rect width="36" height="36" rx="72" fill="#FFFFFF"></rect></mask><g mask="url(#:rov:)"><rect width="36" height="36" fill="#FF7D10"></rect><rect x="0" y="0" width="36" height="36" transform="translate(4 4) rotate(340 18 18) scale(1.1)" fill="#0A0310" rx="36"></rect><g transform="translate(-4 -1) rotate(-0 18 18)"><path d="M15 20c2 1 4 1 6 0" stroke="#FFFFFF" fill="none" stroke-linecap="round"></path><rect x="14" y="14" width="1.5" height="2" rx="1" stroke="none" fill="#FFFFFF"></rect><rect x="20" y="14" width="1.5" height="2" rx="1" stroke="none" fill="#FFFFFF"></rect></g></g></svg>`,
		},
		{
			name: "Amelia Earhart",
			args: []Option{Variant(Beam)},
			want: `<svg viewBox="0 0 36 36" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":rp0:" maskUnits="userSpaceOnUse" x="0" y="0" width="36" height="36"><rect width="36" height="36" rx="72" fill="#FFFFFF"></rect></mask><g mask="url(#:rp0:)"><rect width="36" height="36" fill="#FF005B"></rect><rect x="0" y="0" width="36" height="36" transform="translate(0 0) rotate(324 18 18) scale(1)" fill="#FFB238" rx="36"></rect><g transform="translate(-4 -4) rotate(-4 18 18)"><path d="M15 19c2 1 4 1 6 0" stroke="#000000" fill="none" stroke-linecap="round"></path><rect x="10" y="14" width="1.5" height="2" rx="1" stroke="none" fill="#000000"></rect><rect x="24" y="14" width="1.5" height="2" rx="1" stroke="none" fill="#000000"></rect></g></g></svg>`,
		},
		{
			name: "Mary Roebling",
			args: []Option{Variant(Beam)},
			want: `<svg viewBox="0 0 36 36" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":rp1:" maskUnits="userSpaceOnUse" x="0" y="0" width="36" height="36"><rect width="36" height="36" rx="72" fill="#FFFFFF"></rect></mask><g mask="url(#:rp1:)"><rect width="36" height="36" fill="#FF7D10"></rect><rect x="0" y="0" width="36" height="36" transform="translate(5 -1) rotate(155 18 18) scale(1.2)" fill="#0A0310" rx="6"></rect><g transform="translate(3 -4) rotate(-5 18 18)"><path d="M15 21c2 1 4 1 6 0" stroke="#FFFFFF" fill="none" stroke-linecap="round"></path><rect x="14" y="14" width="1.5" height="2" rx="1" stroke="none" fill="#FFFFFF"></rect><rect x="20" y="14" width="1.5" height="2" rx="1" stroke="none" fill="#FFFFFF"></rect></g></g></svg>`,
		},
		{
			name: "Sarah Winnemucca",
			args: []Option{Variant(Beam)},
			want: `<svg viewBox="0 0 36 36" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":rp2:" maskUnits="userSpaceOnUse" x="0" y="0" width="36" height="36"><rect width="36" height="36" rx="72" fill="#FFFFFF"></rect></mask><g mask="url(#:rp2:)"><rect width="36" height="36" fill="#FFB238"></rect><rect x="0" y="0" width="36" height="36" transform="translate(3 5) rotate(301 18 18) scale(1.1)" fill="#49007E" rx="36"></rect><g transform="translate(-5 3) rotate(-1 18 18)"><path d="M13,20 a1,0.75 0 0,0 10,0" fill="#FFFFFF"></path><rect x="13" y="14" width="1.5" height="2" rx="1" stroke="none" fill="#FFFFFF"></rect><rect x="21" y="14" width="1.5" height="2" rx="1" stroke="none" fill="#FFFFFF"></rect></g></g></svg>`,
		},
		{
			name: "Margaret Brent",
			args: []Option{Variant(Beam)},
			want: `<svg viewBox="0 0 36 36" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":rp3:" maskUnits="userSpaceOnUse" x="0" y="0" width="36" height="36"><rect width="36" height="36" rx="72" fill="#FFFFFF"></rect></mask><g mask="url(#:rp3:)"><rect width="36" height="36" fill="#FFB238"></rect><rect x="0" y="0" width="36" height="36" transform="translate(6 6) rotate(356 18 18) scale(1.2)" fill="#49007E" rx="6"></rect><g transform="translate(4 1) rotate(6 18 18)"><path d="M13,21 a1,0.75 0 0,0 10,0" fill="#FFFFFF"></path><rect x="13" y="14" width="1.5" height="2" rx="1" stroke="none" fill="#FFFFFF"></rect><rect x="21" y="14" width="1.5" height="2" rx="1" stroke="none" fill="#FFFFFF"></rect></g></g></svg>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.name, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			g, err := getInternals(got)
			if err != nil {
				t.Log(got)
				t.Errorf("received error for rendered: %v", err)
				return
			}
			i, err := getInternals(tt.want)
			if err != nil {
				t.Log(tt.want)
				t.Errorf("received error for wanted: %v", err)
				return
			}

			if g != i {
				t.Errorf("New()\ng: %v\nw: %v", g, i)
			}
		})
	}
}

func TestRender(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Mary Baker",
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r3a:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r3a:)"><rect width="80" height="80" fill="#0A0310"></rect><path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="#49007E" transform="translate(-0 -0) rotate(-320 40 40) scale(1.2)"></path><path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="#FF005B" transform="translate(-4 -4) rotate(-300 40 40) scale(1.2)"></path></g><defs><filter id="prefix__filter0_f" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood flood-opacity="0" result="BackgroundImageFix"></feFlood><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"></feBlend><feGaussianBlur stdDeviation="7" result="effect1_foregroundBlur"></feGaussianBlur></filter></defs></svg>`,
		},
		{
			name: "Amelia Earhart",
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r3b:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r3b:)"><rect width="80" height="80" fill="#FFB238"></rect><path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="#0A0310" transform="translate(-0 -0) rotate(-288 40 40) scale(1.2)"></path><path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="#49007E" transform="translate(4 -4) rotate(252 40 40) scale(1.2)"></path></g><defs><filter id="prefix__filter0_f" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood flood-opacity="0" result="BackgroundImageFix"></feFlood><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"></feBlend><feGaussianBlur stdDeviation="7" result="effect1_foregroundBlur"></feGaussianBlur></filter></defs></svg>`,
		},
		{
			name: "Mary Roebling",
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r3c:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r3c:)"><rect width="80" height="80" fill="#0A0310"></rect><path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="#49007E" transform="translate(6 6) rotate(310 40 40) scale(1.3)"></path><path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="#FF005B" transform="translate(-1 -1) rotate(-105 40 40) scale(1.3)"></path></g><defs><filter id="prefix__filter0_f" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood flood-opacity="0" result="BackgroundImageFix"></feFlood><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"></feBlend><feGaussianBlur stdDeviation="7" result="effect1_foregroundBlur"></feGaussianBlur></filter></defs></svg>`,
		},
		{
			name: "Sarah Winnemucca",
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r3d:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r3d:)"><rect width="80" height="80" fill="#49007E"></rect><path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="#FF005B" transform="translate(-2 -2) rotate(-242 40 40) scale(1.5)"></path><path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="#FF7D10" transform="translate(-7 7) rotate(-183 40 40) scale(1.5)"></path></g><defs><filter id="prefix__filter0_f" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood flood-opacity="0" result="BackgroundImageFix"></feFlood><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"></feBlend><feGaussianBlur stdDeviation="7" result="effect1_foregroundBlur"></feGaussianBlur></filter></defs></svg>`,
		},
		{
			name: "Margaret Brent",
			want: `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="80" height="80"><mask id=":r3e:" maskUnits="userSpaceOnUse" x="0" y="0" width="80" height="80"><rect width="80" height="80" rx="160" fill="#FFFFFF"></rect></mask><g mask="url(#:r3e:)"><rect width="80" height="80" fill="#49007E"></rect><path filter="url(#prefix__filter0_f)" d="M32.414 59.35L50.376 70.5H72.5v-71H33.728L26.5 13.381l19.057 27.08L32.414 59.35z" fill="#FF005B" transform="translate(0 0) rotate(352 40 40) scale(1.2)"></path><path filter="url(#prefix__filter0_f)" style="mix-blend-mode: overlay;" d="M22.216 24L0 46.75l14.108 38.129L78 86l-3.081-59.276-22.378 4.005 12.972 20.186-23.35 27.395L22.215 24z" fill="#FF7D10" transform="translate(-4 -4) rotate(-348 40 40) scale(1.2)"></path></g><defs><filter id="prefix__filter0_f" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood flood-opacity="0" result="BackgroundImageFix"></feFlood><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"></feBlend><feGaussianBlur stdDeviation="7" result="effect1_foregroundBlur"></feGaussianBlur></filter></defs></svg>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Render(tt.name, Size(80, "px"))

			b := bytes.NewBuffer(nil)

			if err := r.Render(nil, b); err != nil {
				t.Errorf("unable to render: %v", err)
				return
			}

			got := b.String()

			g, err := getInternals(got)
			if err != nil {
				t.Log(got)
				t.Errorf("received error for rendered: %v", err)
				return
			}
			i, err := getInternals(tt.want)
			if err != nil {
				t.Log(tt.want)
				t.Errorf("received error for wanted: %v", err)
				return
			}

			if g != i {
				t.Errorf("New()\ng: %v\nw: %v", g, i)
			}

		})
	}
}

func TestClasses(t *testing.T) {
	tests := []struct {
		name        string
		want        []string
		wantClasses bool
	}{
		{
			name:        "Mary Baker",
			want:        nil,
			wantClasses: false,
		},
		{
			name:        "Mary Baker",
			want:        []string{"single"},
			wantClasses: true,
		},
		{
			name:        "Mary Baker",
			want:        []string{"single", "double"},
			wantClasses: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			raw := Render(tt.name, Classes(tt.want...))
			r := raw.String()

			if tt.wantClasses && !strings.HasPrefix(r, fmt.Sprintf(`<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="40" height="40" class="%s">`, strings.Join(tt.want, " "))) {
				t.Errorf("%s does not have classes %s", tt.name, tt.want)
				return
			}

			if !tt.wantClasses && !strings.HasPrefix(r, `<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="40" height="40">`) {
				t.Errorf("%s has classes when they shouldn't", tt.name)
				return
			}

		})
	}

}

func TestErr(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Mary Baker",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := Render(tt.name, returnErr("testing"))

			b := bytes.NewBuffer(nil)

			if err := r.Render(nil, b); err == nil {
				t.Errorf("%s did not return an error", tt.name)
			}

			if r.String() != "" {
				t.Errorf("%s returned data", tt.name)
			}

		})
	}

}

func TestSizing(t *testing.T) {
	tests := []struct {
		name string
		size float64
		unit string
		want string
	}{
		{
			name: "Mary Baker",
			size: 40,
			unit: "px",
			want: "40px",
		},
		{
			name: "Mary Baker",
			size: 1.500,
			unit: "rem",
			want: "1.5rem",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r, err := New(tt.name, Size(tt.size, tt.unit))
			if err != nil {
				t.Errorf("%s got %v", tt.name, err)
				return
			}

			if !strings.HasPrefix(r, fmt.Sprintf(`<svg viewBox="0 0 80 80" fill="none" role="img" xmlns="http://www.w3.org/2000/svg" width="%s" height="%s">`, tt.want, tt.want)) {

				b, _, _ := strings.Cut(r, "><")

				t.Errorf("%s does not have %s as sizes\n%s>", tt.name, tt.want, b)
				return
			}

		})
	}

}
