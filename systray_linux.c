#include <stdlib.h>
#include <stdio.h>
#include "systray.h"

void registerSystray(void) {
}

int nativeLoop(void) {
	printf("Linux is currently unsupported, failing to show system tray icon\n");
	return 0;
}

void setIcon(const char* iconBytes, int length, bool template) {
}

void setTitle(char* ctitle) {
}

void setTooltip(char* ctooltip) {
}

void setMenuItemIcon(const char* iconBytes, int length, int menuId, bool template) {
}

void add_or_update_menu_item(int menu_id, int parent_menu_id, char* title, char* tooltip, short disabled, short checked, short isCheckable) {
}

void add_separator(int menu_id) {
}

void hide_menu_item(int menu_id) {
}

void show_menu_item(int menu_id) {
}

void quit() {
}
