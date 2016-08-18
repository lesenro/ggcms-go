/* Setup general page controller */
angular.module('GgcmsApp').controller('GeneralPageController', ['$rootScope', '$scope', 'settings', '$stateParams', function($rootScope, $scope, settings, $stateParams) {
    $scope.$on('$viewContentLoaded', function() {
        // initialize core components
        App.initAjax();
        // set default layout mode
        $rootScope.settings.layout.pageContentWhite = true;
        $rootScope.settings.layout.pageBodySolid = false;
        $rootScope.settings.layout.pageSidebarClosed = false;
    });

}]);