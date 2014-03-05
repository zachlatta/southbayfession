'use strict';

angular.module('southbayfession')
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/tweets', {
        templateUrl: 'views/tweet/tweets.html',
        controller: 'TweetController',
        resolve:{
          resolvedTweet: ['Tweet', function (Tweet) {
            return Tweet.query();
          }]
        }
      })
    }]);
