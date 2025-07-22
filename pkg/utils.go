package pkg

// ANSI escape codes
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

// colorForStatus devuelve el código ANSI según el rango de status
func ColorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return ColorGreen
	case code >= 300 && code < 400:
		return ColorWhite
	case code >= 400 && code < 500:
		return ColorYellow
	default: // 500+
		return ColorRed
	}
}

// colorForMethod devuelve un color fijo para cada verbo HTTP
func ColorForMethod(method string) string {
	switch method {
	case "GET":
		return ColorCyan
	case "POST":
		return ColorBlue
	case "PUT":
		return ColorYellow
	case "DELETE":
		return ColorRed
	case "PATCH":
		return ColorGreen
	default:
		return ColorWhite
	}
}
