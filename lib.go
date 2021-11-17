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
	"regexp"
	"strconv"
	"strings"
	"os"
	"bytes"
)

func remove_comments(s string) string {
	var r1 = regexp.MustCompile("(?m)^\t.*?(\n|$)")
	var r2 = regexp.MustCompile("\t.*?(\n|$)")
	s = r1.ReplaceAllString(s, "")
	s = r2.ReplaceAllString(s, "\n")
	return s
}

func only_payLoad(v []string) []string {
	var ret []string
	for _, e := range v {
		if e != "" {
			ret = append(ret, e)
		}
	}
	return ret
}

func generate_color_fragment(color Color, count int) string {
	ret := ""
	var i int = 0
	for ; i < count; i++ {
		ret += color.ToString()
	}
	return ret
}

type RowFragment struct {
    Text string
    FgColor Color
    BgColor Color
}

type Row []RowFragment

type Frame []Row

type Header struct {
	Width uint16
	Height uint16
	Delay uint16
	LoopEnable bool
	ColorMod ColorMod
	Utf8 bool
	Datacols uint16
	Preview uint16
	Audio string
	Title string
	Author string
}

func (header Header) ToString() string {
	ret := ""
	ret += "width "
	ret += strconv.Itoa(int(header.Width))
	ret += "\nheight "
	ret += strconv.Itoa(int(header.Height))
	if header.Delay != DEFAULT_DELAY {
		ret += "\ndelay "
		ret += strconv.Itoa(int(header.Delay))
	}
	if header.LoopEnable != DEFAULT_LOOP {
		ret += "\nloop "
		if header.LoopEnable {
			 ret += "true"
		}else{
			ret += "false"
		}
	}
	if header.ColorMod != DEFAULT_COLORS {
		ret += "\ncolors "
		ret += header.ColorMod.ToString()
	}
	if header.Utf8 {
		ret += "\nutf8"
	}
	if header.ColorMod.ToDatacols() != header.Datacols {
		ret += "\ndatacols "
		ret += strconv.Itoa(int(header.Datacols))
	}
	if header.Preview != DEFAULT_PREVIEW {
		ret += "\npreview "
		ret += strconv.Itoa(int(header.Preview))
	}
	if header.Audio != "" {
		ret += "\naudio "
		ret += header.Audio
	}
	if header.Title != "" {
		ret += "\ntitle "
		ret += header.Title
	}
	if header.Author != "" {
		ret += "\nauthor "
		ret += header.Author
	}
	ret += "\n\n"
	return ret
}

func HeaderFromString(s string) (*Header, error) {
	var (
		width uint16 = 0
		width_set bool = false
		height uint16 = 0
		height_set bool = false
		delay uint16 = DEFAULT_DELAY
		loop_enable bool = DEFAULT_LOOP
		color_mod ColorMod = DEFAULT_COLORS
		utf8 bool = false
		datacols uint16 = 0
		datacols_set bool = false
		preview uint16 = DEFAULT_PREVIEW
		audio string = ""
		title string = ""
		author string = ""
	)
	rows := strings.Split(s, "\n")
	for _, row := range rows {
		tokens := strings.Split(row, " ")
		tokens = only_payLoad(tokens)
		if tokens[0] == "utf8" {
			utf8 = true
		}else{
			if len(tokens) < 2 {
				continue
			}
			switch tokens[0] {
				case "width": {
					var w, err = strconv.ParseInt(tokens[1], 0, 16)
					if err == nil {
						width =  uint16(w)
						width_set = true
					}
				}
				case "height": {
					var h, err = strconv.ParseInt(tokens[1], 0, 16)
					if err == nil {
						height =  uint16(h)
						height_set = true
					}
				}
				case "delay": {
					var d, err = strconv.ParseInt(tokens[1], 0, 16)
					if err == nil {
						delay = uint16(d)
					}
				}
				case  "loop": {
					if tokens[1] == "true" {
						loop_enable = true
					}else if tokens[1] == "false" {
						loop_enable = false
					}
				}
				case "colors": {
					var cm, err = ColorModFromString(tokens[1])
					if err == nil {
						color_mod = cm
					}
				}
				case "datacols": {
					var d, err = strconv.ParseInt(tokens[1], 0, 16)
					if err == nil {
						datacols = uint16(d)
						datacols_set = true
					}
				}
				case "preview": {
					var p, err = strconv.ParseInt(tokens[1], 0, 16)
					if err == nil {
						preview = uint16(p)
					}
				}
				case "audio": {
					if tokens[1] != "" {
						audio = tokens[1]
					}
				}
				case "title": {
					s := ""
					for i, token := range tokens {
						if i > 0 {
							if i > 1 {
								s += " "
							}
							s += token
						}
					}
					if s != "" {
						title = s
					}
				}
				case "author": {
					s := ""
					for i, token := range tokens {
						if i > 0 {
							if i > 1 {
								s += " "
							}
							s += token
						}
					}
					if s != "" {
						author = s
					}
				}
			}
		}
	}
	if !width_set {
		return nil, InvalidWidth{}
	}
	if !height_set {
		return nil, InvalidHeight{}
	}
	if !datacols_set {
		datacols = color_mod.ToDatacols()
	}
	ret := new(Header)
	*ret = Header{width, height, delay, loop_enable, color_mod, utf8, datacols, preview, audio, title, author}
	return ret, nil
}

