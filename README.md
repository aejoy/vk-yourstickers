<img align="center" src="preview.png">
<p align="center">Бот ВКонтакте для просмотра наборов стикеров пользователя</p>

<div align="center">
    <a href="http://www.opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/license-MIT-fuchsia.svg" />
    </a>
    <a href="https://goreportcard.com/report/github.com/aejoy/vk-yourstickers">
        <img src="https://goreportcard.com/badge/github.com/aejoy/vk-yourstickers" />
    </a>
</div>

<h3 align="center">Инстансы</h3>

| Название | Создатель      |
| --- |----------------|
| [СтикерБот](https://vk.com/yourstickersbot) | [Алмаз Вайс](https://vk.com/shizumico) |

<h3 align="center">Технологии</h3>

- Язык программирования: **Go**
- База данных: **PostgresQL**
- Кэш-хранилище: **Redis**
- Контейнеризация: **Docker**
- Оркестрация: **Docker Compose**

<h3 align="center">Секреты</h3>

- **SSH_HOST** - IP-адрес VDS
- **SSH_USER** - root пользователь VDS
- **SSH_PASSWORD** - пароль от SSH VDS
- **WORKDIR** - каталог, где находится docker-compose.yaml из [Deployments](/deployments)
- **ALBUM_ID** - ID альбома сообщества ВКонтакте
- **POSTGRES_URL** -  URL подключения PostgresQL (обычно "postgres://vk-yourstickers:ReAVECncNy5hpkQG09wk9yFJ6NYTrc0A@vk-yourstickers-db:5432?sslmode=disable")
- **REDIS_URL** - URL подключения Redis (обычно vk-yourstickers-cache:6379)
- **TOKEN** - токен сообщества ВКонтакте
- **USER_TOKEN** - токен пользователя ВКонтакте (с возможностью запрашивать чужие стикеры)
- **GH_SECRET** - Classic персональный секрет GitHub ([\*тык*](https://github.com/settings/tokens))