'use strict';

// Declare app level module which depends on filters, and services
angular.module('quotee', ['ngResource', 'ngRoute', 'ui.bootstrap', 'ui.date'])
  .config(['$routeProvider', '$locationProvider', function ($routeProvider, $locationProvider) {
    $routeProvider
    .when('/', {
        templateUrl: 'views/home/home.html',
        controller: 'HomeController'
    })
    .otherwise({redirectTo: '/'});
    $locationProvider.html5Mode(true);
}]);
