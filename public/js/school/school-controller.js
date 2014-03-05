'use strict';

angular.module('southbayfession')
  .controller('SchoolController', ['$scope', '$modal', '$routeParams', 'School',
    function ($scope, $modal, $routeParams, School) {
      $scope.school = School.get({id: $routeParams.id});
      $scope.twitterUrl = function (tweet) {
        return "https://twitter.com/southbayfession/status/" + tweet.twitterId;
      };
    }]);
