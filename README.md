# pos-wrapper

This is a wrapper for different POS taggers.

## input format

This tools reads from a database:

```
mysql> DESCRIBE sentences_tagged;
+----------+------------------+------+-----+---------+-------+
| Field    | Type             | Null | Key | Default | Extra |
+----------+------------------+------+-----+---------+-------+
| s_id     | int(10) unsigned | NO   | MUL | NULL    |       |
| sentence | text             | YES  |     | NULL    |       |
+----------+------------------+------+-----+---------+-------+

mysql> DESCRIBE words_pos_base;
+-----------+------------------+------+-----+---------+-------+
| Field     | Type             | Null | Key | Default | Extra |
+-----------+------------------+------+-----+---------+-------+
| w_id      | int(10) unsigned | NO   | MUL | NULL    |       |
| word      | varchar(256)     | YES  | MUL | NULL    |       |
| pos       | varchar(64)      | YES  |     | NULL    |       |
| base_form | varchar(256)     | YES  |     | NULL    |       |
| freq      | int(8) unsigned  | YES  |     | NULL    |       |
| pos_ud17  | varchar(64)      | YES  |     |         |       |
+-----------+------------------+------+-----+---------+-------+
```

## implemented taggers

* [TreeTagger](http://www.cis.uni-muenchen.de/~schmid/tools/TreeTagger/)
* [RDRPOSTagger](http://rdrpostagger.sourceforge.net) (there are some things to do)


## usage

```
./pos-wrapper -tagger the_tagger_of_your_choice -db-name database_name
```

## options

* -h, -help: display help and exit
* -tagger: the tagger that should be used
* -db-name: the database that should be used

## be careful

If the TreeTagger lexicon file is longer than ~3mio lines, there is a segmentation fault. Also TreeTagger fails, if the training corpus is too large. ("ERROR: Can't open for reading")
