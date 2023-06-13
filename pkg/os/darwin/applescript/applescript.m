#import "applescript.h"

void Keystroke(const char* keystroke){
    NSString *nKeystroke = [NSString stringWithUTF8String:keystroke];
    NSString *aScript = [NSString stringWithFormat:@"tell application \"System Events\"\n\
                                                    keystroke \"%@\" & return \n\
                                                    end tell", nKeystroke];
    // NSLog(@"aScript is %@ \n", aScript);
    NSAppleScript *script = [[NSAppleScript alloc] initWithSource:aScript];
    NSDictionary *errorInfo = nil;
    [script executeAndReturnError:&errorInfo];
    
    if (errorInfo) {
        NSLog(@"Error executing AppleScript: %@", errorInfo);
    }
}