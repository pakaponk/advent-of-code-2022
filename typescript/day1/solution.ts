import fs from 'fs';
import path from 'path';

function readFile() {
  return fs.readFileSync(path.resolve(__dirname, './input.txt')).toString();
}

function solvePart1(text: string) {
  const elves = text.split('\n\n');

  const weights = elves.map((str) =>
    str
      .split('\n')
      .map((str) => parseInt(str))
      .reduce((sum, curr) => sum + curr)
  );

  return Math.max(...weights);
}

function solvePart2(text: string) {
  const elves = text.split('\n\n');

  const weights = elves.map((str) =>
    str
      .split('\n')
      .map((str) => parseInt(str))
      .reduce((sum, curr) => sum + curr)
  );

  const [one, two, three] = weights.sort((a, z) => z - a);

  return one + two + three;
}

const text = readFile();
console.log(`Part 1: ${solvePart1(text)}`);
console.log(`Part 2: ${solvePart2(text)}`);
