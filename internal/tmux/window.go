package tmux

import (
	"strconv"
)

const WINDOW_FORMAT = "#{window_id}|#{window_name}|#{window_index}|#{window_height}|#{window_width}|#{window_offset_x}|#{window_offset_y}|#{window_layout}"

type WindowField int

const (
	WindowId WindowField = iota
	WindowName
	WindowIndex
	WindowHeight
	WindowWidth
	WindowOffsetX
	WindowOffsetY
	WindowLayout
)

// Move the window related functions and models here.
type Window struct {
	// window_name
	Name string `json:"name" yaml:"name"`
	// window_id
	Id int64 `json:"id" yaml:"id"`
	// window_index
	Index int64 `json:"index" yaml:"index"`
	// window_height
	Height int64 `json:"height" yaml:"height"`
	// window_width
	Width int64 `json:"width" yaml:"width"`
	// window_offset_x
	Xoff int64 `json:"xoff" yaml:"xoff"`
	// window_offset_y
	Yoff int64 `json:"yoff" yaml:"yoff"`
	// window_layout
	Layout string `json:"layout" yaml:"layout"`
	// Will replace the session_path.
	WorkDir   string   `json:"work_dir" yaml:"work_dir"`
	SessionId int64    `json:"session_id" yaml:"session_id"`
	Session   *Session `json:"session" yaml:"session"`
	// window_panes.
	Panes []Pane `json:"panes" yaml:"panes"`
}

func WindowsFromString(input string) *[]Window {
	var windows []Window

	windowStrs := splitLines(input)

	for _, windowStr := range windowStrs {
		windows = append(windows, windowFromString(windowStr))
	}

	return &windows
}

func windowFromString(input string) Window {
	fields := splitFields(input)

	// Skip the first character of the
	// window id.
	fields[WindowId] = fields[WindowId][1:]

	wid, _ := strconv.ParseInt(fields[WindowId], 10, 64)
	index, _ := strconv.ParseInt(fields[WindowIndex], 10, 64)
	height, _ := strconv.ParseInt(fields[WindowHeight], 10, 64)
	width, _ := strconv.ParseInt(fields[WindowWidth], 10, 64)
	xoff, _ := strconv.ParseInt(fields[WindowOffsetX], 10, 64)
	yoff, _ := strconv.ParseInt(fields[WindowOffsetY], 10, 64)

	return Window{
		Id:     wid,
		Index:  index,
		Height: height,
		Width:  width,
		Xoff:   xoff,
		Yoff:   yoff,
		Name:   fields[WindowName],
		Layout: fields[WindowLayout],
	}
}