type Body []Frame

func (frames Body) ToString(pretify bool) string {
	ret := ""
	for frm, frame := range frames {
		for _, row := range frame{
			text_col := ""
			color1_col := ""
			color2_col := ""
			for _, fragment := range row {
				text_col += fragment.Text
				color1_col += generate_color_fragment(fragment.FgColor, len(fragment.Text))
				color2_col += generate_color_fragment(fragment.BgColor, len(fragment.Text))
			}
			ret += text_col
			ret += color1_col
			ret += color2_col
			if pretify {
				ret += "\n"
			}
		}
		if frm < len(frames)-1 {
			ret += "\n"
		}
	}
	return ret
}

func BodyFromString(s string, h Header) (Body, error) {
	r := regexp.MustCompile("(\n|\t)")
	s = r.ReplaceAllString(s, "")
	char_vec := []rune(s)
	length := uint16(len(char_vec))
	var frm uint16 = 0
	width := h.Width
	height := h.Height
	datacols := h.Datacols
	frames := []Frame{}
	nxt := true
	brk := false
	for nxt {
		frame := []Row{}
		var y uint16 = 0
		for ; y < height; y++ {
			var row Row
			row_fragment := RowFragment{"", NoColor, NoColor}
			var x uint16 = 0
			for ; x < width; x++ {
				symbol_pos := (frm*width*datacols*height)+(y*width*datacols)+x
				if symbol_pos >= length {
					break
				}
				symbol := char_vec[symbol_pos]
				fg_color := NoColor
				bg_color := NoColor
				if h.ColorMod == ColorModFg {
					fg_color_position := (frm*width*datacols*height)+(y*width*datacols)+width+x
					if fg_color_position >= length {
						nxt = false
						break
					}
					var err error
					fg_color, err = color_from_string(string(char_vec[fg_color_position]))
					if err != nil{
						return nil, err
					}
				}else if h.ColorMod == ColorModBg {
					bg_color_position := (frm*width*datacols*height)+(y*width*datacols)+width+x
					if bg_color_position >= length {
						nxt = false
						break
					}
					var err error
					bg_color, err = color_from_string(string(char_vec[bg_color_position]))
					if err != nil{
						return nil, err
					}
				}else if h.ColorMod == ColorModFull {
					fg_color_position := (frm*width*datacols*height)+(y*width*datacols)+width+x
					bg_color_position := (frm*width*datacols*height)+(y*width*datacols)+width*2+x
					if fg_color_position >= length || bg_color_position >= length {
						nxt = false
						break
					}
					var err error
					fg_color, err = color_from_string(string(char_vec[fg_color_position]))
					bg_color, err = color_from_string(string(char_vec[bg_color_position]))
					if err != nil{
						return nil, err
					}
				}
				if x == 0 {
					row_fragment.FgColor = fg_color
					row_fragment.BgColor = bg_color
				}else if row_fragment.FgColor != fg_color || row_fragment.BgColor != bg_color {
						row = append(row, row_fragment)
						row_fragment = RowFragment{
							string(symbol),
							fg_color,
							bg_color,
						}
						continue
				}
				row_fragment.Text += string(symbol)
			}
			if len(row_fragment.Text) > 0 {
				row = append(row, row_fragment)
			}
			if len(row) < 1 {
				brk = true
				break
			}
			frame = append(frame, row)
		}
		if brk {
			break
		}
		frames = append(frames, frame)
		frm += 1
	}
	return frames, nil
}

type Art struct {
	Header Header
	Body Body
}

func Load(s string) (*Art, error) {
	fragments := strings.SplitN(s, "\n\n", 2)
	if len(fragments) < 2 {
		return nil, ThereIsNoBody{}
	}
	header_string := fragments[0]
	body_string := fragments[1]
	header, h_err := HeaderFromString(header_string)
	if h_err != nil {
		return nil, h_err
	}
	body, b_err := BodyFromString(body_string, *header)
	if b_err != nil {
		return nil, b_err
	}
	ret := new(Art)
	*ret = Art{*header, body}
	return ret, nil
}

func Save(art Art, pretify bool) string {
	ret := ""
	ret += art.Header.ToString()
	ret += "\n"
	ret += art.Body.ToString(pretify)
	return ret
}

func LoadFile(path string) (*Art, error) {
	file, err := os.Open("hello.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	text := buf.String()
	return Load(text)
}

func SaveFile(art Art, pretify bool, path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write([]byte(Save(art, pretify)))
	return nil
}
