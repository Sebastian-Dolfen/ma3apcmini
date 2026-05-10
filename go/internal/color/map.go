package color

import (
	"math"
	"strconv"
	"strings"
)

// Subset of colorToVelocity from the JS, plus a few extras; can be extended.
var colorToVelocity = map[string]int{
	"#000000": 0, "#1E1E1E": 1, "#7F7F7F": 2, "#FFFFFF": 3,
	"#FF4C4C": 4, "#FF0000": 5, "#590000": 6, "#190000": 7,
	"#FFBD6C": 8, "#FF5400": 9, "#591D00": 10, "#271B00": 11,
	"#FFFF4C": 12, "#FFFF00": 13, "#595900": 14, "#191900": 15,
	"#88FF4C": 16, "#54FF00": 17, "#1D5900": 18, "#142B00": 19,
	"#4CFF4C": 20, "#00FF00": 21, "#005900": 22, "#001900": 23,
	"#4CFF5E": 24, "#00FF19": 25, "#00590D": 26, "#001902": 27,
	"#4CFF88": 28, "#00FF55": 29, "#00591D": 30, "#001F12": 31,
	"#4CFFB7": 32, "#00FF99": 33, "#005935": 34, "#001912": 35,
	"#4CC3FF": 36, "#00A9FF": 37, "#004152": 38, "#001019": 39,
	"#4C88FF": 40, "#0055FF": 41, "#001D59": 42, "#000819": 43,
	"#4C4CFF": 44, "#0000FF": 45, "#000059": 46, "#000019": 47,
	"#874CFF": 48, "#5400FF": 49, "#190064": 50, "#0F0030": 51,
	"#FF4CFF": 52, "#FF00FF": 53, "#590059": 54, "#190019": 55,
	"#FF4C87": 56, "#FF0054": 57, "#59001D": 58, "#220013": 59,
	"#FF1500": 60, "#993500": 61, "#795100": 62, "#436400": 63,
	"#033900": 64, "#005735": 65, "#00547F": 66, "#0000FE": 67,
	"#00454F": 68, "#2500CC": 69, "#7F7F70": 70, "#202020": 71,
	"#FF0001": 72, "#BDFF2D": 73, "#AFED06": 74, "#64FF09": 75,
	"#108B00": 76, "#00FF87": 77, "#00A8FF": 78, "#002AFF": 79,
	"#3F00FF": 80, "#7A00FF": 81, "#B21A7D": 82, "#402100": 83,
	"#FF4A00": 84, "#88E106": 85, "#72FF15": 86, "#00FF01": 87,
	"#3BFF26": 88, "#59FF71": 89, "#38FFCC": 90, "#5B8AFF": 91,
	"#3151C6": 92, "#877FE9": 93, "#D31DFF": 94, "#FF005D": 95,
	"#FF7F00": 96, "#B9B001": 97, "#90FF00": 98, "#835D07": 99,
	"#392b00": 100, "#144C10": 101, "#0D5038": 102, "#15152A": 103,
	"#16205A": 104, "#693C1C": 105, "#A8000A": 106, "#DE513D": 107,
	"#D86A1C": 108, "#FFE126": 109, "#9EE12F": 110, "#67B50F": 111,
	"#1E1E30": 112, "#DCFF6B": 113, "#80FFBD": 114, "#9A99FF": 115,
	"#8E66FF": 116, "#404040": 117, "#757575": 118, "#E0FFFF": 119,
	"#A00000": 120, "#350000": 121, "#1AD000": 122, "#074200": 123,
	"#B9B000": 124, "#3F3100": 125, "#B35F00": 126, "#4B1502": 127,
}

type rgb struct{ r, g, b int }

func hexToRGB(hex string) (rgb, bool) {
	if len(hex) != 7 || hex[0] != '#' {
		return rgb{}, false
	}
	r, err1 := strconv.ParseInt(hex[1:3], 16, 32)
	g, err2 := strconv.ParseInt(hex[3:5], 16, 32)
	b, err3 := strconv.ParseInt(hex[5:7], 16, 32)
	if err1 != nil || err2 != nil || err3 != nil {
		return rgb{}, false
	}
	return rgb{int(r), int(g), int(b)}, true
}

func manhattan(a, b rgb) int {
	return int(math.Abs(float64(a.r-b.r)) + math.Abs(float64(a.g-b.g)) + math.Abs(float64(a.b-b.b)))
}

// LookupVelocity returns exact or closest velocity for given #RRGGBB hex.
func LookupVelocity(hex string) int {
	if v, ok := colorToVelocity[strings.ToUpper(hex)]; ok { // exact
		return v
	}
	// find closest by Manhattan distance
	target, ok := hexToRGB(hex)
	if !ok {
		return 0
	}
	bestKey := "#000000"
	bestDist := math.MaxInt32
	for k := range colorToVelocity {
		c, ok := hexToRGB(k)
		if !ok {
			continue
		}
		d := manhattan(target, c)
		if d < bestDist {
			bestDist = d
			bestKey = k
		}
	}
	return colorToVelocity[bestKey]
}
