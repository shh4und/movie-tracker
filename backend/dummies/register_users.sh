for user in $(cat users.json | jq -c '.[]'); do
  curl -X POST http://localhost:8080/api/v1/register \
    -H "Content-Type: application/json" \
    -d "$user"
  echo
  sleep 1
done