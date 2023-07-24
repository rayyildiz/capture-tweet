import {FormEvent, useState} from "react";
import './HomePage.css';
import {gql, useMutation} from "@apollo/client";
import {Navigate} from "react-router-dom";
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


const HomePage = () => {
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
      } catch (ex:unknown) {
        if ( ex instanceof Error) {
          console.log(ex.toString());
        }
        
      }
    }
  }

  if (data && data.capture && data.capture.id.length > 0) {
    return <Navigate replace to={`/tweet/${data.capture.id}`}/>
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
            {validation.length > 0 && <div className="alert alert-warning mt-1 alert-box mx-auto alert-dismissible" role="alert">
              <p className="mb-0">{validation}</p>
              <button type="button" className="btn-close" data-dismiss="alert" aria-label="Close" onClick={() => setValidation('')}> </button>
            </div>
            }
            <form onSubmit={handleSummit} className="mt-2">
              <h5>Sunsetting: capturetweet is now read-only. It will be closed on <u>31.12.2023.</u> </h5>
              <p>The source code is open and you can find it <a href="https://github.com/rayyildiz/capture-tweet">here</a></p>
              <label htmlFor="url" hidden={true}>Enter tweet url</label>
              <br/>
              <input readOnly autoFocus autoComplete="off" type="text" id="url" name="url" placeholder="Enter Twitter URL (disabled)" onChange={event => setUrl(event.target.value)}/>
              {loading ?
                  <button className="form-button" type="submit" disabled>
                    <span className="spinner-border spinner-border-sm mr-2" role="status" aria-hidden="true"> </span>
                    Capturing tweet...
                  </button> :
                  <button disabled className="form-button" type="submit" onClick={handleSummit}>CAPTURE</button>
              }
            </form>
          </div>
        </div>
      </>
  )
};


export default HomePage;
