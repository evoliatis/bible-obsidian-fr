package main

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
)

type Verse struct {
	ID   int
	Text string
}

type Chapter struct {
	ID     int
	Verses []Verse
}

type Book struct {
	ID       int
	Chapters []Chapter
	Text     string
}

type Testament struct {
	Books []Book
	Text  string
}

type Bible struct {
	Abbreviation string
	Language     string
	Publisher    string
	Description  string
	VersionDate  string
	Text         string
	IsCompressed int8
	IsProtected  int8
	Guid         string
	Testaments   []Testament
}

type BookAliases struct {
	id      int
	aliases []string
}

// Nom du répertoire de destination
const dirname string = "BIBLE"
const dirintroname string = "INTRODUCTION"

// Tag pour le contenu
const tag = "Bible"

// Niveau pour les versets
const versesH string = "######"

// Tableau de traduction et aliases
var aliases []BookAliases = []BookAliases{
	{
		id: 1,
		aliases: []string{
			"Genèse",
			"Ge",
			"Genesis",
			"Gen",
		},
	},
	{
		id: 2,
		aliases: []string{
			"Exode",
			"Ex",
			"Exodus",
			"Exod",
		},
	},
	{
		id: 3,
		aliases: []string{
			"Lévitique",
			"Lé",
			"Leveticus",
			"Lev",
		},
	},
	{
		id: 4,
		aliases: []string{
			"Nombres",
			"No",
			"Numbers",
			"Num",
		},
	},
	{
		id: 5,
		aliases: []string{
			"Deutéronome",
			"De",
			"Deuteronomy",
			"Deut",
		},
	},
	{
		id: 6,
		aliases: []string{
			"Josué",
			"Jos",
			"Joshua",
			"Josh",
		},
	},
	{
		id: 7,
		aliases: []string{
			"Juges",
			"Jg",
			"Judges",
			"Judg",
		},
	},
	{
		id: 8,
		aliases: []string{
			"Ruth",
			"Ru",
		},
	},
	{
		id: 9,
		aliases: []string{
			"1 Samuel",
			"1 S",
			"1 Sam",
		},
	},
	{
		id: 10,
		aliases: []string{
			"2 Samuel",
			"2 S",
			"2 Sam",
		},
	},
	{
		id: 11,
		aliases: []string{
			"1 Rois",
			"1 R",
			"1 Kings",
		},
	},
	{
		id: 12,
		aliases: []string{
			"2 Rois",
			"2 R",
			"2 Kings",
		},
	},
	{
		id: 13,
		aliases: []string{
			"1 Chroniques",
			"1 Ch",
			"1 Chronicles",
			"1 Chron",
		},
	},
	{
		id: 14,
		aliases: []string{
			"2 Chroniques",
			"2 Ch",
			"2 Chronicles",
			"2 Chron",
		},
	},
	{
		id: 15,
		aliases: []string{
			"Esdras",
			"Esd",
			"Ezra",
			"Ezr",
		},
	},
	{
		id: 16,
		aliases: []string{
			"Néhémie",
			"Né",
			"Nehemiah",
			"Neh",
		},
	},
	{
		id: 17,
		aliases: []string{
			"Esther",
			"Est",
			"Esther",
			"Esth",
		},
	},
	{
		id: 18,
		aliases: []string{
			"Job",
		},
	},
	{
		id: 19,
		aliases: []string{
			"Psaumes",
			"Ps",
			"Psalms",
		},
	},
	{
		id: 20,
		aliases: []string{
			"Proverbes",
			"Pr",
			"Proverbs",
			"Prov",
		},
	},
	{
		id: 21,
		aliases: []string{
			"Ecclésiaste",
			"Ec",
			"Ecclesiastes",
			"Eccles",
		},
	},
	{
		id: 22,
		aliases: []string{
			"Cantique des Cantiques",
			"Ca",
			"Song of Solomon",
			"Song",
		},
	},
	{
		id: 23,
		aliases: []string{
			"Esaïe",
			"Es",
			"Isaiah",
			"Isa",
		},
	},
	{
		id: 24,
		aliases: []string{
			"Jérémie",
			"Jé",
			"Jeremiah",
			"Jer",
		},
	},
	{
		id: 25,
		aliases: []string{
			"Lamentations",
			"La",
			"Lamentations",
			"Lam",
		},
	},
	{
		id: 26,
		aliases: []string{
			"Ezéchiel",
			"Ez",
			"Ezekiel",
			"Ezek",
		},
	},
	{
		id: 27,
		aliases: []string{
			"Daniel",
			"Da",
			"Dan",
		},
	},
	{
		id: 28,
		aliases: []string{
			"Osée",
			"Os",
			"Hosea",
			"Hos",
		},
	},
	{
		id: 29,
		aliases: []string{
			"Joël",
			"Joë",
			"Joel",
		},
	},
	{
		id: 30,
		aliases: []string{
			"Amos",
			"Am",
		},
	},
	{
		id: 31,
		aliases: []string{
			"Abdias",
			"Ab",
			"Obadiah",
			"Obad",
		},
	},
	{
		id: 32,
		aliases: []string{
			"Jonas",
			"Jon",
			"Jonah",
		},
	},
	{
		id: 33,
		aliases: []string{
			"Michée",
			"Mi",
			"Micah",
		},
	},
	{
		id: 34,
		aliases: []string{
			"Nahoum",
			"Na",
			"Nahum",
			"Nah",
		},
	},
	{
		id: 35,
		aliases: []string{
			"Habaquq",
			"Ha",
			"Habakuk",
			"Hab",
		},
	},
	{
		id: 36,
		aliases: []string{
			"Sophonie",
			"So",
			"Zephaniah",
			"Zeph",
		},
	},
	{
		id: 37,
		aliases: []string{
			"Aggée",
			"Ag",
			"Haggai",
			"Hag",
		},
	},
	{
		id: 38,
		aliases: []string{
			"Zacharie",
			"Za",
			"Zechariah",
			"Zech",
		},
	},
	{
		id: 39,
		aliases: []string{
			"Malachie",
			"Mal",
			"Malachi",
			"Mal",
		},
	},
	{
		id: 40,
		aliases: []string{
			"Matthieu",
			"Mt",
			"Matthew",
			"Matt",
		},
	},
	{
		id: 41,
		aliases: []string{
			"Marc",
			"Mc",
			"Mark",
		},
	},
	{
		id: 42,
		aliases: []string{
			"Luc",
			"Lu",
			"Luke",
		},
	},
	{
		id: 43,
		aliases: []string{
			"Jean",
			"Jn",
			"John",
		},
	},
	{
		id: 44,
		aliases: []string{
			"Actes",
			"Ac",
			"Acts",
		},
	},
	{
		id: 45,
		aliases: []string{
			"Romains",
			"Ro",
			"Romans",
			"Rom",
		},
	},
	{
		id: 46,
		aliases: []string{
			"1 Corinthiens",
			"1 Co",
			"1 Corinthians",
			"1 Cor",
		},
	},
	{
		id: 47,
		aliases: []string{
			"2 Corinthiens",
			"2 Co",
			"2 Corinthians",
			"2 Cor",
		},
	},
	{
		id: 48,
		aliases: []string{
			"Galates",
			"Ga",
			"Galatians",
			"Gal",
		},
	},
	{
		id: 49,
		aliases: []string{
			"Ephésiens",
			"Ep",
			"Ephesians",
			"Ephes",
		},
	},
	{
		id: 50,
		aliases: []string{
			"Philippiens",
			"Ph",
			"Philippians",
			"Phil",
		},
	},
	{
		id: 51,
		aliases: []string{
			"Colossiens",
			"Col",
			"Colossians",
		},
	},
	{
		id: 52,
		aliases: []string{
			"1 Thessaloniciens",
			"1 Th",
			"1 Thessalonians",
			"1 Thess",
		},
	},
	{
		id: 53,
		aliases: []string{
			"2 Thessaloniciens",
			"2 Th",
			"2 Thessalonians",
			"2 Thess",
		},
	},
	{
		id: 54,
		aliases: []string{
			"1 Timothée",
			"1 Ti",
			"1 Timothy",
			"1 Tim",
		},
	},
	{
		id: 55,
		aliases: []string{
			"2 Timothée",
			"2 Ti",
			"2 Timothy",
			"2 Tim",
		},
	},
	{
		id: 56,
		aliases: []string{
			"Tite",
			"Tit",
			"Titus",
		},
	},
	{
		id: 57,
		aliases: []string{
			"Philémon",
			"Phm",
			"Philemon",
			"Philem",
		},
	},
	{
		id: 58,
		aliases: []string{
			"Hébreux",
			"Hé",
			"Hebrews",
			"Heb",
		},
	},
	{
		id: 59,
		aliases: []string{
			"Jacques",
			"Ja",
			"James",
		},
	},
	{
		id: 60,
		aliases: []string{
			"1 Pierre",
			"1 Pi",
			"1 Peter",
			"1 Pet",
		},
	},
	{
		id: 61,
		aliases: []string{
			"2 Pierre",
			"2 Pi",
			"2 Peter",
			"2 Pet",
		},
	},
	{
		id: 62,
		aliases: []string{
			"1 Jean",
			"1 Jn",
			"1 John",
		},
	},
	{
		id: 63,
		aliases: []string{
			"2 Jean",
			"2 Jn",
			"2 John",
		},
	},
	{
		id: 64,
		aliases: []string{
			"3 Jean",
			"3 Jn",
			"3 John",
		},
	},
	{
		id: 65,
		aliases: []string{
			"Jude",
			"Jud",
		},
	},
	{
		id: 66,
		aliases: []string{
			"Apocalypse",
			"Ap",
			"Revelation",
			"Rev",
		},
	},
}

