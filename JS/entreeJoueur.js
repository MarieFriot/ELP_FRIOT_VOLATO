var prompt  = require('prompt');
const modPioche = require('./pioche.js');



function motJuste(word, oldWord, playerPile) {
    const newLetters = modPioche.newLetter(word, oldWord);
    // Vérifiez chaque nouvelle lettre et sa quantité dans playerPile
    return newLetters.every(letter => {
        const newLetterCount = newLetters.filter(l => l === letter).length;
        const playerPileCount = playerPile.filter(l => l === letter).length;
        return playerPileCount >= newLetterCount;
    });
}

function getLineWord(grille, playerPile) {
    return new Promise((resolve) => {
        prompt.get(['ligne', 'mot'], function (err, result) {
            var lineNumber = parseInt(result.ligne);
            if ((lineNumber > 0 && lineNumber <= grille.length) && result.mot.length >= 3 && motJuste(result.mot, grille[lineNumber-1], playerPile)){
                var oldWord = grille[lineNumber - 1];
                grille[lineNumber - 1] = result.mot;
                playerPile = modPioche.removePioche(result.mot, playerPile, oldWord);
                let mot = result.mot;
                resolve({ grille : grille, playerPile :playerPile, mot : mot , ligne : lineNumber });
            }else {
                console.log("Le numéro de ligne doit être compris entre 1 et 8 et vous devez utilisez les lettres disponibles. Veuillez réessayer.");
                // Appel récursif pour demander à nouveau à l'utilisateur de saisir le numéro de ligne.
                resolve(getLineWord(grille, playerPile));
            }
        });
    });
}

function getLineWordJarnac(grille, lettres) {
    return new Promise((resolve) => {
        prompt.get(['ligne', 'mot'], function (err, result) {
            var lineNumber = parseInt(result.ligne);
            var mot = result.mot;
            if ((lineNumber > 0 && lineNumber <= grille.length) && grille[lineNumber -1]== '' && motJuste(mot,'',lettres)){
                resolve({ ligne : lineNumber,  mot :mot});
            }else {
                console.log("Le numéro de ligne doit être compris entre 1 et 8 et vous devez utilisez les volées. Veuillez réessayer.");
                // Appel récursif pour demander à nouveau à l'utilisateur de saisir le numéro de ligne.
                resolve(getLineWordJarnac(grille, lettres));
            }
        });
    });
}

function getLineJarnac(grille) {
    return new Promise((resolve) => {
        prompt.get(['ligne'], function (err, result) {
            var lineNumber = parseInt(result.ligne);
            if (lineNumber > 0 && lineNumber <= grille.length){
                resolve(lineNumber);
            }else {
                console.log("Le numéro de ligne doit être compris entre 1 et 8 et vous devez utilisez les lettres disponibles. Veuillez réessayer.");
                // Appel récursif pour demander à nouveau à l'utilisateur de saisir le numéro de ligne.
                resolve(getLineJarnac(grille));
            }
        });
    });
}

function getLettresJarnac(playerPile){
	return new Promise((resolve) => {
        prompt.get(['lettres'], function (err, result) {
            var lettres = result.lettres
			if (motJuste(lettres, '', playerPile)) {
                resolve(lettres);
            }else{
                console.log("Tu dois indiquée des lettres qui sont dans la pioche de l'autres joueurs.");
                // Appel récursif pour demander à nouveau à l'utilisateur de saisir le numéro de ligne.
                resolve(getLettresJarnac(playerPile));
            }
        });
    });

}

module.exports = {
    getLineWord,
    getLineWordJarnac,
    getLineJarnac,
    getLettresJarnac
};