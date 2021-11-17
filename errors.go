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


import "fmt"

type UnknownColor struct {
	color string
}

func (e UnknownColor) Error() string {
    return fmt.Sprintf("UnknownColor: %s", e.color)
}

type UnknownColorMod struct {
	color_mod string
}

func (e UnknownColorMod) Error() string {
    return fmt.Sprintf("UnknownColorMod: %s", e.color_mod)
}

type ThereIsNoBody struct {}

func (e ThereIsNoBody) Error() string {
    return fmt.Sprintf("ThereIsNoBody")
}

type InvalidWidth struct {}

func (e InvalidWidth) Error() string {
    return fmt.Sprintf("InvalidWidth")
}

type InvalidHeight struct {}

func (e InvalidHeight) Error() string {
    return fmt.Sprintf("InvalidHeight")
}
