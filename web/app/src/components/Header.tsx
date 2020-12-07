import {FC} from "react";
import {Link} from "react-router-dom";
import {Search} from "./Search";

export const Header: FC = () => (
    <div className="d-flex flex-column flex-md-row align-items-center bg-dark p-3 px-md-4 mb-3 border-bottom shadow">
      <h5 className="my-0 mr-md-auto font-weight-normal"><Link to='/' className='p-2 text-light navbar-brand'>Capture Tweet</Link></h5>
      <nav className="my-2 my-md-0 mr-md-3">
        <Link to='/' className="p-2 text-light text-decoration-none mr-2">Capture</Link>
        <Link to='/privacy' className="p-2 text-light text-decoration-none mr-2">Privacy</Link>
        <Link to='/contact' className="p-2 text-light text-decoration-none mr-2">Contact Us</Link>
      </nav>
      <Search/>
    </div>
)
