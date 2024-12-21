package main

import (
	"fmt"
	"strings"
)

const (
	Reset = "\033[0m"
)

// Color codes for text and background
var (
	Styles = map[string]string{
		"reset":     "0",
		"bold":      "1",
		"underline": "4",
		"invert":    "7",
	}

	ForegroundColors = map[string]string{
		"black":   "30",
		"red":     "31",
		"green":   "32",
		"yellow":  "33",
		"blue":    "34",
		"magenta": "35",
		"cyan":    "36",
		"white":   "37",
	}

	BackgroundColors = map[string]string{
		"black":   "40",
		"red":     "41",
		"green":   "42",
		"yellow":  "43",
		"blue":    "44",
		"magenta": "45",
		"cyan":    "46",
		"white":   "47",
	}
)

type OptFunc func(*Opts)

type Opts struct {
	style     string
	textColor string
	bgColor   string
	padding   int16
}

func defaultOpts() Opts {
	return Opts{
		style:     Styles["reset"],
		textColor: ForegroundColors["white"],
		bgColor:   BackgroundColors["black"],
		padding:   4,
	}
}

func main() {
	text := "Lorum ipsum dolar sit amet!"

	Block(text, textRed)
}

func textRed(opts *Opts) {
	opts.textColor = ForegroundColors["red"]
}

func styleText(s string, opts ...OptFunc) string {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}

	ansi := "\033["
	codes := []string{}

	codes = append(codes, o.style)
	codes = append(codes, o.textColor)
	codes = append(codes, o.bgColor)

	if len(codes) > 0 {
		ansi += strings.Join(codes, ";")
	}
	ansi += "m"

	return ansi + s + Reset
}

func Block(s string, opts ...OptFunc) {
	i := len(s) + opts.padding

	fmt.Println(styleText(strings.Repeat(" ", i), opts...))
	fmt.Println(styleText(s, opts...))
	fmt.Println(styleText(strings.Repeat(" ", i), opts...))
}
