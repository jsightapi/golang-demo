# JSight API Validator demo (golang)

1. Copy from the `./lib/{plugin-version}/go-{version}` folder appropriate plugin binary
   `jsightplugin.so` to the `./src` folder.
2. Go to `./src` folder.
3. Run server with `go run main.go jsight.go`.
4. Run "good" test request: `curl --request POST localhost:8000/users -d '{"id":123,"login":"john"}'`.
5. Run "bad" test request and get a validation error: `curl --request POST localhost:8000/monsters -d '{"id":123,"login":"john"}'`.

NOTE: for running on Windows read `./windows/README.md`.
