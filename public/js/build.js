({
  baseUrl: ".",
  packages: ['main', 'router'],
  shim: {
      'backbone-1.1.2': {
          //These script dependencies should be loaded before loading
          //backbone.js
          deps: ['underscore-1.5.1', 'jquery-1.11.1'],
          //Once loaded, use the global 'Backbone' as the
          //module value.
          exports: 'Backbone'
      },
      'underscore-1.5.1': {
          exports: '_'
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
  name: "app.js",
  out: "app.min.js"
})