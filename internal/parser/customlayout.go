// TODO: Incomplete need to wrap up the parser dependencies.
//
// The parser package is an implementation of the tmux custom-layout parser.
//
// Most of the code was directly inspired by the tmux source code here:
// https://github.com/tmux/tmux
//
// Primarily the https://github.com/tmux/tmux/blob/master/layout-custom.c file
// and any dependencies.
// IMPROVEMENTS:
// - Go in and double check the types in tmux.h to make sure we have everything
// that we need.
// - Revisit the datastructures I especially around the LinkedList and
// using a slice instead of tailq.
// - Confirm double pointers are working how we expect.

package parser

import (
	"errors"
	"fmt"
	"strings"
)

type LayoutCellType int

var (
	ErrorInvalidLayout = errors.New("invalid layout")
)

const (
	LayoutWindowPane LayoutCellType = iota
	LayoutLeftRight
	LayoutTopBottom
)

const (
	DEFAULT_XPIXEL = 16
	DEFAULT_YPIXEL = 32
)

type LayoutCell struct {
	Type   LayoutCellType
	Sx     uint        // size in the x-direction (width)
	Sy     uint        // size in the y-direction (height)
	Xoff   uint        // x offset from the parent
	Yoff   uint        // y offset from the parent
	PaneID uint        // ID of the pane, if this cell is a pane
	Wp     *WindowPane // Pane, if this cell is a pane
	Cells  []*LayoutCell
	Parent *LayoutCell
}

type WindowPane struct {
	ID uint
	Lc *LayoutCell
	// Other pane properties
}

type WindowPaneIterator struct {
	panes   []*WindowPane
	current int64
}

func NewWindowPaneIterator(panes []*WindowPane) *WindowPaneIterator {
	return &WindowPaneIterator{
		panes:   panes,
		current: -1,
	}
}

func (it *WindowPaneIterator) Next() *WindowPane {
	it.current++
	if it.current >= int64(len(it.panes)) {
		return nil
	}
	return it.panes[it.current]
}

type Window struct {
	Root  *LayoutCell
	Panes []*WindowPane
}

func layoutChecksum(layoutStr string) uint16 {
	var csum uint16
	for _, c := range layoutStr {
		csum = (csum >> 1) + ((csum & 1) << 15)
		csum += uint16(c)
	}
	return csum
}

func ParseTmuxWindowLayoutString(w *Window, layoutStr string) (*Window, error) {
	var nPanes, nCells, sx, sy uint
	var csum uint16

	n, err := fmt.Sscanf(layoutStr, "%hx,", &csum)
	if n != 1 {
		return nil, ErrorInvalidLayout
	}
	if err != nil {
		return nil, err
	}

	// advance the string 5 characters.
	layoutStr = layoutStr[5:]
	if csum != layoutChecksum(layoutStr) {
		return nil, ErrorInvalidLayout
	}

	// Build the layout
	lc, err := layoutConstruct(nil, &layoutStr)
	if err != nil {
		return nil, err
	}

	if lc == nil {
		return nil, ErrorInvalidLayout
	}

	if len(layoutStr) > 0 {
		return nil, ErrorInvalidLayout
	}

	// Check this window will fit into the layout
	for {
		nPanes = uint(len(w.Panes))
		nCells = uint(len(lc.Cells))

		if len(w.Panes) > len(lc.Cells) {
			return nil, fmt.Errorf("have %d panes but need %d", nPanes, nCells)
		}

		if nPanes == nCells {
			break
		}

		lcchild := findBottomRight(lc)

		// destroy cell
		destroyCell(w, lcchild, &lc)
	}

	switch lc.Type {
	case LayoutWindowPane:
		break
	case LayoutLeftRight:
		for _, lcchild := range lc.Cells {
			sy = lcchild.Sy + 1
			sx += lcchild.Sx + 1
		}
	case LayoutTopBottom:
		for _, lcchild := range lc.Cells {
			sx = lcchild.Sx + 1
			sy += lcchild.Sy + 1
		}
	}

	if lc.Type == LayoutWindowPane && (lc.Sx != sx || lc.Sy != sy) {
		fmt.Printf("fix layout %d,%d to %d,%d", lc.Sx, lc.Sy, sx, sy)
		lc.Sx = sx - 1
		lc.Sy = sy - 1
	}

	// layout check
	if !layoutCheck(lc) {
		return nil, errors.New("size mismatch after applying layout")
	}

	// resize to the layout size
	resizeWindow(w, lc.Sx, lc.Sy, -1, -1)

	// destory old layout and swap to the new
	layoutFreeCell(w.Root)
	w.Root = lc

	// assign the panes into the cells
	wpIter := NewWindowPaneIterator(w.Panes)

	layoutAssign(wpIter, lc)

	// Update pane offsets and sizes
	fixLayoutOffsets(w)
	// TODO: Finish implementation.
	fixLayoutPanes(w, nil)

	// TODO: after finishing the fixLayoutPanes implementation and it's child implementations
	// - recalculate the pane offsets and sizes

	// TODO: update the window model so we can accurately log the resize
	return nil, nil
}

