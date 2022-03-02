# test-q-change

question2 Postman curl
curl --location --request POST 'localhost:1323/cashier' \
--header 'Content-Type: application/json' \
--data-raw '{
    "productPrice":179,
    "customerPays": 23000
}'