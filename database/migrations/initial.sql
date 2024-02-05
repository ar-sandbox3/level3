-- This file contains the initial SQL schema for the database.

-- drop tables if they exist. Since this is for testing, we don't want to have any existing tables.
DROP TABLE IF EXISTS employee;
DROP TABLE IF EXISTS department;

-- department table.
CREATE TABLE department (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) UNIQUE
);

-- employee table with foreign key.
CREATE TABLE employee (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100),
  department_id INT NOT NULL,
  -- foreign key constraint
  FOREIGN KEY (department_id) REFERENCES department(id)
);
