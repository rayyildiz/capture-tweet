import React, {FC, useState} from "react";
import {Link} from "react-router-dom";
import {Search} from "./Search";

/*
<button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarColor01" aria-controls="navbarColor01" aria-expanded="false" aria-label="Toggle navigation">
          <span className="navbar-toggler-icon"></span>
        </button>
 */
export const Header: FC = () => {
  const [link, setLink] = useState('home');

  return (
      <nav className="navbar navbar-expand-lg fixed-top navbar-dark bg-primary">
        <Link to='/' className='navbar-brand'>Capture Tweet</Link>

        <div className="collapse navbar-collapse" id="navbarColor01">
          <ul className="navbar-nav mr-auto">
            <li className={`nav-item ${link === "home" ? "active" : ""}`}>
              <Link to='/' className="nav-link" onClick={() => setLink("home")}>Home</Link>
            </li>
            <li className={`nav-item ${link === "privacy" ? "active" : ""}`}>
              <Link to='/privacy' className="nav-link" onClick={() => setLink("privacy")}>Privacy</Link>
            </li>
          </ul>
          <Search/>
        </div>
      </nav>
  );
}
