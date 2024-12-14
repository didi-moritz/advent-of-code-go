package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

type v struct {
	x, y int
}

type robot struct {
	p, v v
}

func main() {
	part, realData := utils.GetRunConfig(2, true)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2024, 14, realData))

	var robots []robot

	for _, line := range data {
		var px, py, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		robots = append(robots, robot{v{px, py}, v{vx, vy}})
	}

	var result int
	if part == 1 {
		width := 11
		height := 7
		if realData {
			width = 101
			height = 103
		}
		result = part1(robots, width, height)
	} else {
		width := 11
		height := 7
		if realData {
			width = 101
			height = 103
		}
		result = part2(robots, width, height)
	}

	fmt.Println(result)
}

func part1(robots []robot, width int, height int) int {
	mx := width / 2
	my := height / 2

	score := make([]int, 4)

	for _, r := range robots {
		x := (r.p.x + r.v.x*100) % width
		y := (r.p.y + r.v.y*100) % height

		if x < 0 {
			x += width
		}

		if y < 0 {
			y += height
		}

		if x == mx || y == my {
			continue
		}

		i := 0
		if x > mx {
			i++
		}
		if y > my {
			i += 2
		}

		score[i]++
	}

	result := 1
	for _, s := range score {
		result *= s
	}

	return result
}

func part2(robots []robot, width int, height int) int {
	m := make([][]bool, height)
	for y := range height {
		m[y] = make([]bool, width)
	}

	for number := 0; number < 1000000000000; number++ {
		if number%100000 == 0 {
			fmt.Println(number)
		}

		for iy := range height {
			for ix := range width {
				m[iy][ix] = false
			}
		}

		for _, r := range robots {
			x := (r.p.x + r.v.x*number) % width
			y := (r.p.y + r.v.y*number) % height

			if x < 0 {
				x += width
			}

			if y < 0 {
				y += height
			}

			m[y][x] = true
		}

		show := true

		show = false
		bs := 3
		for iy := 0; !show && iy < height-bs+1; iy += bs {
			for ix := 0; !show && ix < width-bs+1; ix += bs {
				bc := true
				for cy := iy; bc && cy < iy+bs; cy++ {
					for cx := ix; bc && cx < ix+bs; cx++ {
						if !m[cy][cx] {
							bc = false
						}
					}
				}
				if bc {
					show = true
					break
				}
			}
		}

		if show {
			fmt.Print("\033[H\033[2J")
			for iy := 0; iy < height; iy += 2 {
				for ix := 0; ix < width; ix += 2 {
					c1 := m[iy][ix]
					c2 := ix < width-1 && m[iy][ix+1]
					c3 := iy < height-1 && m[iy+1][ix]
					c4 := ix < width-1 && iy < height-1 && m[iy+1][ix+1]

					if !c1 && !c2 && !c3 && !c4 {
						fmt.Print(" ")
					} else if c1 && !c2 && !c3 && !c4 {
						fmt.Print("▘")
					} else if !c1 && c2 && !c3 && !c4 {
						fmt.Print("▝")
					} else if !c1 && !c2 && c3 && !c4 {
						fmt.Print("▖")
					} else if !c1 && !c2 && !c3 && c4 {
						fmt.Print("▗")
					} else if c1 && c2 && !c3 && !c4 {
						fmt.Print("▀")
					} else if !c1 && !c2 && c3 && c4 {
						fmt.Print("▄")
					} else if c1 && !c2 && c3 && !c4 {
						fmt.Print("▌")
					} else if !c1 && c2 && !c3 && c4 {
						fmt.Print("▐")
					} else if c1 && !c2 && !c3 && c4 {
						fmt.Print("▚")
					} else if !c1 && c2 && c3 && !c4 {
						fmt.Print("▞")
					} else if c1 && c2 && !c3 && c4 {
						fmt.Print("▜")
					} else if c1 && c2 && c3 && !c4 {
						fmt.Print("▛")
					} else if c1 && !c2 && c3 && c4 {
						fmt.Print("▙")
					} else if !c1 && c2 && c3 && c4 {
						fmt.Print("▟")
					} else if c1 && c2 && c3 && c4 {
						fmt.Print("█")
					} else {
						fmt.Println(c1, c2, c3, c4)
						return 666
					}
				}
				fmt.Println()
			}

			return number
		}
	}

	return 0
}
