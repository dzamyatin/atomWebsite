openssl genpkey -algorithm RSA -out private_key.pem && \
openssl req -new -x509 -key private_key.pem -out certificate.pem -days 365 && \
cat private_key.pem certificate.pem > combined.pem && \
openssl x509 -in certificate.pem -out certificate.crt;



openssl genpkey -algorithm RSA -out private_key.pem
openssl req -new -key private_key.pem -out request.csr
openssl x509 -req -days 365 -in request.csr -signkey private_key.pem -out certificate.crt