func fixLayoutPanes(w *Window, skip *WindowPane) {
	// 	var wp *WindowPane
	// 	var lc *LayoutCell
	// 	var status int64
	//
	// TODO: setup window options.

}

// update cell offsets based on their sizes
func fixLayoutOffsets(w *Window) {
	lc := w.Root

	lc.Xoff = 0
	lc.Yoff = 0

	// Need to do something with the children..
	fixLayoutOffsetsDFS(lc)

}

func fixLayoutOffsetsDFS(lc *LayoutCell) {
	var lcchild *LayoutCell
	if lc.Type == LayoutLeftRight {
		xoff := lc.Xoff
		for _, lcchild = range lc.Cells {
			lcchild.Xoff = xoff
			lcchild.Yoff = lc.Yoff
			if lcchild.Type != LayoutWindowPane {
				fixLayoutOffsetsDFS(lcchild)
			}
			xoff += lcchild.Sx + 1
		}
	} else {
		yoff := lc.Yoff
		for _, lcchild = range lc.Cells {
			lcchild.Xoff = lc.Xoff
			lcchild.Yoff = yoff
			if lcchild.Type != LayoutWindowPane {
				fixLayoutOffsetsDFS(lcchild)
			}
			yoff += lcchild.Sy + 1
		}
	}
}

func layoutMakeLeaf(lc *LayoutCell, wp *WindowPane) {
	if lc != nil && wp != nil {
		lc.Type = LayoutWindowPane
		lc.Wp = wp
		wp.Lc = lc
	}
}

func layoutAssign(it *WindowPaneIterator, lc *LayoutCell) {
	if lc == nil {
		return
	}

	switch lc.Type {
	case LayoutWindowPane:
		layoutMakeLeaf(lc, it.Next())
	case LayoutLeftRight, LayoutTopBottom:
		for _, lcchild := range lc.Cells {
			layoutAssign(it, lcchild)
		}
	}
}

func resizeWindow(w *Window, sx, sy uint, xpixel, ypixel int) {

	if xpixel == 0 {
		xpixel = DEFAULT_XPIXEL
	}

	if ypixel == 0 {
		ypixel = DEFAULT_YPIXEL
	}

	// TODO: Update the window model so we can accurately log the resize
	// replcing it with either the gofig window or adding an id/name to
	// the parser window definition.
	// fmt.Printf("%s: @%d resize %dx%d (%dx%d)", w.Root.

	// TODO: Figure out if this implementation works how I expect it to.
	// Otherwise, update the window model.
	w.Root.Sx = sx
	w.Root.Sy = sy

}

func layoutCheck(lc *LayoutCell) bool {
	var n uint

	switch lc.Type {
	case LayoutWindowPane:
		break
	case LayoutLeftRight:
		for _, lcchild := range lc.Cells {
			if lcchild.Sy != lc.Sy {
				return false
			}
			if !layoutCheck(lcchild) {
				return false
			}
			n += lcchild.Sx + 1
		}
		if n-1 != lc.Sx {
			return false
		}
	case LayoutTopBottom:
		for _, lcchild := range lc.Cells {
			if lcchild.Sx != lc.Sx {
				return false
			}
			if !layoutCheck(lcchild) {
				return false
			}
			n += lcchild.Sy + 1
		}
		if n-1 != lc.Sy {
			return false
		}
	}

	return true
}

