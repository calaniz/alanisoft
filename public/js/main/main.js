define('main', 
	["text!main/main.html", "logger", "backbone", "mustache"], 
	function(Template, Logger, Backbone, Mustache) {
		return {View: Backbone.View.extend({
			el: "#app",
			template: Template,
			events: {
				"click #blog": "blog"
			},
			render: function() {
				this.$el.html(Mustache.to_html(this.template, {}))
				return this;
			},
			blog: function(e) {
				e.preventDefault();

				Backbone.trigger("nav", "/blog");
			}
		})}
	}
)