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


func TestHeaderToStringAllParams(t *testing.T) {
	header := Header{
        1,
        2,
        DEFAULT_DELAY+1,
        !DEFAULT_LOOP,
        ColorModFull,
        true,
        123,
        1,
        "1234567",
        "",
        "",
    }
    s_ref := "width 1\nheight 2\ndelay 51\nloop false\ncolors full\nutf8\ndatacols 123\npreview 1\naudio 1234567\n\n"
    if s_ref != header.ToString() {
    	t.Errorf("Header to text convertion incorrect \n%s--\n%s--", s_ref, header.ToString())
    }
}

func TestHeaderToStringDefaultParams(t *testing.T) {
    header := Header{
        1,
        2,
        DEFAULT_DELAY,
        DEFAULT_LOOP,
        DEFAULT_COLORS,
        DEFAULT_UTF8,
        DEFAULT_COLORS.ToDatacols(),
        DEFAULT_PREVIEW,
        "",
        "",
        "",
    }
    s_ref := "width 1\nheight 2\n\n"
    if s_ref != header.ToString() {
    	t.Errorf("Header to text convertion incorrect \n%s--\n%s--", s_ref, header.ToString())
    }
}

func TestHeaderToStringDatacols(t *testing.T) {
    header := Header{
        1,
        2,
        DEFAULT_DELAY,
        DEFAULT_LOOP,
        DEFAULT_COLORS,
        DEFAULT_UTF8,
        DEFAULT_COLORS.ToDatacols()+1,
        DEFAULT_PREVIEW,
        "",
        "",
        "",
    }
    s_ref := "width 1\nheight 2\ndatacols 2\n\n"
    if s_ref != header.ToString() {
    	t.Errorf("Header to text convertion incorrect \n%s--\n%s--", s_ref, header.ToString())
    }
}

func TestHeaderFromStringFull(t *testing.T) {
	s := "width 1\nheight 2\ndelay 3\nloop false\ncolors full\nutf8\ndatacols 5\npreview 1\naudio 12345"
    refernce := Header{
        1,
        2,
        3,
        false,
        ColorModFull,
        true,
        5,
        1,
        "12345",
        "",
        "",
    }
    h, err := HeaderFromString(s)
    if err != nil {
    	t.Errorf("Error while parcing header: %d", err)
    	return
    }
    if *h != refernce {
    	t.Errorf("Header from text convertion incorrect \n%s--\n%s--", s, h.ToString())
    }
}

func TestHeaderFromStringOnlyRequired(t *testing.T) {
	s := "width 1\nheight 2"
    refernce := Header{
        1,
        2,
        DEFAULT_DELAY,
        true,
        ColorModNone,
        false,
        1,
        0,
        "",
        "",
        "",
    }
    h, err := HeaderFromString(s)
    if err != nil {
    	t.Errorf("Error while parcing header: %d", err)
    	return
    }
    if *h != refernce {
    	t.Errorf("Header from text convertion incorrect \n%s--\n%s--", s, h.ToString())
    }
}

func TestHeaderFromStringOptionalIncorrect(t *testing.T) {
	s := "width 1\nheight 2\ndelay safdsfsdf\nloop dsfsdf\ncolors dfdfdf\ndatacols dfsfsddf"
    refernce := Header{
        1,
        2,
        50,
        true,
        ColorModNone,
        false,
        1,
        0,
        "",
        "",
        "",
    }
    h, err := HeaderFromString(s)
    if err != nil {
    	t.Errorf("Error while parcing header: %d", err)
    	return
    }
    if *h != refernce {
    	t.Errorf("Header from text convertion incorrect \n%s--\n%s--", s, h.ToString())
    }
}

func TestHeadeFromStringWidthIncorrect(t *testing.T) {
	s := "width sdfsfsdf\nheight 2\ndelay 3\nloop false\ncolors full\nutf8\ndatacols 5\naudio 12345"
	h, err := HeaderFromString(s)
	if err == nil {
		t.Errorf("Header from text must return error but: \n%s", h.ToString())
	}
}

func TestHeaderFromStringDatacols(t *testing.T) {
    s := "width 1\nheight 2\ncolors full"
    refernce := Header{
        1,
        2,
        50,
        true,
        ColorModFull,
        false,
        3,
        0,
        "",
        "",
        "",
    }
    h, err := HeaderFromString(s)
    if err != nil {
    	t.Errorf("Error while parcing header: %d", err)
    	return
    }
    if *h != refernce {
    	t.Errorf("Header from text convertion incorrect \n%s--\n%s--", s, h.ToString())
    }
    s1 := "width 1\nheight 2\ncolors full\ndatacols 0"
    refernce1 := Header{
        1,
        2,
        50,
        true,
        ColorModFull,
        false,
        0,
        0,
        "",
        "",
        "",
    }
    h1, err1 := HeaderFromString(s1)
    if err1 != nil {
    	t.Errorf("Error while parcing header: %d", err1)
    	return
    }
    if *h1 != refernce1 {
    	t.Errorf("Header from text convertion incorrect \n%s--\n%s--", s1, h1.ToString())
    }
}

func TestHeaderFromStringExtraSpaces(t *testing.T) {
    s := "width    1\nheight    2\ndelay    3\nloop    false\ncolors    full \nutf8   \ndatacols    5\naudio    12345"
    refernce := Header{
        1,
        2,
        3,
        false,
        ColorModFull,
        true,
        5,
        0,
        "12345",
        "",
        "",
    }
    h, err := HeaderFromString(s)
    if err != nil {
    	t.Errorf("Error while parcing header: %d", err)
    	return
    }
    if *h != refernce {
    	t.Errorf("Header from text convertion incorrect \n%s--\n%s--", s, h.ToString())
    }
}
