package main

const (
	prefix = "\033["

	black  = prefix + "0;30m"
	red    = prefix + "0;31m"
	green  = prefix + "0;32m"
	brown  = prefix + "0;33m"
	blue   = prefix + "0;34m"
	purple = prefix + "0;35m"
	cyan   = prefix + "0;36m"
	gray   = prefix + "0;37m"

	none = prefix + "0m"
)

func Black(text string) string {
	return black + text + none
}

func Red(text string) string {
	return red + text + none
}

func Green(text string) string {
	return green + text + none
}

func Brown(text string) string {
	return brown + text + none
}

func Blue(text string) string {
	return blue + text + none
}

func Purple(text string) string {
	return purple + text + none
}

func Cyan(text string) string {
	return cyan + text + none
}

func Gray(text string) string {
	return gray + text + none
}
