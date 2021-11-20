import {FC, lazy, Suspense} from "react";
import {Route, Routes} from "react-router-dom";
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
        <Routes>
          <Route path='/privacy' element={<PrivacyPage />}/>
          <Route path='/search' element={<SearchPage />}/>
          <Route path='/contact' element={<ContactPage />}/>
          <Route path='/tweet/:id' element={<TweetPage />}/>
          <Route path='/user/:id' element={<UserPage />}/>
          <Route path="/" element={<HomePage />}/>
        </Routes>
      </Suspense>
  )
};

