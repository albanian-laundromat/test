package main

import "fmt"

func main() {
	ruleset := "RRRLLLLR"
	rulelength := len(ruleset)

	var grid [33][33]int

	pos_x, pos_y, vel_x, vel_y := 16, 16, 0, -1

	for range 1000 {
		dir := ruleset[grid[pos_x][pos_y]]
		if dir == 'L' {
			vel_x, vel_y = vel_y, -vel_x
		} else {
			vel_x, vel_y = -vel_y, vel_x
		}
		grid[pos_x][pos_y] = (grid[pos_x][pos_y] + 1) % rulelength
		pos_x += vel_x
		pos_y += vel_y

		if pos_x < 0 || pos_x > 32 || pos_y < 0 || pos_y > 32 {
			break
		}
	}

	outp := ""

	for _, row := range grid {
		if outp != "" {
			outp += "\n"
		}
		for _, cell := range row {
			outp += fmt.Sprint(cell)
		}
	}

	fmt.Println(outp)
}
