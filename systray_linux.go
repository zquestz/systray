package systray

// SetTemplateIcon sets the systray icon as a template icon (on macOS), falling back
// to a regular icon on other platforms.
// templateIconBytes and iconBytes should be the content of .ico for windows and
// .ico/.jpg/.png for other platforms.
func SetTemplateIcon(templateIconBytes []byte, regularIconBytes []byte) {
}


// SetIcon sets the systray icon.
// iconBytes should be the content of .ico for windows and .ico/.jpg/.png
// for other platforms.
func SetIcon(iconBytes []byte) {
}

// SetTitle sets the systray title, only available on Mac and Linux.
func SetTitle(title string) {
}

// SetTooltip sets the systray tooltip to display on mouse hover of the tray icon,
// only available on Mac and Windows.
func SetTooltip(tooltip string) {
}

// SetIcon sets the icon of a menu item. Only works on macOS and Windows.
// iconBytes should be the content of .ico/.jpg/.png
func (item *MenuItem) SetIcon(iconBytes []byte) {
}

// SetTemplateIcon sets the icon of a menu item as a template icon (on macOS). On Windows, it
// falls back to the regular icon bytes and on Linux it does nothing.
// templateIconBytes and regularIconBytes should be the content of .ico for windows and
// .ico/.jpg/.png for other platforms.
func (item *MenuItem) SetTemplateIcon(templateIconBytes []byte, regularIconBytes []byte) {
}

func addOrUpdateMenuItem(item *MenuItem) {
}

func setInternalLoop(internal bool) {
}

func registerSystray() {
}

func addSeparator(id uint32) {
}

func hideMenuItem(item *MenuItem) {
}

func showMenuItem(item *MenuItem) {
}
func nativeLoop() int {
	return 0
}

func nativeEnd() {
	systrayExit()
}

func quit() {
}

func nativeStart() {
	systrayReady()
}
