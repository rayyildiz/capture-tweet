package service

type UserService interface {
	FindById(id string) (*UserModel, error)
	FindOrCreate(id, userName, screenName string) (*UserModel, error)
}

type TweetService interface {
	FindById(id string) (*TweetModel, error)
	Store(tweet *TweetModel, user *UserModel, resources []ResourceModel) error
	Search(term string, size int, cursorId *string) ([]TweetModel, error)
}

type ResourceService interface {
	FindResourceByTweetId(tweetId string) ([]ResourceModel, error)
}
