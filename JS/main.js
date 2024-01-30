var prompt  = require('prompt');
var fs = require('fs');
const modPioche = require('./pioche.js');
const modFin = require('./finPartie.js')

prompt.start();

var grille1 = Array(2).fill("")
var grille2 = Array(2).fill("")
var playerPile1 = []
var playerPile2 = []


function nextPlayer(nameJoueur){
	tourNumber = tourNumber +1;
	if(nameJoueur == "1"){
		Jarnac("2", function(){tour("2")});
	}else{
		Jarnac("1", function(){tour("1")});
	}
}

function playAgain(nameJoueur){
		if (nameJoueur === "1") {
			modPioche.pioche1(playerPile1, nameJoueur, askWord);
		}else {
			modPioche.pioche1(playerPile2, nameJoueur, askWord);
		}
}

function Jarnac(nameJoueur, callback){
	let fin
	console.log('-------------------------------------------------------\n')
	console.log("Au tour du joueur " + nameJoueur);
	console.log("Veux tu faire un Jarnac ?[oui/non]");
	prompt.get(['answer'], function(err,result){
		if(result.answer == "oui"){
			console.log("Quelle ligne veux tu lui voler ? (Donne son numéro)")
			prompt.get(['ligne'], function(err,result){
				let ligne = result.ligne;
				console.log("Quelles lettre(s) de sa pioche veux tu lui voler ?")
				prompt.get(['lettres'], function(err, result){
					let lettres = result.lettres
					console.log("Pour rapel voici ta grille")
					if (nameJoueur == "1"){
						console.log(grille1);
					}else{
						console.log(grille2);
					}
					console.log("Quel mot et à quelle position veux-tu l'écrire?")
					prompt.get(['mot', 'position'], function(err,result){
					if (nameJoueur == "1"){
						grille2[parseInt(ligne)-1]= "";
						grille1[parseInt(result.position)-1] = result.mot;
						fin = modFin.finPartie("1", grille1, grille2);
						playerPile2 = modPioche.removePioche(lettres, playerPile2, "");
					}else{
						grille1[parseInt(ligne)-1]= "";
						grille2[parseInt(result.position)-1] = result.mot;
						fin = modFin.finPartie("2", grille1, grille2);
						playerPile1 = modPioche.removePioche(lettres, playerPile1, "");
					}
					if (!fin){
						if (callback){
							callback();
					}}

					});
				});
			});
		}else{
		if (!fin){
			if (callback){
				callback();
			}}
		}

	});

}

function askWord(nameJoueur, playerPile){
	let grille;
	let fin;
	if (nameJoueur == "1"){
		grille = grille1;
	}
	else{
		grille = grille2;
	}
	console.log('')
	console.log("Voici ta pioche")
	console.log(playerPile)
	console.log("Voici ta grille")
	console.log(grille);
	console.log("Veux tu jouer [oui/non] ?")
	prompt.get(['answer'], function(err, result){
		if(result.answer == "non"){
			if (nameJoueur == "1"){
				playerPile1 = playerPile;
			}
			else{
				playerPile2 = playerPile;
			}
			nextPlayer(nameJoueur);
		}else{
			console.log("Dis moi sur quelle ligne tu veux écrire ton mot.")
			prompt.get(['ligne', 'mot'], function(err, result){
				oldWord = grille[parseInt(result.ligne)-1];
				grille[parseInt(result.ligne)-1] = result.mot;
				console.log(grille);
				playerPile = modPioche.removePioche(result.mot, playerPile, oldWord);
				console.log(playerPile)

				if (nameJoueur == "1"){
					grille1 = grille;
					playerPile1 = playerPile;
					console.log(playerPile1)
					fin = modFin.finPartie("1",grille1, grille2);
				}
				else{
					grille2 = grille;
					playerPile2 = playerPile;
					console.log(playerPile2)
					fin = modFin.finPartie("2", grille1, grille2);
				}
				let log = `Au tour ${tourNumber}, le joueur ${nameJoueur} a écrit le mot '${result.mot}' sur la ligne ${result.ligne}\n`;
				fs.appendFile('test.txt',log , (err) => {
					if (err) {
						console.error(err);
					}
				});
				console.log("")
				if (!fin){
					console.log("Tu peux rejouer si tu veux !")
					playAgain(nameJoueur);
				}
			});
		}
	});
}


function tour(nameJoueur,callback){
	console.log('---------------------------------------------------\n');
	if (nameJoueur == "1"){
		modPioche.pioche(tourNumber, playerPile1, nameJoueur, askWord);
	}
	else{
		modPioche.pioche(tourNumber, playerPile2, nameJoueur, askWord);
	}
	if (callback){
		callback();
	}
}

var tourNumber = 0;
console.log("Le jeu commence ! N'oubliez pas d'écrire les lettres en majuscule !")
console.log("C'est au premier joueur de jouer !")
tour("1");
