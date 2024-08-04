go-prod: 
	DEPLOY_PLATFORM="prod" air -c .air.toml
go-test: 
	DEPLOY_PLATFORM="test" air -c .air.test.toml
go-dev: 
	DEPLOY_PLATFORM="dev" air -c .air.dev.toml

zip:
	zip ../site.zip -r * .[^.]* -x "./web/node_modules/*" -x "./web/tailwind.config.js" -x "./web/bun.lockb" -x "./web/package.json" -x ".git/*"

bun-watch:
	bun --watch build --minify --splitting --outdir=./web/static/js/dist ./web/static/js/index.js

dev:
	make -j 2 go-dev bun-watch

bin = bin/site
build:
	go build -o $(bin)
run: build
	DEPLOY_PLATFORM="prod" ./$(bin)

# docker tag IMAGE_ID NEW_TAG_NAME
# docker build -t ii .
# docker-run:
# 	docker run -d -p 80:80 --name go-website site
