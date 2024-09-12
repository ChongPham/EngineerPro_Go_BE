curl -X POST http://localhost:8080/register -d "{\"username\":\"trong5\", \"password\":\"123456\", \"profile\":\"\"}" -H "Content-Type: application/json"
curl -X POST http://localhost:8080/login -d "{\"username\":\"trong5\", \"password\":\"123456\"}" -H "Content-Type: application/json"
curl -X PUT http://localhost:8080/profile -d "{\"profile\":\"Cập nhật thông tin cá nhân\", \"password\":\"trong123\"}" -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjY0MjcxMzgsInVzZXJuYW1lIjoidHJvbmc1In0.f-mwyL3o1NgPm-C87URItvgK8V7SwRwxy6Prul_Knck"
