MATCH (i:Ingredient)-[:INGREDIENT_OF]->(r:Recipe)
WHERE r.id = {rid}
RETURN i.id as id
     , i.name as name
     , i.url as url
     , i.amount as amount
     , i.amountunit as amountunit
ORDER BY i.id ASC
