# This is example of usage HTTPS with Self-signed Certificate (but browser warns client that is untrusted source) and with mkcert that allows us to create local CA !

## Firstly, INSTRUCTION !!!

### How to make self-signed Certificate (but browser will warn client with untrusted connection!!!)

* install openssl

* and then run
`openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout localhost.key -out localhost.crt -subj "/C=US/ST=State/L=City/O=YourOrg/CN=localhost" `

#### This will create a self signed certs localhost.key and  localhost.crt
#### If needed you need to move these files to project folder

#### then use them with srv.ListenAndServerTLS("localhost.crt ", "localhost.key")

### How to make Local CA for making the server Trustable for Browser !!!

* install choco
* with `choco install mkcert` that is Locally CA (like GoDaddy)
* with `mkcert -install`
* move to project folder
* with `mkcert localhost 127.0.0.1 ::1` added certs to folder your are currently on
* and use this created `"localhost+2.pem"`, `"localhost+2-key.pem"` certs in server

## And how to use them, you can find in server/main.go
