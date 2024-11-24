## PPC

| Событие | Название | Категория | Сложность |
| :------ | ---- | ---- | ---- |
| VKAKIDS 2024 | WordSearcher | misc | easy |

  
### Описание


> Автор: [one!tea]
>
Вроде считать слова легко... Хм, кажется их тут не больше 30. СТОП, ещё так 50 раз?! Не, не осилил...


### Решение
Подключаясь, мы получаем строку, на первый взгляд просто набор случайных букв. Однако тут же замечаем слова, которые окружены кучей случайных букв.

bxz`pentest`f`process`rsknj`information`cycqb`phone`inlw`computer`v

А теперь, постоянно переподключаясь, пытаемся собрать полный словарь. В описании задания есть намёк на то, что их `30 штук`. 

Вуаля
```words_list = ["vkactf", "security", "coding", "chill", "flag", "game", "process", "music", "create", "cookie", "python", "java", "computer", "science", "network", "lucky", "exploit", "attack", "defense", "pentest", "proxy", "random", "server", "virus", "firewall", "data", "information", "browser", "phone", "byte"]```

Теперь остаётся написать код, который будет подключаться, считывать строку, считать кол-во слов и отправлять ответ. И так 50 раз. Пример реализован в [`solution.py`](./solution.py).

Как только мы отправим 50-ый ответ на задачу, получим флаг!

### Флаг

```
vka{perhaps_we_should_have_used_a_smaller_list}
```
