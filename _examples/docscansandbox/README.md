# IDV (Doc Scan) Sandbox Example

An example of configuring a session response can be found in the [Demo Test](./demo_test.go)

## Configuring the application
- Copy the [.env.example](.env.example) file and rename it to be `.env`, then
  modify the `YOTI_SANDBOX_CLIENT_SDK_ID` and `YOTI_KEY_FILE_PATH` environment variables.

## Running the test

1. Run `docker-compose up -d` to start the test app
1. Run `godog`
