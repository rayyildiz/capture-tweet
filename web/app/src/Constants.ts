const prod = {
  enableSW: false,
  apiURL: 'https://beta-api.capturetweet.com/api/query'
};

const dev = {
  enableSW: false,
  apiURL: '/api/query'
};

const config = process.env.NODE_ENV === 'development' ? dev : prod;

export const BASE_API = config.apiURL;
export const ENABLE_SW = config.enableSW;
export const WEB_BASE_URL= 'https://beta.capturetweet.com';
