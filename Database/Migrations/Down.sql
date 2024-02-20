-- NOTE: always drop the tables referencing others first!
DROP TABLE IF EXISTS Company_StockPrice;
DROP TABLE IF EXISTS Company_BalanceSheet;

DROP TABLE IF EXISTS Player_Firm;

DROP TABLE IF EXISTS Bank_Accounts;
DROP TABLE IF EXISTS Bank_Reserves;

-- entities
DROP TABLE IF EXISTS News;
DROP TABLE IF EXISTS Market;
DROP TABLE IF EXISTS Companies;
DROP TABLE IF EXISTS Banks;
DROP TABLE IF EXISTS Players;
DROP TABLE IF EXISTS Firms;

DROP TABLE IF EXISTS oauth2_clients;
DROP TABLE IF EXISTS oauth2_tokens;