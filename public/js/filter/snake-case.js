angular.module('southbayfession')
  .filter('snakeCase', function () {
    return function (input) {
      return input.toLowerCase().split(' ').join('_');
    };
  });
