-- all large numbers should be at least NUMERIC(102, 2)
-- enough for 100 - 0,01 duotrigintillion

CREATE TABLE Players (
    ID SERIAL PRIMARY KEY,
    Username varchar(24) NOT NULL,
    NetWorth NUMERIC(102, 2) NOT NULL,

    UNIQUE (Username),
    CHECK (Username <> '')
);

-- //

CREATE TABLE Firms (
    ID SERIAL PRIMARY KEY,
    FirmName varchar(32) NOT NULL,
    FirmLead integer NOT NULL, -- ID of the player that created the firm

    UNIQUE (FirmName),
    UNIQUE (FirmLead), -- Player can only own one firm at once

    -- if a player deletes their account, delete their firm too
    FOREIGN KEY (FirmLead) REFERENCES Players (ID) ON DELETE CASCADE
);

CREATE TABLE Player_Firm (
    PlayerID integer,
    FirmID integer,

    UNIQUE (PlayerID), -- Player can only be in one firm at once

    FOREIGN KEY (PlayerID) REFERENCES Players (ID) ON DELETE CASCADE,
    FOREIGN KEY (FirmID) REFERENCES Firms (ID) ON DELETE CASCADE
);

-- //

CREATE TABLE Companies (
    ID SERIAL PRIMARY KEY,
    CompanyName varchar(30) NOT NULL, -- NOTE: should all be normalized to lowercase!
    StockName varchar(5) NOT NULL,

    -- prevent duplicates
    UNIQUE (CompanyName),
    UNIQUE (StockName)
);

-- history of stock prices
CREATE TABLE Company_StockPrice (
    CompanyID integer NOT NULL,
    StockPrice NUMERIC(102, 2) NOT NULL,
    Updated timestamp DEFAULT NOW(),

    -- should never be updated twice at the same time!
    UNIQUE (CompanyID, Updated),

    FOREIGN KEY (CompanyID) REFERENCES Companies (ID) ON DELETE RESTRICT,

    CHECK (StockPrice >= 0.01)
);

CREATE TABLE Company_BalanceSheet (
    CompanyID integer PRIMARY KEY,
    CurrentAssets NUMERIC(102, 2) NOT NULL,
    NonCurrentAssets NUMERIC(102, 2) NOT NULL,
    CurrentLiabilities NUMERIC(102, 2) NOT NULL,
    NonCurrentLiabilities NUMERIC(102, 2) NOT NULL,
    Equity NUMERIC(102, 2) NOT NULL,
    Reserves NUMERIC(102, 2) NOT NULL,

    FOREIGN KEY (CompanyID) REFERENCES Companies (ID) ON DELETE CASCADE
);

-- //

CREATE TABLE Banks (
    ID SERIAL PRIMARY KEY,
    BankName varchar(30) NOT NULL,
    LoanInterest NUMERIC(5,4) NOT NULL,
    AccountInterest NUMERIC(5,4) NOT NULL,
    -- surely we should never charge more than 2 billion to open an account
    AccountFee integer NOT NULL,

    UNIQUE (BankName)
);

CREATE TABLE Bank_Accounts (
    AccountNumber SERIAL PRIMARY KEY,
    BankID integer NOT NULL,
    PlayerID integer NOT NULL,
    Balance NUMERIC(102,2) DEFAULT 0,

    FOREIGN KEY (BankID) REFERENCES Banks (ID) ON DELETE CASCADE,
    FOREIGN KEY (PlayerID) REFERENCES Players (ID) ON DELETE CASCADE
);

-- how much cash reserves the bank has for loans
CREATE TABLE Bank_Reserves (
    BankID integer PRIMARY KEY,
    Balance NUMERIC (102,2) NOT NULL,

    FOREIGN KEY (BankID) REFERENCES Banks (ID) ON DELETE CASCADE
);

-- //

CREATE TABLE News (
    Template text, -- sprintf format string
    Targeted varchar(8) NOT NULL, -- the entity that is targeted
    Positive boolean -- can be null, at that point let it be random
);

-- //

-- order history of the ingame year
-- TODO: add more markets
CREATE TABLE Market (
    PlayerID integer NOT NULL,
    CompanyID integer NOT NULL,
    Amount integer DEFAULT 1,
    Price NUMERIC (102,2) NOT NULL, -- price of the stock, not the total
    TransactedAt timestamp DEFAULT NOW(),
    Bought boolean, -- false = sold, true = bought, unknown = error 

    FOREIGN KEY (PlayerID) REFERENCES Players (ID) ON DELETE RESTRICT,
    FOREIGN KEY (CompanyID) REFERENCES Companies (ID) ON DELETE RESTRICT
);

-- //

CREATE TABLE IF NOT EXISTS oauth2_clients (
    ID text PRIMARY KEY,
    Secret text NOT NULL,
    Domain text NOT NULL,
    UserID integer NOT NULL,

    CHECK (ID <> ''),
    FOREIGN KEY (UserID) REFERENCES Players(ID)
);

CREATE TABLE IF NOT EXISTS oauth2_tokens (
    ID text PRIMARY KEY,
    Code text NOT NULL,
    CreatedAt timestamp DEFAULT NOW(),
    ExpiresAt timestamp NOT NULL,

    FOREIGN KEY (ID) REFERENCES oauth2_clients (ID),
    CHECK (ExpiresAt > NOW())
)