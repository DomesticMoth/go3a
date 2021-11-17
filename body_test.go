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


import (
    "testing"
)

func TestBodyToText(t *testing.T) {
	body := Body{
		Frame{
			Row{
				RowFragment{
					"AA",
					ColorBrightGreen,
					ColorBlue,
				},
				RowFragment{
					"AA",
					ColorBrightCyan,
					ColorGreen,
				},
			},
			Row{
				RowFragment{
					"BB",
					ColorBrightGreen,
					ColorBlue,
				},
				RowFragment{
					"B",
					ColorBrightCyan,
					ColorGreen,
				},
				RowFragment{
					"B",
					ColorBrightRed,
					ColorGreen,
				},
			},
			Row{
				RowFragment{
					"CCCC",
					ColorBrightGreen,
					ColorBlue,
				},
			},
			Row{
				RowFragment{
					"D",
					ColorBrightGreen,
					ColorBlue,
				},
				RowFragment{
					"D",
					ColorBrightCyan,
					ColorBlue,
				},
				RowFragment{
					"D",
					ColorBrightRed,
					ColorBlue,
				},
				RowFragment{
					"D",
					ColorBrightMagenta,
					ColorBlue,
				},
			},
		},
		Frame{
			Row{
				RowFragment{
					"AA",
					ColorBrightGreen,
					ColorBlue,
				},
				RowFragment{
					"AA",
					ColorBrightCyan,
					ColorGreen,
				},
			},
			Row{
				RowFragment{
					"BB",
					ColorBrightGreen,
					ColorBlue,
				},
				RowFragment{
					"B",
					ColorBrightCyan,
					ColorGreen,
				},
				RowFragment{
					"B",
					ColorBrightRed,
					ColorGreen,
				},
			},
			Row{
				RowFragment{
					"CCCC",
					ColorBrightGreen,
					ColorBlue,
				},
			},
			Row{
				RowFragment{
					"D",
					ColorBrightGreen,
					ColorBlue,
				},
				RowFragment{
					"D",
					ColorBrightCyan,
					ColorBlue,
				},
				RowFragment{
					"D",
					ColorBrightRed,
					ColorBlue,
				},
				RowFragment{
					"D",
					ColorBrightMagenta,
					ColorBlue,
				},
			},
		},
	}
	text_reference := "AAAAaabb1122\nBBBBaabc1122\nCCCCaaaa1111\nDDDDabcd1111\n\nAAAAaabb1122\nBBBBaabc1122\nCCCCaaaa1111\nDDDDabcd1111\n"
    if text_reference != body.to_string(true) {
    	t.Errorf("Body to text convertion incorrect \n%s--\n%s--", text_reference, body.to_string(true))
    }
}

func TestBodyFromTextCorrectFullcolor(t *testing.T) {
    header := Header{
        4,
        4,
        200,
        true,
        ColorModFull,
        false,
        3,
        DEFAULT_PREVIEW,
        "",
        "",
        "",
    }
	body := Body{
		Frame{
			Row{
				RowFragment{
					"AA",
					ColorBrightGreen,
					ColorBlue,
				},
				RowFragment{
					"AA",
					ColorBrightCyan,
					ColorGreen,
				},
			},
			Row{
				RowFragment{
					"BB",
					ColorBrightGreen,
					ColorBlue,
				},
				RowFragment{
					"B",
					ColorBrightCyan,
					ColorGreen,
				},
				RowFragment{
					"B",
					ColorBrightRed,
					ColorGreen,
				},
			},
			Row{
				RowFragment{
					"CCCC",
					ColorBrightGreen,
					ColorBlue,
				},
			},
			Row{
				RowFragment{
					"D",
					ColorBrightGreen,
					ColorBlue,
				},
				RowFragment{
					"D",
					ColorBrightCyan,
					ColorBlue,
				},
				RowFragment{
					"D",
					ColorBrightRed,
					ColorBlue,
				},
				RowFragment{
					"D",
					ColorBrightMagenta,
					ColorBlue,
				},
			},
		},
		Frame{
			Row{
				RowFragment{
					"AA",
					ColorBrightGreen,
					ColorBlue,
				},
				RowFragment{
					"AA",
					ColorBrightCyan,
					ColorGreen,
				},
			},
			Row{
				RowFragment{
					"BB",
					ColorBrightGreen,
					ColorBlue,
				},
				RowFragment{
					"B",
					ColorBrightCyan,
					ColorGreen,
				},
				RowFragment{
					"B",
					ColorBrightRed,
					ColorGreen,
				},
			},
			Row{
				RowFragment{
					"CCCC",
					ColorBrightGreen,
					ColorBlue,
				},
			},
			Row{
				RowFragment{
					"D",
					ColorBrightGreen,
					ColorBlue,
				},
				RowFragment{
					"D",
					ColorBrightCyan,
					ColorBlue,
				},
				RowFragment{
					"D",
					ColorBrightRed,
					ColorBlue,
				},
				RowFragment{
					"D",
					ColorBrightMagenta,
					ColorBlue,
				},
			},
		},
	}
	text := "AAAAaabb1122\nBBBBaabc1122\nCCCCaaaa1111\nDDDDabcd1111\n\nAAAAaabb1122\nBBBBaabc1122\nCCCCaaaa1111\nDDDDabcd1111\n"
	result, err := body_from_string(text, header)
	if err != nil {
		t.Errorf("Error while parcing body %d", err)
	}
	if body.to_string(true) != result.to_string(true) {
		t.Errorf("Body from text convertion incorrect \n%s--\n%s--", body.to_string(true), result.to_string(true))
	}
}
