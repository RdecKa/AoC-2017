// Sort members on time needed to get first/second star on selected day
function sortOnDayStar(members, day, star) {
	return members.filter(x => x.completion_day_level.hasOwnProperty(day) && x.completion_day_level[day].hasOwnProperty(star))
		.sort((a, b) => new Date(a.completion_day_level[day][star].get_star_ts) - new Date(b.completion_day_level[day][star].get_star_ts))
		.map(x => x.id);
}

// Get rankings for numBest members for each day & each star
function preprocessData(membersSorted, numRanks, numBest) {
	const maxDay = Math.max(...membersSorted.map(x => Math.max(...Object.keys(x.completion_day_level))));
	const best = membersSorted.slice(0, numBest);
	let rankings = best.map(x => new Object({id: x.id, name: (x.name != null && x.name || "#" + x.id), star1: new Array(maxDay).fill(numRanks), star2: new Array(maxDay).fill(numRanks)}));
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

function drawChart(membersSorted, numRanks, numBest) {
	const rankings = preprocessData(membersSorted, numRanks, numBest);

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
		let color = "#" + ((1 << 24) * Math.random() | 0).toString(16);
		dataToPlot[2 * i].color = color;
		dataToPlot[2 * i + 1].color = color;
		dataToPlot[2 * i].lineDashType = "dash";
	}

	let chart = new CanvasJS.Chart("chartContainer", {
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

window.onload = function() {
	let app = new Vue({
		el: '#stats',

		data: {
			json: JSON.parse(results),
			membs : null,
			numRankings: 5,
			numBest: 5
		},

		mounted: function() {
			this.membs = Object.keys(this.json.members)
				.map(k => this.json.members[k])
				.filter(x => x.local_score > 0)
				.sort((a, b) => b.local_score - a.local_score);
			this.draw();
		},

		watch: {
			numRankings: function(val) {
				if (val > 50) {
					this.numRankings = 50;
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
			}
		},

		components: {
			"chart": {
				template: '<div id="chartContainer" style="min-height: 600px; width: 100%;"></div>'
			}
		},

		methods: {
			draw: function() {
				drawChart(this.membs, this.numRankings, this.numBest);
			}
		}
	});
};
