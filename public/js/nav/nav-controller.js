angular.module('southbayfession')
  .controller('NavController', ['$scope', 'School', function ($scope, School) {
    $scope.schools = School.query();
    $scope.twitter = 'https://twitter.com/southbayfession';
  }]);
