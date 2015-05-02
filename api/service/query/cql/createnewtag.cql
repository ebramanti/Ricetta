MATCH (r:Recipe)
WHERE r.id = {rid}
CREATE (t:Tag {
    name: {name}
})
CREATE (r)-[:TYPE_OF]->(t)
RETURN t.name as name
