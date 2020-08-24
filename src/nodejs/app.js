'use strict';

const express = require('express');
const request = require('request')
const fetch = require('node-fetch')
// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

// App
const app = express();

// Use environment variables
//require('./extras/load_env')(app);

// read config from json
//require('./extras/config_file')(app);

require("./extras/health")(app)
//? MID

var MSG=process.env.MSG || " 📦 Container ";

const hostname=process.env.HOSTNAME || "";



app.get('/api',  async function(req, res, next) {
	let NEXT_URL=process.env.NEXT_URL
	if (NEXT_URL){
		const resp= await fetch(NEXT_URL ,{
			headers: {'Content-Type': 'application/json'}
		}
		)
		const yeison = await resp.json();
		res.json({"msg": `${MSG+yeison.msg}`, "hostname":`${hostname}`} )
	}
	res.end
})



app.get('/',(req,res) => {
	console.log("/root")
  res.send(`<div style="text-align: center;"> <h1>${MSG}</h1><br><h3>NodeJS</h3> </div>`);
})



app.listen(PORT, HOST);
console.log(`Running NodeJS Server on http://${HOST}:${PORT}`);
