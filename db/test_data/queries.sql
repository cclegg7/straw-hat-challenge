SELECT 
    COUNT(category) as num,
    user_id, 
    WEEK(date) as week,
    category, 
    difficulty
FROM climbs
GROUP BY user_id, week, category, difficulty;


SELECT 
    COUNT(user_id) as num,
    user_id, 
    WEEK(date) as week
FROM runs
GROUP BY user_id, week;