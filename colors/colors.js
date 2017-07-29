var images = [
  ['blue', 'green', 'red'],
  ['red', 'black', 'white'],
  ['white', 'black', 'yellow'],
];

var d = {};
images.forEach(colors => colors.forEach(c => {
  if(d[c] == null) {
    d[c] = 0;
  }
  d[c]++;
}));

console.log(d);

var topColor = 0;
var color = [];

Object.keys(d).forEach(c => {
  if(d[c] >= topColor) {
    topColor = d[c];
    if(color[topColor-1] == null) {
      color[topColor-1] = []
    }
    color[topColor-1].push(c);
  }
});

console.log(color.pop().sort())

