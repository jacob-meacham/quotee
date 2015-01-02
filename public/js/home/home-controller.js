'use strict';

angular.module('quotee')
  .controller('HomeController', ['$scope', '$log', 'Quotes', function ($scope, $log, Quotes) {
    Quotes.get(function (quote) {
      $scope.quote = quote.body;
      $scope.author = quote.author;
    });
}]);
