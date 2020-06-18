package service

type UserService interface {
	FindById(id string) (*UserModel, error)
	FindOrCreate(id, userName, screenName string) (*UserModel, error)
}

type TweetService interface {
	FindById(id string) (*TweetModel, error)
	Store(tweet *TweetModel, user *UserModel, resources []ResourceModel) error
	Search(term string, size, start, page int) ([]TweetModel, error)
}

type ResourceService interface {
	FindResourceByTweetId(tweetId string) ([]ResourceModel, error)
}

type SearchService interface {
	Search(term string, size int) ([]SearchModel, error)
	Put(tweetId, fullText, author string) error
}
