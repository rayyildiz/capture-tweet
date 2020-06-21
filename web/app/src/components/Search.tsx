import React, {FC, FormEvent, useState} from "react";
import {useHistory} from "react-router-dom";
import './Search.css';

export const Search: FC = () => {
  const [term, setTerm] = useState<string>('');
  const history = useHistory();
  const handleSummit = (event: FormEvent) => {
    event.preventDefault();
    if (term.length >= 2) {
      history.push("/search?q=" + term);
    }
  }
  return (
      <form className="form-inline my-2 my-lg-0" onSubmit={handleSummit}>
        <div className="p-2 rounded rounded-pill shadow-sm mb-6 bg-secondary">
          <div className="input-group bg-secondary">
            <input type="search" placeholder="Search for Tweets" aria-describedby="button-addon1"
                   className="form-control border-0 bg-secondary text-white"
                   onChange={(event => setTerm(event.target.value))}/>
            <div className="input-group-append">
              <button id="button-addon1" type="submit" className="btn btn-link text-primary"><i className="fa fa-search"></i></button>
            </div>
          </div>
        </div>
      </form>
  )
}
