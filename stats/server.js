const http = require('http');
const https = require('https');
const Vue = require('vue');
const VueResource = require('vue-resource');
const url = require('url');

Vue.use(VueResource);

let sourceFile;
try {
	require.resolve( './.env' )
	sourceFile = require('./.env');
} catch( e ) {
	console.log('Cannot find file .env.')
	process.exit(1);
}

const port = sourceFile.port;
const host = sourceFile.host;

if ([port, host].some(x => x == null)) {
	console.log('Make sure that the following variables are declared in your .env file:\n- port\n- host');
	process.exit(1);
}

const server = http.createServer((req, res) => {
	const url_parts = url.parse(req.url, true);
	const query = url_parts.query;
	if (url_parts.pathname === '/aoc-data') {
		getData(query.leaderboard_id, query.session_cookie, data => sendData(res, data));
	}
});

server.listen(port, () => {
	console.log(`Server running at http://${host}:${port}/`);
});

function sendData(res, data) {
	res.writeHead(200, {'Content-Type': 'application/json', 'Access-Control-Allow-Origin': '*'});
	res.write(JSON.stringify(data));
	res.end();
}

function getData(leaderboard_id, session_cookie, callback) {
	const options = {
		hostname: 'adventofcode.com',
		path: `/2017/leaderboard/private/view/${leaderboard_id}.json`,
		headers: {'Cookie': `session=${session_cookie}`}
	};

	https.get(options, (res) => {
		const { statusCode } = res;
		const contentType = res.headers['content-type'];

		let error;
		if (statusCode !== 200) {
			error = new Error('Request Failed.\n' +
							`Status Code: ${statusCode}`);
		} else if (!/^application\/json/.test(contentType)) {
			error = new Error('Invalid content-type.\n' +
							`Expected application/json but received ${contentType}`);
		}
		if (error) {
			// consume response data to free up memory
			res.resume();
			console.log(error.message);
			return
		}

		res.setEncoding('utf8');

		let rawData = '';
		res.on('data', (chunk) => { rawData += chunk; });
		res.on('end', () => {
			try {
				callback(JSON.parse(rawData));
			} catch (e) {
				console.log(e.message);
			}
		});
	}).on('error', (e) => {
		console.log(`Got error: ${e.message}`);
	});
}
