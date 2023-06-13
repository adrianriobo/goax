#import <Cocoa/Cocoa.h>
#import <CoreFoundation/CoreFoundation.h>
#import <AppKit/AppKit.h>

// Check if an axuielement has children
const char* HasChildren(CFTypeRef axuielement);

// Get an array holding all children for an auxuielement
CFArrayRef GetChildren(CFTypeRef axuielement);

// Get the axuielement child at position index from array of children
CFTypeRef GetChild(CFArrayRef children, CFIndex index);

// Get title attribute
const char* GetTitle(CFTypeRef axuielement);

// Get value attribute
const char* GetValue(CFTypeRef axuielement);

// Set value attribute
void SetValue(CFTypeRef axuielement, const char* value);

// Get description attribute
const char* GetDescription(CFTypeRef axuielement);

// Get description attribute
const char* GetRole(CFTypeRef axuielement);

// Show actions to perfom for an axuielement
void ShowActions(CFTypeRef axuielement);

// Execute press action on an axuielement
void Press(CFTypeRef axuielement);

