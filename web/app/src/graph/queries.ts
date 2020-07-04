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
    search(input: $input) {
      id
      fullText
      lang
      postedAt
      captureURL
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
  mutation Contact($input:ContactInput!) {
    contact(input:$input)
  }`;
