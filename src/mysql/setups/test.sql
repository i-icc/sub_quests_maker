-- EXPLAIN SELECT w.*
-- FROM whats AS w
-- INNER JOIN (
--   SELECT CEIL(RAND() * (SELECT MAX(`id`) FROM whats)) AS `id`
-- ) AS tmp ON w.id >= tmp.id
-- ORDER BY w.id
-- LIMIT 1;

-- SELECT * FROM whats ORDER BY RAND() LIMIT 1;