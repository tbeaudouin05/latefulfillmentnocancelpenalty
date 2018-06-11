
-- stored view for user purpose

CREATE TABLE baa_application.vendor_service.lfnc_penalty_exception (
  uid_lfnc_penalty_exception INT IDENTITY(1,1) UNIQUE
  ,id_lfnc_penalty_exception VARCHAR(20) PRIMARY KEY
  ,timestamp DATETIME NOT NULL DEFAULT (GETDATE())
  ,initiator  VARCHAR(50)
  ,seller_name  NVARCHAR(200)
  ,id_seller INT
  ,start_date  DATE
  ,end_date  DATE
  --,order_nr  BIGINT
  ,bob_id_sales_order_item BIGINT
  ,amount REAL
  ,reason NVARCHAR(600) NOT NULL
  ,lfnc_penalty_exception_status VARCHAR(20)
  ,fk_lfnc_penalty_exception_status INT
);

-- tables used for backend calculation purpose

-- to add exceptional date
CREATE TABLE baa_application.vendor_service.lfnc_penalty_date_exception (
  id_lfnc_penalty_date_exception INT IDENTITY(1,1) PRIMARY KEY -- cannot make natural unique key due to start / end dates
  ,timestamp DATETIME NOT NULL DEFAULT (GETDATE())
  ,start_date  DATE NOT NULL
  ,end_date  DATE NOT NULL
  ,fk_lfnc_penalty_exception_status INT
  ,fk_lfnc_penalty_exception VARCHAR(20) FOREIGN KEY REFERENCES baa_application.vendor_service.lfnc_penalty_exception(id_lfnc_penalty_exception)
  
);

-- to add exceptional seller-date
CREATE TABLE baa_application.vendor_service.lfnc_penalty_seller_date_exception (
  id_lfnc_penalty_seller_date_exception INT IDENTITY(1,1) PRIMARY KEY -- cannot make natural unique key due to start / end dates
  ,timestamp DATETIME NOT NULL DEFAULT (GETDATE())
  ,seller_name NVARCHAR(200) NOT NULL
  ,id_seller INT NOT NULL
  ,start_date  DATE NOT NULL
  ,end_date  DATE NOT NULL
  ,fk_lfnc_penalty_exception_status INT
  ,fk_lfnc_penalty_exception VARCHAR(20) FOREIGN KEY REFERENCES baa_application.vendor_service.lfnc_penalty_exception(id_lfnc_penalty_exception)
  
);

-- to add exceptional order_nr
CREATE TABLE baa_application.vendor_service.lfnc_penalty_order_exception (
  timestamp DATETIME NOT NULL DEFAULT (GETDATE())
  ,order_nr  BIGINT PRIMARY KEY
  ,fk_lfnc_penalty_exception_status INT
  ,fk_lfnc_penalty_exception VARCHAR(20) FOREIGN KEY REFERENCES baa_application.vendor_service.lfnc_penalty_exception(id_lfnc_penalty_exception)
  
);

-- to add exceptional order bob_id_sales_order_item
CREATE TABLE baa_application.vendor_service.lfnc_penalty_soi_exception (
  timestamp DATETIME NOT NULL DEFAULT (GETDATE())
  ,bob_id_sales_order_item BIGINT PRIMARY KEY
  ,fk_lfnc_penalty_exception_status INT
  ,fk_lfnc_penalty_exception VARCHAR(20) FOREIGN KEY REFERENCES baa_application.vendor_service.lfnc_penalty_exception(id_lfnc_penalty_exception)
  
);


CREATE TABLE baa_application.vendor_service.user_access (
  email VARCHAR(50) PRIMARY KEY
  ,name VARCHAR(50) NOT NULL
  ,access VARCHAR(50) NOT NULL
);

CREATE TABLE baa_application.vendor_service.user_access (
  email VARCHAR(50) PRIMARY KEY
  ,name VARCHAR(50) NOT NULL
  ,access VARCHAR(50) NOT NULL
);



INSERT INTO baa_application.vendor_service.user_access (
email
,name
,access
  )
VALUES ('foruzan.moghadam@bamilo.com','Foruzan Moghadam','admin')

	UPDATE baa_application.vendor_service.user_access
	SET baa_application.vendor_service.user_access.access = 'lfnc_user' 
	WHERE baa_application.vendor_service.user_access.email = 'mohammad.goudarzi@bamilo.com'


DELETE FROM baa_application.vendor_service.lfnc_penalty_exception;
DELETE FROM baa_application.vendor_service.lfnc_penalty_date_exception;

DROP TABLE baa_application.vendor_service.lfnc_penalty_exception;
DROP TABLE baa_application.vendor_service.lfnc_penalty_date_exception;
DROP TABLE baa_application.vendor_service.lfnc_penalty_seller_date_exception;
DROP TABLE baa_application.vendor_service.lfnc_penalty_order_exception;
DROP TABLE baa_application.vendor_service.lfnc_penalty_soi_exception;
