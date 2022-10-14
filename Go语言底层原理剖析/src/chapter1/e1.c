#include <stdio.h>
 
#include <unistd.h>
 
int main(int argc, char **argv) {
     
    char* argument_list[] = {"hello", "-g", NULL}; // NULL terminated array of char* strings
 
    execvp("/Users/Ben/work/test/hello", argument_list);
}