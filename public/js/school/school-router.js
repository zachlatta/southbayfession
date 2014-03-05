'use strict';

angular.module('southbayfession')
  .config(['$routeProvider', function ($routeProvider, $routeParams) {
    $routeProvider
      .when('/schools/:id', {
        templateUrl: 'views/school/school.html',
        controller: 'SchoolController'
      })
    }]);
