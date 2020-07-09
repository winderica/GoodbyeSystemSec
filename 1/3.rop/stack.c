#include <stdlib.h>
#include <stdio.h>
#include <string.h>

int bof(char *str)
{
    char buffer[12];
    /* The following statement has a buffer overflow problem */
    strcpy(buffer, str);
    return 1;
}
int main(int argc, char **argv)
{
    char str[200];
    FILE *badfile;
    int count = 0;
    badfile = fopen("badfile", "rb");
    count = fread(str, sizeof(char), 200, badfile);
    printf("%d\n", count);
    bof(str);
    printf("Returned Properly\n");
    return 0;
}
