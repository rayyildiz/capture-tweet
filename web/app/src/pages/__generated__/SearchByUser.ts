/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: SearchByUser
// ====================================================

export interface SearchByUser_searchByUser_author {
  __typename: "Author";
  userName: string;
  screenName: string | null;
  profileImageURL: string | null;
}

export interface SearchByUser_searchByUser {
  __typename: "Tweet";
  id: string;
  fullText: string;
  lang: string | null;
  postedAt: any | null;
  captureThumbURL: string | null;
  author: SearchByUser_searchByUser_author | null;
}

export interface SearchByUser {
  searchByUser: SearchByUser_searchByUser[] | null;
}

export interface SearchByUserVariables {
  userID: string;
}
