package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"os"
	"strings"
)

var images []*image.Paletted
var delays []int

type direction rune

const (
	up    direction = 't'
	right direction = 'r'
	down  direction = 'b'
	left  direction = 'l'
)

type beamTip struct {
	dir direction
	x   int
	y   int
}

type block struct {
	char  rune
	beams []direction
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {

	bytes, err := os.ReadFile("input.txt")
	check(err)
	input := string(bytes)

	originalMatrix := formatInput(input)
	width := len(originalMatrix[0])
	height := len(originalMatrix)
	bestSum := 0
	var bestTipsConfig []beamTip
	for i := 0; i < width*2+2*height; i++ {
		tips := getTipConfig(i, width, height)
		var matrix [][]*block
		for y, row := range originalMatrix {
			matrix = append(matrix, nil)
			for _, blockRef := range row {
				blockValue := *blockRef
				matrix[y] = append(matrix[y], &blockValue)
			}
		}
		for {
			tips = processBeam(matrix, tips)
			if len(tips) == 0 {
				break
			}
		}
		sum := 0
		for _, row := range matrix {
			for _, theBlock := range row {
				if len(theBlock.beams) > 0 {
					sum += 1
				}
			}
		}
		if sum > bestSum {
			bestSum = sum
			bestTipsConfig = getTipConfig(i, width, height)
		}
	}
	fmt.Println(bestSum)

	gifFile := openGifFile()
	defer closeGifFile(gifFile)
	painNewFrame(originalMatrix, make([]beamTip, 0))
	tips := bestTipsConfig
	for {
		tips = processBeam(originalMatrix, tips)
		painNewFrame(originalMatrix, tips)
		if len(tips) == 0 {
			break
		}
	}
	sum := 0
	for _, row := range originalMatrix {
		for _, theBlock := range row {
			if len(theBlock.beams) > 0 {
				sum += 1
			}
		}
	}
	painNewFrame(originalMatrix, tips)
	encodeImagesToGifFile(gifFile)
}

func formatInput(input string) (matrix [][]*block) {
	lines := strings.Split(input, "\n")

	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		matrix = append(matrix, nil)
		for _, char := range line {
			matrix[y] = append(matrix[y], &block{
				char: char,
			})
		}
	}

	return
}

