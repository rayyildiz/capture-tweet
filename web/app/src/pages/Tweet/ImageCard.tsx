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
    startPolling(2000);

    return () => stopPolling();
  }, [id])

  if (data && data?.tweet && data.tweet.captureURL) {
    stopPolling();

    return <>
      image : {JSON.stringify(data?.tweet)}
      <br/>
    </>
  }

  return <img src={folderImage} alt={"No image"} className="img-fluid"/>
}

