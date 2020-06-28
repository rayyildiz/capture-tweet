import React, {FC} from "react";
import {useParams} from "react-router-dom";
import {useQuery} from "@apollo/client";
import {TWEET_GQL} from "../../graph/queries";
import {Tweet, Tweet_tweet, TweetVariables} from "../../graph/Tweet";
import {ImageCard} from "./ImageCard";
import moment from "moment";


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
        <div className="row">
          <div className="col-12">
            {error && <div className="alert alert-warning margin-top-1 alert-box mx-auto">
              <p className="mb-0">{error.message}</p>
            </div>
            }
          </div>
          {data && data.tweet && <TweetDetail key={`detail_${data.tweet.id}`} tweet={data.tweet}/>}
        </div>

      </div>
  )
};

type TweetDetailProps = {
  tweet: Tweet_tweet
}
const TweetDetail: FC<TweetDetailProps> = ({tweet}) => (
    <>
      <div className="col-md-6 col-lg-6 col-sm-12">
        <ImageCard key={`img_${tweet.id}`} id={tweet.id}/>
      </div>
      <div className="col-md-6 col-lg-6 col-sm-12">
        <h4>Tweet by <a target="_blank" rel="noopener noreferrer" className="text-decoration-none" href={`https://twitter.com/${tweet.author?.userName}`}>{tweet.author?.screenName}</a></h4>
        <br/>
        <a target="_blank" rel="noopener noreferrer" className=" text-muted text-black text-justify text-wrap text-decoration-none" href={`https://twitter.com/${tweet.author?.userName}/status/${tweet.id}`}>{tweet.fullText}</a>
        <br/>
        <div className="col-12 text-right text-secondary">
          <br/>

          <p>Posted at {moment(tweet.postedAt).format("DD-MM-YYYY HH:MM")}</p>
        </div>
        <div className="col-12 text-align-left">
          {tweet.resources?.map(res => (
              <a key={`media_${tweet.id}_${res?.id}`} href={res?.url}>
                <img style={{maxHeight: '15rem'}} className={"img-thumbnail rounded float-left"} src={res?.url} alt="" key={`img_res_${res?.id}`}/>
              </a>
          ))}
        </div>
      </div>

    </>)

export default TweetPage;
