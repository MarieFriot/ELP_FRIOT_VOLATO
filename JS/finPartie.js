function score(grille){
	// Utiliser map pour créer un tableau de points pour chaque mot (taille au carré)
	const pointsParMot = grille.map(mot => Math.pow(mot.length, 2));
	// Utiliser reduce pour calculer la somme totale des points
	const sommePoints = pointsParMot.reduce((total, points) => total + points, 0);
	return sommePoints;
}

function finPartie(nameJoueur, grille1, grille2){
	let score1, score2, grille;
    if (nameJoueur == "1"){
        grille = grille1;
    }else{
        grille = grille2;
    }
    console.log(grille)
	if ( grille.every(mot => mot.length >= 3)){
		console.log("C'est la fin de la partie !")
		score1 = score(grille1);
		score2 = score(grille2);
		console.log("Le joueur 1 a " + score1 + " points et le joueur 2 en a " + score2)
		if (score1 >score2){
			console.log("Le joueur 1 a gagné !")
		}
		else if(score2 > score1){
			console.log("Le joueur 2 a gagné!")
		}else{
			console.log("Egalité!")
		}
		return true;
	}else{
		return false
	}
}

module.exports = {
    finPartie
};