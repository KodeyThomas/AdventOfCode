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
// 0 => a,b,c,e,f,g
// 1 => c,f
// 2 => a,c,d,e,g
// 3 => a,c,d,f,g
// 4 => b,c,d,f
// 5 => a,b,d,f,g
// 6 => a,b,d,e,f,g
// 7 => a,c,f
// 8 => a,b,c,d,e,f,g
// 9 => a,b,c,d,f,g

const getSegmentFreqMap = () => { // i mean mathematically there is a better way to do this, but i don't know how to do it.
	let segMap = {
		a: 8, // a is used in 8 digits
		b: 6, // b is used in 6 digits
		c: 8, // you get the gist
		d: 7,
		e: 4,
		f: 9,
		g: 7
	};
	
	let segmentFreqMap = { // These numbers will all be unique
		0: segMap['a'] + segMap['b'] + segMap['c'] + segMap['e'] + segMap['f'] + segMap['g'],
		1: segMap['c'] + segMap['f'],
		2: segMap['a'] + segMap['c'] + segMap['d'] + segMap['e'] + segMap['g'],
		3: segMap['a'] + segMap['c'] + segMap['d'] + segMap['f'] + segMap['g'],
		4: segMap['b'] + segMap['c'] + segMap['d'] + segMap['f'],
		5: segMap['a'] + segMap['b'] + segMap['d'] + segMap['f'] + segMap['g'],
		6: segMap['a'] + segMap['b'] + segMap['d'] + segMap['e'] + segMap['f'] + segMap['g'],
		7: segMap['a'] + segMap['c'] + segMap['f'],
		8: segMap['a'] + segMap['b'] + segMap['c'] + segMap['d'] + segMap['e'] + segMap['f'] + segMap['g'],
		9: segMap['a'] + segMap['b'] + segMap['c'] + segMap['d'] + segMap['f'] + segMap['g']
	}

	// Why on earth does this work?
	//
	// Luckily, the shape of the segment's create 10 Unique numbers
	// As we know the digit shapes are unique, we can compute the answer
	// No matter the input, if we assign an 'c' to 8 and 'f' to 9 
	// If the total sum of the input wires for that segment is the same as the total sum of the segments wires
	// We can conclude that c and f output the digit 1
	//
	// Input = (c,f) = (8,9) = (8 + 9) = 17
	// Output = (f,c) = (9,8) = (9 + 8) = 17
	// Input == Output
	//
	// Digit 1 has a score of 17, so 1 is the decoded answer
	//
	// So any combination of 'a' and 'b' will output the digit 1 for that line
	//
	// If this doesn't make sense drop me a message on twitter @ThomasKodey and i'll try and explain it.

	return segmentFreqMap;
}

let totalOutputsSum = 0;

const segmentFreqMap = getSegmentFreqMap();
console.log('Segment Frequency Map:', segmentFreqMap);
console.log('\n');

lines.forEach((line, index) => {
	let freqMap = {
		a: 0,
		b: 0,
		c: 0,
		d: 0,
		e: 0,
		f: 0,
		g: 0
	};

	line.USP.forEach(wiringCombo => {
		wiringCombo.split('').forEach(char => { // Calculate frequency of each character
			freqMap[char]++;
		});
	});

	console.log('Line', index + 1);
	console.log('Frequency Map: ', freqMap);

	let outputDigits = [];
	line.output.forEach(segment => {
		let outputFreqMap = { // Maps the output into a nice little object containing the frequency of each digit (0-1)
			a: 0,
			b: 0,
			c: 0,
			d: 0,
			e: 0,
			f: 0,
			g: 0
		};

		segment.split('').forEach(char => {
			outputFreqMap[char]++;
		});
		
		let digitScore = 0;
		for (const [key, value] of Object.entries(outputFreqMap)) {
			if (value == 1) {
				digitScore += freqMap[key];
			}
		}

		outputDigits.push(Object.keys(segmentFreqMap)[Object.values(segmentFreqMap).indexOf(digitScore)]); // Gets the correct digit from the segmentFreqMap
	});
	console.log('Segment Display:', outputDigits.join(''));
	console.log();

	totalOutputsSum += parseInt(outputDigits.join('')); // Sum up all the outputs
});

console.log('Puzzle Two Answer:', totalOutputsSum);