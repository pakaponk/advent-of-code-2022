import fs from 'fs';
import path from 'path';

function readFile() {
  return fs.readFileSync(path.resolve(__dirname, './input.txt')).toString();
}

function getWeights(text: string) {
  const elves = text.split('\n\n');

  return elves.map((str) =>
    str
      .split('\n')
      .map((str) => parseInt(str))
      .reduce((sum, curr) => sum + curr)
  );
}

function solvePart1(text: string) {
  const weights = getWeights(text);

  return Math.max(...weights);
}

function solvePart2(text: string) {
  const weights = getWeights(text);

  const [one, two, three] = weights.sort((a, z) => z - a);

  return one + two + three;
}

const text = readFile();
console.log(`Part 1: ${solvePart1(text)}`);
console.log(`Part 2: ${solvePart2(text)}`);
