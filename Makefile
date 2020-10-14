run-dynamodb:
	docker kill dynamodb-lookup || true
	docker run --rm -d --name dynamodb-lookup -p 8000:8000 amazon/dynamodb-local

build:
	sh ./build.sh

deploy: build
	sls deploy -s dev

deploy-local: build
	AWS_PROFILE=personal sls deploy -s dev