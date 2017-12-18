const http = require('http');
const https = require('https');
const Vue = require('vue');
const VueResource = require('vue-resource');

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
const session_cookie = sourceFile.session_cookie
const leaderboard_id = sourceFile.leaderboard_id

if ([port, host, session_cookie, leaderboard_id].some(x => x == null)) {
	console.log('Make sure that the following variables are declared in your .env file:\n- port\n- host\n- session_cookie\n- leaderboard_id');
	process.exit(1);
}

const server = http.createServer((req, res) => {
	if (req.url === '/aoc-data') {
		getData(data => sendData(res, data));
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

function getData(callback) {
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
				const parsedData = JSON.parse(rawData);
				callback(parsedData);
			} catch (e) {
				console.log(e.message);
			}
		});
	}).on('error', (e) => {
		console.log(`Got error: ${e.message}`);
	});
}
