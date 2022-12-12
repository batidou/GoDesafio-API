CREATE TABLE `goexpert`.`Cotacao` (
                                      `idCotacao` INT NOT NULL AUTO_INCREMENT,
                                      `Moeda` CHAR(7) NOT NULL,
                                      `ValorCotacao` FLOAT NOT NULL,
                                      `DataHora` DATETIME NOT NULL,
                                      PRIMARY KEY (`idCotacao`))
    COMMENT = 'Entidade responsável por armazenar as informações das cotações realizar no desafio - Client-Server-API da FullCycle - GoExpert';

