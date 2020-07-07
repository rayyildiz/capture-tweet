import React, {FC, useEffect} from "react";
import {useParams} from "react-router-dom";
import {useQuery} from "@apollo/client";
import {TWEET_BY_ID, TWEET_IMAGE} from "../graph/queries";
import {Tweet, Tweet_tweet, TweetVariables} from "../graph/Tweet";
import {Helmet} from "react-helmet";
import moment from "moment";
import {TweetImage, TweetImageVariables} from "../graph/TweetImage";
import folderImage from './../assets/folder.svg';
import {WEB_BASE_URL} from "../Constants";

const TweetPage: FC = () => {
  const {id} = useParams();

  const {data, loading, error} = useQuery<Tweet, TweetVariables>(TWEET_BY_ID, {
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
          {data && data.tweet && <TweetDetail key={`detail_${data.tweet.id}`} tweet={data.tweet}/>
          }
        </div>

      </div>
  )
};

type TweetDetailProps = {
  tweet: Tweet_tweet
}
const TweetDetail: FC<TweetDetailProps> = ({tweet}) => (
    <>
      <Helmet>
        <meta property="og:title" content={tweet.fullText}/>
        <meta property="og:url" content={`${WEB_BASE_URL}/tweet/${tweet?.id}`}/>
        <meta property="og:image" content={`${WEB_BASE_URL}/${tweet?.captureURL}`}/>
        <title>Capture Tweet | Tweet by {tweet.author?.userName}</title>
      </Helmet>

      <div className="col-md-6 col-lg-6 col-sm-12">
        <TweetImageCard key={`img_${tweet.id}`} id={tweet.id}/>
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

    </>);


type TweetImageCardProps = {
  id: string
}
const TweetImageCard: FC<TweetImageCardProps> = ({id}) => {
  const {data, startPolling, stopPolling} = useQuery<TweetImage, TweetImageVariables>(TWEET_IMAGE, {
    variables: {
      id
    },
  })

  useEffect(() => {
    startPolling(1900);
    return () => stopPolling();
  }, [id, startPolling, stopPolling])

  if (data && data?.tweet && data.tweet.captureURL) {
    stopPolling();

    return (<img src={`/${data.tweet.captureURL}`} alt="" className="img-fluid"/>);

  }

  return (
      <>
        <img src={folderImage} alt="" className="img-fluid"/>
        <br/><br/>
        <div className="spinner-border" role="status">
          <span className="sr-only">Loading...</span>
        </div>
      </>
  );
}


export default TweetPage;
