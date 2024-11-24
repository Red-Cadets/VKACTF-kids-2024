#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# This exploit template was generated via:
# $ pwn template --host 212.193.56.123 --port 61111 minesweeper.py
from copy import copy
from pwn import *

# Set up pwntools for the correct architecture
context.update(arch='i386')
exe = './minesweeper.py'

# Many built-in settings can be controlled on the command-line and show up
# in "args".  For example, to dump all data sent/received, and disable ASLR
# for all created processes...
# ./exploit.py DEBUG NOASLR
# ./exploit.py GDB HOST=example.com PORT=4141 EXE=/tmp/executable
host = args.HOST or '212.193.56.123'
port = int(args.PORT or 61111)


def start_local(argv=[], *a, **kw):
    '''Execute the target binary locally'''
    if args.GDB:
        return gdb.debug([exe] + argv, gdbscript=gdbscript, *a, **kw)
    else:
        return process([exe] + argv, *a, **kw)

def start_remote(argv=[], *a, **kw):
    '''Connect to the process on the remote host'''
    io = connect(host, port)
    if args.GDB:
        gdb.attach(io, gdbscript=gdbscript)
    return io

def start(argv=[], *a, **kw):
    '''Start the exploit against the target.'''
    if args.LOCAL:
        return start_local(argv, *a, **kw)
    else:
        return start_remote(argv, *a, **kw)

# Specify your GDB script here for debugging
# GDB will be launched if the exploit is run via e.g.
# ./exploit.py GDB
gdbscript = '''
continue
'''.format(**locals())

#===========================================================
#                    EXPLOIT GOES HERE
#===========================================================
def Sample(possible, mines, a, c, limit=1000):
    possible = list(possible)
    m = mines
    samp = []
    mina = 1
    for i in range(min(limit, m)):
        random.seed(i)
        random.shuffle(possible)
        mina = (a*mina+c)%m
        samp.append(possible[mina])
    return samp
DD = 83
for DD in range(83, 84):
    retry = 0
    print(f'Try {DD = }')
    while retry < 3:
        io = start()
        try:
            possible = [(0, 3), (0, 4), (0, 5), (0, 6), (0, 7), (0, 8), (0, 9), (1, 3), (1, 4), (1, 5), (1, 6), (1, 7), (1, 8), (1, 9), (2, 3), (2, 4), (2, 5), (2, 6), (2, 7), (2, 8), (2, 9), (3, 0), (3, 1), (3, 2), (3, 3), (3, 4), (3, 5), (3, 6), (3, 7), (3, 8), (3, 9), (4, 0), (4, 1), (4, 2), (4, 3), (4, 4), (4, 5), (4, 6), (4, 7), (4, 8), (4, 9), (5, 0), (5, 1), (5, 2), (5, 3), (5, 4), (5, 5), (5, 6), (5, 7), (5, 8), (5, 9), (6, 0), (6, 1), (6, 2), (6, 3), (6, 4), (6, 5), (6, 6), (6, 7), (6, 8), (6, 9), (7, 0), (7, 1), (7, 2), (7, 3), (7, 4), (7, 5), (7, 6), (7, 7), (7, 8), (7, 9), (8, 0), (8, 1), (8, 2), (8, 3), (8, 4), (8, 5), (8, 6), (8, 7), (8, 8), (8, 9), (9, 0), (9, 1), (9, 2), (9, 3), (9, 4), (9, 5), (9, 6), (9, 7), (9, 8), (9, 9)]
            io.sendline(b'1 1')
            io.recvuntil('мины: '.encode())
            mines = io.recvline().strip()
            mines = re.findall(r'\d+', mines.decode())
            mines = [
                (int(mines[i*2]), int(mines[i*2+1]))
                for i in range(len(mines)//2)
            ]
            possible_variants = set()
            for a in range(100):
                for c in range(100):
                    r = Sample(possible[:], DD, a=a, c=c, limit=2)
                    if tuple(r[:2]) == tuple(mines):
                        possible_variants.add((a, c))
                        break
            print(possible_variants)
            assert len(possible_variants) >= 1
            a, c = possible_variants.pop()
            samp = Sample(possible[:], DD, a=a, c=c)
            samp = set(samp)
            print(f'{len(samp) = }')
            # print(samp)
            for x in range(10):
                for y in range(10):
                    if (x, y) in samp:
                        continue
                    io.sendline(f'{x} {y}'.encode())
            io.interactive()


        except AssertionError as e:
            print(e)
        retry += 1
        io.close()
    