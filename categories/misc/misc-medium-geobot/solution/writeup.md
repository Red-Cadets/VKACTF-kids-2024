## MISC

| Событие | Название | Категория | Сложность |
| :------ | ---- | ---- | ---- |
| VKAKIDS 2024 | geobot | misc | medium |


### Описание


> Автор: [Inssurg3nt]
>
Наши осведомители донесли информацию от некоего иностранного разработчика под псевдонимом ***P414-gtx***, что на одной из зарубежных военных баз США начали производить секретное чудо-оружие. Однако, он не сказал, на какой именно, лишь упомянул о том, что база носит имя какого то сержанта...


### Решение
- Из описания, самое важное, нужно было понять, что если дан ник разработчика, то следует искать его на [гитхабе](https://github.com/P414-gtx).
- На аккаунте найденного пользователя всего один репозиторий - geobot. Изучаем его, не забываем про историю коммитов.
- В коммитах главное найти две зацепки: ссылку на самого бота и пулл координат разных баз.
- Изучив функционал бота по сорцам и взаимодействию с ним понимаем, что из этого пулла баз нам необходимо найти базу Кэмп Бондстил, которая находится в Сербии. (Именно в Сербии, так как наша страна не признала независимость Косово)


- Итак, нужна строчка для бота: 42.362756, 21.247567, Сербия

### Флаг

```
vka{Ru5Su1aN_IN73LLi63NcE_5erv1CE}
```