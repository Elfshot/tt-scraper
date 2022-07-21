docker build --tag ttscraper .
docker container stop ttscraper
docker container rm ttscraper
docker run -d --restart unless-stopped --name ttscraper ttscraper
docker container logs ttscraper -f