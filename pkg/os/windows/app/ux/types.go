//go:build windows

package ux

import (
	wa "github.com/openstandia/w32uiautomation"
)

const (
	WINDOW    = "window"
	BUTTON    = "button"
	LIST      = "list"
	LISTITEM  = "listitem"
	GROUP     = "group"
	TEXT      = "text"
	MENU      = "menu"
	MENUITEM  = "menuitem"
	CHECKBOX  = "checkbox"
	EDIT      = "edit"
	PANE      = "pane"
	COMBOBOX  = "combobox"
	DOCUMENT  = "document"
	HYPERLINK = "hyperlink"

	windowId    = wa.UIA_WindowControlTypeId
	buttonId    = wa.UIA_ButtonControlTypeId
	listId      = wa.UIA_ListControlTypeId
	listitemId  = wa.UIA_ListItemControlTypeId
	groupId     = wa.UIA_GroupControlTypeId
	textId      = wa.UIA_TextControlTypeId
	menuId      = wa.UIA_MenuControlTypeId
	menuitemId  = wa.UIA_MenuItemControlTypeId
	checkboxId  = wa.UIA_CheckBoxControlTypeId
	editId      = wa.UIA_EditControlTypeId
	paneId      = wa.UIA_PaneControlTypeId
	comboboxId  = wa.UIA_ComboBoxControlTypeId
	documentId  = wa.UIA_DocumentControlTypeId
	hyperlinkId = wa.UIA_HyperlinkControlTypeId
)

var elementTypes map[string]int64 = map[string]int64{
	WINDOW:    windowId,
	BUTTON:    buttonId,
	LIST:      listId,
	LISTITEM:  listitemId,
	GROUP:     groupId,
	TEXT:      textId,
	MENU:      menuId,
	MENUITEM:  menuitemId,
	CHECKBOX:  checkboxId,
	EDIT:      editId,
	PANE:      paneId,
	COMBOBOX:  comboboxId,
	DOCUMENT:  documentId,
	HYPERLINK: hyperlinkId}

type UXElement struct {
	name        string
	elementType string
	Ref         *wa.IUIAutomationElement
	Parent      *UXElement
	children    []*UXElement
}
