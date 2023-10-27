# 183. Customers Who Never Order
select name as "Customers"
from customers as c1
where c1.id not in 
    (select c2.id
    from customers as c2
    inner join orders as o on c2.id = o.customerId )