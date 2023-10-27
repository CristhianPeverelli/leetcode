# 184. Department Highest Salary
select d.name as "Department", e.name as "Employee", e.salary as "Salary"
from employee as e
inner join department as d on e.departmentId = d.id
where (e.salary, e.departmentId) in (
    select max(salary), departmentId
    from employee
    group by departmentId
);