import sys
import random

def Sample(possible, mines):
    print()
    possible = list(possible)
    a = int(random.random() * 1337)
    c = int(random.random() * 69068)
    m = mines
    samp = []
    mina = 1
    print('Тут точно есть мины: ', end='')
    for i in range(m):
        random.seed(i)
        random.shuffle(possible)
        mina = (a*mina+c)%m
        if i == 0:
            print(possible[mina], end  = ' ')
        if i == 1:
            print(possible[mina], end  = ' ')
        samp.append(possible[mina])
    print()
    print()
    return samp


class Minesweeper:
    def __init__(self, rows, cols, mines):
        self.size = rows * cols
        self.rows = rows
        self.cols = cols
        self.mines = mines
        self.board = [-1] * self.size
        self.first_step = True

    def play(self):
        while True:
            print("\033[H\033[J", end="")  # очищаем терминал
            self.print_board()
            try:
                row, col = map(int, input("Введи строку и столбец (0-9): ").split())
                assert 0 <= row < self.rows and 0 <= col < self.cols
            except (ValueError, AssertionError):
                print("Ошибка. Значения должны быть между (0 0) и (9 9).")
                continue
            idx = row * self.cols + col
            if self.first_step:
                self.place_mines(row, col)
                self.first_step = False

            if self.board[idx] == 0:
                print("Взрыв! Игра окончена!")
                self.print_board(reveal_mines=True)
                break

            self.reveal(idx)

            if all(cell >= 0 for cell in self.board):
                flag = open('flag.txt', 'r')
                print("Вот флаг: ", flag.read())
                break

    def print_board(self, reveal_mines=False):
        max_col_digits = len(str(self.cols - 1))
        max_row_digits = len(str(self.rows - 1))

        print(" " * (max_row_digits + 2) + " ".join(f"{i:>{max_col_digits}}" for i in range(self.cols)))

        for i in range(self.rows):
            row = ['X' if reveal_mines and value == 0 else
                   str(abs(value) - 1) if reveal_mines and value < 0 else
                   '*' if value <= 0 else 0
                   for value in (self.board[i * self.cols + j] for j in range(self.cols))]
            print(f"{i:>{max_row_digits}}: " + " ".join(f"{cell:>{max_col_digits}}" for cell in row))




    def place_mines(self, safe_row, safe_col):
        def count_mines(idx):
            return sum(1 for ni in self.get_neighbors(idx) if self.board[ni] == 0)

        safe = set(self.get_neighbors(safe_row * self.cols + safe_col) + [safe_row * self.cols + safe_col])
        possible = [(i, j) for i in range(self.rows) for j in range(self.cols) if (i * self.cols + j) not in safe]
        list(map(lambda pos: self.board.__setitem__(pos[0] * self.cols + pos[1], 0), Sample(possible, self.mines)))
        self.board = [-1 * (count_mines(idx) + 1) if cell != 0 else cell for idx, cell in enumerate(self.board)]

    def get_neighbors(self, idx):
        r, c = divmod(idx, self.cols)
        return [i * self.cols + j for i in range(r - 1, r + 2)
                for j in range(c - 1, c + 2)
                if 0 <= i < self.rows and 0 <= j < self.cols and (i != r or j != c)]

    def reveal(self, idx):
        stack = [idx]
        while stack:
            current_idx = stack.pop()
            self.board[current_idx] = abs(self.board[current_idx])
            if self.board[current_idx] - 1 == 0:
                stack.extend(ni for ni in self.get_neighbors(current_idx) if self.board[ni] < 0)


if __name__ == "__main__":
    print("Слепой сапёр!")
    rows, cols, mines = (int(arg) for arg in sys.argv[1:]) if len(sys.argv) > 1 else (10, 10, 83)
    game = Minesweeper(rows, cols, mines)
    game.play()