#!/usr/bin/env python3

from pwn import *

IP = sys.argv[1] if len(sys.argv) > 1 else 'localhost'
PORT = 17777

def is_valid(board, row, col, num):
    for i in range(9):
        if board[row][i] == num or board[i][col] == num:
            return False
    start_row, start_col = 3 * (row // 3), 3 * (col // 3)
    for i in range(start_row, start_row + 3):
        for j in range(start_col, start_col + 3):
            if board[i][j] == num:
                return False

    return True

def solve_sudoku(board):
    for row in range(9):
        for col in range(9):
            if board[row][col] == 0:
                for num in range(1, 10):
                    if is_valid(board, row, col, num):
                        board[row][col] = num

                        if solve_sudoku(board):
                            return True

                        board[row][col] = 0
                return False
    return True

with remote(f"{IP}", PORT) as r:
    r.recvuntil(b']]\n')
    r.recvline()
    for c in range(30):
        matrix = []
        for i in range(9):
            string = r.recvline().decode().replace('\n','').replace('-','0').split(' ')
            # print(string)
            row = [int(x) for x in string]
            matrix.append(row)
        r.recvuntil(b'answer: ')
        if solve_sudoku(matrix):
            answer = matrix
        r.sendline(str(answer).encode())
        print(r.recvuntil(b']\n').decode())
    