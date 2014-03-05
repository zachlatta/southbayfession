'use strict';

angular.module('southbayfession')
  .factory('School', ['$resource', function ($resource) {
    return $resource('southbayfession/schools/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'}
    });
  }]);
