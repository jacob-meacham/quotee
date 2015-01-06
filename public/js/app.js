'use strict';

// Declare app level module which depends on filters, and services
angular.module('quotee', ['ngResource', 'ngRoute'])
  .config(['$routeProvider', '$locationProvider', function ($routeProvider, $locationProvider) {
    
    $routeProvider
    .when('/', {
        templateUrl: 'views/home/home.html',
        controller: 'HomeController'
    })
    .otherwise({redirectTo: '/'});
    
    $locationProvider.html5Mode(true);
}]);
