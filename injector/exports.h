#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>

#define EXPORTED __attribute__((used)) __attribute__ ((visibility ("default")))

#ifdef __cplusplus
extern "C" {
#endif

    EXPORTED
    uint32_t SuperFastHash(const void * data, int len);

#ifdef __cplusplus
}
#endif