func openGifFile() *os.File {
	file, err := os.OpenFile("gold_view.gif", os.O_WRONLY|os.O_CREATE, 0644)
	check(err)
	return file
}
func painNewFrame(matrix [][]*block, tips []beamTip) {
	var w, h = len(matrix[0]) * 4, len(matrix) * 4
	img := image.NewPaletted(image.Rect(0, 0, w, h), palette.Plan9)
	images = append(images, img)
	delays = append(delays, 0)

	for y, row := range matrix {
		for x, theBlock := range row {
			switch theBlock.char {
			case '.':
				if len(theBlock.beams) > 0 {
					for i := 0; i < 4; i++ {
						addX := i % 2
						addY := i / 2
						img.Set(x*4+1+addX, y*4+1+addY, color.RGBA{0, 255, 0, 255})
					}
				}
			case '\\', '/':
				theColor := color.RGBA{0, 0, 255, 255}
				if theBlock.char == '/' {
					for i := 0; i < 4; i++ {
						img.Set(x*4+(3-i), y*4+i, theColor)
					}
				} else {
					for i := 0; i < 4; i++ {
						img.Set(x*4+i, y*4+i, theColor)
					}
				}
			case '|', '-':
				theColor := color.RGBA{255, 0, 0, 255}
				theColorFade := color.RGBA{255, 123, 123, 255}
				if theBlock.char == '|' {
					for i := 0; i < 4; i++ {
						lColor := theColor
						if i == 0 || i == 3 {
							lColor = theColorFade
						}
						img.Set(x*4+1, y*4+i, lColor)
						img.Set(x*4+2, y*4+i, lColor)
					}
				} else {
					for i := 0; i < 4; i++ {
						lColor := theColor
						if i == 0 || i == 3 {
							lColor = theColorFade
						}
						img.Set(x*4+i, y*4+1, lColor)
						img.Set(x*4+i, y*4+2, lColor)
					}
				}
			}
		}
	}

	tipColor := color.RGBA{0, 255, 255, 255}
	blackColor := color.RGBA{0, 0, 0, 255}
	for _, theBeamTip := range tips {
		tip := ""
		switch theBeamTip.dir {
		case up:
			tip = "" +
				" @@ " +
				"@@@@" +
				"@  @" +
				"    "
		case right:
			tip = "" +
				" @@ " +
				"  @@" +
				"  @@" +
				" @@ "
		case down:
			tip = "" +
				"    " +
				"@  @" +
				"@@@@" +
				" @@ "
		case left:
			tip = "" +
				" @@ " +
				"@@  " +
				"@@  " +
				" @@ "

		}

		for i := 0; i < 16; i++ {
			x := i % 4
			y := i / 4
			if tip[i] == ' ' {
				img.Set(theBeamTip.x*4+x, theBeamTip.y*4+y, blackColor)
			} else {
				img.Set(theBeamTip.x*4+x, theBeamTip.y*4+y, tipColor)
			}
		}
	}
}
func encodeImagesToGifFile(file *os.File) {
	err := gif.EncodeAll(file, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	check(err)
}
func closeGifFile(file *os.File) {
	err := file.Close()
	check(err)
}

func processBeam(matrix [][]*block, tips []beamTip) (newTips []beamTip) {
	maxX := len(matrix[0]) - 1
	maxY := len(matrix) - 1
	for _, tip := range tips {
		theBlock := matrix[tip.y][tip.x]
		skip := false
		for _, beam := range theBlock.beams {
			if beam == tip.dir {
				skip = true
				continue
			}
		}
		if skip {
			continue
		}
		theBlock.beams = append(theBlock.beams, tip.dir)
		if theBlock.char == '.' {
			x, y := move(tip.x, tip.y, tip.dir)
			if x >= 0 && x <= maxX && y >= 0 && y <= maxY {
				newTips = append(newTips, beamTip{
					dir: tip.dir,
					x:   x,
					y:   y,
				})
			}
		} else if theBlock.char == '/' {
			var newDirection direction
			switch tip.dir {
			case up:
				newDirection = right
			case right:
				newDirection = up
			case down:
				newDirection = left
			case left:
				newDirection = down
			}
			x, y := move(tip.x, tip.y, newDirection)
			if x >= 0 && x <= maxX && y >= 0 && y <= maxY {
				newTips = append(newTips, beamTip{
					dir: newDirection,
					x:   x,
					y:   y,
				})
			}
		} else if theBlock.char == '\\' {
			var newDirection direction
			switch tip.dir {
			case up:
				newDirection = left
			case right:
				newDirection = down
			case down:
				newDirection = right
			case left:
				newDirection = up
			}
			x, y := move(tip.x, tip.y, newDirection)

			if x >= 0 && x <= maxX && y >= 0 && y <= maxY {
				newTips = append(newTips, beamTip{
					dir: newDirection,
					x:   x,
					y:   y,
				})
			}
		} else if theBlock.char == '|' {
			last := false
			for {
				var newDirection direction
				if last {
					newDirection = up
				} else {
					newDirection = down
				}
				if !isOpposite(newDirection, tip.dir) {
					x, y := move(tip.x, tip.y, newDirection)
					if x >= 0 && x <= maxX && y >= 0 && y <= maxY {
						newTips = append(newTips, beamTip{
							dir: newDirection,
							x:   x,
							y:   y,
						})
					}
				}
				if last {
					break
				} else {
					last = true
				}
			}
		} else if theBlock.char == '-' {
			last := false
			for {
				var newDirection direction
				if last {
					newDirection = right
				} else {
					newDirection = left
				}
				if !isOpposite(newDirection, tip.dir) {
					x, y := move(tip.x, tip.y, newDirection)
					if x >= 0 && x <= maxX && y >= 0 && y <= maxY {
						newTips = append(newTips, beamTip{
							dir: newDirection,
							x:   x,
							y:   y,
						})
					}
				}
				if last {
					break
				} else {
					last = true
				}
			}
		}
	}
	return newTips
}

func move(x int, y int, dir direction) (_x int, _y int) {
	switch dir {
	case up:
		y--
	case right:
		x++
	case down:
		y++
	case left:
		x--
	}
	return x, y
}

func isOpposite(dir1 direction, dir2 direction) bool {
	return dir1 == up && dir2 == down ||
		dir1 == down && dir2 == up ||
		dir1 == left && dir2 == right ||
		dir1 == right && dir2 == left
}

func getTipConfig(index int, width int, height int) []beamTip {
	// []beamTip{{dir: right, x: 0, y: 0}}
	if index < width {
		i := index
		return []beamTip{{dir: down, x: i, y: 0}}
	}
	if index < width+height {
		i := index - width
		return []beamTip{{dir: left, x: width - 1, y: i}}
	}
	if index < 2*width+height {
		i := index - (width + height)
		return []beamTip{{dir: up, x: width - 1 - i, y: height - 1}}
	}
	if index < 2*width+2*height {
		i := index - (2*width + height)
		return []beamTip{{dir: right, x: 0, y: height - 1 - i}}
	}
	panic("Overflow index")
}
