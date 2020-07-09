#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
void main()
{
    int fd;
    fd = open("./zzz", O_RDWR | O_APPEND);
    if (fd == -1)
    {
        printf("Cannot open ./zzz\n");
        exit(0);
    }

    write(fd, "Malicious Data\n", 15);
    close(fd);
}