use WebHosting;

-- WebHosting.Admins definition

CREATE TABLE `Admins` (
  `ID` int NOT NULL,
  `UserName` varchar(100) NOT NULL,
  `Password` varchar(100) NOT NULL,
  PRIMARY KEY (`ID`)
) ;

-- WebHosting.Plans definition

CREATE TABLE `Plans` (
  `ID` int NOT NULL,
  `Name` varchar(255) DEFAULT NULL,
  `Price` float DEFAULT NULL,
  `Details` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ;

-- WebHosting.Users definition

CREATE TABLE `Users` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Nombre` varchar(100) NOT NULL,
  `Apellidos` varchar(100) NOT NULL,
  `UserName` varchar(100) NOT NULL,
  `Email` varchar(100) NOT NULL,
  `Password` varchar(100) NOT NULL,
  PRIMARY KEY (`ID`)
) ;

-- WebHosting.Contracts definition

CREATE TABLE `Contracts` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `PlanID` int NOT NULL,
  `UserID` int NOT NULL,
  `DateOfContract` date NOT NULL,
  `DateOfExpiration` date NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `Contracts_FK` (`PlanID`),
  KEY `Contracts_FK_1` (`UserID`),
  CONSTRAINT `Contracts_FK` FOREIGN KEY (`PlanID`) REFERENCES `Plans` (`ID`),
  CONSTRAINT `Contracts_FK_1` FOREIGN KEY (`UserID`) REFERENCES `Users` (`ID`)
) ;



-- WebHosting.InSuggestions definition

CREATE TABLE `InSuggestions` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `UserID` int NOT NULL,
  `Suggestion` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `InSuggestions_FK` (`UserID`),
  CONSTRAINT `InSuggestions_FK` FOREIGN KEY (`UserID`) REFERENCES `Users` (`ID`)
) ;


CREATE TABLE `ExSuggestions` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Email` varchar(100) NOT NULL,
  `Name` varchar(100) NOT NULL,
  `Suggestion` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`ID`)
);

INSERT INTO Admins (ID, UserName, Password) VALUES(1, 'Admin1', '987654321');

INSERT INTO Plans (ID, Name, Price, Details) VALUES(1, "Starter", 250.0, "4 features");
INSERT INTO Plans (ID, Name, Price, Details) VALUES(2, "Basico", 455.0, "6 features");
INSERT INTO Plans (ID, Name, Price, Details) VALUES(3, "No limite", 799.0, "8 features");
INSERT INTO Plans (ID, Name, Price, Details) VALUES(4, "Avanzado", 1039.0, "9 features");