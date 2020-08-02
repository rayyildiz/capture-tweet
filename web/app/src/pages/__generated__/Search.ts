/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { SearchInput } from "./../../../__generated__/globalTypes";

// ====================================================
// GraphQL query operation: Search
// ====================================================

export interface Search_search_author {
  __typename: "Author";
  userName: string;
  screenName: string | null;
  profileImageURL: string | null;
}

export interface Search_search {
  __typename: "Tweet";
  id: string;
  fullText: string;
  lang: string | null;
  postedAt: any | null;
  captureThumbURL: string | null;
  author: Search_search_author | null;
}

export interface Search {
  search: (Search_search | null)[] | null;
}

export interface SearchVariables {
  input: SearchInput;
}
