----https://www.programiz.com/sql/online-compiler/
CREATE DATABASE COVID;
CREATE table COVID.PATIENT_COVID_STATUS(
   HNID VARCHAR(80),
   COVID_STATUS VARCHAR(80),
  CONSTRAINT HNID PRIMARY KEY (HNID)
  
);

INSERT INTO PATIENT_COVID_STATUS(HNID, COVID_STATUS) VALUES
('0001', 'Positive'),
('0002', 'Positive'),
('0003', 'Negative'),
('0004', 'Positive'),
('0005', 'Negative');
-------------------foereign 1
CREATE table PATIENT(
   HNID VARCHAR(80),
   Firstname VARCHAR(255),
   Lastname VARCHAR(255),
   HID VARCHAR(80) REFERENCES HOSPITAL(HID),
  CONSTRAINT HNID PRIMARY KEY (HNID)
);


-------------------foereign 2
CREATE table PATIENT(
   HNID VARCHAR(80),
   Firstname VARCHAR(255),
   Lastname VARCHAR(255),
  HID VARCHAR(80),
  CONSTRAINT HNID PRIMARY KEY (HNID)
);


ALTER TABLE PATIENT
ADD FOREIGN KEY (HID) REFERENCES HOSPITAL(HID);

INSERT INTO PATIENT(HNID, Firstname, Lastname, HID) VALUES
('0001', 'Ayesha', 'Jenkins', '0001'),
('0002',  'Savannah', 'Mcintosh', '0002'),
('0003',  'Harley', 'Greene', '0002'),
('0004',  'Tilly', 'Frazier', '0001'),
('0005',  'Gloria', 'Hull', '0003');

CREATE table HOSPITAL(
 HID VARCHAR(80),
 Title VARCHAR(255)
);

INSERT INTO HOSPITAL(HID, Title) VALUES
('0001', 'St. Violet Hospital'),
('0002', 'The Unity Health Center'),
('0003', 'John Medical Center');

