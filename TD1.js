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
var grille1 = Array(8).fill("")
var grille2 = Array(8).fill("")
var playerPile1 = []
var playerPile2 = []

const randomItem = arr => {
	const index = (Math.random() * arr.length) | 0;
	return arr.splice(index,1)[0];};

function pioche(arr, n, callback) {
    Array.from({ length: n }).forEach(() => {
        arr.push(randomItem(Letters));
    });
   console.log("Voici ta pioche")
   console.log(arr);
   if (callback){
	callback();
   }

}

function newLetter(word, oldWord) {
    const newLetters = [];
    const wordArray = word.split('');
    const oldWordSet = new Set(oldWord);;

    wordArray.forEach(letter => {
        if (!oldWordSet.has(letter)) {
            newLetters.push(letter);
        }
    });

    return newLetters;
}

function removePioche(word, playerPile, oldword) {
    let newLetters = newLetter(word, oldword);	
    let piocheOccurrences = {};
    let lettresOccurrences = {};

    playerPile.forEach(lettre => {
        piocheOccurrences[lettre] = (piocheOccurrences[lettre] || 0) + 1;
    });

    newLetters.forEach(lettre => {
        lettresOccurrences[lettre] = (lettresOccurrences[lettre] || 0) + 1;
    });

    Object.keys(lettresOccurrences).forEach(lettre => {
        if (piocheOccurrences[lettre]) {
            piocheOccurrences[lettre] -= lettresOccurrences[lettre];
        }
    });

    let nouvellePioche = '';
    Object.keys(piocheOccurrences).forEach(lettre => {
        nouvellePioche += lettre.repeat(piocheOccurrences[lettre]);
    });

    return nouvellePioche.split('');
}

function playAgain(nameJoueur){
	let playerPile
	if (nameJoueur === "1") {
        	playerPile = playerPile1;
    	}else {
        	playerPile = playerPile2;
    	}
	console.log("Veux tu continuer ? [oui/non]");
	prompt.get(['answer'], function(err, result){
		if (result.answer == "non"){
			if(nameJoueur == "1"){
				tour("2");}
			else{
				tour("1");
			}
		}
		else{
			pioche(playerPile, 1 , function(){askWord(nameJoueur);});
		}
	});
}


function askWord(nameJoueur,callback){
	let grille, playerPile;
	if (nameJoueur == "1"){
		grille = grille1;
		playerPile = playerPile1;
	}
	else{
		grille = grille2;
		playerPile = playerPile2;
	}
	console.log("Voici ta grille")
	console.log(grille);
	prompt.get(['ligne', 'mot'], function(err, result){
		oldWord = grille[parseInt(result.ligne)-1];
		grille[parseInt(result.ligne)-1] = result.mot;
		console.log(grille);
		playerPile = removePioche(result.mot, playerPile, oldWord);
		console.log(playerPile);
	 	fs.writeFile('test.txt', result.mot, (err) => {
        	if (err) {
            	console.error(err);
        	} else {
            	console.log('Le fichier a été correctement écrit.');
		playAgain(nameJoueur);
		if (callback){
			callback();
		};
        	}
    	});
})
}

function tour(nameJoueur,callback){
	let playerPile;
	console.log("C'est au joueur " + nameJoueur + "de jouer");
	if (nameJoueur == "1"){
		playerPile = playerPile1;
	}
	else{
	       playerPile = playerPile2;
	}
	pioche(playerPile, 6 , function(){askWord(nameJoueur);});
	if (callback){
		callback();
	}
}

tour("1");
