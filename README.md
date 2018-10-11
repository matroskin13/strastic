## Strastic

Simple way to serve your static application as service.

### Example:

```
docker run -p 80:80 -v static:/var/www -e ENV_API_URL=http://example.com -spa matroskin13/strastic
```


```
$ curl http://localhost
<!doctype html><html lang="en"><head><title>Hello world</title></head><body></body>
```

```
$ curl http://localhost/config.json
{"API_URL": "http://example.com"}
```

### Via Dockerfile

```
FROM node:10 as builder

WORKDIR /usr/app

COPY package.json yarn.lock ./
RUN yarn install
COPY . .
RUN yarn run build

FROM matroskin13/strastic
COPY --from=builder /usr/app/build /var/www

ENTRYPOINT ["/strastic", "--spa"]
```
