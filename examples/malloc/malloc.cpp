#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <iostream>
#include <math.h>

using namespace std;
#define BLOCK_SIZE (100*1024*1024)

extern "C" {
    char* allocMemory() {
        char* out = (char*)malloc(BLOCK_SIZE);
        memset(out, 'A', BLOCK_SIZE);
        return out;
    }
}