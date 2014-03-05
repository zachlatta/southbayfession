'use strict';

angular.module('southbayfession')
  .factory('Tweet', ['$resource', function ($resource) {
    return $resource('southbayfession/tweets/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);
