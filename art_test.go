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


func TestSaveLoad(t *testing.T) {
	text := "width 12\nheight 5\ndelay 300\ncolors fg\ntitle just an apple\nauthor DomesticMoth\n\n\n  ,--./,-.  444444444444\n / //     \\ 444cc4444444\n|          |444444444444\n \\        / 444444444444\n  '._,._,'  444444444444\n\n  ,--./,-.  444444444444\n / //    _\\ 444cc4444444\n|       /   4444444ffff4\n \\      `-, 4444444ffff4\n  '._,._,'  444444444444\n\n  ,--./,-.  444444444444\n / //   ,-' 444cc4444444\n|      (    4444444f4444\n \\      `-, 4444444ffff4\n  '._,._,'  444444444444\n"
	art, err := load(text)
	if err != nil {
		t.Errorf("Error while loading art %d", err)
		return
	}
	text2 := save(*art, true)
	if text != text2 {
		t.Errorf("Art load/save incorrect \n%s--\n%s--", text, text2)
	}
}

