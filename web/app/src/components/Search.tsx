import {FormEvent, useState} from "react";
import {useNavigate} from "react-router-dom";

export const Search = () => {
    const [term, setTerm] = useState<string>('');
    const navigate = useNavigate();
    const handleSummit = (event: FormEvent) => {
        event.preventDefault();
        if (term.length >= 2) {
            navigate(`/search?q=${term}`);
        }
    }
    return (
        <form className="d-flex" onSubmit={handleSummit}>
            <div className="input-group">
                <input className="form-control" type="search" placeholder="Search for tweet" aria-label="Search"
                       onChange={(event => setTerm(event.target.value))}/>
                <div className="input-group-append">
                    <button className="btn btn-secondary" name="btnSearch" type="submit"
                            aria-label="search capture tweet">
                        <svg width="1em" height="1em" viewBox="0 0 16 16" className="bi bi-search" fill="currentColor"
                             xmlns="http://www.w3.org/2000/svg">
                            <path fillRule="evenodd"
                                  d="M10.442 10.442a1 1 0 0 1 1.415 0l3.85 3.85a1 1 0 0 1-1.414 1.415l-3.85-3.85a1 1 0 0 1 0-1.415z"/>
                            <path fillRule="evenodd"
                                  d="M6.5 12a5.5 5.5 0 1 0 0-11 5.5 5.5 0 0 0 0 11zM13 6.5a6.5 6.5 0 1 1-13 0 6.5 6.5 0 0 1 13 0z"/>
                        </svg>
                    </button>
                </div>
            </div>
        </form>
    )
}
