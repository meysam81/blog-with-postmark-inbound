GET http://localhost:8000/
HTTP 200

POST http://localhost:8000/webhook
content-type: application/json
[BasicAuth]
postmark: secret
file,file1.json;
HTTP 200

POST http://localhost:8000/webhook
content-type: application/json
[BasicAuth]
postmark: secret
file,file2.json;
HTTP 200

POST http://localhost:8000/webhook
content-type: application/json
[BasicAuth]
postmark: secret
file,file3.json;
HTTP 200

POST http://localhost:8000/webhook
content-type: application/json
[BasicAuth]
postmark: secret
file,file4.json;
HTTP 200
