
create:
	curl --request POST 'localhost:3000/clients' \
	--header 'Content-Type: application/json' \
	--data-raw '{  "name":"mahmoud jadan", "birthdate":"1983-01-08", "cpf":"7030000000" }' |jq

getall:
	curl --request GET 'localhost:3000/clients' | jq

getone:
	curl  --request GET 'localhost:3000/clients/1' |jq

delete:
	curl  --request DELETE 'localhost:3000/clients/1' 

update:
	curl --request POST 'localhost:3000/clients/1' \
	--header 'Content-Type: application/json' \
	--data-raw '{     "name":"mahmoud new",   "birthdate":"1983-01-01",   "cpf":"7030000011" }' |jq