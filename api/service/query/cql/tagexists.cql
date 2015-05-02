MATCH (t:Tag)
WHERE t.name = {name}
RETURN t.name as name
