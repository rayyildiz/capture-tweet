import React, {FC, lazy, Suspense} from "react";
import {Route, Switch} from "react-router-dom";
import HomePage from "./pages/HomePage";

const PrivacyPage = lazy(() => import("./pages/PrivacyPage"));
const SearchPage = lazy(() => import("./pages/SearchPage"));
const TweetPage = lazy(() => import("./pages/TweetPage"));
const ContactPage = lazy(() => import('./pages/ContactPage'));

const Loading = () => (<span>Loading...</span>)

export const AppRoutes: FC = () => {
  return (
      <Suspense fallback={<Loading/>}>
        <Switch>
          <Route path='/privacy' component={PrivacyPage}/>
          <Route path='/search' component={SearchPage}/>
          <Route path='/contact' component={ContactPage}/>
          <Route path='/tweet/:id' component={TweetPage}/>
          <Route exact={true} path="/" component={HomePage}/>
        </Switch>
      </Suspense>
  )
};

