MATCH (p:PublicCurator)
WHERE p.name = "PublicCurator"
CREATE UNIQUE (r)-[cur:CURATED_BY]->(p)
