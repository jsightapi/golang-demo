# golang-demo

## How to start

1. Go to `./docker/go-18/nginx-fpm/debian`, run `docker-compose up` there.
2. Create response mock files in the `./mock` folder:  `server.jst`, `response_code`,
   `response_headers`, `response_body`.
3. Send test requests to server `http://localhost:8000`.