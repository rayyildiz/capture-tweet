import React from 'react';
import ReactDOM from 'react-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import './index.css';
import App from './components/App';
import * as serviceWorker from './serviceWorker';
import {ENABLE_MONITORING, ENABLE_SW} from "./Constants";
import reportWebVitals from './reportWebVitals';


ReactDOM.render(
    <React.StrictMode>
      <App/>
    </React.StrictMode>,
    document.getElementById('root')
);

if (ENABLE_SW) {
  serviceWorker.register();
} else {
  serviceWorker.unregister();
}

if (ENABLE_MONITORING) {
  reportWebVitals();
}
