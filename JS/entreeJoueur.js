var prompt  = require('prompt');
const modPioche = require('./pioche.js');

function newLetter(word, oldWord) {
    const newLetters = [];
    const oldWordSet = new Set(oldWord);

    for (let i = 0; i < word.length; i++) {
        const letter = word[i];
        if (!oldWordSet.has(letter)) {
            newLetters.push(letter);
        }
    }

    return newLetters;
}

function motJuste(word, oldWord, playerPile) {
    const newLetters = newLetter(word, oldWord);
    console.log(newLetters)
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
            if ((lineNumber >= 0 && lineNumber < grille.length) && motJuste(result.mot, grille[lineNumber-1], playerPile)){
                var oldWord = grille[lineNumber - 1];
                grille[lineNumber - 1] = result.mot;
                console.log(grille);
                playerPile = modPioche.removePioche(result.mot, playerPile, oldWord);
                console.log(playerPile);
                resolve({ grille, playerPile });
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
            if ((lineNumber >= 0 && lineNumber < grille.length) && grille[lineNumber -1]== '' && motJuste(mot,'',lettres)){
                console.log(mot);
                console.log(lineNumber);
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
            if (lineNumber >= 0 && lineNumber < grille.length){
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
                console.log(lettres)
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