(function() {
    var app = angular.module('moddex', []);

    app.controller("SearchController", function() {
        this.query = "";
        this.search = function() {
            // TODO: hook to a sevrice and a rest call.
            console.log("query = "+this.query);
        }
    });
})();
