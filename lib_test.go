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

const (
	WITH_COMMENTS =    `	0000
aaaa
bbbb
	
	cccc

dddd
`
	WITHOUT_COMMENTS = `aaaa
bbbb

dddd
`
	
)

func TestRemoveComments(t *testing.T) {
	if WITHOUT_COMMENTS != remove_comments(WITH_COMMENTS) {
		t.Errorf("Comments removing incorrect \n%s--\n%s--", WITHOUT_COMMENTS, remove_comments(WITH_COMMENTS))
	}
}
