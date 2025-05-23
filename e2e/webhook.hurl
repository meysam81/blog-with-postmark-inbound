GET {{base_url}}/
HTTP 200

POST {{base_url}}/webhook
content-type: application/json
[BasicAuth]
postmarkapp: Secr3t!
file,file1.json;
HTTP 200

POST {{base_url}}/webhook
content-type: application/json
[BasicAuth]
postmarkapp: Secr3t!
file,file2.json;
HTTP 403

POST {{base_url}}/webhook
content-type: application/json
[BasicAuth]
postmarkapp: Secr3t!
file,file3.json;
HTTP 403

POST {{base_url}}/webhook
content-type: application/json
[BasicAuth]
postmarkapp: Secr3t!
file,file4.json;
HTTP 200

POST {{base_url}}/webhook
content-type: application/json
[BasicAuth]
postmarkapp: Secr3t!
file,file5.json;
HTTP 200
