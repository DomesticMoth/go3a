# go3a
This library provides a structural representation of [3a ascii animations format](https://github.com/DomesticMoth/3a) and methods for reading and writing it.  
- [Usage](#usage)
- [Short API description](#short-api-description)
  - [Structs](#structs)
  - [Functions](#functions)
## Usage
Install  
```
$ go get github.com/DomesticMoth/go3a
```
Here's a simple example that parsing a string in 3a format and displaying title, author and body:  
```go
package main

import  (
	"github.com/DomesticMoth/go3a"
	"fmt"
)

func main() {
	text := "width 12\nheight 5\ndelay 300\ncolors fg\ntitle just an apple\nauthor DomesticMoth\n\n\n  ,--./,-.  444444444444\n / //     \\ 444cc4444444\n|          |444444444444\n \\        / 444444444444\n  '._,._,'  444444444444\n\n  ,--./,-.  444444444444\n / //    _\\ 444cc4444444\n|       /   4444444ffff4\n \\      `-, 4444444ffff4\n  '._,._,'  444444444444\n\n  ,--./,-.  444444444444\n / //   ,-' 444cc4444444\n|      (    4444444f4444\n \\      `-, 4444444ffff4\n  '._,._,'  444444444444\n"
	art, _ := go3a.Load(text)
	fmt.Println("title: ", art.Header.Title)
	fmt.Println("author: ", art.Header.Author)
	fmt.Println(art.Body.ToString(true))
}
```
## Short API description
### Structs
The core of the library is the Art struct, which implements the 3a structure:  
```go
type Art struct {
	Header Header
	Body Body
}
```
Header struct contains information about the header of 3a file:  
```go
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
```
Body type is a list of frames, where each frame is a list of rows, and each row is a list of row fragments:  
```go
type Row []RowFragment
type Frame []Row
type Body []Frame
```
Each RowFragment is a set of consecutive symbols with the same values of foreground and background colors:  
```go
type RowFragment struct {
    Text string
    FgColor Color
    BgColor Color
}
```
### Functions
`Load(s string) (*Art, error)` and `Save(art Art, pretify bool) string` functions allow you to convert strings to `Art` and back.  
`LoadFile(path string) (*Art, error)` and `SaveFile(art Art, pretify bool, path string) error` functions allow you to read 3a files to `Art` and write `Art` to 3a files. 
