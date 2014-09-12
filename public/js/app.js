require.config({
  //Remember: only use shim config for non-AMD scripts,
  //scripts that do not already call define(). The shim
  //config will not work correctly if used on AMD scripts,
  //in particular, the exports and init config will not
  //be triggered, and the deps config will be confusing
  //for those cases.
  urlArgs: "bust=" + (new Date()).getTime(),
  shim: {
      "backbone-1.1.2": {
        //These script dependencies should be loaded before loading
        //backbone.js
        deps: ["underscore-1.5.1", "jquery-1.11.1"],
        //Once loaded, use the global 'Backbone' as the
        //module value.
        exports: "Backbone"
      },
      "underscore-1.5.1": {
        exports: "_"
      }
  },
  map: {
    "*": {
      "backbone": "backbone-1.1.2",
      "underscore": "underscore-1.5.1",
      "logger": "logger-0.9.2",
      "mustache": "mustache-1.0.0",
      "text": "text-2.0.10",
    }
  },
  packages: [
    "main",
    "router"
  ]
})
    
require([
  "router",
  "logger",  
  ], function(Router, Logger) {
    Logger.useDefaults(Logger.DEBUG);
    Logger.info("[alanisoft] loaded");

    new Router();
    Backbone.history.start({pushState: true});


  }
)