/*
    This file is part of go3a.

    go3a is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    go3a is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with go3a.  If not, see <https://www.gnu.org/licenses/>.
*/
package go3a


type ColorMod int64

const (
	ColorModNone ColorMod = iota
	ColorModFg
	ColorModBg
	ColorModFull
)

func (cm ColorMod) ToDatacols() uint16 {
	switch cm {
		case ColorModNone:
			return 1
		case ColorModFg:
			return 2
		case ColorModBg:
			return 2
		case ColorModFull:
			return 3
	}
	return 1
}

func (cm ColorMod) ToString() string {
	switch cm {
		case ColorModNone:
			return "none"
		case ColorModFg:
			return "fg"
		case ColorModBg:
			return "bg"
		case ColorModFull:
			return "full"
	}
	return "none"
}

func ColorModFromString(s string) (ColorMod, error) {
	switch s {
		case "none":
			return ColorModNone, nil
		case "fg":
			return ColorModFg, nil
		case "bg":
			return ColorModBg, nil
		case "full":
			return ColorModFull, nil
	}
	return ColorModNone, UnknownColorMod{s}
}

type Color int64

const (
	ColorBlack Color        = iota
	ColorBlue
	ColorGreen
	ColorCyan
	ColorRed
	ColorMagenta
	ColorYellow
	ColorWhite
	ColorGray
	ColorBrightBlue
	ColorBrightGreen
	ColorBrightCyan
	ColorBrightRed
	ColorBrightMagenta
	ColorBrightYellow
	ColorBrightWhite
	NoColor
)

func (c Color) ToString() string {
	switch c {
		case ColorBlack:
			return "0"
		case ColorBlue:
			return "1"
		case ColorGreen:
			return "2"
		case ColorCyan:
			return "3"
		case ColorRed:
			return "4"
		case ColorMagenta:
			return "5"
		case ColorYellow:
			return "6"
		case ColorWhite:
			return "7"
		case ColorGray:
			return "8"
		case ColorBrightBlue:
			return "9"
		case ColorBrightGreen:
			return "a"
		case ColorBrightCyan:
			return "b"
		case ColorBrightRed:
			return "c"
		case ColorBrightMagenta:
			return "d"
		case ColorBrightYellow:
			return "e"
		case ColorBrightWhite:
			return "f"
	}
	return ""
}

func color_from_string(s string) (Color, error) {
	switch s {
		case "0":
			return ColorBlack, nil
		case "1":
			return ColorBlue, nil
		case "2":
			return ColorGreen, nil
		case "3":
			return ColorCyan, nil
		case "4":
			return ColorRed, nil
		case "5":
			return ColorMagenta, nil
		case "6":
			return ColorYellow, nil
		case "7":
			return ColorWhite, nil
		case "8":
			return ColorGray, nil
		case "9":
			return ColorBrightBlue, nil
		case "a":
			return ColorBrightGreen, nil
		case "b":
			return ColorBrightCyan, nil
		case "c":
			return ColorBrightRed, nil
		case "d":
			return ColorBrightMagenta, nil
		case "e":
			return ColorBrightYellow, nil
		case "f":
			return ColorBrightWhite, nil
		case "":
			return NoColor, nil
	}
	return NoColor, UnknownColor{s}
}
