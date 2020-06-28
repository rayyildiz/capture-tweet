import React, {FC, FormEvent, useState} from "react";
import './HomePage.css';
import {useMutation} from "@apollo/client";
import {Redirect} from "react-router-dom";
import {CAPTURE_TWEET_GQL} from "../graph/queries";
import {Capture, CaptureVariables} from "../graph/Capture";

const HomePage: FC = () => {
  const [url, setUrl] = useState('');
  const [validation, setValidation] = useState('');

  const [doQuery, {data, loading, error}] = useMutation<Capture, CaptureVariables>(CAPTURE_TWEET_GQL)

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
      <div className="wrapper">
        <div id="formContent">
          <h3>Enter a twitter URL and click <code>CAPTURE</code> button</h3>
          {error && <div className="alert alert-warning mt-1 alert-box mx-auto alert-dismissible">
            <p className="mb-0">{error.message}</p>
            <button type="button" className="close" data-dismiss="alert" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
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
            <input type="submit" value="CAPTURE"/>
          </form>
          {loading && <div className="spinner-border" role="status">
            <span className="sr-only">Loading...</span>
          </div>
          }
        </div>
      </div>
  )
};


export default HomePage;
