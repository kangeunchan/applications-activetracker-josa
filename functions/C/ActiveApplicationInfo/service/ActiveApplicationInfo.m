#import "../interface/ActiveApplicationInfo.h"
#import <Cocoa/Cocoa.h>

@implementation ActiveApplicationInfo

+ (NSString *)name {
    NSWorkspace *workspace = [NSWorkspace sharedWorkspace];
    NSRunningApplication *activeApp = [workspace frontmostApplication];
    return [activeApp localizedName];
}

+ (NSString *)runningTime {
    NSWorkspace *workspace = [NSWorkspace sharedWorkspace];
    NSRunningApplication *activeApp = [workspace frontmostApplication];
    
    NSDate *launchDate = [activeApp launchDate];
    if (launchDate == nil) {
        return @"Unknown";
    }

    NSDate *currentDate = [NSDate date];
    NSTimeInterval runningTime = [currentDate timeIntervalSinceDate:launchDate];
    
    NSDateComponentsFormatter *formatter = [[NSDateComponentsFormatter alloc] init];
    formatter.unitsStyle = NSDateComponentsFormatterUnitsStyleFull;
    formatter.allowedUnits = NSCalendarUnitHour | NSCalendarUnitMinute | NSCalendarUnitSecond;
    
    return [formatter stringFromTimeInterval:runningTime];
}

@end

const char* getActiveAppName() {
    return [[ActiveApplicationInfo name] UTF8String];
}

const char* getActiveAppRunningTime() {
    return [[ActiveApplicationInfo runningTime] UTF8String];
}