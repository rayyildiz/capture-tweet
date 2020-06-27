import React, {FC, useEffect} from "react";
import {useQuery} from "@apollo/client";
import {TWEET_IMAGE_QUERY} from "../../graph/queries";
import {TweetImage, TweetImageVariables} from "../../graph/TweetImage";
import folderImage from '../../assets/folder.svg';

type ImageCardProps = {
  id: string
}

export const ImageCard: FC<ImageCardProps> = ({id}) => {
  const {data, startPolling, stopPolling} = useQuery<TweetImage, TweetImageVariables>(TWEET_IMAGE_QUERY, {
    variables: {
      id
    },
  })

  useEffect(() => {
    startPolling(1500);

    return () => stopPolling();
  }, [id])

  if (data && data?.tweet && data.tweet.captureURL) {
    stopPolling();

    return (<img src={`/${data.tweet.captureURL}`} alt={"No image"} className="img-fluid"/>);

  }

  return (
      <div className="d-flex justify-content-center">
        <img src={folderImage} alt={"No image"} className="img-fluid"/>
        <br/>
        <div className="spinner-border" role="status">
          <span className="sr-only">Loading...</span>
        </div>

      </div>
  );
}

