MATCH (u:User)
WHERE u.handle = {handle}
CREATE (r:Recipe {
    id: {id}
  , created: {now}
  , last_modified: {now}
  , title: {title}
  , notes: {notes}
  , cooktime: {cooktime}
  , cooktimeunit: {cooktimeunit}
  , preptime: {preptime}
  , preptimeunit: {preptimeunit}
})
CREATE (r)-[rel:CREATOR]->(u)
RETURN r.id AS id
     , r.title AS title
     , r.notes AS notes
     , r.cooktime AS cooktime
     , r.cooktimeunit AS cooktimeunit
     , r.preptime AS preptime
     , r.preptimeunit AS preptimeunit
