MATCH (r:Recipe), (t:Tag)
WHERE r.id = {rid}
AND   t.name = {name}
CREATE (r)-[:TYPE_OF]->(t)
RETURN t.name as name
