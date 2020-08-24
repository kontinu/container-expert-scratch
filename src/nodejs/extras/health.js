module.exports = function(app){

	app.get('/health', function(req, res){
		console.log("i'm healthy")
		res.send("healthy");
	});
}
