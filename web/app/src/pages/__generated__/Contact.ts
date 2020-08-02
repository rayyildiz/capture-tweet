/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { ContactInput } from "./../../../__generated__/globalTypes";

// ====================================================
// GraphQL mutation operation: Contact
// ====================================================

export interface Contact {
  contact: string;
}

export interface ContactVariables {
  input: ContactInput;
  id?: string | null;
  captcha: string;
}
