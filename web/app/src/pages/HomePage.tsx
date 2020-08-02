import React, {FC, FormEvent, useState} from "react";
import './HomePage.css';
import {gql, useMutation} from "@apollo/client";
import {Redirect} from "react-router-dom";
import {Helmet} from "react-helmet";
import {WEB_BASE_URL} from "../Constants";
import {Capture, CaptureVariables} from "./__generated__/Capture";


const CAPTURE_TWEET = gql`
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


const HomePage: FC = () => {
  const [url, setUrl] = useState('');
  const [validation, setValidation] = useState('');

  const [doQuery, {data, loading, error}] = useMutation<Capture, CaptureVariables>(CAPTURE_TWEET)

  const handleSummit = async (e: FormEvent) => {
    e.preventDefault();
    setValidation('');
    if (url.length <= 5 || !url.startsWith("http")) {
      setValidation("Please enter a valid URL. Example: https://twitter.com/jack/status/20");
    } else {
      try {
        await doQuery({
          variables: {url: url}
        });
      } catch (ex) {
        console.log(ex.toString());
      }
    }
  }

  if (data && data.capture && data.capture.id.length > 0) {
    return <Redirect to={`/tweet/${data.capture.id}`}/>
  }

  return (
      <>
        <Helmet>
          <meta property="og:title" content="Capture Tweet"/>
          <meta property="og:url" content={WEB_BASE_URL}/>
          <meta property="og:image" content={`${WEB_BASE_URL}/logo192.png`}/>
          <title>Capture Tweet | Home</title>
        </Helmet>
        <div className="wrapper">
          <div id="formContent">
            <h3>Enter a twitter URL and click <code>CAPTURE</code> button</h3>
            {error && <div className="alert alert-warning mt-1 alert-box mx-auto">
              <p className="mb-0">{error.message}</p>
            </div>
            }
            {validation.length > 0 && <div className="alert alert-warning mt-1 alert-box mx-auto alert-dismissible">
              <p className="mb-0">{validation}</p>
              <button type="button" className="close" data-dismiss="alert" aria-label="Close" onClick={() => setValidation('')}>
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            }
            <form onSubmit={handleSummit} className="mt-2">
              <input autoFocus autoComplete="off" type="text" id="url" name="url" placeholder="Enter Twitter URL" onChange={event => setUrl(event.target.value)}/>
              {loading ?
                  <button className="form-button" type="submit" disabled>
                    <span className="spinner-border spinner-border-sm mr-2" role="status" aria-hidden="true"></span>
                    Capturing tweet...
                  </button> :
                  <button className="form-button" type="submit" onClick={handleSummit}>CAPTURE</button>
              }
            </form>
          </div>
        </div>
      </>
  )
};


export default HomePage;
