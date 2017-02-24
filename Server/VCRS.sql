-- MySQL dump 10.13  Distrib 5.7.17, for Linux (x86_64)
--
-- Host: localhost    Database: VCRS
-- ------------------------------------------------------
-- Server version	5.7.17-0ubuntu0.16.04.1

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
-- Table structure for table `MutiDimentionalArrays`
--

DROP TABLE IF EXISTS `MutiDimentionalArrays`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `MutiDimentionalArrays` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class` text,
  `ques` text,
  `option1` text,
  `option2` text,
  `option3` text,
  `option4` text,
  `ans` text,
  `level` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `MutiDimentionalArrays`
--

LOCK TABLES `MutiDimentionalArrays` WRITE;
/*!40000 ALTER TABLE `MutiDimentionalArrays` DISABLE KEYS */;
INSERT INTO `MutiDimentionalArrays` VALUES (1,'','An array has the below mentioned properties','datatype','rank','size ','all of thhe above','option4',''),(2,'','In C language, elements of two-dimensional arrays are stored in memory as','Random order','Row major order','Linear order','None of these','option3',''),(3,'','3. What is the output of this C code?\n1. #include <stdio.h>\n2. int main()\n3. {\n4. int ary[2][3];\n5. ary[][]={{1,2,3},{4,5,6}};\n6. printf(\"%d\\n\",ary[1][0]);\n7. }','Compile time error','4','1','2','option2',''),(4,'','4. What is the output of this C code?\n1. #include <stdio.h>\n2. void main()    {\n3. int a[2][3]={1,2,3,4,5};\n4. int i =0, j =0;\n5. for(i =0; i <2; i++)\n6. for(j =0; j <3; j++)\n7. printf(\"%d\", a[i][j]);     }','1 2 3 4 5 0','1 2 3 4 5 junk','1 2 3 4 5 5','Run time error','option3',''),(5,'','5. A multidimensional array can be expressed in terms of \na) \nb) c)  \nd) ','Array of pointers rather than as pointers to a group of contiguous array ','Array without the group of contiguous array ','Data type arrays ','None of these','option4',''),(6,'','6. What is the output of this C code?\n1. #include <stdio.h>\n2. void main()\n3. {\n4. int a[2][3]={1,2,3,,4,5};\n5. int i =0, j =0;\n6. for(i =0; i <2; i++)\n7. for(j =0; j <3; j++)\n8. printf(\"%d\", a[i][j]);\n9. }','1 2 3 junk 4 5','Compile time error','1 2 3 0 4 5','1 2 3 3 4 5','option2',''),(7,'','C allows arrays of greater than two dimensions, who will determined this ','programmer','Compiler','Parameter','None of These','option4',''),(8,'','Which of the following correctly declares an array?','int anarray[10];','nt anarray;','anarray{10};','array anarray[10];','option1',''),(9,'','What is the index number of the last element of an array with 29 elements?','29','28','0','Programmer-defined','option2',''),(10,'','10. Which of the following is a two-dimensional array?',' array anarray[20][20];','int anarray[20][20];','int array[20, 20];','char array[20];','option2','');
/*!40000 ALTER TABLE `MutiDimentionalArrays` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `MutiDimentionalArraysresponse`
--

DROP TABLE IF EXISTS `MutiDimentionalArraysresponse`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `MutiDimentionalArraysresponse` (
  `sid` varchar(10) DEFAULT NULL,
  `qno` int(11) DEFAULT NULL,
  `ans` varchar(30) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `MutiDimentionalArraysresponse`
--

LOCK TABLES `MutiDimentionalArraysresponse` WRITE;
/*!40000 ALTER TABLE `MutiDimentionalArraysresponse` DISABLE KEYS */;
INSERT INTO `MutiDimentionalArraysresponse` VALUES ('CRS-1234',1,'option1'),('CRS-1234',2,'option2'),('CRS-1234',3,'option3'),('CRS-1234',4,'option3'),('CRS-1234',5,'option2'),('CRS-1234',6,'option2'),('CRS-1234',7,'option3'),('CRS-1234',8,'option2'),('CRS-1234',9,'option1'),('CRS-1234',10,'option2');
/*!40000 ALTER TABLE `MutiDimentionalArraysresponse` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin` (
  `username` text,
  `password` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES ('sk','kishore');
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `newTest2`
--

DROP TABLE IF EXISTS `newTest2`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `newTest2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class` text,
  `ques` text,
  `option1` text,
  `option2` text,
  `option3` text,
  `option4` text,
  `ans` text,
  `level` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `newTest2`
--

LOCK TABLES `newTest2` WRITE;
/*!40000 ALTER TABLE `newTest2` DISABLE KEYS */;
/*!40000 ALTER TABLE `newTest2` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `newTest2response`
--

DROP TABLE IF EXISTS `newTest2response`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `newTest2response` (
  `sid` varchar(10) DEFAULT NULL,
  `qno` int(11) DEFAULT NULL,
  `ans` varchar(30) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `newTest2response`
--

LOCK TABLES `newTest2response` WRITE;
/*!40000 ALTER TABLE `newTest2response` DISABLE KEYS */;
/*!40000 ALTER TABLE `newTest2response` ENABLE KEYS */;
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
  `level` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sampleTest1`
--

LOCK TABLES `sampleTest1` WRITE;
/*!40000 ALTER TABLE `sampleTest1` DISABLE KEYS */;
INSERT INTO `sampleTest1` VALUES (1,'class1','what comes after 1','2','1','3','-1','option1','1'),(2,'class1','what comes after 2','0','2','3','-1','option3','1'),(3,'class1','what is 1','odd','even','neither odd nor eveb','none of the above','option1','1'),(4,'','sample q uews','qqqq','wwwww','eeeee','rrrrrr','option2','1');
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
INSERT INTO `sampleTest1response` VALUES ('2',1,'option2'),('2',2,'option2'),('2',3,'option3'),('2',4,'option4'),('1',1,'option1'),('1',2,'option2'),('1',3,'option3'),('1',4,'option4'),('3',1,'option1'),('3',2,'option2'),('3',3,'option3'),('3',4,'option4'),('4',1,'option3'),('4',2,'option2'),('4',3,'option3'),('4',4,'option4'),('5',1,'option3'),('5',2,'option1'),('5',3,'option2'),('5',4,'option4'),('6',1,'option3'),('6',2,'option1'),('6',3,'option1'),('6',4,'option4'),('CRS-1234',1,'option2'),('CRS-1234',2,'option4'),('CRS-1234',3,'option2'),('CRS-1234',4,'option4');
/*!40000 ALTER TABLE `sampleTest1response` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sk`
--

DROP TABLE IF EXISTS `sk`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sk` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class` text,
  `ques` text,
  `option1` text,
  `option2` text,
  `option3` text,
  `option4` text,
  `ans` text,
  `level` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sk`
--

LOCK TABLES `sk` WRITE;
/*!40000 ALTER TABLE `sk` DISABLE KEYS */;
INSERT INTO `sk` VALUES (1,'','j','jkkj','1jkhkj','hkj','hkjh','option4','2'),(2,'','lj','lkjl','kjlj','lkj','ljl','option4','3');
/*!40000 ALTER TABLE `sk` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `skresponse`
--

DROP TABLE IF EXISTS `skresponse`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `skresponse` (
  `sid` varchar(10) DEFAULT NULL,
  `qno` int(11) DEFAULT NULL,
  `ans` varchar(30) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `skresponse`
--

LOCK TABLES `skresponse` WRITE;
/*!40000 ALTER TABLE `skresponse` DISABLE KEYS */;
/*!40000 ALTER TABLE `skresponse` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `staff`
--

DROP TABLE IF EXISTS `staff`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `staff` (
  `username` text,
  `password` text,
  `name` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `staff`
--

LOCK TABLES `staff` WRITE;
/*!40000 ALTER TABLE `staff` DISABLE KEYS */;
INSERT INTO `staff` VALUES ('kishore','kishore','NAME2'),('sk','new user','NAME2');
/*!40000 ALTER TABLE `staff` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stuLogin`
--

DROP TABLE IF EXISTS `stuLogin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `stuLogin` (
  `username` varchar(20) DEFAULT NULL,
  `password` varchar(20) DEFAULT NULL,
  `name` varchar(30) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stuLogin`
--

LOCK TABLES `stuLogin` WRITE;
/*!40000 ALTER TABLE `stuLogin` DISABLE KEYS */;
INSERT INTO `stuLogin` VALUES ('CRS-1234','rhodio','Sam Anderson'),('shiva','kishore','shiva s'),('kishiore','kishire','kishore');
/*!40000 ALTER TABLE `stuLogin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `testDetails`
--

DROP TABLE IF EXISTS `testDetails`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `testDetails` (
  `testName` varchar(100) NOT NULL,
  `noOfQues` text,
  `status` text,
  `level` text,
  PRIMARY KEY (`testName`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `testDetails`
--

LOCK TABLES `testDetails` WRITE;
/*!40000 ALTER TABLE `testDetails` DISABLE KEYS */;
INSERT INTO `testDetails` VALUES ('MutiDimentionalArrays','10','ONLINE','3'),('newTest2','10','','2'),('sampleTest1','4','OFFLINE','4'),('sk','10','ONLINE','3');
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

-- Dump completed on 2017-02-23 23:19:57
