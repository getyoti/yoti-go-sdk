module digitalidentity

go 1.19

require (
	github.com/getyoti/yoti-go-sdk/v3 v3.0.0
	github.com/joho/godotenv v1.3.0
)

require google.golang.org/protobuf v1.30.0 // indirect

replace github.com/getyoti/yoti-go-sdk/v3 => ../../
