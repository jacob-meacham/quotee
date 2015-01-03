'use strict';

angular.module('quotee').factory('Quotes', ['$resource', function($resource) {
  return $resource('/api/quote');
}]);