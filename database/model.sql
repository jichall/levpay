CREATE TABLE IF NOT EXISTS TB_SUPER(
    SUPER_ID SERIAL PRIMARY KEY,
    SUPER_UUID VARCHAR(128),
    SUPER_NAME VARCHAR(30),
    SUPER_FULLNAME VARCHAR(100),
    SUPER_INTELLIGENCE VARCHAR(100),
    SUPER_POWER VARCHAR(100),
    SUPER_OCCUPATION VARCHAR(100),
    SUPER_IMAGE VARCHAR(2048)
);