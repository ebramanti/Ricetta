MATCH (u:User)<-[:CREATOR]-(r:Recipe)
WHERE r.id = {rid}
RETURN u.handle as handle
