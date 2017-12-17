const http = require('http');
const https = require('https');

const port = 8080;
const host = 'localhost';

const session_cookie = 'session_cookie'; // UPDATE!
const leaderboard_id = 'lb_id' // UPDATE!

const server = http.createServer((req, res) => {
	if (req.url === '/aoc-data') {
		getData((data) => {
			res.writeHead(200, {'Content-Type': 'application/json'});
			res.write(JSON.stringify(data));
			res.end();
		});
	}
});

server.listen(port, () => {
	console.log(`Server running at http://${host}:${port}/`);
});

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
