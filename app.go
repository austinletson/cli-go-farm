package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	info := widgets.NewParagraph()
	var world = [][]CellType{
		{WATER, GRASS},
		{WATER, WATER},
	}

	drawCellsOnWorld(world, 1, 1)
	p.Text = getWorldString()
	p.Title = "World"
	p.SetRect(0, 0, worldX+2, worldY+2)

	info.Text = getCellInfo(GRASS)
	info.SetRect(worldX+5, 0, 30, 3)
	ui.Render(p, info)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			if e.ID == "j" {
				moveCellPointer(SOUTH)
			} else if e.ID == "k" {
				moveCellPointer(NORTH)
			} else if e.ID == "l" {
				moveCellPointer(EAST)
			} else if e.ID == "h" {
				moveCellPointer(WEST)
			} else {
				break
			}

			p.Text = getWorldString()
			ui.Render(p)
		}
	}
}
