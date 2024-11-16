//go:build (linux || freebsd || openbsd || netbsd) && !android

package systray

import (
	"fyne.io/systray/internal/generated/notifier"

	"github.com/godbus/dbus/v5"
)

const (
	InterfaceStatusNotifierItem = "org.kde.StatusNotifierItem"
)

// NewCustomStatusNotifierItem creates and allocates org.kde.StatusNotifierItem.
func NewCustomStatusNotifierItem(object dbus.BusObject) *CustomStatusNotifierItem {
	return &CustomStatusNotifierItem{object}
}

// CustomStatusNotifierItem can be embedded to have forward compatible server implementations.
type CustomStatusNotifierItem struct {
	object dbus.BusObject
}

func (o *CustomStatusNotifierItem) iface() string {
	return notifier.InterfaceStatusNotifierItem
}

func (o *CustomStatusNotifierItem) ContextMenu(x int32, y int32) (err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (o *CustomStatusNotifierItem) Activate(x int32, y int32) (err *dbus.Error) {
	if activateFunc == nil {
		err = &dbus.ErrMsgUnknownMethod
		return
	}

	activateFunc()
	return
}

func (o *CustomStatusNotifierItem) SecondaryActivate(x int32, y int32) (err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (o *CustomStatusNotifierItem) Scroll(delta int32, orientation string) (err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}
