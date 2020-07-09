#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
int main()
{
    printf("uid:%d euid:%d\n", getuid(), geteuid());
    setuid(0);
    printf("uid:%d euid:%d\n", getuid(), geteuid());
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