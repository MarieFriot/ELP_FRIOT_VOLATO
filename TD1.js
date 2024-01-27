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


function pioche1(arr,  callback){
	Array.from({ length: 1 }).forEach(() => {
		arr.push(randomItem(Letters));
	});
	console.log("Voici ta pioche")
	console.log(arr);
	if (callback){
		callback();
	}

}

function pioche(arr, nameJoueur, callback) {
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
		console.log("Voici l'état de ta pioche:")
		console.log(arr)
		console.log("Veux tu piocher 1 lettres [oui/non]. Si non dis moi les 3 lettres que tu veux échanger")
		prompt.get(['Answer'], function(err, result){
			let card;
			if (result.Answer == "oui"){
				card = 1;
			}else{
				card = 3;
				arr = removePioche(result.Answer, arr, ""); // cette modification se fait que localement 
			}
			Array.from({ length: card }).forEach(() => {
				arr.push(randomItem(Letters));
			});
			if (nameJoueur == "1"){
				playerPile1 = arr;
			}
			else{
				playerPile2 = arr;
			}
			console.log("Voici ta pioche")
			console.log(arr);
			if (callback){
				callback();
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
    let piocheOccurrences = {}; // Dictionnaire contenant le nombre d'occurence de chaques lettres qui étaient dans la playerPile
    let lettresOccurrences = {}; //Dictionnaire contenant le nombre d'occurence des lettres ajoutées

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

function Jarnac(nameJoueur, callback){
	console.log('\n')
	console.log("Au tour du joueur " + nameJoueur);
	console.log("Veux tu faire un Jarnac ?[oui/non]");
	prompt.get(['answer'], function(err,result){
		if(result.answer == "oui"){
			console.log("Quelle ligne veux tu lui voler ? (Donne son numéro)")
			prompt.get(['ligne'], function(err,result){
				let ligne = result.ligne;
				console.log("Quelles lettre(s) de sa pioche veux tu lui voler ?")
				prompt.get(['lettres'], function(err, result){
					console.log("Quel mot et à quelle position veux-tu l'écrire?")
					prompt.get(['mot', 'position'], function(err,result){
					if (nameJoueur == "1"){
						grille2[parseInt(ligne)-1]= "";
						grille1[parseInt(result.position)-1] = result.mot;
						removePioche(result.lettres, playerPile2, "");
					}else{
						grille1[parseInt(ligne)-1]= "";
						grille2[parseInt(result.position)-1] = result.mot;
						removePioche(result.lettres, playerPile1, "");
					}
					if (callback){
						callback();
					}

					});
				});
			});
		}else{
		if (callback){
			callback();
		}}

	});


}



function playAgain(nameJoueur){
	console.log("Veux tu continuer ? [oui/non]");
	prompt.get(['answer'], function(err, result){
		if (result.answer == "non"){
			tourNumber = tourNumber +1;
			if(nameJoueur == "1"){
				Jarnac("2", function(){tour("2")});
			}else{
				Jarnac("1", function(){tour("1")});
			}
		}
		else{
			if (nameJoueur === "1") {
				pioche1(playerPile1, function(){askWord(nameJoueur);});
			}else {
				pioche1(playerPile2, function(){askWord(nameJoueur);});
			}

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
	console.log('')
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
		let log = `Au tour ${tourNumber}, le joueur ${nameJoueur} a écrit le mot '${result.mot}' sur la ligne ${result.ligne}\n`;
	 	fs.appendFile('test.txt',log , (err) => {
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
	console.log('\n');
	if (nameJoueur == "1"){
		pioche(playerPile1, nameJoueur, function(){askWord(nameJoueur);});
	}
	else{
		pioche(playerPile2, nameJoueur, function(){askWord(nameJoueur);});
	    
	}
	if (callback){
		callback();
	}
}

var tourNumber = 0;
console.log("Le jeu commence ! N'oubliez pas d'écrire les lettres en majuscule !")
console.log("C'est au premier joueur de jouer !")
tour("1");
