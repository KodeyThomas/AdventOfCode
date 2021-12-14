const fs = require('fs');

const lines = fs
	.readFileSync('input.txt', 'utf8')
	.split('\n')
	.map(line => {
		const [USP, output] = line.split(' | ').map(x => x.split(" "));

		return {
			USP,
			output
		};
	});

// Mappings
// 0 => 6
// 1 => 2 UNIQUE
// 2 => 5
// 3 => 5
// 4 => 4 UNIQUE
// 5 => 5
// 6 => 6
// 7 => 3 UNIQUE
// 8 => 7 UNIQUE
// 9 => 6

let counter = 0;
for (const line of lines) {
	const matches = line.output.filter(x => [2,3,4,7].includes(x.length));
	counter += matches.length;
}

console.log(counter);