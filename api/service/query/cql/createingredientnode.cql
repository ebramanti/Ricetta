MATCH (r:Recipe)
WHERE r.id = {rid}
CREATE (i:Ingredient {
    id: {id}
  , created: {now}
  , last_modified: {now}
  , name: {name}
  , amount: {amount}
  , amountunit: {amountunit}
  , url: {url}
})
CREATE (i)-[rel:INGREDIENT_OF]->(r)
RETURN i.id AS id
     , i.name AS name
     , i.amount AS amount
     , i.amountunit AS amountunit
     , i.url AS url
