import {FC, lazy, Suspense} from "react";
import {Route, Switch} from "react-router-dom";
import HomePage from "./pages/HomePage";
import Loading from "./components/Loading";

const PrivacyPage = lazy(() => import("./pages/PrivacyPage"));
const SearchPage = lazy(() => import("./pages/SearchPage"));
const TweetPage = lazy(() => import("./pages/TweetPage"));
const ContactPage = lazy(() => import('./pages/ContactPage'));
const UserPage = lazy(() => import('./pages/UserPage'));

export const AppRoutes: FC = () => {
  return (
      <Suspense fallback={<Loading/>}>
        <Switch>
          <Route path='/privacy' component={PrivacyPage}/>
          <Route path='/search' component={SearchPage}/>
          <Route path='/contact' component={ContactPage}/>
          <Route path='/tweet/:id' component={TweetPage}/>
          <Route path='/user/:id' component={UserPage}/>
          <Route exact={true} path="/" component={HomePage}/>
        </Switch>
      </Suspense>
  )
};

