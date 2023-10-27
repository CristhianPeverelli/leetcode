# 182. Duplicate Emails
select distinct email 
from person
group by email
having count(email) > 1