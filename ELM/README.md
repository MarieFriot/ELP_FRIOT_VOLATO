Ce code permet de générer une page web contenant un jeu : "Guess it!".  On choisi aléatoirement parmis les
1000 mots les plus courants utilisés dans le livre Thing Explainer: Complicated Stuff in Simple Words de Randall Munroe. Puis on affiche sa définition issue de Free Dictionary API. 
Ensuite, le joueur doit essayer de trouver le mot correspondant puis doit écrire sa réponse dans la zone d'écriture.
Si il a juste  on affiche Got it! It is indeed XXX. Si il a faux, rien ne se passe.
Pour finir, le joueur peut demander de voir la réponse à l'aide du bouton "show answer" puis de la cacher à l'aide du bouton "hide answer".

Le model peut avoir trois états : Loading quand il charge, Error si il y a une erreur sur la récupération du mot ou de sa définition et Guessing quand le joueur peut jouer.
