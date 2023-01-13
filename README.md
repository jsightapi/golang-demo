# golang-demo

## How to start

1. Go to `./docker/go-{version}/debian`, run `docker-compose up` there.
2. Create `.mock` folder. Create response mock files in the `./mock` folder:  `server.jst`,
   `response_code`, `response_headers`, `response_body`.
3. Send test requests to server `http://localhost:8000`.

## How to test

1. Go to `./docker/go-{version}/debian`, run `docker-compose up` there.
2. Run `docker exec -it golang-demo bash` in other console.
3. Run `go test main.go jsight.go mock.go jsight_test.go` inside the container.