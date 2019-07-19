# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.23)
# Database: skripsi
# Generation Time: 2019-07-19 02:36:57 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table chat_logs
# ------------------------------------------------------------

DROP TABLE IF EXISTS `chat_logs`;

CREATE TABLE `chat_logs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `question` varchar(255) DEFAULT NULL,
  `answer` varchar(255) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table contexts
# ------------------------------------------------------------

DROP TABLE IF EXISTS `contexts`;

CREATE TABLE `contexts` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `client_key` varchar(255) DEFAULT NULL,
  `sentence` varchar(255) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table entities
# ------------------------------------------------------------

DROP TABLE IF EXISTS `entities`;

CREATE TABLE `entities` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table entity_datasets
# ------------------------------------------------------------

DROP TABLE IF EXISTS `entity_datasets`;

CREATE TABLE `entity_datasets` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `entity_id` int(11) unsigned DEFAULT NULL,
  `word` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `entity_dataset` (`entity_id`),
  CONSTRAINT `entity_dataset` FOREIGN KEY (`entity_id`) REFERENCES `entities` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table entity_synonyms
# ------------------------------------------------------------

DROP TABLE IF EXISTS `entity_synonyms`;

CREATE TABLE `entity_synonyms` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `entity_dataset_id` int(11) unsigned DEFAULT NULL,
  `word` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `entity_synonym_entity_dataset_id` (`entity_dataset_id`),
  CONSTRAINT `entity_synonym_entity_dataset_id` FOREIGN KEY (`entity_dataset_id`) REFERENCES `entity_datasets` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table intent_datasets
# ------------------------------------------------------------

DROP TABLE IF EXISTS `intent_datasets`;

CREATE TABLE `intent_datasets` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `intent_id` int(11) unsigned DEFAULT NULL,
  `sentence` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `intent_datasets` (`intent_id`),
  CONSTRAINT `intent_datasets` FOREIGN KEY (`intent_id`) REFERENCES `intents` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `intent_datasets` WRITE;
/*!40000 ALTER TABLE `intent_datasets` DISABLE KEYS */;

INSERT INTO `intent_datasets` (`id`, `intent_id`, `sentence`)
VALUES
	(1,1,'sisa bayaran saya'),
	(4,1,'info pembayaran uang kuliah saya'),
	(5,2,'berapa nomer briva saya '),
	(6,3,'lihat transkip nilai'),
	(7,4,'siapa pembimbing akademik saya'),
	(8,4,'dosen pa saya'),
	(9,5,'lihat krs saya'),
	(10,6,'lihat khs saya');

/*!40000 ALTER TABLE `intent_datasets` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table intent_responses
# ------------------------------------------------------------

DROP TABLE IF EXISTS `intent_responses`;

CREATE TABLE `intent_responses` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `intent_id` int(11) unsigned DEFAULT NULL,
  `sentence` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `intent_responses_intent_id` (`intent_id`),
  CONSTRAINT `intent_responses_intent_id` FOREIGN KEY (`intent_id`) REFERENCES `intents` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table intents
# ------------------------------------------------------------

DROP TABLE IF EXISTS `intents`;

CREATE TABLE `intents` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `category` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `intents` WRITE;
/*!40000 ALTER TABLE `intents` DISABLE KEYS */;

INSERT INTO `intents` (`id`, `name`, `category`)
VALUES
	(1,'cek pembayaran',1),
	(2,'tanya no briva',2),
	(3,'transkip nilai',3),
	(4,'dosen pembimbing akademik',4),
	(5,'krs',5),
	(6,'khs',6);

/*!40000 ALTER TABLE `intents` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
