var prompt  = require('prompt');
var fs = require('fs');
const modPioche = require('./pioche.js');
const modFin = require('./finPartie.js');
const modEntree = require('./entreeJoueur.js');
const { log } = require('console');

prompt.start();

var grille1 = Array(8).fill("")
var grille2 = Array(8).fill("")
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
	let fin, playerPile; 
	console.log('-------------------------------------------------------\n')
	console.log("Au tour du joueur " + nameJoueur);
	console.log("Veux tu faire un Jarnac ?[oui/non]");
	prompt.get(['answer'], function(err,result){
		let grille;
		if(result.answer == "oui"){
			console.log("Quelle ligne veux tu lui voler ? (Donne son numéro)")
			modEntree.getLineJarnac(grille1).then((ligneVolée) => {
				if (nameJoueur == "1"){
					playerPile= playerPile2;
				}else{
					playerPile = playerPile1;
				}
				console.log("Quelles lettre(s) de sa pioche veux tu lui voler ?")
				modEntree.getLettresJarnac(playerPile).then((lettres) => {
					let newlettres;
					console.log("Pour rapel voici ta grille")
					if (nameJoueur == "1"){
						grille = grille1;
						newlettres = lettres + grille2[ligneVolée - 1];
					}else{
						grille = grille2;
						newlettres = lettres + grille1[ligneVolée - 1];
						console.log(ligneVolée)
						console.log(grille[ligneVolée -1]);
					}
					console.log("Pour rapel voici ta grille")
					console.log(grille)
					console.log("Quel mot et à quelle position veux-tu l'écrire?")
					console.log(lettres)
					modEntree.getLineWordJarnac(grille, newlettres.split('')).then((result) => {
						const {ligne, mot} = result;
						if (nameJoueur == "1"){
							grille2[ligneVolée-1]= "";
							grille1[ligne -1] = mot;
							fin = modFin.finPartie("1", grille1, grille2);
							playerPile2 = modPioche.removePioche(lettres, playerPile2, "");
						}else{
							grille1[ligneVolée-1]= "";
							grille2[ligne-1] = mot;
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
                modEntree.getLineWord(grille, playerPile).then((result) => {
					const {grille , playerPile ,ligne, mot} = result;
                    if (nameJoueur == "1") {
                        grille1 = grille;
                        playerPile1 = playerPile;
                        console.log(playerPile1);
                        fin = modFin.finPartie("1", grille1, grille2);
                    } else {
                        grille2 = grille;
                        playerPile2 = playerPile;
                        console.log(playerPile2);
                        fin = modFin.finPartie("2", grille1, grille2);
                    }
                    let log = `Au tour ${tourNumber}, le joueur ${nameJoueur} a écrit le mot '${mot}' sur la ligne ${ligne}\n`;
                    fs.appendFile('test.txt', log, (err) => {
                        if (err) {
                            console.error(err);
                        }
                    });
                    console.log("");

                    if (!fin) {
                        console.log("Tu peux rejouer si tu veux !");
                        playAgain(nameJoueur);
                    }
                })
                .catch(error => {
                    console.error(error);
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
