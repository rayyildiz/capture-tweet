/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: Tweet
// ====================================================

export interface Tweet_tweet_author {
  __typename: "Author";
  userName: string;
  screenName: string | null;
  profileImageURL: string | null;
}

export interface Tweet_tweet_resources {
  __typename: "Resource";
  id: string;
  url: string;
  mediaType: string | null;
  width: number | null;
  height: number | null;
}

export interface Tweet_tweet {
  __typename: "Tweet";
  id: string;
  fullText: string;
  lang: string | null;
  postedAt: any | null;
  captureURL: string | null;
  retweetCount: number | null;
  favoriteCount: number | null;
  author: Tweet_tweet_author | null;
  resources: (Tweet_tweet_resources | null)[] | null;
}

export interface Tweet {
  tweet: Tweet_tweet | null;
}

export interface TweetVariables {
  id: string;
}
