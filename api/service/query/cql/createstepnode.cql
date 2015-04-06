MATCH (r:Recipe)
WHERE r.id = {rid}
CREATE (s:Step {
    id: {id}
  , created: {now}
  , last_modified: {now}
  , instruction: {instruction}
  , time: {time}
  , timeunit: {timeunit}
})
CREATE (s)-[rel:STEP_OF]->(r)
RETURN s.id AS id
     , s.instruction AS instruction
     , s.time AS time
     , s.timeunit AS timeunit
