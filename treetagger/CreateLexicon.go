package treetagger

import (
	"../utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/fatih/set.v0"
	"log"
	"os"
)

func CreateLexicon(db *sql.DB, lexiconFileName string) {
	// TODO: DISTINCT, ORDER BY
	rows, err := db.Query("SELECT word,pos_ud17,base_form FROM words_pos_base")
	//rows, err := db.Query("SELECT word,pos_ud17,base_form FROM words_pos_base LIMIT 10000")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var lexiconRow [3]string

	// remove file if existing
	os.Remove(lexiconFileName)
	lexiconFile, _ := os.OpenFile(lexiconFileName,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		0660)
	defer lexiconFile.Close()

	firstColumnSet := set.New()
	tagSetAll := set.New()
	tagSetAll.Add("ADJ", "ADP", "ADV", "AUX", "CONJ", "DET", "INTJ", "NOUN", "NUM", "PART", "PRON", "PROPN", "PUNCT", "SCONJ", "SYM", "VERB", "X")
	tagSetLexicon := set.New()

	if err != nil {
		log.Fatal(err)
	}

	// train-tree-tagger can only handle ~3.3 million lines, otherwise
	// there is a segmentation fault
	rowCounter := 3300000
	for rows.Next() {
		rows.Scan(&lexiconRow[0], &lexiconRow[1], &lexiconRow[2])

		// skip when it's a cardinal or there is no ud17_tag
		if (lexiconRow[1] == "NUM") || (lexiconRow[1] == "") {
			continue
		}

		// skip when word exists already in first column
		if firstColumnSet.Has(lexiconRow[0]) {
			continue
		} else {
			firstColumnSet.Add(lexiconRow[0])
		}

		// replace "<unknown>" with "-"
		if lexiconRow[2] == "<unknown>" {
			lexiconRow[2] = "-"
		}

		preparedRow := lexiconRow[0] + "\t" +
			lexiconRow[1] + "\t" +
			lexiconRow[2] + "\n"

		_, err = lexiconFile.Write([]byte(preparedRow))
		if err != nil {
			log.Fatal(err)
		}

		tagSetLexicon.Add(lexiconRow[1])
		rowCounter--
		if rowCounter == 0 {
			break
		}
	}

	//check which of the tags aren't used in database
	tagSetDiff := set.Difference(tagSetAll, tagSetLexicon)
	tagSetDiffList := set.StringSlice(tagSetDiff)

	for i := range tagSetDiffList {
		tag := string(tagSetDiffList[i])
		dummyEntry := utils.ReturnStringMd5Hash(tag) + "\t" + tag + "\t-\n"
		_, err = lexiconFile.Write([]byte(dummyEntry))
		if err != nil {
			log.Fatal(err)
		}

	}
	// nobody knows why that has to be in there…
	dummyEntrySENT := utils.ReturnStringMd5Hash("SENT") + "\tSENT\t-\n"
	_, err = lexiconFile.Write([]byte(dummyEntrySENT))
	if err != nil {
		log.Fatal(err)
	}

}
