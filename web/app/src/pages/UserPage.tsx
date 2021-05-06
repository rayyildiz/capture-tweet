import {gql, useQuery} from "@apollo/client";
import notFound from '../assets/not_found.svg';
import './SearchPage.css';
import {Helmet} from "react-helmet";
import {WEB_BASE_URL} from "../Constants";
import {TweetCard} from "./TweetCard";
import {useParams} from 'react-router-dom';
import {SearchByUser, SearchByUserVariables} from "./__generated__/SearchByUser";


const SEARCH_BY_USER = gql`
  query SearchByUser($userID: ID!) {
    searchByUser(userID: $userID) {
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

const UserPage = () => {
  const {id} = useParams<{ id: string }>();

  const {data, loading, error} = useQuery<SearchByUser, SearchByUserVariables>(SEARCH_BY_USER, {
    variables: {
      userID: id
    }
  });
  if (loading) {
    return <span>Loading</span>
  }

  return <>
    <Helmet>
      <meta property="og:title" content={`Search results for ${id}`}/>
      <meta property="og:url" content={`${WEB_BASE_URL}/user/${id}`}/>
      <meta property="og:image" content={`${WEB_BASE_URL}${notFound}`}/>
      <title>Capture Tweet</title>
    </Helmet>

    <div>
      <div className="row">
        <div className="col-8">
          <h4>User Tweets</h4>
        </div>
      </div>
      {error && <div className="alert alert-dismissible alert-warning">
        <p className="mb-0">{error.message}</p>
      </div>}

      <div className="row">
        {data && data.searchByUser?.map(t => (t && <TweetCard key={t.id}
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


export default UserPage;
