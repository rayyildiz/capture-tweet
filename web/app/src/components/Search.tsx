import React, {FC, FormEvent, useState} from "react";
import {useHistory} from "react-router-dom";
import './Search.css';

export const Search: FC = () => {
  const [term, setTerm] = useState<string>('');
  const history = useHistory();
  const handleSummit = (event: FormEvent) => {
    event.preventDefault();
    history.push("/search?q=" + term);
  }

  return (
      <form onSubmit={handleSummit}>
        <div className="searchbar">
          <input className="search_input" type="text" name="" placeholder="Search..." onChange={(event => setTerm(event.target.value))}/>
          <a href="javascript:void(0)" className="search_icon"><i className="fas fa-search"></i></a>
        </div>
      </form>
  )
}
