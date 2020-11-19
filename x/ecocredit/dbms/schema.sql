
CREATE TABLE accounts (
	address     varchar(255) PRIMARY KEY
);

CREATE TABLE credit_classes (
	class_id    varchar(255)     PRIMARY KEY,
	designer    varchar(255)     REFERENCES accounts ,
	issuers     varchar(255)[]   NOT NULL,
	metadata    bytea            NOT NULL,
	name        varchar(255)     NOT NULL DEFAULT '',
	typ         int2             NOT NULL
);

CREATE TABLE credit_batches (
	batch_denom  varchar(255)   PRIMARY KEY,
	class_id     varchar(255)   REFERENCES credit_classes,
	issuer       varchar(255)   REFERENCES accounts,
	total_units  decimal        NOT NULL,
	metadata     bytea          NOT NULL,
	start_date   timestamp      NOT NULL,
	end_date     timestamp      NOT NULL,
	point        point,
	polygon      polygon
);
