import React, {FC} from "react";
import {Link} from "react-router-dom";
import {Search} from "./Search";


export const Header: FC = () => (
    <nav className="navbar navbar-expand-lg fixed-top navbar-dark bg-primary">
      <Link to='/' className='navbar-brand'>Capture Tweet</Link>
      <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarColor01" aria-controls="navbarColor01" aria-expanded="false" aria-label="Toggle navigation">
        <span className="navbar-toggler-icon"></span>
      </button>

      <div className="collapse navbar-collapse" id="navbarColor01">
        <ul className="navbar-nav mr-auto">
          <li className="nav-item active">
            <Link to='/' className='nav-link'>Home</Link>
          </li>
          <li className="nav-item">
            <Link to='/privacy' className='nav-link'>Privacy</Link>
          </li>
        </ul>
          <Search />


      </div>
    </nav>
);
