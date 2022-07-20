set GOOS=linux
set GOARCH=amd64

go build
@REM docker build --tag ttscraper .
@REM docker container stop ttscraper
@REM docker container rm ttscraper
@REM docker run -d --restart unless-stopped --name ttscraper ttscraper