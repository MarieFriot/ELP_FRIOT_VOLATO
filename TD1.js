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

function pioche(arr, callback) {
	if(tourNumber <2){
		Array.from({ length: 6 }).forEach(() => {
			arr.push(randomItem(Letters));
		});
		console.log("Voici ta pioche")
		console.log(arr);
		if (callback){
			callback();
		  }
	}else{
		console.log("Veux tu piocher 1 lettres [oui/non]. Si non dis moi les 3 lettres que tu veux échanger")
		prompt.get(['Answer'], function(err, result){
			let card;
			if (result.Answer == "oui"){
				card = 1;
				Array.from({ length: card }).forEach(() => {
					arr.push(randomItem(Letters));
				})
				console.log("Voici ta pioche")
				console.log(arr);
				if (callback){
					callback();
		  		}
			}else{
				card = 3;
				let lettres = result.Answer.split('');
				arr = arr.filter(letter => !lettres.includes(letter));
				Array.from({ length: card }).forEach(() => {
					arr.push(randomItem(Letters));
				})
				console.log("Voici ta pioche")
				console.log(arr);
				if (callback){
					callback();
		  		}
			}
		});
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
			tourNumber = tourNumber +1;
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
		if (nameJoueur == "1"){
			grille1 = grille;
			playerPile1 = playerPile;
		}
		else{
			grille2 = grille;
			playerPile2 = playerPile;
		}
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
	
	pioche(playerPile, function(){askWord(nameJoueur);});
	if (callback){
		callback();
	}
}

var tourNumber = 0;
tour("1");
