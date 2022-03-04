package systray

import (
	"log"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/prop"
)

var rootMenu = menuLayout{}

// SetIcon sets the icon of a menu item. Only works on macOS and Windows.
// iconBytes should be the content of .ico/.jpg/.png
func (item *MenuItem) SetIcon(iconBytes []byte) {
}

func (t *tray) GetLayout(parentId int32, recursionDepth int32, propertyNames []string) (revision uint32, layout menuLayout, err *dbus.Error) {
	return 1, rootMenu, nil
}

// GetGroupProperties is com.canonical.dbusmenu.GetGroupProperties method.
func (t *tray) GetGroupProperties(ids []int32, propertyNames []string) (properties []struct {
	V0 int32
	V1 map[string]dbus.Variant
}, err *dbus.Error) {
	return
}

// GetProperty is com.canonical.dbusmenu.GetProperty method.
func (t *tray) GetProperty(id int32, name string) (value dbus.Variant, err *dbus.Error) {
	return
}

// Event is com.canonical.dbusmenu.Event method.
func (t *tray) Event(id int32, eventId string, data dbus.Variant, timestamp uint32) (err *dbus.Error) {
	if eventId == "clicked" {
		item, ok := menuItems[uint32(id)]
		if !ok {
			log.Printf("Failed to look up clicked menu item")
			return
		}

		item.ClickedCh <- struct{}{}
	}
	return
}

// EventGroup is com.canonical.dbusmenu.EventGroup method.
func (t *tray) EventGroup(events []struct {
	V0 int32
	V1 string
	V2 dbus.Variant
	V3 uint32
}) (idErrors []int32, err *dbus.Error) {
	return
}

// AboutToShow is com.canonical.dbusmenu.AboutToShow method.
func (t *tray) AboutToShow(id int32) (needUpdate bool, err *dbus.Error) {
	return
}

// AboutToShowGroup is com.canonical.dbusmenu.AboutToShowGroup method.
func (t *tray) AboutToShowGroup(ids []int32) (updatesNeeded []int32, idErrors []int32, err *dbus.Error) {
	return
}

func createMenuPropSpec() map[string]map[string]*prop.Prop {
	return map[string]map[string]*prop.Prop{
		"com.canonical.dbusmenu": {
			"Version": {
				uint32(1),
				false,
				prop.EmitTrue,
				nil,
			},
			"TextDirection": {
				"ltr",
				false,
				prop.EmitTrue,
				nil,
			},
			"Status": {
				"active",
				false,
				prop.EmitTrue,
				nil,
			},
			"IconThemePath": {
				[]string{},
				false,
				prop.EmitTrue,
				nil,
			},
		},
	}
}

type menuLayout = struct {
	V0 int32
	V1 map[string]dbus.Variant
	V2 []dbus.Variant
}

func addOrUpdateMenuItem(item *MenuItem) {
	layout := menuLayout{
		V0: int32(item.id),
		V1: map[string]dbus.Variant{
			"label": dbus.MakeVariant(item.title),
		},
		V2: []dbus.Variant{},
	}

	rootMenu.V2 = append(rootMenu.V2, dbus.MakeVariant(layout))
}

func addSeparator(id uint32) {
	layout := menuLayout{
		V0: int32(id),
		V1: map[string]dbus.Variant{
			"type": dbus.MakeVariant("separator"),
		},
		V2: []dbus.Variant{},
	}

	rootMenu.V2 = append(rootMenu.V2, dbus.MakeVariant(layout))
}

func hideMenuItem(item *MenuItem) {
}

func showMenuItem(item *MenuItem) {
}
