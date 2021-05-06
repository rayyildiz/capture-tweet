import {ApolloClient, ApolloProvider, InMemoryCache} from "@apollo/client";
import {BrowserRouter} from "react-router-dom";
import {Header} from "./Header";
import {AppRoutes} from "../App.Routes";
import {BASE_API} from "../Constants";


const client = new ApolloClient({
  uri: BASE_API,
  cache: new InMemoryCache(),
});

const App = () => {
  return (
      <ApolloProvider client={client}>
        <BrowserRouter>
          <div className="flex-shrink-0">
            <Header/>
            <div className="container mt-lg-5 mt-sm-2">
              <AppRoutes/>
            </div>
          </div>
        </BrowserRouter>
      </ApolloProvider>
  );
}

export default App;
