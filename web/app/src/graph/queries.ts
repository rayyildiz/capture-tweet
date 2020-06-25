import {gql} from '@apollo/client';


export const CAPTURE_TWEET_GQL = gql`
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

export const SEARCH_GQL = gql`
  query Search($input: SearchInput!) {
    search(input: $input) {
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
export const TWEET_GQL = gql`
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

export const TWEET_IMAGE_QUERY = gql`
  query TweetImage($id:ID!) {
    tweet(id:$id) {
      id
      captureURL
    }
  }
`;
