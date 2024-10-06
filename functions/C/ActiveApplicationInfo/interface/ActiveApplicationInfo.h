#ifndef ActiveApplicationInfo_h
#define ActiveApplicationInfo_h

#ifdef __OBJC__
#import <Foundation/Foundation.h>

@interface ActiveApplicationInfo : NSObject

+ (NSString *)name;
+ (NSString *)runningTime;

@end
#endif

#ifdef __cplusplus
extern "C" {
#endif

char* getActiveAppName(void);
const char* getActiveAppRunningTime(void);

#ifdef __cplusplus
}
#endif

#endif // ActiveApplicationInfo_h
