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
  , image: {image}
})
CREATE (r)-[rel:CREATOR]->(u)
