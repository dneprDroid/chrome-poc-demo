#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>

#define DLL_PUBLIC __attribute__((used)) __attribute__ ((visibility ("default")))

#ifdef __cplusplus
extern "C" {
#endif

    DLL_PUBLIC
    uint32_t SuperFastHash(const void * data, int len);

#ifdef __cplusplus
}
#endif