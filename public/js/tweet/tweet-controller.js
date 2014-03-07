'use strict';

angular.module('southbayfession')
  .controller('LatestTweetsController', ['$scope', '$modal', 'Tweet',
    function ($scope, $modal, Tweet) {
      $scope.tweets = Tweet.query();
      $scope.twitterUrl = function (tweet) {
        return "https://twitter.com/southbayfession/status/" + tweet.twitterId;
      };
    }]);
