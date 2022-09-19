include .env
export

timestamp := `date +%s`
folder := './migrations/'
filename := ${folder}${timestamp}_`echo ${f}.up.sql | tr ' ' _`

test:
	echo ${filename}

file:
	touch ${filename}

db:
	docker compose up -d neem-db && docker compose exec neem-db bash

migrate:
	docker compose run --rm neem-migrate -path ${folder} -database '$(DB_URL)' up ${n}

rollback:
	docker compose run --rm neem-migrate -path ${folder} -database '$(DB_URL)' down ${n}

seeder:
	docker compose run -it --rm neem-seed bash