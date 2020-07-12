import React, {FC, FormEvent} from 'react';
import * as qs from 'query-string';
import {useQuery} from "@apollo/client";
import {SEARCH_TWEET} from "../graph/queries";
import {Search, Search_search, SearchVariables} from "../graph/Search";
import {useHistory} from "react-router-dom";
import moment from 'moment';
import notFound from '../assets/not_found.svg';
import './SearchPage.css';
import {Helmet} from "react-helmet";
import {WEB_BASE_URL} from "../Constants";
import algoliaLogo from '../assets/search-by-algolia-light-background.svg';

const getQueryStringValue = (key: string, queryString = window.location.search): string => {
  const values = qs.parse(queryString);
  return values[key] as string;
};


const SearchPage: FC = () => {
  const q = getQueryStringValue('q');
  const {data, loading, error} = useQuery<Search, SearchVariables>(SEARCH_TWEET, {
    variables: {
      input: {
        term: q
      }
    }
  });
  if (loading) {
    return <span>Loading</span>
  }

  return <>
    <Helmet>
      <meta property="og:title" content={`Search results for ${q}`}/>
      <meta property="og:url" content={`${WEB_BASE_URL}/search?q=${q}`}/>
      <meta property="og:image" content={`${WEB_BASE_URL}${notFound}`} />
      <title>Capture Tweet | Search</title>
    </Helmet>

    <div>
      <div className="row">
        <div className="col-8">
          <h4>Search results for <b>{q} </b></h4>
        </div>
        <div className="col-4 text-right">
          <img src={algoliaLogo} alt="" style={{height:'1rem'}}/>
        </div>
      </div>
      {error && <div className="alert alert-dismissible alert-warning">
        <p className="mb-0">{error.message}</p>
      </div>}

      <div className="row">
        {data && data.search?.map(t => <TweetCard key={t?.id} tweet={t}/>)}
      </div>
    </div>
  </>
}

type TweetCardProps = {
  tweet: Search_search | null
}

const TweetCard: FC<TweetCardProps> = ({tweet}) => {
  const history = useHistory();

  if (!tweet) {
    return <span>Error</span>
  }

  const handleClick = (e: FormEvent) => {
    e.preventDefault();
    history.push("/tweet/" + tweet.id);
  }

  return (
      <div className="col-sm-4">
        <div className="card mb-3 cursor tweet-card" onClick={handleClick}>
          <h3 className="card-header">Tweet by {tweet.author?.userName}</h3>
          <div className="card-body">
            <h5 className="card-title">Posted at {moment(tweet.postedAt).format("DD-MM-YYYY HH:MM")} </h5>
            <h6 className="card-subtitle text-muted">Language {tweet.lang}</h6>
          </div>
          <div style={{
            width: '100%',
            textAlign: 'center',
            overflow: 'hidden',
            height: '10rem'
          }}>
            {tweet.captureThumbURL ? <img style={{maxWidth: '20rem'}} src={`/${tweet.captureThumbURL}`} alt=""/>
                : <img style={{maxWidth: '20rem'}} src={notFound} alt=""/>}
          </div>
          <div className="card-body">
            <p className="card-text">{tweet.fullText} </p>
          </div>
        </div>
      </div>
  )
}

export default SearchPage;
