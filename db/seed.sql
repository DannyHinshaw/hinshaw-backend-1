DROP TABLE IF EXISTS app_users, customers, credit_scores CASCADE;

CREATE TABLE app_users
(
    id         VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    email      VARCHAR(255),
    password   VARCHAR(255)
);

CREATE TABLE customers
(
    id         VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    address    VARCHAR(255),
    last_name  VARCHAR(35),
    first_name VARCHAR(35)
);

CREATE TABLE credit_scores
(
    id          VARCHAR PRIMARY KEY,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    equifax     SMALLINT,
    experian    SMALLINT,
    trans_union SMALLINT,
    CONSTRAINT fk_customer
        FOREIGN KEY (id)
            REFERENCES customers (id)
            ON DELETE CASCADE
);

INSERT INTO customers(id, address, last_name, first_name)
VALUES ('9f1aff3c-e742-49b1-acc2-d60394242c74', '2339 Hood Avenue, New Orleans, LA 70190', 'Rios', 'Kimberley'),
       ('96e3077d-6048-463a-890e-ce955781b0db', '2908  Hill Haven Drive, Waco, TX 76701', 'Robinson', 'Yee'),
       ('e9f57297-1d5f-4753-a710-5badf3b86b64', '629  Brooke Street, Sugar Land, TX 77478', 'Lainez', 'Charles'),
       ('9f0dd935-6bb4-4e84-ae57-244c731d5569', '2158  Corbin Branch Road, Knoxville TN 37917', 'Cannon', 'James');

INSERT INTO credit_scores(id, equifax, experian, trans_union)
VALUES ('9f1aff3c-e742-49b1-acc2-d60394242c74', 500, 550, 525),
       ('96e3077d-6048-463a-890e-ce955781b0db', 631, 617, 620),
       ('e9f57297-1d5f-4753-a710-5badf3b86b64', 700, 770, 725),
       ('9f0dd935-6bb4-4e84-ae57-244c731d5569', 600, 615, 630);
