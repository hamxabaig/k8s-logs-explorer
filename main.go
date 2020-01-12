package main

import (
	"github.com/hamxabaig/k8s-pod-logs-explorer/adapters"
	"github.com/rivo/tview"
)

func main() {
	// spinner := utils.NewSpinner("Please wait while we load pods...")
	// spinner.Start()
	adapter := adapters.New()
	pods := adapter.GetPods()
	// spinner.Stop()

	app := tview.NewApplication()
	text := tview.NewTextView().
		SetDynamicColors(false).
		SetWrap(true).
		SetWordWrap(true).
		SetText("How are you?")

	list := tview.NewList()
	for idx, pod := range pods {
		list = list.InsertItem(idx, pod, "", 0, func(podName string) func() {
			return func() {
				log := adapter.GetLogs(podName)
				text.SetText(tview.TranslateANSI(log)).SetWrap(true).SetWordWrap(true)
			}
		}(pod))
	}
	list.AddItem("Quit", "Press to exit", 'q', func() {
	})

	grid := tview.NewGrid().
		SetRows(2).
		SetColumns(30, 30).
		SetBorders(true)

	grid.AddItem(list, 0, 0, 50, 1, 0, 0, true)
	grid.AddItem(text, 0, 1, 60, 1, 0, 0, false)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
