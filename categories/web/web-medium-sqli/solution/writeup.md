## Обмани систему доставки!


| Событие | Название | Категория | Сложность |
| :------ | ---- | ---- | ---- |
| VKAKIDS 2024 | Shop&Show | Web | medium |


### Описание


> Автор: [d4nd3l10n]
>
*Самый быстрый курьер в мире доставляет абсолютно все заказы за 5 секунд. Он никогда не опаздывает, ведь если опоздает - придется сделать подарок за счет заведения...*


### Решение
Анализируем код и видим в функции `UpdateAddressHandler` запрос к базе данных, отличающийся от остальных, который как раз оказывается уязвимым для *SQL-инъекции*:
```
query := fmt.Sprintf("UPDATE orders SET address = '%s' \n", newAddress)
    query = query + "WHERE id = ? AND is_delivered = 0 AND user_id = ?"
	_, err = database.DB.ExecContext(ctx, query, orderID, userID)
```
В данном запросе используется незащищенный ввод данных и происходит обновление адреса доставки товаров.
Также в исходниках, уже в функции `OrdersHistoryHandler` мы видим очень интересующий нас блок кода, потому что заветное слово **flag** прямо намекает на ход решения задания:
```
if order.DeliveryTime > 5 {
			order.Comment = readFlag()
		}
```
Через уязвимый запрос мы SQL-инъекцией обновляем не только адрес, но и время доставки `delivery_time`.
Захотев это сделать вручную, мы понимаем, что заказ висит во вкладке **Активные заказы** только 5 секунд => приходим к тому, что данный процесс необходимо автоматизировать, после чего, прокинув SQL-инъекцию, получаем заветный флаг

[solution.py](./solution.py)

### Флаг

```
vka{4utom4ted_sql_1nj3ct10n_5ucc3ss}
```