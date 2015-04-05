MATCH   (u:User)<-[:SESSION_OF]-(a:AuthToken)
WHERE   a.value   = {token}
AND     a.expires > {now}
RETURN  u.handle
