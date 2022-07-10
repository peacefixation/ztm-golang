package pxcanvas

import (
	"pixl/pxcanvas/brush"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (pxCanvas *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.Scale(int(ev.Scrolled.DY))
	pxCanvas.Refresh()
}

// Hoverable
func (pxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	if x, y := pxCanvas.MouseToCanvasXY(ev); x != nil && y != nil {
		brush.TryBrush(pxCanvas.appState, pxCanvas, ev)

		// draw the cursor
		cursor := brush.Cursor(pxCanvas.renderer.pxCanvas.PxCanvasConfig, pxCanvas.appState.BrushType, ev, *x, *y)
		pxCanvas.renderer.SetCursor(cursor)
	} else {
		// hide the cursor
		pxCanvas.renderer.SetCursor(make([]fyne.CanvasObject, 0))
	}

	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)

	pxCanvas.Refresh()

	pxCanvas.mouseState.previousCoord = &ev.PointEvent
}

// Hoverable
func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {}

// Hoverable
func (pxCanvas *PxCanvas) MouseOut() {}

// Mouseable

func (pxCanvas *PxCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
}

func (pxCanvas *PxCanvas) MouseUp(ev *desktop.MouseEvent) {

}
