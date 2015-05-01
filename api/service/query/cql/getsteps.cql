MATCH (s:Step)-[:STEP_OF]->(r:Recipe)
WHERE r.id = {rid}
RETURN s.id as id
     , s.instruction as instruction
     , s.time as time
     , s.timeunit as timeunit
ORDER BY s.id ASC
