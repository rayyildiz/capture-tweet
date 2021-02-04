import {FC, FormEvent} from "react";
import {useHistory} from "react-router-dom";
import moment from "moment";
import notFound from "../assets/not_found.svg";


type TweetCardProps = {
  id: string;
  fullText: string;
  lang?: string;
  postedAt?: any;
  captureThumbURL?: string;
  author?: string;
}

export const TweetCard: FC<TweetCardProps> = (props) => {
  const history = useHistory();

  if (!props) {
    return <span>Error</span>
  }

  const handleClick = (e: FormEvent) => {
    e.preventDefault();
    history.push("/tweet/" + props.id);
  }

  return (
      <div className="col-sm-4">
        <div className="card mb-3 cursor tweet-card" onClick={handleClick}>
          <h3 className="card-header">Tweet by {props.author}</h3>
          <div className="card-body">
            <h6 className="card-title">Posted at {moment(props.postedAt).format("DD-MM-YYYY HH:MM")} <span className="badge rounded-pill bg-success">{props.lang}</span></h6>
          </div>
          <div style={{
            width: '100%',
            textAlign: 'center',
            overflow: 'hidden',
            height: '10rem'
          }}>
            {props.captureThumbURL ? <img style={{maxWidth: '20rem'}} src={`/${props.captureThumbURL}`} alt=""/>
                : <img style={{maxWidth: '20rem'}} src={notFound} alt=""/>}
          </div>
          <div className="card-body">
            <p className="card-text">{props.fullText} </p>
          </div>
        </div>
      </div>
  )
};
