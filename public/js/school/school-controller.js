'use strict';

angular.module('southbayfession')
  .controller('SchoolController', ['$scope', '$modal', '$routeParams', 'School',
    function ($scope, $modal, $routeParams, School) {
      $scope.school = School.get({id: $routeParams.id});
    }]);
