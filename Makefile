create:
	curl -X POST http://localhost:3000/v1/user/create \
     -H "Content-Type: application/json" \
     -d '{"password": "12345", "email": "kkk@gmail.com"}'

find:
	curl -X GET http://localhost:3000/v1/user/find-by-pk/1 
     
update:
	curl -X POST http://localhost:3000/v1/user/update \
     -H "Content-Type: application/json" \
     -d '{"id":"1","password": "12345666"}'

delete:
	curl -X POST http://localhost:3000/v1/user/delete \
     -H "Content-Type: application/json" \
     -d '{"id":"1"}'