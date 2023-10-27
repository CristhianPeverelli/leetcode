# 185. Department Top Three Salaries
select d.name as "Department", 
       e.name as "Employee", 
       e.salary as "Salary"
from employee as e
inner join department as d on e.departmentId = d.id
where (
    select count(distinct salary)
    from employee as e2
    where e2.departmentId = e.departmentId and e2.salary > e.salary
) < 3;

