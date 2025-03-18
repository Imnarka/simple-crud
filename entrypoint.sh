#!/bin/bash
set -e
echo "Запуск миграций"
make migrate

echo "Старт приложения"
exec "$@"
