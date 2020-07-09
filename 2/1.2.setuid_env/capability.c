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
    sleep(1);
    setuid(getuid());
    if (fork())
    {
        close(fd);
        exit(0);
    }
    else
    {
        write(fd, "Malicious Data\n", 15);
        close(fd);
    }
}