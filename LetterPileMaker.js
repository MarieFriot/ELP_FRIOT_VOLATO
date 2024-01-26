let Apparitions = [
    [14, "A"], [4, "B"], [7, "C"], [5, "D"], [19, "E"], [2, "F"], [4, "G"],
    [2, "H"], [11, "I"], [1, "J"], [1, "K"], [6, "L"], [5, "M"], [9, "N"],
    [8, "O"], [4, "P"], [1, "Q"], [10, "R"], [7, "S"], [9, "T"], [8, "U"],
    [2, "V"], [1, "W"], [1, "X"], [1, "Y"], [2, "Z"]
];

let Letters = Apparitions.flatMap(([amount, letter]) => Array.from({ length: amount }, () => letter));

console.log( Letters[3])

const randomItem = arr => arr.splice((Math.random() * arr.length) | 0, 1);

console.log(randomItem(Letters));
console.log(Letters);


playerPile = []

function drawLetters(arr, n) {
    arr.push(...Array.from({ length: n }, () => randomItem(Letters)));
}

drawLetters(playerPile, 4)
console.log(playerPile)
