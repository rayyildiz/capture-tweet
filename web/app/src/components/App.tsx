import React, {FC} from 'react';
import {ApolloClient, ApolloProvider, InMemoryCache} from "@apollo/client";
import {BrowserRouter} from "react-router-dom";
import {Header} from "./Header";
import {AppRoutes} from "../App.Routes";
import {Footer} from "./Footer";


const client = new ApolloClient({
  uri: "/api/query",
  cache: new InMemoryCache(),
});

const App: FC = () => {
  return (
      <ApolloProvider client={client}>
        <BrowserRouter>
          <Header/>
          <div className="container margin-top-7">
            <AppRoutes />
          </div>
          <Footer />
        </BrowserRouter>
      </ApolloProvider>
  );
}

export default App;
