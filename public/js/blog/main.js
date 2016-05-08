define('blog', 
	["text!blog/main.html", "logger", "backbone", "mustache"], 
	function(Template, Logger, Backbone, Mustache) {
		return {View: Backbone.View.extend({
			el: "#app",
			template: Template,
			events: {

			},
			render: function() {
				this.$el.html(Mustache.to_html(this.template, {}))
				return this;
			},
		})}
	}
)