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

func TestColorMod(t *testing.T) {
	var mods []string = []string{"none", "fg", "bg", "full"}
	for _, mod := range mods {
		var color_mod, err = ColorModFromString(mod)
		if err != nil {
			t.Errorf("ColorMod parcing error str: %s  err: %d", mod, err)
		}
		if color_mod.ToString() != mod {
			t.Errorf("ColorMod parcing error str: %s ColorMod: %s ", mod, color_mod.ToString())
		}
	}
}

func TestColo(t *testing.T) {
	var colors []string = []string{"0","1","2","3","4","5","6","7","8","9","a","b","c","d","e","f"}
	for _, clr := range colors {
		var color, err = color_from_string(clr)
		if err != nil {
			t.Errorf("Color parcing error str: %s  err: %d", clr, err)
		}
		if color.ToString() != clr {
			t.Errorf("Color parcing error str: %s Colo: %s ", clr, color.ToString())
		}
	}
}
