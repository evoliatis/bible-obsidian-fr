# bible-obsidian-fr

## Introduction

Vous trouverez dans ce dépôt les dossiers suivants :

- **BIBLE** : la bible en version libre de droits : Louis Segond 1910, Darby et Martin au format Obsidian
- **INTRODUCTION** : vous pouvez y ajouter ici une introduction pour chaque livre
- **originaux** : les fichiers JSON qui ont servi de base à cette génération
- **src** : les fichiers sources adaptés en JSON stricts utilisable par le programme golang

A la racine du dossier vous trouverez le source Golang de l'outil qui génère les dossiers BIBLE et INTRODUCTION. 

## Utilisation

Les liens générés et utilisables sous Obsidian se basent sur le projet : [Bible Study Kit v1](https://forum.obsidian.md/t/bible-study-in-obsidian-kit-including-the-bible-in-markdown/12503)

Pour un même verset vous pouvez utiliser

```
[[Jean 3#16]] : version Louis Segond
[[Jean 3#16[DRB]]] : version Darby
[[Jean 3#16[FMAR]]] : version Martin
![[Jean 3#Jean 3 16]] : pour la version inline avec la référence complète
```
En plus des alias des noms de livres de la version anglaise et la version abrégée du nom du livre : Jn / John...

## Compilation

Pour ceux qui veulent jouer avec voici les commandes principales assez standard :

```bash
go mod tidy
go run main.go
```

Si vous souhaitez ajouter des versions avec d'autres fichiers JSON vous pouvez décommenter/adapter la section `initVersionName` ligne 783

## Copyright

> Attention
> Les versions proposées ici sont libres de droits mais toutes autres versions, sauf erreur de ma part, ne le sont pas : donc je ne peux pas les publier ici !

En ce qui concerne ce code, je le laisse sous licence [GPL v3](https://www.gnu.org/licenses/quick-guide-gplv3.fr.html). 
N'hésitez pas à la modifier. Si vous avez des amélioration à me proposer sentez-vous libre !

Fraternellement