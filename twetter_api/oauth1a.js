// https://github.com/googleworkspace/apps-script-oauth1

const SCRIPT_PROPERTIES = PropertiesService.getScriptProperties();
const TWITTER_CONSUMER_KEY = SCRIPT_PROPERTIES.getProperty('TWITTER_CONSUMER_KEY');
const TWITTER_CONSUMER_SECRET = SCRIPT_PROPERTIES.getProperty('TWITTER_CONSUMER_SECRET');

function getTwitterService() {
  return OAuth1.createService('Twitter')
    .setAccessTokenUrl('https://api.twitter.com/oauth/access_token')
    .setRequestTokenUrl('https://api.twitter.com/oauth/request_token')
    .setAuthorizationUrl('https://api.twitter.com/oauth/authorize')
    .setConsumerKey(TWITTER_CONSUMER_KEY)
    .setConsumerSecret(TWITTER_CONSUMER_SECRET)
    .setCallbackFunction('authCallback')
    .setPropertyStore(PropertiesService.getUserProperties());
}

function authCallback(request) {
  const service = getTwitterService();
    const isAuthorized = service.handleCallback(request);
    if (isAuthorized) {
        return HtmlService.createHtmlOutput('Success! You can close this tab.');
    } else {
        return HtmlService.createHtmlOutput('Denied. You can close this tab');
    }
}

function authorize() {
  const service = getTwitterService();
  const authorizationUrl = service.authorize();
  Logger.log('Open the following URL and re-run the script: %s', authorizationUrl);
}

function reset() {
  const service = getTwitterService();
  service.reset();
}

function postTweet() {
  const service = getTwitterService();

  if (service.hasAccess()) {
    const url = 'https://api.twitter.com/2/tweets';
    const payload = JSON.stringify({
      text: 'Hello, world!'
    });
    const options = {
      method: 'post',
      contentType: 'application/json',
      payload: payload,
      muteHttpExceptions: true
    };
    const response = service.fetch(url, options);
    Logger.log(response.getContentText());
  } else {
    Logger.log('No access yet. Run the authorize function first.');
  }
}