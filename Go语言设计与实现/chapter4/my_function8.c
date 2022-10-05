#include <stdio.h>

// ch04/my_function.c
int my_function(int arg1, int arg2,int arg3, int arg4,int arg5, int arg6,int arg7, int arg8) {
    return arg1 + arg2 + arg3 + arg4 + arg5 + arg6 + arg7 + arg8;
}

int main() {
    int i = my_function(1, 2, 3, 4, 5, 6, 7, 16);
}