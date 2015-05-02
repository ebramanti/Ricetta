MATCH (t:Tag)<-[:TYPE_OF]-(r:Recipe)
WHERE r.id = {rid}
RETURN t.name as name
ORDER BY t.name ASC
