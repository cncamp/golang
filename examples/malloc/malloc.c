#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#define BLOCK_SIZE (100*1024*1024)
char* allocMemory() {
    char* out = (char*)malloc(BLOCK_SIZE);
    memset(out, 'A', BLOCK_SIZE);
    return out;
}