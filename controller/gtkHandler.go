package controller

import (
	"strconv"
	"time"

	"github.com/FelipeAlafy/Flex/handler"
	"github.com/gotk3/gotk3/gtk"
)

func getDataFromEntry(e *gtk.Entry) string {
	v, err := e.GetText()
	handler.Error("controller/gtkHandler.go >> getDataFromEntry", err)
	return v
}

func getDataFromComboBox(c *gtk.ComboBoxText) string {
	v := c.GetActiveText()
	return v
}

func getDataFromTextView(t *gtk.TextView) string {
	v, err := t.GetBuffer()
	handler.Error("controller/gtkHandler.go >> getDataFromTextView while trying to get textBuffer", err)
	s, err  := v.GetText(v.GetStartIter(), v.GetEndIter(), true)
	handler.Error("controller/gtkHandler.go >> getDataFromTextView while trying to get text", err)
	return s
}

func getDataFromCalendar(c *gtk.Calendar) time.Time {
	year, month, day := c.GetDate()
	return time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.Local)
}

func getDataFromCheckBox(c *gtk.CheckButton) bool {
	return c.GetActive()
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {return 0}
	return i
}