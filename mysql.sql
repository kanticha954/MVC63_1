CREATE DATABASE COVID;
CREATE table COVID.PATIENT_COVID_STATUS(
   HNID VARCHAR(80),
   COVID_STATUS VARCHAR(80),
  CONSTRAINT HNID PRIMARY KEY (HNID)
  
);

INSERT INTO COVID.PATIENT_COVID_STATUS(HNID, COVID_STATUS) VALUES
('0001', 'Positive'),
('0002', 'Positive'),
('0003', 'Negative'),
('0004', 'Positive'),
('0005', 'Negative');

CREATE table COVID.HOSPITAL(
 HID VARCHAR(80),
 Title VARCHAR(255)
);

INSERT INTO COVID.HOSPITAL(HID, Title) VALUES
('0001', 'St. Violet Hospital'),
('0002', 'The Unity Health Center'),
('0003', 'John Medical Center');

CREATE table COVID.PATIENT(
   hnid VARCHAR(80),
   firstname VARCHAR(80),
   lastname VARCHAR(80),
   hid VARCHAR(80) REFERENCES HOSPITAL(hid),
  CONSTRAINT hnid PRIMARY KEY (hnid)
);

INSERT INTO COVID.PATIENT(hnid, firstname, lastname, HID) VALUES
('0001', 'Ayesha', 'Jenkins', '0001'),
('0002',  'Savannah', 'Mcintosh', '0002'),
('0003',  'Harley', 'Greene', '0002'),
('0004',  'Tilly', 'Frazier', '0001'),
('0005',  'Gloria', 'Hull', '0003');


INSERT INTO COVID.PATIENT_COVID_STATUS(HNID, COVID_STATUS) VALUES
('0006', 'Positive'),
('0007', 'Positive'),
('0008', 'Negative'),
('0009', 'Positive'),
('00010', 'Negative');

INSERT INTO COVID.PATIENT(hnid, firstname, lastname, HID) VALUES
('0006', 'Ayesha', 'Jenkins', '0003'),
('0007',  'Savannah', 'Mcintosh', '0003'),
('0008',  'Harley', 'Greene', '0002'),
('0009',  'Tilly', 'Frazier', '0001'),
('0010',  'Gloria', 'Hull', '0003');

------find patient covid positive
SELECT PATIENT.hnid, firstname, lastname, hid, PATIENT_COVID_STATUS.covid_status
FROM COVID.PATIENT_COVID_STATUS CROSS JOIN COVID.PATIENT
WHERE PATIENT_COVID_STATUS.covid_status = 'Positive'
AND COVID.PATIENT.hnid = COVID.PATIENT_COVID_STATUS.hnid;


SELECT p.hnid, firstname, lastname, hid, s.covid_status
FROM COVID.PATIENT_COVID_STATUS as s , COVID.PATIENT as p
WHERE s.covid_status = 'Positive' AND p.hnid = s.hnid
GROUP BY p.hnid;


------count
SELECT COUNT(hnid) FROM COVID.PATIENT_COVID_STATUS
WHERE PATIENT_COVID_STATUS.covid_status = 'Positive';


------count per hospital
SELECT p.hnid, firstname, lastname, s.covid_status, h.hid
FROM COVID.PATIENT_COVID_STATUS as s , COVID.patient as p, COVID.HOSPITAL as h
WHERE s.covid_status = 'Positive' AND h.hid = p.hid
GROUP BY h.hid;


SELECT h.hid, title, p.firstname, lastname, COUNT(h.hid)
FROM COVID.PATIENT_COVID_STATUS as s , COVID.patient as p, COVID.HOSPITAL as h
WHERE s.covid_status = 'Positive' AND p.hnid = s.hnid AND h.hid = p.hid
GROUP BY h.hid;



------find top 3
SELECT h.hid, h.title, COUNT(h.hid)
FROM COVID.PATIENT_COVID_STATUS as s , COVID.patient as p, COVID.HOSPITAL as h
WHERE s.covid_status = 'Positive' AND p.hnid = s.hnid AND h.hid = p.hid
GROUP BY h.hid
ORDER BY COUNT(h.hid) DESC;