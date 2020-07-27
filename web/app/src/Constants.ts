const prod = {
  enableSW: true,
  apiURL: 'https://api.capturetweet.com/api/query',
  captchaKey: '6LeH-rYZAAAAAE23jskkX4U2_oJoXvcreHg2n2ic'
};

const dev = {
  enableSW: false,
  apiURL: '/api/query',
  captchaKey: '6LeH-rYZAAAAAE23jskkX4U2_oJoXvcreHg2n2ic'
};

const config = process.env.NODE_ENV === 'development' ? dev : prod;

export const BASE_API = config.apiURL;
export const ENABLE_SW = config.enableSW;
export const CAPTCHA_KEY = config.captchaKey;
export const WEB_BASE_URL = 'https://capturetweet.com';
