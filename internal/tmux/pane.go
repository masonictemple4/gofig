package tmux

import "strconv"

type PaneField int

const PANE_FORMAT = "#{pane_id}|#{pane_index}|#{pane_title}|#{pane_height}|#{pane_width}|#{pane_start_path}|#{pane_layout}"

const (
	PaneId PaneField = iota
	PaneIndex
	PaneTitle
	PaneHeight
	PaneWidth
	PaneStartPath
	PaneLayout
)

type Pane struct {
	// pane_id
	Id int64 `json:"id" yaml:"id"`
	// pane_index
	Index int64 `json:"index" yaml:"index"`
	// pane_title
	Title string `json:"title" yaml:"title"`
	// pane_height
	Height int64 `json:"height" yaml:"height"`
	// pane_width
	Width int64 `json:"width" yaml:"width"`
	// pane_start_path: When creating with this value
	// or
	// pane_path: When reading this value/setting it.
	WorkdDir string `json:"work_dir" yaml:"work_dir"`
	Layout   string `json:"layout" yaml:"layout"`
}

func PanesFromString(input string) *[]Pane {
	var panes []Pane

	paneStrs := splitLines(input)

	for _, paneStr := range paneStrs {
		panes = append(panes, paneFromString(paneStr))
	}

	return &panes
}

func paneFromString(input string) Pane {

	paneParts := splitFields(input)

	// Skip leading character in the pane id.
	paneParts[PaneId] = paneParts[PaneId][1:]

	pid, _ := strconv.ParseInt(paneParts[PaneId], 10, 64)
	pIndex, _ := strconv.ParseInt(paneParts[PaneIndex], 10, 64)
	height, _ := strconv.ParseInt(paneParts[PaneHeight], 10, 64)
	width, _ := strconv.ParseInt(paneParts[PaneWidth], 10, 64)

	return Pane{
		Id:       pid,
		Index:    pIndex,
		Height:   height,
		Width:    width,
		Title:    paneParts[PaneTitle],
		WorkdDir: paneParts[PaneStartPath],
	}
}
