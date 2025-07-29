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

#### It is just quite simple example, and next things can be:

* HSTS (HTTP Strict Transport Security): learn to send the Strict-Transport-Security header to force browsers onto HTTPS and prevent protocol downgrades.

* OCSP stapling & CRL: how to check certificate revocation in real time.

* HTTP/2 & HTTP/3 (QUIC): enabling them over TLS and measuring performance gains.

* Perfect Forward Secrecy (PFS): choosing cipher suites (e.g. ECDHE) to ensure session keys can’t be retroactively decrypted.

* Mutual TLS (mTLS): requiring client certs for two‑way authentication (common in microservices)!

* Security headers: CSP, HPKP (deprecated but instructive), X-Frame-Options, etc.

* TLS configuration testing: using tools like SSL Labs or testssl.sh to harden your server’s config.
