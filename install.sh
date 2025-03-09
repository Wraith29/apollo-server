go build ./cmd/apollo
mv ./apollo /usr/sbin/apollo

go build ./cmd/apollo-server
mv ./apollo-server /usr/sbin/apollo-server
