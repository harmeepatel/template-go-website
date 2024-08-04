# docker tag IMAGE_ID NEW_TAG_NAME
# docker build -t ii .
docker-run:
	docker run -d -p 80:80 --name go-website site
