'use strict';

angular.module('quotee')
  .controller('HomeController', ['$scope', '$interval', '$routeParams', 'Quotes', function ($scope, $interval, $routeParams, Quotes) {
    if ($routeParams.style === 'inverse' || Math.random() > 0.5) {
      $scope.bodyClass = "inverse"
    }

    var intervalPromise;
    var showQuote = function(autoplay) {
      Quotes.get(function (quote) {
        $scope.quote = quote.body;
        $scope.author = "- " + quote.author;
      });

      if (autoplay) {
        intervalPromise = $interval(function() {
          showQuote(true); }, 5 * 1000, 1);
      }
    };

    if ($routeParams.autoplay === 'true') {
      showQuote(true);
    } else {
      showQuote(false);
    }

    $scope.$on('$destroy', function() {
      if (angular.isDefined(intervalPromise)) {
        $interval.cancel(intervalPromise);
        intervalPromise = undefined;
      }
    });
}]);
