package main

import (
	"log"
	"time"
	"github.com/schollz/progressbar/v2"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func Progress(){
	bar := progressbar.NewOptions(uc)
	//bar := progressbar.NewOptions(uc, progressbar.OptionSetRenderBlankState(true))
	//bar.RenderBlank()
	for i := 0; i < uc; i++ {
		bar.Add(1)
	    time.Sleep(1 * time.Second)
	}
}

func Ui(){
	for i := 0; i < uc; i++ {
		g1 := widgets.NewGauge()
		g1.Title = "process bar"
		g1.SetRect(0, 6, 50, 11)
		g1.Percent = i
		g1.BarColor = ui.ColorGreen
		g1.LabelStyle = ui.NewStyle(ui.ColorYellow)
		g1.TitleStyle.Fg = ui.ColorMagenta
		g1.BorderStyle.Fg = ui.ColorWhite
		ui.Render(g1)
		time.Sleep(1 * time.Second)
	}
}

func Qps(){
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	go Ui()
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}