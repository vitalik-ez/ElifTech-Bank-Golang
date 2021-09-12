CREATE TABLE IF NOT EXISTS banks (
    id serial not null UNIQUE,
    name varchar(255) not null,
    interestRate  real not null,
    maximumLoan  real not null,
    minimumDownPayment  real not null,
    loanTermInMonths  integer not null,
);