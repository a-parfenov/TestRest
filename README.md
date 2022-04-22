# Тестовое задание на Golang
### Написать локальный сервер

Условия:
- Первый эндпойнт принимает файл и делит его на части по мегабайту,
если какая-то часть получается меньше 1 мб, то дописывает туда байты (нулями) и сохраняет у себя 
- Все части файла не должны быть меньше 1 мб, даже если файл сам весит несколько КБ

- Второй эндпойнт отдает файл по названию, перед этим собирает все части ранее загруженного файла обратно в один 
и отдает пользователю 
- Обязательно убрать нули из части, которая ими дописана

То есть получается у клиента два действия: он отправляет файл на загрузку и скачивает этот файл по названию

Дополнительно:
- Файлы, из которых состоит входящий файл лежат в отдельной папке (создается при сохранении)
- Собранный файл сохраняется также в отдельную папку


Установка и запуск
------------
```
git clone https://github.com/a-parfenov/TestRest.git && cd TestRest && make run
```

```
http://localhost:8080/upload
```