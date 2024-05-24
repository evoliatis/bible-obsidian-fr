# Source pour récupérer les Bibles

https://videopsalm.weebly.com/bibles-json-zip.html

Ces fichiers ne sont pas exploitables tels quels par Golang car il ne s'agit pas de fichiers JSON stricts, il faut donc les réadapter en quelques opérations simples.

# Pour le bon parsing JSON compatible GoLang

* Depuis un éditeur de texte réencoder le fichier en UTF-8 (pas UTF8-BOM)
* ajouter à tous les noms de champs des guillemets : (Attention Guid / Verses / Text!) 
ex :  Text: -> "Text":

Le fichier doit être intégre pour votre navigateur
 
* supprimer tous les retours à la ligne (sous Linux/Mac non testé sous Windows): 

```bash
tr '\n' ' ' < LSG.json > LSG2.json  
```

* Créer le verset 1 en remplaçant "Verses":[{ "Text" par "Verses":[{ "ID":1, "Text" -> 1189 fois
* Créer le livre 0 en replaçant [{"Books":[{"Chapters" par [{"Books":[{ "ID":0,"Chapters" -> 1 fois
* Créer le chapitre 1 en remplaçant "Chapters":[{"Verses" par "Chapters":[{"ID":1 ,"Verses" -> 66 fois