build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tmp/app .

docker:
	docker build -f Dockerfile --tag=eu.gcr.io/capturetweet/beta-capture:local .

push:
	docker push eu.gcr.io/capturetweet/beta-capture:local

run:
	gcloud run deploy beta-capture --image eu.gcr.io/capturetweet/beta-capture:local --platform managed --region europe-west1 --project capturetweet
