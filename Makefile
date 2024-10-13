# SQLをコンテナに流す
exec-schema:
	cat ./DDL/*.up.sql > ./DDL/schema.sql
	docker cp DDL/schema.sql auth-demo-db:/ && docker exec -it auth-demo-db psql -U postgres -d auth_demo -f /schema.sql
	rm ./DDL/schema.sql
exec-dummy:
	docker cp DDL/insert_dummy_data.sql auth-demo-db:/ && docker exec -it auth-demo-db psql -U postgres -d auth_demo -f /insert_dummy_data.sql

# テーブルをリフレッシュ
refresh-schema:
	@make exec-schema
	@make exec-dummy