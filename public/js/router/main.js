define("router", [
	"main",
	"blog",
	"logger",
	"backbone"], 
	function(Main, Logger, Backbone) {
		var logger = Logger.get("router")
		return Backbone.Router.extend({
			initialize: function(options) {
				Backbone.on('nav', _.bind(function() {
					this.navigate(arguments[0], {trigger: true});
				}, this));
			},
			routes: { 
				"blog": "blog",
				"blog/:year/:month/:day/:title": "fullPost",
				"blog/:year/:month": "monthlyPosts",
				"blog/:year": "yearlyPosts",
				"*defaultRoute": "main"				
			},
		  show: function(callback, always) {
		     if (always == true ) {
		        _.bind(callback, this)();
		    } else {
		      this.navigate('/account', {trigger: true});
		    }   
		  },			
			main: function() {
				logger.debug("main routed");
				this.mainView = this.mainView == undefined ? new Main.View() : this.mainView.render();
			},
			blog: function() {
				logger.debug("blog routed");
				this.blogView == undefined ? new Blog.View().render() : this.blogView.render();
			},
			fullPost: function(year, month, day, title) {
				logger.debug("full post routed");
				this.postView == undefined ? new Blog.PostView().render(year, month, day, title) : this.postView.render(year, month, day, title);
			},
			monthlyPosts: function(year, month) {
				logger.debug("monthly posts routed");
				this.monthlyView == undefined ? new Blog.MonthlyListView().render(year, month) : this.monthlyView.render(year, month);
			},
			yearlyPosts: function(year) {
				logger.debug("yearly posts routed");
				this.yearlyView == undefined ? new Blog.YearlyListView().render(year) : this.yearlyView.render(year);
		})
	}
)