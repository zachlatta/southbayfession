# Ye Be Warned ![Analytics](https://ga-beacon.appspot.com/UA-34529482-6/southbayfession/readme?pixel)

Abandon all hope, ye who enter here. The code is bad. The site is bad. But it
_works_. I say that tentatively.

## Getting Started

To run the application you'll need to set a few environment variables.

* `TWITTER_CONSUMER_KEY` - Consumer key for Twitter API
* `TWITTER_CONSUMER_SECRET` - Consumer secret for Twitter API
* `TWITTER_ACCESS_TOKEN` - Access token for Twitter API
* `TWITTER_ACCESS_TOKEN_SECRET` - Access token secret for Twitter API
* `DATABASE_URL` - Destination of the database.
  * In the debug environment this should be a filesystem path for SQLite. Ex.
    `/tmp/my.db`.
  * In the production environment this should be a Postgres URL. Ex.
    `postgres://username:password@host:port/db`
* `ENV` - Set this to `PRODUCTION` when in a production environment.
* `PRERENDER_URL` - URL for http://prerender.io service. Ex.
  `https://service.prerender.io/http.yourwebsite.com`. Must not have a trailing
  slash.
