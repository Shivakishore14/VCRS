-- MySQL dump 10.13  Distrib 5.5.50, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: VCRS
-- ------------------------------------------------------
-- Server version	5.5.50-0+deb8u1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `chemistry`
--

DROP TABLE IF EXISTS `chemistry`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `chemistry` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class` text,
  `ques` text,
  `option1` text,
  `option2` text,
  `option3` text,
  `option4` text,
  `ans` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `chemistry`
--

LOCK TABLES `chemistry` WRITE;
/*!40000 ALTER TABLE `chemistry` DISABLE KEYS */;
/*!40000 ALTER TABLE `chemistry` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `chemistryresponse`
--

DROP TABLE IF EXISTS `chemistryresponse`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `chemistryresponse` (
  `sid` varchar(10) DEFAULT NULL,
  `qno` int(11) DEFAULT NULL,
  `ans` varchar(30) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `chemistryresponse`
--

LOCK TABLES `chemistryresponse` WRITE;
/*!40000 ALTER TABLE `chemistryresponse` DISABLE KEYS */;
INSERT INTO `chemistryresponse` VALUES ('CRS-1234',1,'option3');
/*!40000 ALTER TABLE `chemistryresponse` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `physics`
--

DROP TABLE IF EXISTS `physics`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `physics` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class` text,
  `ques` text,
  `option1` text,
  `option2` text,
  `option3` text,
  `option4` text,
  `ans` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `physics`
--

LOCK TABLES `physics` WRITE;
/*!40000 ALTER TABLE `physics` DISABLE KEYS */;
INSERT INTO `physics` VALUES (1,'','odd one out','sun','atom','nucleous','electron','option1'),(2,'','which is more elastic','wood','rubberBand','water','steel rod','option4'),(3,'','current is proportional to','resistance','volts','heat','none','option2'),(4,'','which is most colorful','/images/physics4option1','/images/physics4option2','/images/physics4option3','/images/physics4option4','option1');
/*!40000 ALTER TABLE `physics` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `physicsresponse`
--

DROP TABLE IF EXISTS `physicsresponse`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `physicsresponse` (
  `sid` varchar(10) DEFAULT NULL,
  `qno` int(11) DEFAULT NULL,
  `ans` varchar(30) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `physicsresponse`
--

LOCK TABLES `physicsresponse` WRITE;
/*!40000 ALTER TABLE `physicsresponse` DISABLE KEYS */;
INSERT INTO `physicsresponse` VALUES ('CRS-1234',1,'option2'),('CRS-1234',2,'option2'),('CRS-1234',3,'option2');
/*!40000 ALTER TABLE `physicsresponse` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sampleTest1`
--

DROP TABLE IF EXISTS `sampleTest1`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sampleTest1` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class` text,
  `ques` text,
  `option1` text,
  `option2` text,
  `option3` text,
  `option4` text,
  `ans` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sampleTest1`
--

LOCK TABLES `sampleTest1` WRITE;
/*!40000 ALTER TABLE `sampleTest1` DISABLE KEYS */;
INSERT INTO `sampleTest1` VALUES (1,'class1','what comes after 1','2','1','3','-1','option1'),(2,'class1','what comes after 2','0','2','3','-1','option3'),(3,'class1','what is 1','odd','even','neither odd nor eveb','none of the above','option1'),(4,'','sample q uews','qqqq','wwwww','eeeee','rrrrrr','option2');
/*!40000 ALTER TABLE `sampleTest1` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sampleTest1response`
--

DROP TABLE IF EXISTS `sampleTest1response`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sampleTest1response` (
  `sid` varchar(10) DEFAULT NULL,
  `qno` int(11) DEFAULT NULL,
  `ans` varchar(30) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sampleTest1response`
--

LOCK TABLES `sampleTest1response` WRITE;
/*!40000 ALTER TABLE `sampleTest1response` DISABLE KEYS */;
INSERT INTO `sampleTest1response` VALUES ('2',1,'option1'),('2',2,'option2'),('2',3,'option3'),('2',4,'option4'),('1',1,'option1'),('1',2,'option2'),('1',3,'option3'),('1',4,'option4'),('3',1,'option1'),('3',2,'option2'),('3',3,'option3'),('3',4,'option4'),('4',1,'option3'),('4',2,'option2'),('4',3,'option3'),('4',4,'option4'),('5',1,'option3'),('5',2,'option1'),('5',3,'option2'),('5',4,'option4'),('6',1,'option3'),('6',2,'option1'),('6',3,'option1'),('6',4,'option4');
/*!40000 ALTER TABLE `sampleTest1response` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `social`
--

DROP TABLE IF EXISTS `social`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `social` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class` text,
  `ques` text,
  `option1` text,
  `option2` text,
  `option3` text,
  `option4` text,
  `ans` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `social`
--

LOCK TABLES `social` WRITE;
/*!40000 ALTER TABLE `social` DISABLE KEYS */;
/*!40000 ALTER TABLE `social` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `socialresponse`
--

DROP TABLE IF EXISTS `socialresponse`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `socialresponse` (
  `sid` varchar(10) DEFAULT NULL,
  `qno` int(11) DEFAULT NULL,
  `ans` varchar(30) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `socialresponse`
--

LOCK TABLES `socialresponse` WRITE;
/*!40000 ALTER TABLE `socialresponse` DISABLE KEYS */;
INSERT INTO `socialresponse` VALUES ('CRS-1234',1,'option2'),('CRS-1234',2,'option4'),('CRS-1234',3,'option2'),('CRS-1234',4,'option4');
/*!40000 ALTER TABLE `socialresponse` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `staff`
--

DROP TABLE IF EXISTS `staff`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `staff` (
  `username` text,
  `password` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `staff`
--

LOCK TABLES `staff` WRITE;
/*!40000 ALTER TABLE `staff` DISABLE KEYS */;
INSERT INTO `staff` VALUES ('shiva','shiva'),('shiva','boss');
/*!40000 ALTER TABLE `staff` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stuLogin`
--

DROP TABLE IF EXISTS `stuLogin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `stuLogin` (
  `userId` varchar(20) DEFAULT NULL,
  `password` varchar(20) DEFAULT NULL,
  `name` varchar(30) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stuLogin`
--

LOCK TABLES `stuLogin` WRITE;
/*!40000 ALTER TABLE `stuLogin` DISABLE KEYS */;
INSERT INTO `stuLogin` VALUES ('CRS-1234','rhodio','Sam Anderson');
/*!40000 ALTER TABLE `stuLogin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `testDetails`
--

DROP TABLE IF EXISTS `testDetails`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `testDetails` (
  `testName` text,
  `noOfQues` text,
  `class` text,
  `status` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `testDetails`
--

LOCK TABLES `testDetails` WRITE;
/*!40000 ALTER TABLE `testDetails` DISABLE KEYS */;
INSERT INTO `testDetails` VALUES ('sampleTest1','4',NULL,NULL),('physics','4','',''),('chemistry','4','',''),('social','4','','');
/*!40000 ALTER TABLE `testDetails` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-09-29 17:53:46
