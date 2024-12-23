## Minesweeper

| Событие | Название | Категория | Сложность |
| :------ | ---- | ---- | ---- |
| VKAKIDS 2024 | Minesweeper | crypto | hard |

  
### Описание


> Автор: [maloznal]
>
Слепой сапёр


### Решение
При анализе исходного кода видно, что вся суть задания сводится к пониманию того по какому принципу расставляются мины.
В методе ‘place_mines’ вызывается функция Sample которая и отвечает за положение мин на поле.
Первое что мы должны сделать - задать выборку доступных клеток для размещения мин. После первого хода одназначно определяем эту выборку.
Чтобы восстановить координаты мин нужно восстановить параметры генератора a и c, тот есть решить систему линейных уравнений:
```
x1 = a * 1 + c
x2 = a * x1 + c
```
Это можно сделать с помощью Sagemath:
```
a, c = var('a c')
eq1 = a * 1 + c == mina1
eq2 = a * mina1 + c == mina2
solution = solve_mod([eq1, eq2], 83)
```
После нахождения параметров необходимо запустить генератор локально и отсеить клетки с бомбами.

Также данную задачу можно решить, перебрав все значения и с помощью pwntools автоматизировать процесс взаимодействия с ботом

Решение представленно в [`solution.py`](./solution.py)

### Флаг

```
vkactf{cryptor_vs_minesweeper}
```
