import React, {FC} from 'react';
import * as qs from 'query-string';
import {useQuery} from "@apollo/client";
import {SEARCH_TWEET} from "../graph/queries";
import {Search, SearchVariables} from "../graph/Search";
import notFound from '../assets/not_found.svg';
import './SearchPage.css';
import {Helmet} from "react-helmet";
import {WEB_BASE_URL} from "../Constants";
import algoliaLogo from '../assets/search-by-algolia-light-background.svg';
import {TweetCard} from "./TweetCard";

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
      <meta property="og:image" content={`${WEB_BASE_URL}${notFound}`}/>
      <title>Capture Tweet | Search</title>
    </Helmet>

    <div>
      <div className="row">
        <div className="col-8">
          <h4>Search results for <b>{q} </b></h4>
        </div>
        <div className="col-4 text-right">
          <img src={algoliaLogo} alt="" style={{height: '1rem'}}/>
        </div>
      </div>
      {error && <div className="alert alert-dismissible alert-warning">
        <p className="mb-0">{error.message}</p>
      </div>}

      <div className="row">
        {data && data.search?.map(t => (t && <TweetCard key={t.id}
                                                        author={t.author?.userName}
                                                        fullText={t.fullText ?? ""}
                                                        id={t.id ?? ""}
                                                        captureThumbURL={t.captureThumbURL ?? ""}
                                                        lang={t.lang ?? ""}
                                                        postedAt={t.postedAt}/>
        ))}
      </div>
    </div>
  </>
}

export default SearchPage;
