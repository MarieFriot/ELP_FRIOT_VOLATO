var prompt  = require('prompt');
let Apparitions = [
    [14, "A"], [4, "B"], [7, "C"], [5, "D"], [19, "E"], [2, "F"], [4, "G"],
    [2, "H"], [11, "I"], [1, "J"], [1, "K"], [6, "L"], [5, "M"], [9, "N"],
    [8, "O"], [4, "P"], [1, "Q"], [10, "R"], [7, "S"], [9, "T"], [8, "U"],
    [2, "V"], [1, "W"], [1, "X"], [1, "Y"], [2, "Z"]
];

let Letters = Apparitions.flatMap(([amount, letter]) => Array.from({ length: amount }, () => letter))

const randomItem = arr => {
	const index = (Math.random() * arr.length) | 0;
	return arr.splice(index,1)[0];};


function pioche(tourNumber,arr, nameJoueur, callback) {
	if(tourNumber <2){
		Array.from({ length: 6 }).forEach(() => {
			arr.push(randomItem(Letters));
		});
		if (callback){
			callback(nameJoueur, arr);
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
			if (callback){
				callback(nameJoueur, arr);
			}
		});
	}
}

function pioche1(arr, nameJoueur,  callback){
	Array.from({ length: 1 }).forEach(() => {
		arr.push(randomItem(Letters));
	});
	if (callback){
		callback(nameJoueur, arr);
	}
}

function newLetter(word, oldWord) {
    let wordOccurrences = {}; // Dictionnaire contenant le nombre d'occurence de chaques lettres qui étaient dans la playerPile
    let oldWordOccurrences = {}; //Dictionnaire contenant le nombre d'occurence des lettres ajoutées
    let newLetters = [];

    if(oldWord == ''){
        newLetters = word.split('')
    }else{
        word = word.split('')
        oldWord = oldWord.split('')

        word.forEach(lettre => {
            wordOccurrences[lettre] = (wordOccurrences[lettre] || 0) + 1;
        });


        oldWord.forEach(lettre => {
            oldWordOccurrences[lettre] = (oldWordOccurrences[lettre] || 0) + 1;
        });

       
        Object.keys(wordOccurrences).forEach(lettre => {
            if (oldWordOccurrences[lettre]>0) {
                wordOccurrences[lettre] -= oldWordOccurrences[lettre];
            }
        });
        
        Object.keys(wordOccurrences).forEach(lettre => {
            for (let i = 0; i < wordOccurrences[lettre]; i++) {
                newLetters.push(lettre);
            }
        });
    }

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

    playerPile = nouvellePioche.split('');
   
    return playerPile;

}

module.exports = {
    pioche,
    removePioche,
    pioche1,
    randomItem,
    newLetter
};