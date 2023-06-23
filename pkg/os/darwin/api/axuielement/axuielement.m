#import "axuielement.h"

// https://www.electronjs.org/docs/latest/tutorial/accessibility#macos
// CFStringRef kAXManualAccessibility = CFSTR("AXManualAccessibility");

const char* HasChildren(CFTypeRef axuielement) {
    @autoreleasepool {
        AXError err;
        CFArrayRef childrenPtr;
        NSString *result = @"true";
        err = AXUIElementCopyAttributeValue(axuielement, kAXChildrenAttribute, (CFTypeRef *) &childrenPtr);
        if([childrenPtr count]==0){
            result = @"false";
        }
        return strdup([result UTF8String]);
    }
}

CFArrayRef GetChildren(CFTypeRef axuielement) {
    @autoreleasepool {
        AXError err;
        CFArrayRef childrenPtr;
        err = AXUIElementCopyAttributeValue(axuielement, kAXChildrenAttribute, (CFTypeRef *) &childrenPtr);
        return childrenPtr;
    }
}

CFTypeRef GetChild(CFArrayRef children, CFIndex index) {
    @autoreleasepool {
        AXUIElementRef objectChild = CFArrayGetValueAtIndex(children, index);
        return objectChild;
    }
}

const char* GetTitle(CFTypeRef axuielement) {
    @autoreleasepool {
        AXError err;
        NSString *att = nil;
        err = AXUIElementCopyAttributeValue(axuielement, kAXTitleAttribute, (CFTypeRef *) &att);
        if (err != kAXErrorSuccess) 
            return "";
        return strdup([att UTF8String]);
    }
}

const char* GetValue(CFTypeRef axuielement) {
    @autoreleasepool {
        AXError err;
        NSString *att = nil;
        err = AXUIElementCopyAttributeValue(axuielement, kAXValueAttribute, (CFTypeRef *) &att);
        if (err != kAXErrorSuccess) 
            return "";
        // Values holding by elements can be numeric, strings.... 
        // initially we take care for strings then we will add for checkboxes...
        if (CFGetTypeID(att) == CFStringGetTypeID()) {
            return strdup([att UTF8String]);
        }
        return "";
    }
}

void SetValue(CFTypeRef axuielement, const char* value) {
    @autoreleasepool {
        AXError err;
        NSString *nValue = [NSString stringWithUTF8String:value];
        err = AXUIElementSetAttributeValue(axuielement, kAXValueAttribute, CFBridgingRetain(nValue));
        assert(kAXErrorSuccess == err);
    }
}

const char* GetDescription(CFTypeRef axuielement) {
    @autoreleasepool {
        AXError err;
        NSString *att = nil;
        err = AXUIElementCopyAttributeValue(axuielement, kAXDescriptionAttribute, (CFTypeRef *) &att);
        if (err != kAXErrorSuccess) 
            return "";
        return strdup([att UTF8String]);
    }
}

const char* GetRole(CFTypeRef axuielement) {
    @autoreleasepool {
        AXError err;
        NSString *att = nil;
        err = AXUIElementCopyAttributeValue(axuielement, kAXRoleAttribute, (CFTypeRef *) &att);
        if (err != kAXErrorSuccess) 
            return "";
        return strdup([att UTF8String]);
    }
}



void ShowActions(CFTypeRef axuielement) {
    @autoreleasepool {
        // AXUIElementRef a = (AXUIElementRef)axuielement;
        AXError err;
        CFArrayRef actions = nil;
        err = AXUIElementCopyActionNames(axuielement, &actions);
        if (err == kAXErrorSuccess && actions != nil) {
            NSArray *actionsArray = (__bridge NSArray *)(actions);
            for (NSString *action in actionsArray) {
                NSLog(@"%@", action);
            }
        }
    }    
}

void Press(CFTypeRef axuielement) {
    @autoreleasepool {
        AXError err;
        err = AXUIElementPerformAction((AXUIElementRef)axuielement, kAXPressAction);
        if (err == kAXErrorActionUnsupported) 
            NSLog(@"error kAXErrorActionUnsupported \n");
        if (err == kAXErrorIllegalArgument) 
            NSLog(@"error kAXErrorIllegalArgument \n");
        if (err == kAXErrorInvalidUIElement) 
            NSLog(@"error kAXErrorInvalidUIElement \n");
        if (err == kAXErrorCannotComplete) 
            NSLog(@"error kAXErrorCannotComplete \n");
        if (err == kAXErrorNotImplemented) 
            NSLog(@"error kAXErrorNotImplemented \n");
    }
}

