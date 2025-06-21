export GOOS=linux
export GOARCH=arm64
export CGO_ENABLED=0
.DEFAULT_GOAL := deploy

deploy:
	go build -o bootstrap
	zip -r lambda-handler.zip bootstrap
	aws lambda update-function-code --function-name "my-go-lambda" --zip-file fileb://lambda-handler.zip --region="ap-southeast-2" | jq .   

