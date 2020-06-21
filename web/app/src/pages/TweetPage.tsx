import React, {FC} from "react";
import {useParams} from "react-router-dom";
import {useQuery} from "@apollo/client";
import {TWEET_GQL} from "../graph/queries";
import {Tweet, TweetVariables} from "../graph/Tweet";


const TweetPage: FC = () => {
  const {id} = useParams();

  const {data, loading, error} = useQuery<Tweet, TweetVariables>(TWEET_GQL, {
    variables: {id}
  });

  if (loading) {
    return <span>Loading...</span>
  }


  return (
      <div>
        <h3>Tweet Detail</h3>

        {error && <div className="alert alert-warning margin-top-1 alert-box mx-auto">
          <p className="mb-0">{error.message}</p>
        </div>
        }

        {data && JSON.stringify(data.tweet)}
      </div>
  )
};

export default TweetPage;
