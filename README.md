# CaptureTweet

[![build](https://github.com/rayyildiz/capture-tweet/actions/workflows/ci.yml/badge.svg)](https://github.com/rayyildiz/capture-tweet/actions/workflows/ci.yml)
[![pull-request](https://github.com/rayyildiz/capture-tweet/actions/workflows/pr.yml/badge.svg)](https://github.com/rayyildiz/capture-tweet/actions/workflows/pr.yml)
[![Security Scan](https://github.com/rayyildiz/capture-tweet/actions/workflows/security_scan.yml/badge.svg)](https://github.com/rayyildiz/capture-tweet/actions/workflows/security_scan.yml)

I am going to close capturetweet on 31.12.2023, which I started at Starbucks to try [cloud run](https://cloud.google.com/run) features, 
in accordance with [Twitter's pricing policy](https://www.engadget.com/twitter-announces-new-api-pricing-including-a-limited-free-tier-for-bots-005251253.html).
Anyone who wishes can deploy the application to the GCP cloud run. All necessary [CI /CD pipelines](.github/workflows/ci.yml) are ready.

## How

![](docs/CaptureTweet.png)

## Tutorial

[Golang Tutorial](./golang.md)

## Configuration

```dotenv
DOCSTORE_TWEETS=mongo://capturetweet/tweets?id_field=id
DOCSTORE_USERS=mongo://capturetweet/authors?id_field=id
MONGO_SERVER_URL=mongodb://root:123456@localhost:27017
TOPIC_CAPTURE=mem://captureRequest
BLOB_BUCKET=file:///tmp/capture
GRAPHQL_ENABLE_PLAYGROUND=true
TWITTER_ACCESS_SECRET=
TWITTER_ACCESS_TOKEN=
TWITTER_CONSUMER_KEY=
TWITTER_CONSUMER_SECRET=
ALGOLIA_SECRET=
ALGOLIA_CLIENT_ID=
ALGOLIA_INDEX=tweets-LOCAL
```

## RoadMap

- [x] Create a skeleton project
- [x] Add graphql support.
- [x] Tweet Service
  - [x] CRUD for tweet service
  - [x] Use algolia for search
  - [x] Store user additional data in a different collection.
- [x] Async capture
  - [x] PubSub support
  - [x] Capture and update collection
- [x] New UI design
- [x] Apollo client support
- [x] Capture tweet
  - [ ] Real time image update with firebase js.
  - [ ] graphql subscribe ??? '
  - [x] Report capture
- [x] Search tweet
- [x] Contact us
