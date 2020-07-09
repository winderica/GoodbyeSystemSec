#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
int main()
{
    printf("uid:%d euid:%d\n", getuid(), geteuid());
    setuid(0);
    printf("uid:%d euid:%d\n", getuid(), geteuid());
}
