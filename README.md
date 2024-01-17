# GO_FRIOT_VOLATO
On créer un serveur, qui est en ecoute. Lors de la réception d'une liste d'entier de la part d'un client, le serveur partitionne les valeurs recues au fur et à mesure de leur réception en deux liste par rapport à une valeur pivot. Une fois les deux sous-listes creent, nous appliquons le quicksort en parallele sur les listes. Le serveur réasssemble les deux listes triées en une même liste triée et la renvoie au client.

On créer un client qui génère une grande liste de nombre entier et qui l'envoie au serveur.

la version V1 fonctionne avec des buffer de taille 1 (soit 4 bytes) c'est-a-dire, chaque element de la liste est envoye 1 par 1. La version V2 est legerement optimiser, elle travaille avec des buffer de taille arbitraire (par defaut on a mis une taille de buffer de 100) et les derniers envoie et recpetion adapte leur taille pour etre uniquement de longueur egal au nombre de valeur restant a envoyer/recevoir.
