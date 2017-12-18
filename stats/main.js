// Sort members on time needed to get first/second star on selected day
function sortOnDayStar(members, day, star) {
	return members.filter(x => {
			return x.completion_day_level.hasOwnProperty(day)
			&& x.completion_day_level[day].hasOwnProperty(star)
		})
		.sort((a, b) => {
			return new Date(a.completion_day_level[day][star].get_star_ts).getTime()
			- new Date(b.completion_day_level[day][star].get_star_ts).getTime()
		})
		.map(x => x.id);
}

// Get rankings for numBest members for each day & each star
function preprocessData(membersSorted, numRanks, numBest, maxDay) {
	const best = membersSorted.slice(0, numBest);
	let rankings = best.map(x => new Object({
		id: x.id,
		name: (x.name != null && x.name || "#" + x.id),
		star1: new Array(maxDay).fill(numRanks),
		star2: new Array(maxDay).fill(numRanks)
	}));

	for (let day = 1; day <= maxDay; day++) {
		for (let star = 1; star <= 2; star++) {
			let r = sortOnDayStar(membersSorted, day, star);
			for (let i = 0; i < Math.min(r.length, numRanks); i++) {
				let p = rankings.filter(x => x.id == r[i])[0];
				if (typeof p !== 'undefined') {
					if (star == 1) {
						p.star1[day - 1] = i + 1;
					} else {
						p.star2[day - 1] = i + 1;
					}
				}
			}
		}
	}
	return rankings;
}

function drawChart(membersSorted, numRanks, numBest, maxDay) {
	const rankings = preprocessData(membersSorted, numRanks, numBest, maxDay);

	let dataToPlot = new Array(rankings.length * 2).fill(null).map(
		(x, index) => new Object({
			type:"line",
			axisYType: "secondary",
			showInLegend: index % 2 == 0,
			markerSize: 0,
			dataPoints: []
		})
	);

	for (let i = 0; i < rankings.length; i++) {
		dataToPlot[2 * i].name = rankings[i].name;
		dataToPlot[2 * i + 1].name = rankings[i].name + " (2)";
		dataToPlot[2 * i].dataPoints =
			rankings[i].star1.map((el, index) => new Object({x: index + 1, y: el}));
		dataToPlot[2 * i + 1].dataPoints =
			rankings[i].star2.map((el, index) => new Object({x: index + 1, y: el}));
		let color = getRandomColor();
		dataToPlot[2 * i].color = color;
		dataToPlot[2 * i + 1].color = color;
		dataToPlot[2 * i].lineDashType = "dash";
	}

	let chart = new CanvasJS.Chart("chart-container", {
		title: {
			text: "Daily results",
			fontFamily: "Open Sans, sans-serif"
		},
		axisX: {
			title: "Day",
			fontFamily: "Open Sans, sans-serif",
			valueFormatString: "#",
			interval: 1
		},
		axisY2: {
			title: "Place",
			fontFamily: "Open Sans, sans-serif",
			reversed:  true,
			minimum: 0,
			maximum: numRanks,
			interval: 1
		},
		toolTip: {
			shared: true,
			fontFamily: "Open Sans, sans-serif"
		},
		legend: {
			cursor: "pointer",
			verticalAlign: "top",
			horizontalAlign: "center",
			dockInsidePlotArea: false,
			fontFamily: "Open Sans, sans-serif"
		},
		data: dataToPlot
	});
	chart.render();
}

function getRandomColor() {
    let s = "000000" + ((1 << 24) * Math.random() | 0).toString(16);
	return "#" + s.substr(s.length - 6);
}

// Example of list: [["id"], ["day_star", 7, 1]]
function getNeededFromObject(obj, list) {
	let l = []
	for (let i = 0; i < list.length; i++) {
		if (list[i][0] === "id") {
			l.push(obj.id);
		} else if (list[i][0] === "name") {
			l.push(obj.name != null && obj.name || "#" + obj.id);
		} else if (list[i][0] === "day_star") {
			let day = list[i][1]
			let star = list[i][2]
			let date = formatDate(new Date(obj.completion_day_level[day][star].get_star_ts));
			l.push(date);
		} else {
			l.push(null);
		}
	}
	return l;
}

function formatDate(d) {
	return ("0" + d.getHours()).slice(-2) + ":"
	+ ("0" + d.getMinutes()).slice(-2) + ":"
	+ ("0" + d.getSeconds()).slice(-2) + " ("
	+ ("0" + d.getDate()).slice(-2) + "-"
	+ ("0"+(d.getMonth()+1)).slice(-2) + "-"
	+ d.getFullYear() + ")";
}

function showDay(members, day) {
	let idResults1 = sortOnDayStar(members, day, 1);
	let idResults2 = sortOnDayStar(members, day, 2);

	let objResults1 = idResults1.map(x => members.find(el => el.id === x));
	let objResults2 = idResults2.map(x => members.find(el => el.id === x));

	let out1 = objResults1.map(x => getNeededFromObject(x,  [["name"], ["day_star", day, 1]]));
	let out2 = objResults2.map(x => getNeededFromObject(x,  [["name"], ["day_star", day, 2]]));

	return [out1, out2];
}

window.onload = function() {
	let app = new Vue({
		el: '#stats',

		data: {
			json: null,
			membs : null,
			numMembs: 0,
			numRankings: 5,
			numBest: 5,
			showDay: 1,
			maxDay : 0,
			star1 : [],
			star2: []
		},

		created: function() {
			this.$http.get('http://localhost:8080/aoc-data', {
				params: {
					leaderboard_id: 'leaderboard_id',
					session_cookie: 'session_cookie'
				}
			}).then(response => {
				this.json = response.body;
				this.membs = Object.keys(this.json.members)
					.map(k => this.json.members[k])
					.filter(x => x.local_score > 0)
					.sort((a, b) => b.local_score - a.local_score);
				this.maxDay = this.membs
					.map(x => Object.keys(x.completion_day_level)
						.reduce((a, b) => Math.max(a, b))
					)
					.reduce((a, b) => Math.max(a, b));
				this.numMembs = this.membs.length;

				this.draw();
				this.showSelectedDay();
			}, response => {
				console.error("Oooops. Something went wrong. Did you forget to run the server?");
			});
		},

		watch: {
			numRankings(val) {
				if (val > this.membs.length) {
					this.numRankings = this.membs.length;
				} else if (val < 1) {
					this.numRankings = 1;
				} else {
					this.draw();
				}
			},
			numBest(val) {
				if (val > this.membs.length) {
					this.numBest = this.membs.length;
				} else if (val < 1) {
					this.numBest = 1;
				} else {
					this.draw();
				}
			},
			showDay(val) {
				if (val > this.maxDay) {
					this.showDay = this.maxDay;
				} else if (val < 1) {
					this.showDay = 1;
				} else {
					this.showSelectedDay();
				}
			}
		},

		components: {
			"chart": {
				template: '<div id="chart-container"></div>'
			}
		},

		methods: {
			draw: function() {
				drawChart(this.membs, this.numRankings + 1, this.numBest, this.maxDay);
			},
			showSelectedDay: function() {
				let r = showDay(this.membs, this.showDay);
				this.star1 = r[0];
				this.star2 = r[1];
			}
		}
	});
};
