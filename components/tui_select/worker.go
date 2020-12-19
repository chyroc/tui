package tui_select

import (
	"github.com/sirupsen/logrus"
	"github.com/tj/go-terminput"

	"github.com/chyroc/tui"
)

// 实现 tui.Worker
type worker struct {
	title     string
	options   []string
	size      int
	windowMin int
	windowMax int
	selected  int
	tui       tui.TUI
	err       error
}

// 实现 tui.Worker
func (r *worker) Init() error {
	return nil
}

// 实现 tui.Worker
func (r *worker) Close() error {
	return nil
}

// 实现 tui.Worker
func (r *worker) View() string {
	s := ""
	if r.title != "" {
		s = s + r.title + "\n\n"
	}
	if r.size == 0 {
		for idx, v := range r.options {
			if idx == r.selected {
				s += "(•) " + v
			} else {
				s += "( ) " + v
			}
			s += "\n"
		}
	} else {
		for idx := r.windowMin; idx <= max(r.windowMax, len(r.options)-1); idx++ {
			if idx == r.selected {
				s += "(•) " + r.options[idx]
			} else {
				s += "( ) " + r.options[idx]
			}
			s += "\n"
		}
	}
	return s
}

// 处理输入
func (r *worker) HandleInput(e *terminput.KeyboardInput) {
	logrus.Infof("[input start] e=%s, e.ctrl=%v, rune=%v, mod=%v, key=%v\n", e, e.Ctrl(), e.Rune(), e.Mod(), e.Key())
	defer func() {
		logrus.Infof("[input   end] selected=%d", r.selected)
	}()

	switch {
	case e.Key() == terminput.KeyEscape || e.Rune() == 'q' || e.Key() == tui.KeyCtrlC:
		r.err = ErrNoSelect
		r.tui.Stop()
	case e.Key() == terminput.KeyUp:
		r.ptrUp()
	case e.Key() == terminput.KeyDown:
		r.ptrDown()
	case e.Key() == terminput.KeyEnter:
		r.tui.Stop()
	}
}

// 窗口上移
func (r *worker) ptrUp() {
	logrus.Infof("click up, r.selected=%d", r.selected)
	if r.selected > 0 {
		r.selected--
	}
	if r.size == 0 || r.isPtrInWindow() {
		return
	}
	if r.windowMin > 0 {
		r.windowMin--
		r.windowMax--
	}
}

// 窗口下移
func (r *worker) ptrDown() {
	if r.selected < len(r.options)-1 {
		r.selected++
	}
	if r.size == 0 || r.isPtrInWindow() {
		return
	}
	if r.windowMax < len(r.options)-1 {
		r.windowMin++
		r.windowMax++
	}
}

// 指针是否在窗口中
func (r *worker) isPtrInWindow() bool {
	if r.size == 0 {
		return true
	}
	return r.selected >= r.windowMin && r.selected <= r.windowMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
