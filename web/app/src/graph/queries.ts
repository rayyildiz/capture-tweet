import {gql} from '@apollo/client';


export const CAPTURE_TWEET = gql`
  mutation Capture($url:String!) {
    capture(url:$url) {
      id
      fullText
      favoriteCount
      retweetCount
      postedAt
      author {
        userName
        screenName
        profileImageURL
      }
    }
  }
`;

export const SEARCH_TWEET = gql`
  query Search($input: SearchInput!) {
    search(input: $input, size: 21) {
      id
      fullText
      lang
      postedAt
      captureThumbURL
      author {
        userName
        screenName
        profileImageURL
      }
    }
  }
`;

export const SEARCH_BY_USER = gql `
  query SearchByUser($userID: ID!) {
    searchByUser(userID: $userID) {
      id
      fullText
      lang
      postedAt
      captureThumbURL
      author {
        userName
        screenName
        profileImageURL
      }
    }
  }
`;

export const TWEET_BY_ID = gql`
  query Tweet($id:ID!) {
    tweet(id:$id) {
      id
      fullText
      lang
      postedAt
      captureURL
      retweetCount
      favoriteCount
      author {
        id
        userName
        screenName
        profileImageURL
      }
      resources {
        id
        url
        mediaType
        width
        height
      }
    }
  }
`;

export const TWEET_IMAGE = gql`
  query TweetImage($id:ID!) {
    tweet(id:$id) {
      id
      captureURL
    }
  }
`;


export const CONTACT_US = gql`
  mutation Contact($input:ContactInput! $id:ID, $captcha:String!) {
    contact(input:$input, tweetID: $id, capthca: $captcha)
  }`;
