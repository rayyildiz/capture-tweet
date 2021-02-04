# CaptureTweet

![Build status](https://github.com/rayyildiz/capture-tweet/workflows/build/badge.svg)
![Build status](https://github.com/rayyildiz/capture-tweet/workflows/pull-request/badge.svg)
[![codecov](https://codecov.io/gh/rayyildiz/capture-tweet/branch/main/graph/badge.svg?token=58YR43PZFS)](https://codecov.io/gh/rayyildiz/capture-tweet)
[![Dependabot Status](https://api.dependabot.com/badges/status?host=github&repo=rayyildiz/capture-tweet&identifier=272198663)](https://dependabot.com)


## Local Config

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

- [x] Create skeleton project
- [x] Add graphql support.
- [x] Tweet Service
  - [x] CRUD for tweet service
  - [x] Use algolia for search
  - [x] Store user additional data in a different collection.
- [x] Async capture 
  - [x] Pubsub support
  - [x] Capture and update collection
- [x] New UI design
- [x] Apollo client support
- [ ] Capture tweet 
  - [ ] Real time image update with firebase js.
  - [ ] graphql subscribe ??? 
- [x] Search tweet
