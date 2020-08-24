const readline = require('readline');

const fs = require('fs');

module.exports = function(app){

	app.get('/config', function(req, res){

		const configFile=process.env.CONFIG_FILE || "./example.json";
		const config = require(configFile);

		const data = fs.readFileSync(configFile, 'UTF-8');
		const lines = data.split(/\r?\n/);
		txt=""
		lines.forEach((line) => {
			console.log(line);
			txt+=`${line}<br>`
		});

		res.send(`<div style="text-align: center;"> <h1 >All from json config file: </h1><br> ${txt} <br><h3>- NodeJS</h3> </div>` );
	});

	//other routes..
}


