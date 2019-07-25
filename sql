
CREATE TABLE questions (
   id serial PRIMARY KEY,
   question varchar NOT NULL,
   dimension varchar NOT NULL
);

CREATE TABLE answers (
   id serial PRIMARY KEY,
   question INT NOT NULL,
   answer INT NOT NULL,
   userid INT NOT NULL
);

CREATE TABLE users (
   id serial PRIMARY KEY,
   email varchar NOT NULL
);

INSERT INTO questions (
    question,
    dimension
)

VALUES
(
    'I feel invigorated from my time being around other people.',
    'extraversion'
),
(
    'I feel comfortable working in groups of people and enjoy it.',
    'extraversion'
),
(
    'Others may describe me as ‘reserved’ or ‘reflective.',
    'introversion'
),
(
    'Sometimes I spend too much time reflecting and do not take action quickly enough.',
    'introversion'
),
(
  'I’m a practical person who tends to be concerned with the bottom line.',
  'sensing'
),
(
    'I prefer to start with the facts/details and then form the big picture.',
    'sensing'
),
(
    'I like to solve problems by leaping between different possibilities and ideas.',
    'intuition'
),
(
    'I pride myself on making decisions with my head and being both fair and consistent.',
    'thinking'
),
(
    'I look for logical explanations and inconsistencies.',
    'thinking'
),
(
  'I make decisions with my heart and strive to be compassionate.',
  'feeling'
),
(
    'I like to get my work done before playing.',
    'judging'
),
(
  'I am energized/stimulated by an approaching deadline.',
  'perceiving'
);