'use strict';

angular.module('quotee')
  .controller('HomeController', ['$scope', '$routeParams', 'Quotes', function ($scope, $routeParams, Quotes) {
    if ($routeParams.style === 'inverse' || Math.random() > 0.5) {
      $scope.bodyClass = "inverse"
    }

    Quotes.get(function (quote) {
      $scope.quote = quote.body;
      $scope.author = "- " + quote.author;
    });
}]);