// Modèle de cartouche Livre
var cartoucheLivre = []string{
	"---",
	"aliases:",
	"%Alias%",
	"tags:",
	"%Tags%",
	"BibleType: Livre",
	"---",
	"[[La Sainte Bible|Sommaire]]",
	"",
	"---",
	"",
	"%Introduction%",
	"",
	"---",
	"%Content%",
	"",
	"---",
}

// Modèle de cartouche Chapitre
var cartoucheChapter = []string{
	"---",
	"aliases:",
	"%Alias%",
	"tags:",
	"%Tags%",
	"BibleType: Chapitre",
	"---",
	"%Navbar%",
	"",
	"---",
}

// Modèle de cartouche Introduction
var cartoucheIntro = []string{
	"---",
	"tags:",
	"- Bible",
	"BibleType: Introduction",
	"---",
	"Votre texte ici",
	"",
	"---",
}

func CreateBookIntro(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	// remember to close the file
	defer f.Close()

	// Génération du contenu
	for _, line := range cartoucheIntro {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateChapterFile(dirname, filename string, id, size int) error {
	for i := 1; i < size+1; i++ {
		// create files
		si := strconv.Itoa(i)
		f, err := os.Create(dirname + "/" + filename + " " + si + ".md")
		if err != nil {
			return err
		}
		// remember to close the file
		defer f.Close()

		// create cartouche
		for _, line := range cartoucheChapter {
			var toDisk string = ""
			switch line {
			case "%Alias%":
				for _, s := range aliases[id].aliases {
					toDisk += "  - " + s + "\n"
				}
			case "%Tags%":
				toDisk = "  - " + tag
			case "%Navbar%":
				if i != 1 {
					left := strconv.Itoa(i - 1)
					toDisk += "[[" + filename + " " + left + "]] | "
				}
				toDisk += "[[" + filename + "]]"
				if i != size {
					right := strconv.Itoa(i + 1)
					toDisk += " | [[" + filename + " " + right + "]]"
				}
			default:
				toDisk = line
			}
			_, err := f.WriteString(toDisk + "\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Création du fichier livre
func CreateBookFile(dirname, filename string, id, size int) error {
	// create file
	f, err := os.Create(dirname + "/" + filename + ".md")
	if err != nil {
		return err
	}
	// remember to close the file
	defer f.Close()

	for _, line := range cartoucheLivre {
		var toDisk string = ""
		switch line {
		case "%Alias%":
			for _, s := range aliases[id].aliases {
				toDisk += "  - " + s + "\n"
			}
		case "%Tags%":
			toDisk = "  - " + tag
		case "%Introduction%":
			toDisk = "![[Introduction " + aliases[id].aliases[0] + "]]"
		case "%Content%":
			for i := 1; i < size+1; i++ {
				si := strconv.Itoa(i)
				toDisk += "- [[" + filename + " " + si + "]] \n"
			}

		default:
			toDisk = line
		}
		_, err := f.WriteString(toDisk + "\n")
		if err != nil {
			return err
		}

		err = CreateChapterFile(dirname, filename, id, size)
		if err != nil {
			return err
		}
	}

	return nil
}

// Mise en place des versions
// /!\ : La premère est la version part défaut
func initVersionName() []string {
	var versionName []string
	versionName = append(versionName, "LSG")
	versionName = append(versionName, "DRB")
	versionName = append(versionName, "FMAR")
	//	versionName = append(versionName, "FRC97")
	//	versionName = append(versionName, "BDS")
	//	versionName = append(versionName, "NBS")
	//	versionName = append(versionName, "NEG79")
	//	versionName = append(versionName, "NVS78P")
	//	versionName = append(versionName, "OST")
	//	versionName = append(versionName, "PDV")
	//	versionName = append(versionName, "S21")
	//	versionName = append(versionName, "TOB")
	return versionName
}

// Chargement de la version choisie
func loadVersion(name string) (Bible, error) {

	// Ouvrir le fichier JSON
	jsonFile, err := os.Open("src/" + name + ".json")
	if err != nil {
		return Bible{}, err
	}
	defer jsonFile.Close()

	// Lire le contenu du fichier
	byteValue, _ := io.ReadAll(jsonFile)

	// Un map pour stocker les données JSON de manière générique
	var result Bible

	// Désérialiser le contenu du fichier JSON dans le map
	if err := json.Unmarshal(byteValue, &result); err != nil {
		return Bible{}, err
	}

	return result, nil
}

// Chargement de l'ensemble des versions
func loadVersions(versions []string) ([]Bible, error) {
	// Stockage des bible
	var results []Bible
	for _, version := range versions {
		result, err := loadVersion(version)
		if err != nil {
			return []Bible{}, err
		}
		results = append(results, result)

	}
	return results, nil
}

// Création des dossiers
func CreateBooksDir(result Bible) error {
	// Parcours des testaments
	for _, testament := range result.Testaments {
		// Parcours des livres
		for _, book := range testament.Books {
			// Création de la structure de répertoire
			sBook := strconv.Itoa(book.ID + 1)
			bookname := aliases[book.ID].aliases[0]
			dName := dirname + "/" + sBook + " - " + bookname
			os.Mkdir(dName, 0755)

			// Création du livre
			err := CreateBookFile(dName, bookname, book.ID, len(book.Chapters))
			if err != nil {
				return err
			}

			os.Mkdir(dirintroname, 0755)
			introname := dirintroname + "/Introduction " + bookname + ".md"
			// Création de l'introduction
			err = CreateBookIntro(introname)
			if err != nil {
				return err
			}

		}
	}
	return nil
}

func GenereVersion(bible *Bible, status string) error {
	for _, testament := range bible.Testaments {
		for _, book := range testament.Books {
			for _, chapter := range book.Chapters {

				bid := strconv.Itoa(book.ID + 1)
				cid := strconv.Itoa(chapter.ID)
				bookname := aliases[book.ID].aliases[0]

				filename := dirname + "/" + bid + " - " + bookname + "/" + bookname + " " + cid + ".md"
				f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					return err
				}
				// remember to close the file
				defer f.Close()

				var toFile string = ""
				for _, verse := range chapter.Verses {
					vid := strconv.Itoa(verse.ID)

					// different version
					switch status {
					case "default":
						toFile += versesH + " " + vid + "\n"
					case "complete":
						toFile += versesH + " " + bookname + " " + cid + ":" + vid + "\n"
					default:
						toFile += versesH + " " + vid + "[" + bible.Abbreviation + "]" + "\n"

					}
					toFile += verse.Text + "\n"
				}
				toFile += "\n---\n"
				_, err = f.WriteString(toFile)
				if err != nil {
					return err
				}

			}
		}

	}
	return nil
}

// Ajout des versions dans les chapitres
func UpdateChapterFile(results *[]Bible) error {

	// Premier livre géré comme livre par défaut
	var start bool = true

	// Ajout de toutes les version
	for _, bible := range *results {
		if start {
			err := GenereVersion(&bible, "default")
			if err != nil {
				return err
			}
			err = GenereVersion(&bible, "complete")
			if err != nil {
				return err
			}
		} else {
			err := GenereVersion(&bible, "standard")
			if err != nil {
				return err
			}
		}
		start = false
	}

	return nil
}

func main() {
	// Choix des versions
	versionName := initVersionName()

	// Chargement des versions
	results, err := loadVersions(versionName)
	if err != nil {
		panic(err)
	}

	// Génération du dossier
	os.Mkdir(dirname, 0755)

	// Génération des livres
	err = CreateBooksDir(results[0])
	if err != nil {
		panic(err)
	}

	// Génération des versets
	err = UpdateChapterFile(&results)
	if err != nil {
		panic(err)
	}

}
