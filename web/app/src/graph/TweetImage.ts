/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: TweetImage
// ====================================================

export interface TweetImage_tweet {
  __typename: "Tweet";
  id: string;
  captureURL: string | null;
}

export interface TweetImage {
  tweet: TweetImage_tweet | null;
}

export interface TweetImageVariables {
  id: string;
}