func layoutFreeCell(lc *LayoutCell) {
	var lcchild *LayoutCell

	switch lc.Type {
	case LayoutLeftRight:
	case LayoutTopBottom:
		for len(lc.Cells) > 0 {
			lcchild = lc.Cells[0]
			lc.Cells = lc.Cells[1:]
			layoutFreeCell(lcchild)
		}
	case LayoutWindowPane:
		if lc.Wp != nil {
			lc.Wp.Lc = nil
		}
	}
}

func destroyCell(w *Window, lc *LayoutCell, lcroot **LayoutCell) {
	// If no parent, this is the last pane so window close is imminent and
	// there is no need to resize anything.
	lcparent := lc.Parent
	if lcparent == nil {
		*lcroot = nil
		return
	}

	// Find the position of lc in the parent's cells.
	idx := -1
	for i, cell := range lcparent.Cells {
		if cell == lc {
			idx = i
			break
		}
	}
	if idx == -1 {
		// Cell not found, return error or handle it.
		return
	}

	// Merge the space into the previous or next cell.
	var lcother *LayoutCell
	if idx == 0 {
		lcother = lcparent.Cells[1]
	} else {
		lcother = lcparent.Cells[idx-1]
	}

	if lcother != nil && lcparent.Type == LayoutLeftRight {
		lcother.Sx += lc.Sx + 1 // Adjust the size.
	} else if lcother != nil {
		lcother.Sy += lc.Sy + 1 // Adjust the size.
	}

	// Remove lc from the parent's list of cells.
	lcparent.Cells = append(lcparent.Cells[:idx], lcparent.Cells[idx+1:]...)

	// If the parent now has one cell, remove the parent from the tree and replace it with that cell.
	if len(lcparent.Cells) == 1 {
		singleCell := lcparent.Cells[0]
		singleCell.Parent = lcparent.Parent

		if singleCell.Parent == nil {
			singleCell.Xoff = 0
			singleCell.Yoff = 0
			*lcroot = singleCell
		} else {
			// Replace lcparent with singleCell in the parent's cells.
			for i, cell := range singleCell.Parent.Cells {
				if cell == lcparent {
					singleCell.Parent.Cells[i] = singleCell
					break
				}
			}
		}
	}
}

func findBottomRight(lc *LayoutCell) *LayoutCell {
	if lc.Type == LayoutWindowPane {
		return lc
	}
	lc = lc.Cells[len(lc.Cells)-1]
	return findBottomRight(lc)
}

func layoutConstruct(parent *LayoutCell, layoutStr *string) (*LayoutCell, error) {
	var sx, sy, xoff, yoff uint
	var lc *LayoutCell

	if _, err := fmt.Sscanf(*layoutStr, "%ux%u,%u%u", &sx, &sy, &xoff, &yoff); err != nil {
		return nil, err
	}

	// Advance the layout string past the parsed section
	advanceLayout(layoutStr, sx, sy, xoff, yoff)

	lc = &LayoutCell{
		Sx:     sx,
		Xoff:   xoff,
		Yoff:   yoff,
		Parent: parent,
	}

	switch (*layoutStr)[0] {
	case ',', '}', ']', '\x00':
		return lc, nil
	case '{':
		lc.Type = LayoutLeftRight
	case '[':
		lc.Type = LayoutTopBottom
	default:
		return nil, ErrorInvalidLayout
	}

	*layoutStr = (*layoutStr)[1:] // Skip layout type char

	for {

		if (*layoutStr)[0] != ',' {
			break
		}

		*layoutStr = (*layoutStr)[1:] // skip comma

		lcchild, err := layoutConstruct(lc, layoutStr)
		if err != nil {
			return nil, err
		}

		lc.Cells = append(lc.Cells, lcchild)
	}

	// Check for closing char
	switch lc.Type {
	case LayoutLeftRight:
		if (*layoutStr)[0] != '}' {
			return nil, errors.New("expected '}' in layout")
		}
	case LayoutTopBottom:
		if (*layoutStr)[0] != ']' {
			return nil, errors.New("expected ']' in layout")
		}
	}
	*layoutStr = (*layoutStr)[1:] // Skip the closing char

	return lc, nil
}

func advanceLayout(layoutStr *string, sx, sy, xoff, yoff uint) {
	*layoutStr = strings.SplitN(*layoutStr, "{", 2)[1]
	*layoutStr = strings.SplitN(*layoutStr, "[", 2)[1]
}
