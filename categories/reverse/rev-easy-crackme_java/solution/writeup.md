## Разминка

| Событие | Название | Категория | Сложность |
| :------ | ---- | ---- | ---- |
| VKAKIDS 2024 | Getting started | reverse | easy |

  
### Описание


> Автор: [Invis_one]
>
Хороший реверсер, как и спортсмены, всегда должны начинать с разминки, не так ли?


### Решение
Открывaем предоставленный нам исходный код `Java`

Видим, что введенная пользователем строка должна начинаться с `vka{` и заканчитваться `}`, иначе в консоль выведется `Nope :(`. Смотрим дальше.

```java
String input = userInput.substring("vka{".length(), userInput.length() - 1);
```
В данном блоке из введеной пользователем строки убирается начало `vka{` и убирается последний символ.
Затем измененный `input` передается в метод `CheckPassword`

```java
public boolean checkPassword(String password) { 
    boolean part1 = password.startsWith("java_code_");
    boolean part3 = password.endsWith("sy_to_read");
    boolean part2 = password.contains("is_that_ea");        
    return part1 && part3 && part2;
}
```
В данном методе проверяется, чтобы переданная ему строка начиналась с `java_code_`, заканчивалась `sy_to_read` и содержала `is_that_ea`.
Значит, вся строка должна иметь вид `java_code_is_that_easy_to_read`. Обьединяя с выше указанными проверками получаем флаг, при вводе которого в консоль выведется `Good job!!!`
 
### Флаг

```
vka{java_code_is_that_easy_to_read}
```
