GET {{base_url}}/
HTTP 200

POST {{base_url}}/webhook
content-type: application/json
[BasicAuth]
postmark: secret
file,file1.json;
HTTP 200

POST {{base_url}}/webhook
content-type: application/json
[BasicAuth]
postmark: secret
file,file2.json;
HTTP 200

POST {{base_url}}/webhook
content-type: application/json
[BasicAuth]
postmark: secret
file,file3.json;
HTTP 200

POST {{base_url}}/webhook
content-type: application/json
[BasicAuth]
postmark: secret
file,file4.json;
HTTP 200
