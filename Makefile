create:
	curl -X POST http://localhost:3000/v1/user/create \
     -H "Content-Type: application/json" \
     -d '{"password": "12345", "email": "node@gmail.com"}'

find:
	curl -X GET http://localhost:3000/v1/user/find-by-pk/1 