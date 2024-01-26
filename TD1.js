var prompt  = require('prompt');
var fs = require('fs');
prompt.start();
let Apparitions = [
    [14, "A"], [4, "B"], [7, "C"], [5, "D"], [19, "E"], [2, "F"], [4, "G"],
    [2, "H"], [11, "I"], [1, "J"], [1, "K"], [6, "L"], [5, "M"], [9, "N"],
    [8, "O"], [4, "P"], [1, "Q"], [10, "R"], [7, "S"], [9, "T"], [8, "U"],
    [2, "V"], [1, "W"], [1, "X"], [1, "Y"], [2, "Z"]
];

let Letters = Apparitions.flatMap(([amount, letter]) => Array.from({ length: amount }, () => letter))
var listeLettres = [ "A","B", "C","D"];
var grille1 = Array(8).fill("")
var grille2 = Array(8).fill("")
var playerPile1 = []
var playerPile2 = []

const randomItem = arr => arr.splice((Math.random() * arr.length) | 0, 1);

function pioche(arr, n, callback) {
   arr.push(...Array.from({ length: n }, () => randomItem(Letters)));
   console.log("Voici ta pioche")
   console.log(arr);
   if (callback){
	callback();
   }

}

function removePioche( word, playerPile){
	let piocheArray = playerPile.slice();
	piocheArray = piocheArray.filter(lettre => !word.includes(lettre));
	return piocheArray;
}

function playAgain(){
	console.log("Veux tu continuer ? [oui/non]");
	prompt.get(['answer'], function(err, result){
		if (result.answer == "non"){
			return;
		}
		else{
			pioche(playerPile1, 1 , function(){askWord();});
		}
	});
}


function askWord(callback){
	console.log("Voici ta grille")
	console.log(grille1);
	prompt.get(['ligne', 'mot'], function(err, result){
		grille1[parseInt(result.ligne)-1] = result.mot;
		console.log(grille1);
		playerPile1 = removePioche(result.mot, playerPile1);
		console.log(playerPile1);
	 	fs.writeFile('test.txt', result.mot, (err) => {
        	if (err) {
            	console.error(err);
        	} else {
            	console.log('Le fichier a étécorrectement écrit.');
		playAgain();
		if (callback){
			callback();
		};
        	}
    	});
})
}

function tour(nameJoueur){
	console.log("C'est au joueur " + nameJoueur + "de jouer");
	pioche(playerPile1, 6 , function(){askWord();});
}

tour("Joueur1");
