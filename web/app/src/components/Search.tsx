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
      <form className="d-flex" onSubmit={handleSummit}>
        <input className="form-control mr-2" type="search" placeholder="Search" aria-label="Search" onChange={(event => setTerm(event.target.value))}/>
        <button className="btn btn-outline-success" type="submit">Search</button>
      </form>
  )
}
