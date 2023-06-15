package elements

type ElementType string

var SupporterElementTypes = []ElementType{
	WINDOW, BUTTON, LIST, LISTITEM, GROUP, TEXT, MENU, MENUITEM,
	CHECKBOX, EDIT, PANE, COMBOBOX, DOCUMENT, HYPERLINK, IMAGE, TITLE, AREA}

const (
	WINDOW       ElementType = "window"
	BUTTON       ElementType = "button"
	LIST         ElementType = "list"
	LISTITEM     ElementType = "listitem"
	GROUP        ElementType = "group"
	TEXT         ElementType = "text"
	MENU         ElementType = "menu"
	MENUITEM     ElementType = "menuitem"
	CHECKBOX     ElementType = "checkbox"
	EDIT         ElementType = "edit"
	PANE         ElementType = "pane"
	COMBOBOX     ElementType = "combobox"
	DOCUMENT     ElementType = "document"
	HYPERLINK    ElementType = "hyperlink"
	IMAGE        ElementType = "image"
	TITLE        ElementType = "title"
	AREA         ElementType = "area"
	NONSUPPORTED ElementType = "unknown"
)
