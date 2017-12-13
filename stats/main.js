// Sort members on time needed to get first/second star on selected day
function sortOnDayStar(members, day, star) {
	return members.filter(x => x.completion_day_level.hasOwnProperty(day) && x.completion_day_level[day].hasOwnProperty(star))
		.sort((a, b) => new Date(a.completion_day_level[day][star].get_star_ts) - new Date(b.completion_day_level[day][star].get_star_ts))
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
			r = sortOnDayStar(membersSorted, day, star);
			for (let i = 0; i < Math.min(r.length, numRanks); i++) {
				p = rankings.filter(x => x.id == r[i])[0];
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

	let dataToPlot = new Array(rankings.length * 2).fill().map(
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
		dataToPlot[2 * i].dataPoints = rankings[i].star1.map((el, index) => new Object({x: index + 1, y: el}));
		dataToPlot[2 * i + 1].dataPoints = rankings[i].star2.map((el, index) => new Object({x: index + 1, y: el}));
		let color = getRandomColor();
		dataToPlot[2 * i].color = color;
		dataToPlot[2 * i + 1].color = color;
		dataToPlot[2 * i].lineDashType = "dash";
	}

	let chart = new CanvasJS.Chart("chart-container", {
		title: {
			text: "Daily results"
		},
		axisX: {
			title: "Day",
			valueFormatString: "#",
			interval: 1
		},
		axisY2: {
			title: "Place",
			reversed:  true,
			minimum: 0,
			maximum: numRanks,
			interval: 1
		},
		toolTip: {
			shared: true
		},
		legend: {
			cursor: "pointer",
			verticalAlign: "top",
			horizontalAlign: "center",
			dockInsidePlotArea: false
		},
		data: dataToPlot
	});
	chart.render();
}

function getRandomColor() {
    var s = "000000" + ((1 << 24) * Math.random() | 0).toString(16);
	return "#" + s.substr(s.length - 6);
}

// Example of list: [["id"], ["day_star", 7, 1]]
function getNeededFromObject(obj, list) {
	l = []
	for (i = 0; i < list.length; i++) {
		if (list[i][0] === "id") {
			l.push(obj.id);
		} else if (list[i][0] === "name") {
			l.push(obj.name != null && obj.name || "#" + obj.id);
		} else if (list[i][0] === "day_star") {
			day = list[i][1]
			star = list[i][2]
			date = formatDate(new Date(obj.completion_day_level[day][star].get_star_ts));
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
	idResults1 = sortOnDayStar(members, day, 1);
	idResults2 = sortOnDayStar(members, day, 2);

	// Update!
	objResults1 = idResults1.map(x => members.find(function(el) { return el.id === x; }));
	objResults2 = idResults2.map(x => members.find(function(el) { return el.id === x; }));

	out1 = objResults1.map(x => getNeededFromObject(x,  [["name"], ["day_star", day, 1]]));
	out2 = objResults2.map(x => getNeededFromObject(x,  [["name"], ["day_star", day, 2]]));
	return [out1, out2];
}

window.onload = function() {
	let app = new Vue({
		el: '#stats',

		data: {
			json: JSON.parse(results),
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
			this.membs = Object.keys(this.json.members)
				.map(k => this.json.members[k])
				.filter(x => x.local_score > 0)
				.sort((a, b) => b.local_score - a.local_score);
			this.maxDay = Math.max(...this.membs.map(x => Math.max(...Object.keys(x.completion_day_level))));
			this.numMembs = this.membs.length;
		},

		mounted: function() {
			this.draw();
			this.showSelectedDay();
		},

		watch: {
			numRankings: function(val) {
				if (val > this.membs.length) {
					this.numRankings = this.membs.length;
				} else if (val < 1) {
					this.numRankings = 1;
				} else {
					this.draw();
				}
			},
			numBest: function(val) {
				if (val > this.membs.length) {
					this.numBest = this.membs.length;
				} else if (val < 1) {
					this.numBest = 1;
				} else {
					this.draw();
				}
			},
			showDay: function(val) {
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
				r = showDay(this.membs, this.showDay);
				this.star1 = r[0];
				this.star2 = r[1];
			}
		}
	});
};
