'use strict';

angular.module('southbayfession')
  .controller('TweetController', ['$scope', '$modal', 'resolvedTweet', 'Tweet',
    function ($scope, $modal, resolvedTweet, Tweet) {

      $scope.tweets = resolvedTweet;

      $scope.create = function () {
        $scope.clear();
        $scope.open();
      };

      $scope.update = function (id) {
        $scope.tweet = Tweet.get({id: id});
        $scope.open(id);
      };

      $scope.delete = function (id) {
        Tweet.delete({id: id},
          function () {
            $scope.tweets = Tweet.query();
          });
      };

      $scope.save = function (id) {
        if (id) {
          Tweet.update({id: id}, $scope.tweet,
            function () {
              $scope.tweets = Tweet.query();
              $scope.clear();
            });
        } else {
          Tweet.save($scope.tweet,
            function () {
              $scope.tweets = Tweet.query();
              $scope.clear();
            });
        }
      };

      $scope.clear = function () {
        $scope.tweet = {
          
          "createdAt": "",
          
          "twitterId": "",
          
          "text": "",
          
          "school": "",
          
          "id": ""
        };
      };

      $scope.open = function (id) {
        var tweetSave = $modal.open({
          templateUrl: 'tweet-save.html',
          controller: TweetSaveController,
          resolve: {
            tweet: function () {
              return $scope.tweet;
            }
          }
        });

        tweetSave.result.then(function (entity) {
          $scope.tweet = entity;
          $scope.save(id);
        });
      };
    }]);

var TweetSaveController =
  function ($scope, $modalInstance, tweet) {
    $scope.tweet = tweet;

    

    $scope.ok = function () {
      $modalInstance.close($scope.tweet);
    };

    $scope.cancel = function () {
      $modalInstance.dismiss('cancel');
    };
  };
