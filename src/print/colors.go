package print

type Colors struct {
	systemDefault string
	black         string
	red           string
	green         string
	yellow        string
	blue          string
	purple        string
	cyan          string
	white         string
}

func ANSIColors() Colors {
	return Colors{
		systemDefault: "\033[1;0m",
		black:         "\033[1;30m",
		red:           "\033[1;31m",
		green:         "\033[1;32m",
		yellow:        "\033[1;33m",
		blue:          "\033[1;34m",
		purple:        "\033[1;35m",
		cyan:          "\033[1;36m",
		white:         "\033[1;37m",
	}
}
