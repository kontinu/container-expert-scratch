module.exports = function(app){

	app.get('/env', function(req, res){
		const the_variable=process.env.THE_VAR || "default value";
		res.send(`<div style="text-align: center;"> <h1 >The Variable is:<br> ${the_variable} </h1><br><h3>- NodeJS</h3> </div>` );
	});

}
