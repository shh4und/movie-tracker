# Login to get the JWT token
TOKEN=$(curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username": "charlie_brown", "password": "charliepass123"}' | jq -r '.data.token')

echo "JWT Token: $TOKEN"

# Test /users/profile
curl -X GET http://localhost:8080/api/v1/users/profile?username=charlie_brown \
  -H "Authorization: Bearer $TOKEN"

# Test /users/update
curl -X PUT http://localhost:8080/api/v1/users/update \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"username": "charlie_brown_jr", "email": "charlie_brown@example.com", "password": "newpassword"}'

# Test /users/delete
curl -X DELETE http://localhost:8080/api/v1/users/delete \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"id": "1"}'

# Test /titles/search
curl -X GET http://localhost:8080/api/v1/titles/search?title=Joker \
  -H "Authorization: Bearer $TOKEN"