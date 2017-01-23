# Repo for potolok flask app

Порядок установки:
- установить зависимости:
```
~ $ sudo apt-get install python3-pip python3-virtualenv nginx
```
- создать рабочий каталог и перейти в него:
```
~ $ mkdir project && cd $_
```
- создаём новое окружение:
``` 
~/project $ virtualenv -p python3 .
```
- активируем его:
```
~/project $ source bin/activate
(project) ~/project $ 
```
- клонируем этот репозиторий:
```
(project) ~/project $ git clone git@github.com:conformist-mw/potolok.git
```
- переходим в каталог с репозиторием и устанавливаем все необходимые зависимости:
```
(project) ~/project $ cd potolok/
(project) ~/project/potolok $ pip install -r requirements.txt
```
- проверяем работоспособность:
```
(project) ~/project/potolok $ python server.py
 * Running on http://0.0.0.0:5000/ (Press CTRL+C to quit)
```
- выходим:
```
(project) ~/project/potolok $ deactivate
```
- копируем конфиги nginx и systemd в соответствующие каталоги, создаём ссылку, перезапускаем:
```
~/project/potolok $ sudo cp etc/nginx/sites-available/project /etc/nginx/sites-available/
~/project/potolok $ sudo ln -s /etc/nginx/sites-available/project /etc/nginx/sites-enabled/
~/project/potolok $ sudo cp etc/systemd/system/project.service /etc/systemd/system/
~/project/potolok $ sudo systemctl daemon-reload
~/project/potolok $ sudo systemctl start project.service
~/project/potolok $ sudo systemctl enable project.service
~/project/potolok $ sudo systemctl restart nginx.service
```
Теперь можно переходить по [ссылке](http://localhost), должно работать.
