import React, {FC, useEffect, useRef, useState} from "react";
import {Link, useParams} from "react-router-dom";
import {useQuery} from "@apollo/client";
import {TWEET_BY_ID, TWEET_IMAGE} from "../graph/queries";
import {Tweet, Tweet_tweet, TweetVariables} from "../graph/Tweet";
import {Helmet} from "react-helmet";
import moment from "moment";
import {TweetImage, TweetImageVariables} from "../graph/TweetImage";
import folderImage from './../assets/folder.svg';
import {WEB_BASE_URL} from "../Constants";
import 'bootstrap/js/dist/modal';
import ContactPage from "./ContactPage";
import Loading from "../components/Loading";
import Modal from "bootstrap/js/src/modal";

const TweetPage: FC = () => {
  const {id} = useParams();

  const {data, loading, error} = useQuery<Tweet, TweetVariables>(TWEET_BY_ID, {
    variables: {id}
  });

  if (loading) {
    return <Loading/>
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
const TweetDetail: FC<TweetDetailProps> = ({tweet}) => {
  const [render, setRender] = useState(false);

  const modal = useRef<HTMLDivElement>(null);
  const showModal = () => {
    if (modal.current != null) {
      const m = new Modal(modal.current);
      if (m != null) m.show();
    }
  }

  const openModal = () => {
    if (render) {
      showModal();
    } else {
      setRender(true);
      setTimeout(() => {
        showModal();
      }, 1500);
    }
  }

  return (
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
          <div className="row">
            <div className="col-8 text-left">
              <h4>Tweet by <Link to={`/user/${tweet.author?.id}`} rel="noopener noreferrer" className="text-decoration-none">{tweet.author?.screenName}</Link></h4>
            </div>
            <div className="col-4 text-right">
              <button type="button" className="btn btn-link" onClick={openModal}>
              <span data-toggle="tooltip" data-placement="top" title="Report this tweet">
                 <svg width="1em" height="1em" viewBox="0 0 16 16" className="bi bi-info-circle" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>
                <path d="M8.93 6.588l-2.29.287-.082.38.45.083c.294.07.352.176.288.469l-.738 3.468c-.194.897.105 1.319.808 1.319.545 0 1.178-.252 1.465-.598l.088-.416c-.2.176-.492.246-.686.246-.275 0-.375-.193-.304-.533L8.93 6.588z"/>
                <circle cx="8" cy="4.5" r="1"/>
              </svg>
              </span>
              </button>
            </div>
          </div>

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
        {render &&
        <div className="modal fade" tabIndex={-1} role="dialog" ref={modal}>
          <div className="modal-dialog modal-xl">
            <div className="modal-content">
              <div className="modal-header">
                <h5 className="modal-title" id="reportTweetModalLabel"> Report Tweet </h5>
                <button type="button" className="close" data-dismiss="modal" aria-label="Close">
                  <span aria-hidden="true">&times;</span>
                </button>
              </div>
              <div className="modal-body">
                <ContactPage tweetID={tweet.id}/>
              </div>
            </div>
          </div>
        </div>
        }
      </>);
}


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
    setTimeout(() => {
      stopPolling();
    }, 10 * 1900);

    return () => stopPolling();
  }, [id, startPolling, stopPolling])

  if (data && data?.tweet && data.tweet.captureURL) {
    stopPolling();

    return (<img src={`/${data.tweet.captureURL}`} alt="" className="img-fluid"/>);

  }

  return (
      <>
        <img src={folderImage} alt="" className="img-fluid"/>
        <br/>
        <div className="d-flex align-items-start">
          <div className="spinner-grow text-danger" role="status" aria-hidden="true"></div>
          <span className="text-muted" style={{paddingLeft: '1rem'}}>Tweet screenshot is capturing, please wait...</span>
        </div>

      </>
  );
}


export default TweetPage;
