import React from 'react';
import {render} from '@testing-library/react';
import HomePage from './HomePage';
import {ApolloClient, ApolloProvider, InMemoryCache} from "@apollo/client";
import {BASE_API} from "../Constants";

test('renders capture button', () => {
  const client = new ApolloClient({
    uri: BASE_API,
    cache: new InMemoryCache(),
  });
  const {getAllByText} = render(
      <ApolloProvider client={client}>
        <HomePage/>
      </ApolloProvider>);
  const button = getAllByText("CAPTURE");
  expect(button).toBeDefined();
})
