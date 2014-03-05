// Declare app level module which depends on filters, and services
angular.module('southbayfession', [
    'ngResource',
    'ngRoute',
    'ui.bootstrap',
    'ui.date',
    'angulartics',
    'angulartics.google.analytics'
  ])
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/home/home.html', 
        controller: 'HomeController'})
      .otherwise({redirectTo: '/'});
  }]);
