package window

import (
	"time"

	"github.com/SpaiR/imgui-go"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"sdmm/platform"
)

func (w *Window) Process() {
	ticker := time.NewTicker(time.Second / time.Duration(fps))

	for !w.application.IsClosed() {
		// Override window closing behaviour to enforce our checks.
		if w.handle.ShouldClose() {
			w.application.CloseCheck()
			w.handle.SetShouldClose(false)
		}
		w.runFrame()
		<-ticker.C
	}
}

func (w *Window) runFrame() {
	w.startFrame()
	w.application.Process()
	w.endFrame()
	w.application.PostProcess()
}

func (w *Window) startFrame() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	platform.NewImGuiGLFWFrame()
	imgui.NewFrame()
	runLaterJobs()
	runRepeatJobs()
}

func runLaterJobs() {
	for _, job := range laterJobs {
		job()
	}
	laterJobs = nil
}

func runRepeatJobs() {
	for _, job := range repeatJobs {
		job()
	}
}

func (w *Window) endFrame() {
	imgui.Render()
	platform.Render(imgui.RenderedDrawData())
	w.handle.SwapBuffers()
	glfw.PollEvents()
}
