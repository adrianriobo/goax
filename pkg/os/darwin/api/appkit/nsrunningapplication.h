#import <Cocoa/Cocoa.h>
#import <AppKit/AppKit.h>

// https://developer.apple.com/documentation/appkit/nsworkspace/1532097-frontmostapplication?language=objc
void* FrontmostApplication();

void* FindRunningApplication(const char* bundleID, const char* windowTitle);

// Show all 
void ShowAllApplications();

// https://developer.apple.com/documentation/appkit/nsrunningapplication/1529140-bundleidentifier?language=objc
const char* BundleIdentifier(void* app);

// https://developer.apple.com/documentation/applicationservices/1459374-axuielementcreateapplication?language=objc
CFTypeRef CreateApplicationAXRef(void* app);

// https://developer.apple.com/documentation/applicationservices/kaxfocusedwindowattribute?language=objc
CFTypeRef GetAXFocusedWindow(CFTypeRef appAXRef);
