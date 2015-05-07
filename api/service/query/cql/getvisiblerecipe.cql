MATCH  (u:User)<-[:CREATOR]-(r:Recipe)-[:CURATED_BY]->(pc:PublicCurator)
WHERE  r.id = {rid}
RETURN r.id AS id
     , u.handle as author
     , r.title as title
     , r.notes as notes
     , r.cooktime as cooktime
     , r.cooktimeunit as cooktimeunit
     , r.preptime as preptime
     , r.preptimeunit as preptimeunit
     , r.steps as steps
     , r.private as private
     , r.image as image
     , r.created as created
     , r.last_modified as last_modified
