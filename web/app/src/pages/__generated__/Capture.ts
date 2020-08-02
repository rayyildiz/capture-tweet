/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: Capture
// ====================================================

export interface Capture_capture_author {
  __typename: "Author";
  userName: string;
  screenName: string | null;
  profileImageURL: string | null;
}

export interface Capture_capture {
  __typename: "Tweet";
  id: string;
  fullText: string;
  favoriteCount: number | null;
  retweetCount: number | null;
  postedAt: any | null;
  author: Capture_capture_author | null;
}

export interface Capture {
  capture: Capture_capture | null;
}

export interface CaptureVariables {
  url: string;
}
